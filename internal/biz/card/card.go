package card

import (
	"context"
	"fmt"
	pb "foxcard/api/card/v1"
	"foxcard/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"time"
)

// CardRepo is a Card repo.
type CardRepo interface {
	// 卡
	IsCardExist(ctx context.Context, cardId string) (bool, error)
	SaveCardInfo(ctx context.Context, card *pb.CardInfo) (bool, error)
	DelCardInfo(ctx context.Context, cardId string) (bool, error)
	GetCardInfo(ctx context.Context, cardId string) (*pb.CardInfo, error)
	// 登陆状态
	GetLoginStatus(ctx context.Context, id string) (*LoginStatusCard, error)
	SetLoginStatus(ctx context.Context, status *LoginStatusCard, expiration time.Duration) (bool, error)
	DelLoginStatus(ctx context.Context, id string) (bool, error)
}

// CardUsecase is a Card usecase.
type CardUsecase struct {
	repo CardRepo
	log  *log.Helper
	key  string
}

// NewCardUsecase new a Card usecase.
func NewCardUsecase(repo CardRepo, auth *conf.Auth, logger log.Logger) *CardUsecase {
	return &CardUsecase{repo: repo, log: log.NewHelper(logger), key: auth.JwtKey}
}

func (s *CardUsecase) CardLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	card, err := s.repo.GetCardInfo(ctx, req.CardId)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	// 判断是否过期
	if now.Unix() > card.ExpiresTs {
		// 过期
		if now.Unix() < card.EffectTimestamp {
			// 还未生效
		}
	}

	if req.AppKey == card.AppKey {
		for _, value := range card.Device {
			fmt.Println(value)
			// 验证通过
			if req.Device == value {
				// 判断无重复登陆
				id := GenerateStatusID(req.CardId, req.Device, req.AppKey)
				_, err := s.repo.GetLoginStatus(ctx, id)
				if err != nil && err != redis.Nil {
					return nil, err
				}
				if err == redis.Nil {
					// 生成登陆状态
					cardLoginStatus := LoginStatusCard{
						Id:        id,
						Timestamp: time.Now().Unix(),
						ExpiresTs: card.ExpiresTs,
						CardId:    card.CardId,
						AppKey:    card.AppKey,
						Device:    req.Device,
					}
					// 登陆状态信息	卡密+设备（3分钟过期）如登陆时存在，即为重复登陆，有心跳时，自动续一下过期时间
					ok, _ := s.repo.SetLoginStatus(ctx, &cardLoginStatus, time.Second*time.Duration(HeartbeatTimeout))
					if ok {
						// 生成token
						token := HeartbeatToken{
							LoginStatusID: id,
							Timestamp:     now.Unix(),
							ExpiresTs:     now.Unix() + HeartbeatTimeout,
							AppKey:        card.AppKey,
							Device:        req.Device,
						}

						// 生成心跳信息
						claims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, jwtv4.MapClaims{
							"token": token,
						})
						signedStr, err := claims.SignedString([]byte(s.key))
						if err != nil {
							return nil, err
						}

						// 保存心跳信息 key:value
						// 返回
						return &pb.LoginReply{
							CardType:   card.CardType,
							Token:      signedStr,
							ExpiresTs:  cardLoginStatus.ExpiresTs,
							Config:     card.Config,
							ServerTime: now.Unix(),
						}, nil

					}
				}

				return &pb.LoginReply{}, nil
			}
		}
	}
	return &pb.LoginReply{}, nil
}

