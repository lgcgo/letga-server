// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"letga/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMedia interface {
		// 上传媒体
		Upload(ctx context.Context, uploadFile *ghttp.UploadFile) (*model.Media, error)
		// 获取媒体
		GetMedia(ctx context.Context, key string) (*model.Media, error)
		// 解析媒体地址
		ParseUrl(ctx context.Context, url string) (*model.Media, error)
		// 设置用户状态
		SetMediaStatus(ctx context.Context, in *model.StatusSetInput) (*model.Media, error)
		// 获取媒体分页
		GetMediaPage(ctx context.Context, in *model.MediaPageInput) (*model.MediaPageOutput, error)
		// 删除媒体
		DeleteMedias(ctx context.Context, keys []string) error
	}
)

var (
	localMedia IMedia
)

func Media() IMedia {
	if localMedia == nil {
		panic("implement not found for interface IMedia, forgot register?")
	}
	return localMedia
}

func RegisterMedia(i IMedia) {
	localMedia = i
}
