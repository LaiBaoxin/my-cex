// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"github.com/wwater/my-cex/app/account/api/internal/model"

	"github.com/wwater/my-cex/app/account/api/internal/svc"
	"github.com/wwater/my-cex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogsLogic {
	return &GetLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLogs 获取日志
func (l *GetLogsLogic) GetLogs(req *types.GetLogsReq) (resp *types.GetLogsResp, err error) {
	var logs []model.SystemLog

	// 从数据库查询最近 10 条日志
	if err = l.svcCtx.DB.Order("id desc").Limit(req.Limit).Find(&logs).Error; err != nil {
		return nil, err
	}

	list := make([]types.LogItem, 0)
	for _, v := range logs {
		list = append(list, types.LogItem{
			Id:        v.Id,
			UserId:    v.Uid,
			OpType:    v.OpType,
			Content:   v.Content,
			Amount:    fmt.Sprintf("%.4f", v.Amount),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetLogsResp{
		List: list,
	}, nil
}