func (s *CardUsecase) CardHeartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {

	now := time.Now()
	oldToken := HeartbeatToken{}
	if claims, ok := jwt.FromContext(ctx); ok {
		//oldToken := (*claims.(*jwtv4.MapClaims))["token"]
		convertMapToStruct((*claims.(*jwtv4.MapClaims))["token"].(map[string]interface{}), &oldToken)
	} else {
		return nil, ErrSignInvalid
	}
	// 判断是否超时
	if now.Unix() > oldToken.ExpiresTs {
		// 过期
		return nil, ErrCardExpire
	}
	// 检查并更新 cardLoginStatus
	status, err := s.repo.GetLoginStatus(ctx, oldToken.LoginStatusID)
	if err != nil {
		// 登陆状态失效
		return nil, ErrCartLoginExpire
	}
	// 判断是否超时
	if now.Unix() > status.ExpiresTs {
		//
		return nil, ErrCardExpire
	}

	// 判断检查设备
	if status.Device != req.Device || req.Device != oldToken.Device {
		//
		return nil, ErrDevice
	}

	// 生成新token
	newToken := HeartbeatToken{}
	copier.Copy(&newToken, oldToken)
	newToken.Timestamp = now.Unix()
	newToken.ExpiresTs = now.Unix() + HeartbeatTimeout

	claims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, jwtv4.MapClaims{
		"token": newToken,
	})
	signedStr, err := claims.SignedString([]byte(s.key))
	if err != nil {
		return nil, err
	}

	// 生成新token
	// 返回
	return &pb.HeartbeatReply{Token: signedStr, ExpiresTs: newToken.ExpiresTs, ServerTime: now.Unix()}, nil
}
func (s *CardUsecase) CardLogout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {

	// 删除cardLoginStatus
	return &pb.LogoutReply{}, nil
}
func (s *CardUsecase) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.CardReply, error) {

	if req.EffectDuration >= MaxDuration {
		return &pb.CardReply{Code: 10001}, nil
	}

	card := pb.CardInfo{
		AppKey:          req.AppKey,
		CardId:          "",
		CardType:        req.CardType,
		EffectTimestamp: req.EffectTimestamp,
		EffectDuration:  req.EffectDuration,
		ExpiresTs:       req.EffectTimestamp + req.EffectDuration,
		Config:          req.Config,
		Status:          0,
		Device:          req.Device,
	}

	for i := 0; i < 3; i++ {
		card.CardId = GenerateCardID()
		fmt.Printf(card.CardId)
		// 判断卡密是否存在
		if ok, _ := s.repo.IsCardExist(ctx, card.CardId); !ok && card.CardId != "" {
			// 如不存在直接添加
			info, err := s.repo.SaveCardInfo(ctx, &card)
			if err != nil {
				return nil, err
			}
			if info {
				return &pb.CardReply{
					Card: []*pb.CardInfo{&card},
				}, nil
			}
		}
	}
	return &pb.CardReply{Code: 10000}, nil
}

func (s *CardUsecase) FrozenCard(ctx context.Context, req *pb.FrozenCardRequest) (*pb.FrozenCardReply, error) {
	return &pb.FrozenCardReply{}, nil
}
func (s *CardUsecase) ThawCard(ctx context.Context, req *pb.ThawCardRequest) (*pb.ThawCardReply, error) {
	return &pb.ThawCardReply{}, nil
}
func (s *CardUsecase) DelCard(ctx context.Context, req *pb.DelCardRequest) (*pb.DelCardReply, error) {

	card, err := s.repo.GetCardInfo(ctx, req.CardId)
	if err != nil {
		return nil, err
	}

	ok, err := s.repo.DelCardInfo(ctx, req.CardId)
	if err != nil {
		return nil, err
	}

	if ok {
		return &pb.DelCardReply{
			AppKey: card.AppKey,
			CardId: req.CardId,
		}, nil
	}

	return &pb.DelCardReply{}, nil
}
func (s *CardUsecase) CloneCard(ctx context.Context, req *pb.CloneCardRequest) (*pb.CardReply, error) {
	return &pb.CardReply{}, nil
}
func (s *CardUsecase) DurationCard(ctx context.Context, req *pb.DurationCardRequest) (*pb.CardReply, error) {
	return &pb.CardReply{}, nil
}
