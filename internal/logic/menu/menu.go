// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package auth

import (
	"context"
	"letga/internal/biz"
	"letga/internal/consts"
	"letga/internal/dao"
	"letga/internal/model"
	"letga/internal/model/do"
	"letga/internal/service"
	"letga/utility/hashid"
	"letga/utility/tree"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct {
}

func init() {
	service.RegisterMenu(New())
}

func New() service.IMenu {
	return &sMenu{}
}

// 添加菜单
func (s *sMenu) CreateMenu(ctx context.Context, in *model.MenuCreateInput) (*model.Menu, error) {
	var (
		data      *do.Menu
		menu      *model.Menu
		parent    *model.Menu
		err       error
		available bool
	)
	// 获取父菜单
	if len(in.ParentKey) > 0 {
		if parent, err = s.GetMenu(ctx, in.ParentKey); err != nil {
			return nil, err
		}
	}
	// 标题防重
	if available, err = dao.Menu.IsTitleAvailable(ctx, in.Title); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.MenuNameExists)
	}
	// 写入数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if parent != nil {
		data.ParentId = parent.Id
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		menu, err = dao.Menu.Insert(ctx, data)
		return err
	}); err != nil {
		return nil, err
	}
	return menu, nil
}

// 获取菜单
func (s *sMenu) GetMenu(ctx context.Context, key string) (*model.Menu, error) {
	var (
		keyId uint
		err   error
	)
	// Key解密
	if keyId, err = hashid.Decode(key, consts.TABLE_MENU_SALT); err != nil {
		return nil, err
	}
	return dao.Menu.Get(ctx, keyId)
}

// 修改菜单
func (s *sMenu) UpdateMenu(ctx context.Context, in *model.MenuUpdateInput) (*model.Menu, error) {
	var (
		handler   *tree.Tree
		list      []*model.Menu
		data      *do.Menu
		parent    *model.Menu
		menu      *model.Menu
		err       error
		available bool
	)
	// 扫描数据
	if menu, err = s.GetMenu(ctx, in.Key); err != nil {
		return nil, err
	}
	// 检测父级
	if len(in.ParentKey) > 0 {
		var childKeys []string
		if parent, err = s.GetMenu(ctx, in.ParentKey); err != nil {
			return nil, err
		}
		// 过滤自身
		if in.Key == parent.Key {
			return nil, gerror.NewCode(biz.MenuParentInvalid)
		}
		// 获取所有数据
		if list, err = dao.Menu.GetAll(ctx); err != nil {
			return nil, err
		}
		// 获取树句柄
		if handler, err = tree.NewWithData(list); err != nil {
			return nil, err
		}
		// 过滤子级
		if childKeys, err = handler.GetChildKeys(parent.Key); err != nil {
			return nil, err
		}
		for _, v := range childKeys {
			if in.ParentKey == v {
				return nil, gerror.NewCode(biz.MenuParentInvalid)
			}
		}
	}
	// 标题防重
	if available, err = dao.Menu.IsTitleAvailable(ctx, in.Title, []uint{menu.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.NewCode(biz.MenuNameExists)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if parent != nil {
		data.ParentId = parent.Id
	}
	data.Id = menu.Id
	// 更新实体
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		menu, err = dao.Menu.Update(ctx, data)
		return err
	}); err != nil {
		return nil, err
	}
	return menu, nil
}

// 更改菜单状态
func (s *sMenu) SetRoleStatus(ctx context.Context, in *model.StatusSetInput) (*model.Menu, error) {
	var (
		handler *tree.Tree
		list    []*model.Menu
		menu    *model.Menu
		err     error
	)
	if menu, err = s.GetMenu(ctx, in.Key); err != nil {
		return nil, err
	}
	// 获取所有数据
	if list, err = dao.Menu.GetAll(ctx); err != nil {
		return nil, err
	}
	// 获取树句柄
	if handler, err = tree.NewWithData(list); err != nil {
		return nil, err
	}
	// 获取子Key集
	var (
		keyIds      []uint
		childKeys   []string
		childKeyIds []uint
	)
	if childKeys, err = handler.GetChildKeys(menu.Key); err != nil {
		return nil, err
	}
	// Key集解密
	if childKeyIds, err = hashid.BatchDecode(childKeys, consts.TABLE_MENU_SALT); err != nil {
		return nil, err
	}
	keyIds = append(childKeyIds, menu.Id)
	for _, v := range keyIds {
		if _, err = dao.Menu.SetStatus(ctx, v, in.Value); err != nil {
			return nil, err
		}
	}
	return s.GetMenu(ctx, in.Key)
}

