// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package server

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/internal/logic"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) GetMaxAndMinSeq(ctx context.Context, in *pb.GetMaxAndMinSeqReq) (*pb.GetMaxAndMinSeqResp, error) {
	l := logic.NewGetMaxAndMinSeqLogic(ctx, s.svcCtx)
	return l.GetMaxAndMinSeq(in)
}

func (s *ChatServer) PullMessageBySeqList(ctx context.Context, in *pb.WrapPullMessageBySeqListReq) (*pb.WrapPullMessageBySeqListResp, error) {
	l := logic.NewPullMessageBySeqListLogic(ctx, s.svcCtx)
	return l.PullMessageBySeqList(in)
}

func (s *ChatServer) SendMsg(ctx context.Context, in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	l := logic.NewSendMsgLogic(ctx, s.svcCtx)
	return l.SendMsg(in)
}

func (s *ChatServer) DelMsgList(ctx context.Context, in *pb.WrapDelMsgListReq) (*pb.WrapDelMsgListResp, error) {
	l := logic.NewDelMsgListLogic(ctx, s.svcCtx)
	return l.DelMsgList(in)
}
