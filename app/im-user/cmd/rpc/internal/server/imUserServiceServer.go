// Code generated by goctl. DO NOT EDIT!
// Source: im-user.proto

package server

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/logic"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
)

type ImUserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedImUserServiceServer
}

func NewImUserServiceServer(svcCtx *svc.ServiceContext) *ImUserServiceServer {
	return &ImUserServiceServer{
		svcCtx: svcCtx,
	}
}

//  获取群组成员列表
func (s *ImUserServiceServer) GetGroupMemberIDListFromCache(ctx context.Context, in *pb.GetGroupMemberIDListFromCacheReq) (*pb.GetGroupMemberIDListFromCacheResp, error) {
	l := logic.NewGetGroupMemberIDListFromCacheLogic(ctx, s.svcCtx)
	return l.GetGroupMemberIDListFromCache(in)
}

//  判断用户A是否在B黑名单中
func (s *ImUserServiceServer) IfAInBBlacklist(ctx context.Context, in *pb.IfAInBBlacklistReq) (*pb.IfAInBBlacklistResp, error) {
	l := logic.NewIfAInBBlacklistLogic(ctx, s.svcCtx)
	return l.IfAInBBlacklist(in)
}

//  判断用户A是否在B好友列表中
func (s *ImUserServiceServer) IfAInBFriendList(ctx context.Context, in *pb.IfAInBFriendListReq) (*pb.IfAInBFriendListResp, error) {
	l := logic.NewIfAInBFriendListLogic(ctx, s.svcCtx)
	return l.IfAInBFriendList(in)
}

//  获取单聊会话的消息接收选项
func (s *ImUserServiceServer) GetSingleConversationRecvMsgOpts(ctx context.Context, in *pb.GetSingleConversationRecvMsgOptsReq) (*pb.GetSingleConversationRecvMsgOptsResp, error) {
	l := logic.NewGetSingleConversationRecvMsgOptsLogic(ctx, s.svcCtx)
	return l.GetSingleConversationRecvMsgOpts(in)
}