// 删除菜单
func (s *sMenu) DeleteMenus(ctx context.Context, keys []string) error {
	var (
		handler *tree.Tree
		list    []*model.Menu
		keyIds  []uint
		err     error
	)
	// Key集解密
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_MENU_SALT); err != nil {
		return err
	}
	// 检测ID集
	if err = dao.Menu.CheckIds(ctx, keyIds); err != nil {
		return err
	}
	// 获取所有数据
	if list, err = dao.Menu.GetAll(ctx); err != nil {
		return err
	}
	// 获取树句柄
	if handler, err = tree.NewWithData(list); err != nil {
		return err
	}
	// 合并子Key集
	var temKeys = keys
	for _, v := range keys {
		var tempChildKeys []string
		if tempChildKeys, err = handler.GetChildKeys(v); err != nil {
			return err
		}
		temKeys = append(temKeys, tempChildKeys...)
	}
	keys = append(keys, temKeys...)
	if keyIds, err = hashid.BatchDecode(keys, consts.TABLE_MENU_SALT); err != nil {
		return err
	}
	// 嵌套事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除实体
		if err = tx.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return dao.Menu.Delete(ctx, keyIds)
		}); err != nil {
			return err
		}
		// 重置关联数据menu_id
		if err = dao.AuthRoute.ResetMenuId(ctx, keyIds); err != nil {
			return err
		}
		return nil
	})
}

// 获取菜单树
func (s *sMenu) GetMenuTree(ctx context.Context, in *model.MenuTreeInput) (*model.TreeOutput, error) {
	var (
		handler         *tree.Tree
		list            []*model.Menu
		data            []*tree.TreeNode
		ctxDisabledKeys []string
		err             error
	)
	// 获取所有数据
	if list, err = dao.Menu.GetAll(ctx); err != nil {
		return nil, err
	}
	// 获取树句柄
	if handler, err = tree.NewWithData(list); err != nil {
		return nil, err
	}
	// 获取树数据
	if data, err = handler.Search(in.Search, in.Keys); err != nil {
		return nil, err
	}

	var (
		menu               *model.Menu
		menuChildKeys      []string
		menuChildKeysArray *garray.StrArray
	)
	// 获取menuForm中的菜单上下文
	if in.CtxScene == "menuForm" && len(in.CtxMenuKey) > 0 {
		if menu, err = s.GetMenu(ctx, in.CtxMenuKey); err != nil {
			return nil, err
		}
		if menuChildKeys, err = handler.GetChildKeys(menu.Key); err != nil {
			return nil, err
		}
		menuChildKeysArray = garray.NewStrArrayFrom(menuChildKeys)
	}

	// 根据上下文组装禁用以及选中项
	for _, v := range list {
		switch in.CtxScene {
		case "mainTable":
			// 不做处理
		case "menuForm":
			// 禁用自身与下级
			if menu != nil && (menuChildKeysArray.Contains(v.Key) || v.Key == menu.Key) {
				ctxDisabledKeys = append(ctxDisabledKeys, v.Key)
			}
		}
	}

	return &model.TreeOutput{
		Data:            data,
		CtxDisabledKeys: ctxDisabledKeys,
	}, nil
}

// 路由统计
func (s *sMenu) GetFliterTree(ctx context.Context, in *model.MenuFliterTreeInput) (*model.FliterTree, error) {
	var (
		data            []*model.FliterTreeItem
		err             error
		count           int
		handler         *tree.Tree
		ctxSelectedKeys []string
		ctxDisabledKeys []string
	)
	// 组装菜单组
	var (
		menuList       []*model.Menu
		menuTreeSource []*model.FliterTreeData
		menuKeys       []string
		menuIDs        []uint
	)
	if menuList, err = dao.Menu.GetAll(ctx); err != nil {
		return nil, err
	}
	if handler, err = tree.NewWithData(menuList); err != nil {
		return nil, gerror.NewCode(biz.TreeInitialFail)
	}
	for _, v := range menuList {
		if menuKeys, err = handler.GetChildKeys(v.Key); err != nil {
			return nil, err
		}
		menuKeys = append(menuKeys, v.Key)
		if menuIDs, err = hashid.BatchDecode(menuKeys, consts.TABLE_MENU_SALT); err != nil {
			return nil, err
		}
		switch in.CtxScene {
		case "routeSidebar":
			if count, err = dao.AuthRoute.CountByMenuIds(ctx, menuIDs); err != nil {
				return nil, err
			}
			menuTreeSource = append(menuTreeSource, &model.FliterTreeData{
				Key:       v.Key,
				ParentKey: v.ParentKey,
				Title:     v.Title,
				Count:     count,
				Weight:    v.Weight,
			})
			// 可用根据上下文需求设置选中以及禁用项

		}
	}
	if handler, err = tree.NewWithData(menuTreeSource); err != nil {
		return nil, gerror.NewCode(biz.TreeInitialFail)
	}
	data = append(data, &model.FliterTreeItem{
		Name:  "menu",
		Data:  handler.GetWholeTree(),
		Total: len(menuTreeSource),
	})
	return &model.FliterTree{
		Data:            data,
		CtxSelectedKeys: ctxSelectedKeys,
		CtxDisabledKeys: ctxDisabledKeys,
	}, nil
}
