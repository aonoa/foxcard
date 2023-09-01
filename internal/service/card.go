package service

import (
	"context"
	"foxcard/internal/biz/card"

	pb "foxcard/api/card/v1"
)

type CardService struct {
	pb.UnimplementedCardServer

	uc *card.CardUsecase
}

func NewCardService(uc *card.CardUsecase) *CardService {
	return &CardService{uc: uc}
}

func (s *CardService) CardLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return s.uc.CardLogin(ctx, req)
}
func (s *CardService) CardHeartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	return s.uc.CardHeartbeat(ctx, req)
}
func (s *CardService) CardLogout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *CardService) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.CardReply, error) {
	return s.uc.CreateCard(ctx, req)
}
func (s *CardService) FrozenCard(ctx context.Context, req *pb.FrozenCardRequest) (*pb.FrozenCardReply, error) {
	return &pb.FrozenCardReply{}, nil
}
func (s *CardService) ThawCard(ctx context.Context, req *pb.ThawCardRequest) (*pb.ThawCardReply, error) {
	return &pb.ThawCardReply{}, nil
}
func (s *CardService) DelCard(ctx context.Context, req *pb.DelCardRequest) (*pb.DelCardReply, error) {
	return s.uc.DelCard(ctx, req)
}
func (s *CardService) CloneCard(ctx context.Context, req *pb.CloneCardRequest) (*pb.CardReply, error) {
	return &pb.CardReply{}, nil
}
func (s *CardService) DurationCard(ctx context.Context, req *pb.DurationCardRequest) (*pb.CardReply, error) {
	return &pb.CardReply{}, nil
}
