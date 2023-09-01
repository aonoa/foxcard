package data

import (
	"context"
	"encoding/json"
	"fmt"
	pb "foxcard/api/card/v1"
	"foxcard/internal/biz/card"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

// NewCardRepo .
func NewCardRepo(data *Data, logger log.Logger) card.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cardRepo) IsCardExist(ctx context.Context, card string) (bool, error) {
	val, err := r.data.rdb.Get(ctx, card).Result()
	fmt.Println(val)
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *cardRepo) SaveCardInfo(ctx context.Context, card *pb.CardInfo) (bool, error) {
	strCard, _ := json.Marshal(card)
	fmt.Println(string(strCard))
	val, err := r.data.rdb.Set(ctx, card.CardId, strCard, 0).Result()
	fmt.Println(val)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *cardRepo) DelCardInfo(ctx context.Context, cardId string) (bool, error) {
	val, err := r.data.rdb.Del(ctx, cardId).Result()
	fmt.Println(val)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *cardRepo) GetCardInfo(ctx context.Context, cardId string) (*pb.CardInfo, error) {
	val, err := r.data.rdb.Get(ctx, cardId).Result()
	fmt.Println(val)
	if err != nil {
		return nil, err
	}
	res := &pb.CardInfo{}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *cardRepo) GetLoginStatus(ctx context.Context, id string) (*card.LoginStatusCard, error) {
	val, err := r.data.rdb.Get(ctx, "card-login_"+id).Result()
	fmt.Println(val)
	if err != nil {
		return nil, err
	}
	res := &card.LoginStatusCard{}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *cardRepo) SetLoginStatus(ctx context.Context, status *card.LoginStatusCard, expiration time.Duration) (bool, error) {
	strStatus, _ := json.Marshal(status)
	fmt.Println(string(strStatus))
	val, err := r.data.rdb.Set(ctx, "card-login_"+status.Id, strStatus, expiration).Result()
	fmt.Println(val)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *cardRepo) DelLoginStatus(ctx context.Context, id string) (bool, error) {
	val, err := r.data.rdb.Del(ctx, id).Result()
	fmt.Println(val)
	if err != nil {
		return false, nil
	}
	return true, nil
}
