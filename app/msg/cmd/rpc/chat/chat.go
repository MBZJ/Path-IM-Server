// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package chat

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetMaxAndMinSeqReq           = pb.GetMaxAndMinSeqReq
	GetMaxAndMinSeqResp          = pb.GetMaxAndMinSeqResp
	MsgDataToDB                  = pb.MsgDataToDB
	MsgDataToMQ                  = pb.MsgDataToMQ
	PushMsgDataToMQ              = pb.PushMsgDataToMQ
	SendMsgReq                   = pb.SendMsgReq
	SendMsgResp                  = pb.SendMsgResp
	WrapDelMsgListReq            = pb.WrapDelMsgListReq
	WrapDelMsgListResp           = pb.WrapDelMsgListResp
	WrapPullMessageBySeqListReq  = pb.WrapPullMessageBySeqListReq
	WrapPullMessageBySeqListResp = pb.WrapPullMessageBySeqListResp

	Chat interface {
		GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error)
		PullMessageBySeqList(ctx context.Context, in *WrapPullMessageBySeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error)
		SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
		DelMsgList(ctx context.Context, in *WrapDelMsgListReq, opts ...grpc.CallOption) (*WrapDelMsgListResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.GetMaxAndMinSeq(ctx, in, opts...)
}

func (m *defaultChat) PullMessageBySeqList(ctx context.Context, in *WrapPullMessageBySeqListReq, opts ...grpc.CallOption) (*WrapPullMessageBySeqListResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.PullMessageBySeqList(ctx, in, opts...)
}

func (m *defaultChat) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.SendMsg(ctx, in, opts...)
}

func (m *defaultChat) DelMsgList(ctx context.Context, in *WrapDelMsgListReq, opts ...grpc.CallOption) (*WrapDelMsgListResp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.DelMsgList(ctx, in, opts...)
}
