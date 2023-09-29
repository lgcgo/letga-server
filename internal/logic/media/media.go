// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package media

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"letga/internal/biz"
	"letga/internal/consts"
	"letga/internal/dao"
	"letga/internal/model"
	"letga/internal/model/do"
	"letga/internal/service"
	"letga/utility/hashid"
	"mime/multipart"
	"strings"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/h2non/filetype"
)

func init() {
	service.RegisterMedia(New())
}

type sMedia struct {
	staticServerRoot string
	allowFileType    []string
	allowMimeType    []string
	baseDir          string
	// localAddress     string
}

func New() service.IMedia {
	var (
		ctx = gctx.New()
	)
	staticServerRoot := g.Cfg().MustGet(ctx, "server.static.serverRoot").String()
	allowFileTypeCfg := g.Cfg().MustGet(ctx, "upload.allowFileType").String()
	allowMimeTypeCfg := g.Cfg().MustGet(ctx, "upload.allowMimeType").String()
	return &sMedia{
		staticServerRoot: staticServerRoot,
		allowFileType:    strings.Split(allowFileTypeCfg, ","),
		allowMimeType:    strings.Split(allowMimeTypeCfg, ","),
		baseDir:          g.Cfg().MustGet(ctx, "upload.baseDir").String(),
	}
}

// 上传媒体
func (s *sMedia) Upload(ctx context.Context, uploadFile *ghttp.UploadFile) (*model.Media, error) {
	var (
		user  *model.User
		data  *do.Media
		media *model.Media
		file  multipart.File
		err   error
	)
	// 获取用户
	if user, err = service.User().GetCurrentUser(ctx); err != nil {
		return nil, err
	}
	if file, err = uploadFile.Open(); err != nil {
		return nil, gerror.NewCode(biz.MediaReadFail)
	}
	defer file.Close()
	// filetype读取文件信息只需要261长度
	head := make([]byte, 261)
	file.Read(head)
	// 读取文件类型
	kind, _ := filetype.Match(head)
	if kind == filetype.Unknown {
		return nil, gerror.NewCode(biz.MediaUnknownType)
	}
	// 验证文件类型
	if !gstr.InArray(s.allowFileType, kind.Extension) {
		return nil, gerror.NewCode(biz.MediaFileTypeInvalid)
	}
	// 验证MIME类型
	if !gstr.InArray(s.allowMimeType, kind.MIME.Value) {
		return nil, gerror.NewCode(biz.MediaMimeTypeInvalid)
	}
	// 哈希防重
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, gerror.NewCode(biz.MediaHashCalculationFail)
	}
	hashString := fmt.Sprintf("%x", hash.Sum(nil))
	if media, err = dao.Media.GetByHash(ctx, user.Id, hashString); err != nil {
		return nil, err
	}
	// 已存在则直接返回
	if media != nil {
		return media, nil
	}

	// 组装路径
	basePath := s.staticServerRoot + s.baseDir
	dateDirName := gtime.Now().Format("Ymd")
	reName, err := uploadFile.Save(gfile.Join(basePath, dateDirName), true)
	if err != nil {
		return nil, err
	}
	filePath := s.baseDir + "/" + dateDirName + "/" + reName
	if gstr.HasPrefix(filePath, "./") {
		filePath = gstr.SubStr(filePath, 1, len(filePath)-1)
	}
	// 组装写入数据
	data = &do.Media{
		UserId:   user.Id,
		Name:     uploadFile.Filename,
		Path:     filePath,
		Size:     uploadFile.Size,
		FileType: kind.Extension,
		MimeType: kind.MIME.Value,
		Hash:     hashString,
		Storage:  "local", // 这里暂时固定本地存储
	}
	return dao.Media.Insert(ctx, data)
}

// 获取媒体
func (s *sMedia) GetMedia(ctx context.Context, key string) (*model.Media, error) {
	var (
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_MEDIA_SALT); err != nil {
		return nil, err
	}
	return dao.Media.Get(ctx, keyId)
}

// 解析媒体地址
func (s *sMedia) ParseUrl(ctx context.Context, url string) (*model.Media, error) {
	var (
		media *model.Media
		data  map[string]string
		err   error
	)
	if data, err = gurl.ParseURL(url, -1); err != nil {
		return nil, err
	}
	// map[fragment: host:localhost pass: path:/upload/20230903/cv8o32pcdhn0fcvif7.png port:8000 query: scheme:http user:]
	if media, err = dao.Media.GetByPath(ctx, data["path"]); err != nil {
		return nil, err
	}
	if media == nil {
		return nil, gerror.NewCode(biz.WithCodef(biz.MediaNotExists, url))
	}
	return media, nil
}

// 设置用户状态
func (s *sMedia) SetMediaStatus(ctx context.Context, in *model.StatusSetInput) (*model.Media, error) {
	var (
		media *model.Media
		err   error
	)
	if media, err = s.GetMedia(ctx, in.Key); err != nil {
		return nil, err
	}
	// 限制禁止条件
	// if  {
	// }
	return dao.Media.SetStatus(ctx, media.Id, in.Value)
}

// 获取媒体分页
func (s *sMedia) GetMediaPage(ctx context.Context, in *model.MediaPageInput) (*model.MediaPageOutput, error) {
	var (
		data            []*model.Media
		count           int
		ctxDisabledKeys []string
		err             error
	)
	if data, count, err = dao.Media.PageData(ctx, &in.PageParams, &in.MediaSearch); err != nil {
		return nil, err
	}
	// 禁用项
	// for _, v := range data {
	// }
	return &model.MediaPageOutput{
		Data: data,
		Page: model.Page{
			Total:           count,
			CtxDisabledKeys: ctxDisabledKeys,
		},
	}, nil
}

// 删除媒体
func (s *sMedia) DeleteMedias(ctx context.Context, keys []string) error {
	var (
		keyIds []uint
		err    error
	)
	// Key集解密
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_MEDIA_SALT); err != nil {
		return err
	}
	// 检测ID集
	if err = dao.Media.CheckIds(ctx, keyIds); err != nil {
		return err
	}
	return dao.Media.Delete(ctx, keyIds)
}
