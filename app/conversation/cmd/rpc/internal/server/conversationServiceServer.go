// Code generated by goctl. DO NOT EDIT!
// Source: conversation.proto

package server

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/logic"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/pb"
)

type ConversationServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedConversationServiceServer
}

func NewConversationServiceServer(svcCtx *svc.ServiceContext) *ConversationServiceServer {
	return &ConversationServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ConversationServiceServer) ModifyConversationField(ctx context.Context, in *pb.ModifyConversationFieldReq) (*pb.ModifyConversationFieldResp, error) {
	l := logic.NewModifyConversationFieldLogic(ctx, s.svcCtx)
	return l.ModifyConversationField(in)
}
