package v1

import (
	"letga/api/admin/v1/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 获取媒体
type MediaGetReq struct {
	g.Meta `path:"/media" method:"get" tags:"Admin@MediaService" summary:"Get media"`
	Key    string `json:"key" v:"required"` // 索引
}
type MediaGetRes struct {
	model.Media
}

// 解析媒体
type MediaParseUrlReq struct {
	g.Meta `path:"/media/parser" method:"get" tags:"Admin@MediaService" summary:"Parse media url"`
	Url    string `json:"url" v:"required"`
}
type MediaParseUrlRes struct {
	model.Media
}

// 媒体上传
type MediaUploadReq struct {
	g.Meta `path:"/media" method:"post" mime:"multipart/form-data" tags:"Admin@MediaService" summary:"UpLoad media"`
	File   *ghttp.UploadFile `json:"file" type:"file"`
}
type MediaUploadRes struct {
	model.Media
}

// 设置媒体状态
type MediaSetStatusReq struct {
	g.Meta `path:"/media/status" method:"put" tags:"Admin@MediaService" summary:"Set media status"`
	Key    string `json:"key" v:"required"`
	Value  string `json:"value" v:"required|in:normal,disabled"`
}
type MediaSetStatusRes struct {
	model.Media
}

// 删除媒体
type MediaDeleteReq struct {
	g.Meta `path:"/media" method:"delete" tags:"Admin@MediaService" summary:"Delete media"`
	Keys   []string `json:"keys" v:"required"`
}
type MediaDeleteRes struct{}

// 媒体分页
type MediaGetPageReq struct {
	g.Meta `path:"/medias" method:"get" tags:"Admin@MediaService" summary:"Get media page"`
	model.PageParams
	Key      string `json:"key"`      // 搜索字段,索引
	UserKey  string `json:"userKey"`  // 搜索字段,用户索引
	Name     string `json:"name"`     // 搜索字段,文件名
	Hash     string `json:"hash"`     // 搜索字段,哈希值
	FileType string `json:"fileType"` // 搜索字段,文件类型
	Storage  string `json:"storage"`  // 搜索字段,储存库
	Status   string `json:"status"`   // 搜索字段,状态
}
type MediaGetPageRes struct {
	Data  []*model.Media `json:"data,omitempty"`
	Total int            `json:"total"`
}
