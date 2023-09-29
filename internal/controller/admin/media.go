package admin

import (
	"context"
	v1 "letga/api/admin/v1"
	"letga/internal/model"
	"letga/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type cMedia struct{}

var Media = cMedia{}

// 获取媒体
func (c *cMedia) Get(ctx context.Context, req *v1.MediaGetReq) (res *v1.MediaGetRes, err error) {
	var (
		out *model.Media
	)
	if out, err = service.Media().GetMedia(ctx, req.Key); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 使用路径获取媒体
func (c *cMedia) ParseUrl(ctx context.Context, req *v1.MediaParseUrlReq) (res *v1.MediaParseUrlRes, err error) {
	var (
		out *model.Media
	)
	if out, err = service.Media().ParseUrl(ctx, req.Url); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 上传媒体
func (c *cMedia) Upload(ctx context.Context, req *v1.MediaUploadReq) (res *v1.MediaUploadRes, err error) {
	var (
		out *model.Media
	)
	if out, err = service.Media().Upload(ctx, req.File); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 设置媒体状态
func (c *cMedia) SetStatus(ctx context.Context, req *v1.MediaSetStatusReq) (res *v1.MediaSetStatusRes, err error) {
	var (
		out *model.Media
	)
	if out, err = service.Media().SetMediaStatus(ctx, &model.StatusSetInput{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}

// 删除媒体
func (c *cMedia) Delete(ctx context.Context, req *v1.MediaDeleteReq) (res *v1.MediaDeleteRes, err error) {
	err = service.Media().DeleteMedias(ctx, req.Keys)
	return
}

// 获取媒体分页
func (c *cMedia) GetPage(ctx context.Context, req *v1.MediaGetPageReq) (res *v1.MediaGetPageRes, err error) {
	var (
		in  *model.MediaPageInput
		out *model.MediaPageOutput
	)
	if err = gconv.Struct(req, &in); err != nil {
		return
	}
	if out, err = service.Media().GetMediaPage(ctx, in); err != nil {
		return
	}
	err = gconv.Struct(out, &res)
	return
}
