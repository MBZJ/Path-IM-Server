package logic

import (
	"context"
	"encoding/json"
	imuserpb "github.com/Path-IM/Path-IM-Server/app/im-user/cmd/rpc/pb"
	onlinemessagerelayservice "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	gatewaypb "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	"github.com/Path-IM/Path-IM-Server/app/msg-push/cmd/rpc/internal/svc"
	"github.com/Path-IM/Path-IM-Server/app/msg-push/cmd/rpc/pb"
	chatpb "github.com/Path-IM/Path-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server/common/fastjson"
	"github.com/Path-IM/Path-IM-Server/common/types"
	"github.com/Path-IM/Path-IM-Server/common/utils"
	numUtils "github.com/Path-IM/Path-IM-Server/common/utils/num"
	strUtils "github.com/Path-IM/Path-IM-Server/common/utils/str"
	"github.com/Path-IM/Path-IM-Server/common/xtrace"
	"github.com/zeromicro/go-zero/core/mr"
	"go.opentelemetry.io/otel/attribute"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenIMContent struct {
	SessionType int    `json:"sessionType"`
	From        string `json:"from"`
	To          string `json:"to"`
	Seq         uint32 `json:"seq"`
}
type AtContent struct {
	Text       string   `json:"text"`
	AtUserList []string `json:"atUserList"`
	IsAtSelf   bool     `json:"isAtSelf"`
}

var (
	successCount = uint64(0)
	pushTerminal = []int32{types.IOSPlatformID, types.AndroidPlatformID}
)

type PushMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushMsgLogic {
	return &PushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushMsgLogic) PushMsg(in *pb.PushMsgReq) (*pb.PushMsgResp, error) {
	l.MsgToUser(in)
	return &pb.PushMsgResp{ResultCode: 0}, nil
}

func (l *PushMsgLogic) getAllMsgGatewayService() (services []onlinemessagerelayservice.OnlineMessageRelayService, err error) {
	if l.svcCtx.Config.MsgGatewayRpcK8sTarget == "" {
		return onlinemessagerelayservice.GetAllByEtcd(l.ctx, l.svcCtx.Config.MsgGatewayRpc, l.svcCtx.Config.MsgGatewayRpc.Key)
	} else {
		return onlinemessagerelayservice.GetAllByK8s(l.svcCtx.Config.MsgGatewayRpcK8sTarget)
	}
}

func (l *PushMsgLogic) MsgToUser(pushMsg *pb.PushMsgReq) {
	var wsResult []*gatewaypb.SingleMsgToUser
	isOfflinePush := utils.GetSwitchFromOptions(pushMsg.MsgData.Options, types.IsOfflinePush)

	services, err := l.getAllMsgGatewayService()
	if err != nil {
		l.Errorf("getAllMsgGatewayService error: %v", err)
		err = nil
	}
	var fs []func() error
	for index, msgClient := range services {
		fs = append(fs, func() error {
			var reply *gatewaypb.OnlinePushMsgResp
			var err error
			xtrace.StartFuncSpan(l.ctx, "MsgToUser.OnlinePushMsg", func(ctx context.Context) {
				reply, err = msgClient.OnlinePushMsg(ctx, &gatewaypb.OnlinePushMsgReq{MsgData: pushMsg.MsgData, PushToUserID: pushMsg.PushToUserID})
			}, attribute.Int("index", index))
			if err != nil {
				l.Errorf("OnlinePushMsg error: %v", err)
				return nil
			}
			if reply != nil && reply.Resp != nil {
				wsResult = append(wsResult, reply.Resp...)
			}
			return nil
		})
	}
	_ = mr.Finish(fs...)
	l.Info("push_result ", wsResult, " sendData ", pushMsg.MsgData)
	successCount++
	if isOfflinePush && pushMsg.PushToUserID != pushMsg.MsgData.SendID {
		for _, v := range wsResult {
			if v.ResultCode == 0 {
				continue
			}
			if numUtils.IsContainInt32(v.RecvPlatFormID, pushTerminal) {
				//Use offline push messaging
				var UIDList []string
				UIDList = append(UIDList, v.RecvID)
				customContent := OpenIMContent{
					SessionType: int(pushMsg.MsgData.SessionType),
					From:        pushMsg.MsgData.SendID,
					To:          pushMsg.MsgData.RecvID,
					Seq:         pushMsg.MsgData.Seq,
				}
				bCustomContent, _ := json.Marshal(customContent)
				jsonCustomContent := string(bCustomContent)
				var content string
				if pushMsg.MsgData.OfflinePushInfo != nil {
					content = pushMsg.MsgData.OfflinePushInfo.Title

				} else {
					switch pushMsg.MsgData.ContentType {
					case types.Text:
						content = types.ContentType2PushContent[types.Text]
						if pushMsg.MsgData.AtUserIDList != nil {
							if strUtils.IsContain(v.RecvID, pushMsg.MsgData.AtUserIDList) {
								content = "[有人@你]" + types.ContentType2PushContent[types.Common]
							} else {
								content = types.ContentType2PushContent[types.GroupMsg]
							}
						}
					case types.Picture:
						content = types.ContentType2PushContent[types.Picture]
					case types.Voice:
						content = types.ContentType2PushContent[types.Voice]
					case types.Video:
						content = types.ContentType2PushContent[types.Video]
					case types.File:
						content = types.ContentType2PushContent[types.File]
					default:
						content = types.ContentType2PushContent[types.Common]
					}
				}
				var err error
				xtrace.StartFuncSpan(l.ctx, "MsgToUser.OfflinePushMsg", func(ctx context.Context) {
					_, err = l.svcCtx.GetOfflinePusher().Push(ctx, UIDList, content, jsonCustomContent)
				})
				if err != nil {
					l.Error("offline push error ", pushMsg.String(), err.Error())
				}
				break
			}
		}
	}
}

func (l *PushMsgLogic) PushSuperGroupMsg(in *chatpb.PushMsgToSuperGroupDataToMQ) (*pb.PushMsgResp, error) {
	isOfflinePush := utils.GetSwitchFromOptions(in.MsgData.Options, types.IsOfflinePush)

	a := AtContent{}
	tagAll := false
	// 如果艾特人了
	if in.MsgData.AtUserIDList != nil {
		tagAll = strUtils.IsContain(types.AtAllString, in.MsgData.AtUserIDList)
		_ = fastjson.Unmarshal(in.MsgData.Content, &a)
	}
	// 被艾特的人 先去获取被艾特的人是否屏蔽了群消息
	var atUsers *imuserpb.GetUserListFromSuperGroupWithOptResp
	atPushUserChan := make(chan string, 1)
	go l.listenAtPushUserChan(atPushUserChan, in)
	var err error
	if tagAll {
		xtrace.StartFuncSpan(l.ctx, "PushSuperGroupMsg.GetUserListFromSuperGroupWithOpt", func(ctx context.Context) {
			// 我们去查询这个群的所有接收消息通知的用户
			atUsers, err = l.svcCtx.ImUserService.GetUserListFromSuperGroupWithOpt(l.ctx, &imuserpb.GetUserListFromSuperGroupWithOptReq{
				SuperGroupID: in.SuperGroupID,
				Opts: []imuserpb.RecvMsgOpt{
					imuserpb.RecvMsgOpt_ReceiveMessage,
					imuserpb.RecvMsgOpt_ReceiveNotNotifyMessage,
				},
			})
		})
		if err == nil {
			l.pushSuperGroupMsg(in, atUsers, nil, isOfflinePush, atPushUserChan)
		} else {
			logx.WithContext(l.ctx).Error("GetUserListFromSuperGroupWithOpt failed, err: ", err)
			err = nil
		}
	} else if len(a.AtUserList) > 0 {
		xtrace.StartFuncSpan(l.ctx, "PushSuperGroupMsg.GetUserListFromSuperGroupWithOpt", func(ctx context.Context) {
			// 我们去查询这个群的所有接收消息通知的用户
			atUsers, err = l.svcCtx.ImUserService.GetUserListFromSuperGroupWithOpt(l.ctx, &imuserpb.GetUserListFromSuperGroupWithOptReq{
				SuperGroupID: in.SuperGroupID,
				Opts: []imuserpb.RecvMsgOpt{
					imuserpb.RecvMsgOpt_ReceiveNotNotifyMessage,
				},
				UserIDList: a.AtUserList,
			})
		})
		if err == nil {
			var verifyAtUsers = &imuserpb.GetUserListFromSuperGroupWithOptResp{
				CommonResp:    &imuserpb.CommonResp{},
				UserIDOptList: nil,
			}
			for _, opt := range atUsers.UserIDOptList {
				if strUtils.IsContain(opt.UserID, a.AtUserList) {
					verifyAtUsers.UserIDOptList = append(verifyAtUsers.UserIDOptList, opt)
				}
			}
			l.pushSuperGroupMsg(in, verifyAtUsers, nil, isOfflinePush, atPushUserChan)
		} else {
			logx.WithContext(l.ctx).Error("GetUserListFromSuperGroupWithOpt failed, err: ", err)
			err = nil
		}
	}
	if tagAll {
		return &pb.PushMsgResp{ResultCode: 0}, nil
	}
	var allUsers *imuserpb.GetUserListFromSuperGroupWithOptResp
	offlinePushUserChan := make(chan string, 1)
	xtrace.StartFuncSpan(l.ctx, "PushSuperGroupMsg.GetUserListFromSuperGroupWithOpt", func(ctx context.Context) {
		// 我们去查询这个群的所有接收消息通知的用户
		allUsers, err = l.svcCtx.ImUserService.GetUserListFromSuperGroupWithOpt(l.ctx, &imuserpb.GetUserListFromSuperGroupWithOptReq{
			SuperGroupID: in.SuperGroupID,
			Opts: []imuserpb.RecvMsgOpt{
				imuserpb.RecvMsgOpt_ReceiveMessage,
			},
		})
	})
	if err != nil {
		return nil, err
	}
	l.Info("allUsers.UserIDOptList:", allUsers.UserIDOptList)

	go l.listenOfflinePushUserChan(offlinePushUserChan, in)
	l.pushSuperGroupMsg(in, allUsers, a.AtUserList, isOfflinePush, offlinePushUserChan)
	return &pb.PushMsgResp{ResultCode: 0}, nil
}

func (l *PushMsgLogic) pushSuperGroupMsg(
	in *chatpb.PushMsgToSuperGroupDataToMQ,
	users *imuserpb.GetUserListFromSuperGroupWithOptResp,
	atList []string,
	isOfflinePush bool,
	offlinePushUserChan chan string,
) {
	services, _ := l.getAllMsgGatewayService()
	go func() {
		defer func() {
			offlinePushUserChan <- string([]byte{2})
		}()
		for uIndex, user := range users.UserIDOptList {
			if strUtils.IsContain(user.UserID, atList) {
				// 跳过被艾特的人
				continue
			}
			{
				allServiceFailed := true
				var fs []func() error
				for i, service := range services {
					fs = append(fs, func() error {
						allPlatformIsFailed := true
						xtrace.StartFuncSpan(l.ctx, "PushSuperGroupMsg.PushMsgToUser", func(ctx context.Context) {
							resp, err := service.OnlinePushMsg(ctx, &gatewaypb.OnlinePushMsgReq{
								MsgData:      in.MsgData,
								PushToUserID: user.UserID,
							})
							if err != nil {
								l.Errorf("PushMsgToUser error: %v", err)
								return
							}
							if resp == nil || resp.Resp == nil {
								l.Errorf("PushMsgToUser error: resp == nil")
								return
							}
							for _, res := range resp.Resp {
								// 是否全部平台都失败了
								if res.ResultCode != -1 {
									// 成功了
									allPlatformIsFailed = false
									break
								}
							}
						},
							attribute.Int("user.index", uIndex),
							attribute.Int("services.index", i),
							attribute.String("user.id", user.UserID),
						)
						if !allPlatformIsFailed {
							allServiceFailed = false
						}
						return nil
					})
				}
				_ = mr.Finish(fs...)
				if allServiceFailed {
					// 这条消息要不要离线推送
					if isOfflinePush && in.MsgData.SendID != user.UserID {
						offlinePushUserChan <- user.UserID
					}
				}
			}
		}
	}()
}

func (l *PushMsgLogic) listenOfflinePushUserChan(
	userChan chan string,
	pushMsg *chatpb.PushMsgToSuperGroupDataToMQ,
) {
	var uids []string
	for uid := range userChan {
		bytes := []byte(uid)
		if len(bytes) == 1 && bytes[0] == 2 {
			break
		}
		uids = append(uids, uid)
	}
	logx.WithContext(l.ctx).Info("开始进行离线推送:", uids)
	customContent := OpenIMContent{
		SessionType: int(pushMsg.MsgData.SessionType),
		From:        pushMsg.MsgData.SendID,
		To:          pushMsg.MsgData.RecvID,
		Seq:         pushMsg.MsgData.Seq,
	}
	bCustomContent, _ := json.Marshal(customContent)
	jsonCustomContent := string(bCustomContent)
	var content string
	if pushMsg.MsgData.OfflinePushInfo != nil {
		content = pushMsg.MsgData.OfflinePushInfo.Title

	} else {
		switch pushMsg.MsgData.ContentType {
		case types.Text:
			content = types.ContentType2PushContent[types.Text]
			if pushMsg.MsgData.AtUserIDList != nil {
				content = types.ContentType2PushContent[types.GroupMsg]
			}
		case types.Picture:
			content = types.ContentType2PushContent[types.Picture]
		case types.Voice:
			content = types.ContentType2PushContent[types.Voice]
		case types.Video:
			content = types.ContentType2PushContent[types.Video]
		case types.File:
			content = types.ContentType2PushContent[types.File]
		default:
			content = types.ContentType2PushContent[types.Common]
		}
	}
	var err error
	xtrace.StartFuncSpan(l.ctx, "MsgToUser.OfflinePushMsg", func(ctx context.Context) {
		_, err = l.svcCtx.GetOfflinePusher().Push(ctx, uids, content, jsonCustomContent)
	})
	if err != nil {
		l.Error("offline push error ", pushMsg.String(), err.Error())
	}
}

func (l *PushMsgLogic) listenAtPushUserChan(
	userChan chan string,
	pushMsg *chatpb.PushMsgToSuperGroupDataToMQ,
) {
	var uids []string
	for uid := range userChan {
		bytes := []byte(uid)
		if len(bytes) == 1 && bytes[0] == 2 {
			break
		}
		uids = append(uids, uid)
	}
	logx.WithContext(l.ctx).Info("开始进行at离线推送:", uids)
	customContent := OpenIMContent{
		SessionType: int(pushMsg.MsgData.SessionType),
		From:        pushMsg.MsgData.SendID,
		To:          pushMsg.MsgData.RecvID,
		Seq:         pushMsg.MsgData.Seq,
	}
	bCustomContent, _ := json.Marshal(customContent)
	jsonCustomContent := string(bCustomContent)
	var content string
	if pushMsg.MsgData.OfflinePushInfo != nil {
		content = pushMsg.MsgData.OfflinePushInfo.Title
	} else {
		a := AtContent{}
		_ = fastjson.Unmarshal(pushMsg.MsgData.Content, &a)
		content = types.ContentType2PushContent[types.GroupMsg]
	}
	var err error
	xtrace.StartFuncSpan(l.ctx, "MsgToUser.OfflinePushMsg", func(ctx context.Context) {
		_, err = l.svcCtx.GetOfflinePusher().Push(ctx, uids, content, jsonCustomContent)
	})
	if err != nil {
		l.Error("offline push error ", pushMsg.String(), err.Error())
	}
}
