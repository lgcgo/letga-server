// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package dao

import (
	"context"
	"fmt"
	"letga/internal/biz"
	"letga/internal/global"
	"letga/utility/hashid"
	"log"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
)

// 搜索处理器
func SearchHandler(tableName string, ptr interface{}, search string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		// 组装搜索条件
		searchMeta := gmeta.Get(ptr, "search").String()
		if len(searchMeta) > 0 {
			searchFields := strings.Split(searchMeta, ",")
			for _, searchField := range searchFields {
				if len(search) > 0 {
					m = m.WhereOrLike(searchField, "%"+search+"%")
				}
			}
		}
		// 获取入参的类型
		reType := reflect.TypeOf(ptr)
		// 入参类型校验
		if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
			panic("参数应该为结构体指针")
		}
		// 取指针指向的结构体变量
		v := reflect.ValueOf(ptr).Elem()

		// 解析字段
		for i := 0; i < v.NumField(); i++ {
			// 获取结构体字段信息
			structField := v.Type().Field(i)
			// 取tag
			tag := structField.Tag
			// 解析fliter tag，获取tag值
			f := tag.Get("f")
			if len(f) == 0 {
				continue
			}
			value := fmt.Sprintf("%v", v.Field(i))
			if len(value) == 0 {
				continue
			}
			var (
				tagFiled string
				tagValue string
			)
			if gstr.Contains(f, ":") {
				fs := strings.Split(f, ":")
				tagFiled = fs[0]
				tagValue = fs[1]
			} else {
				tagFiled = f
			}

			switch tagFiled {
			case "hashFc":
				var (
					mapTableName = tableName
					mapFieldName string
					fieldName    string
				)
				// 转换字段名称
				if structField.Name == "Key" {
					fieldName = "id"
				} else {
					fieldName = gstr.TrimRightStr(gstr.CaseSnake(structField.Name), "_key") + "_id"
				}
				if len(tagValue) > 0 {
					tagValues := strings.Split(tagValue, ",")
					if len(tagValues) > 1 {
						mapTableName = tagValues[0]
						mapFieldName = tagValues[1]
					} else {
						mapFieldName = tagValues[0]
					}
				} else {
					mapFieldName = fieldName
				}
				m = hashFc(m, fieldName, value, mapTableName, mapFieldName)
			case "equal":
				m = m.Where(structField.Name, value)
			case "like":
				m = m.WhereLike(structField.Name, "%"+value+"%")
			case "like-start":
				m = m.WhereLike(structField.Name, "%"+value)
			case "like-end":
				m = m.WhereLike(structField.Name, value+"%")
			default:
				// do nothing
			}
		}
		return m
	}
}

// 分页处理器
func PageHandler(pageSize int, pageCurrent int) func(m *gdb.Model) *gdb.Model {
	if pageSize == 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return func(m *gdb.Model) *gdb.Model {
		return m.Page(pageCurrent, pageSize)
	}
}

// 获取Hash转换钩子
func HashSelectHook(ctx context.Context, unionAsTable ...string) gdb.HookHandler {
	var hook = gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			var (
				tableName string
				newField  string
				keyValue  string
			)

			fmt.Println("in.Table ======== ")
			fmt.Println(in.Table)

			if len(in.Table) > 0 {
				// 在多表关联查询时table异常，这里时暂时的修补办法
				tableNameSplit := strings.Split(in.Table, " ")
				// tableName = in.Table
				tableName = gstr.Trim(tableNameSplit[0], "`")
			} else {
				if len(unionAsTable) == 0 {
					return
				}
				tableName = unionAsTable[0]
			}
			// tableNameSplit := strings.Split(in.Table, " ")
			// fmt.Println("tableNameSplit ======== ")
			// fmt.Println(tableNameSplit)
			for i, record := range result {
				fields, ok := global.HashidSaltMap[tableName]
				if !ok {
					return
				}
				for k, v := range fields {
					if !record[k].IsEmpty() {
						if k == "id" {
							newField = "key"
						} else {
							newField = gstr.TrimRightStr(k, "_id") + "_key"
						}
						if keyValue, err = hashid.Encode(record[k].Uint(), v); err != nil {
							return
						}
						record[newField] = gvar.New(keyValue)
					}
				}
				result[i] = record
			}
			return
		},
	}
	return hook
}

// 检测模型IDs
func CheckIds(ctx context.Context, tableName string, keyIds []uint, salt int) error {
	var (
		gs1  = gset.NewFrom(keyIds, true)
		gs2  = gset.New(true)
		vars []*gvar.Var
		err  error
	)

	if vars, err = g.DB().Ctx(ctx).Model(tableName).Fields("id").Where("id IN(?)", keyIds).Array(); err != nil {
		return err
	}
	for _, v := range vars {
		gs2.Add(v.Uint())
	}

	diffSet := gs1.Diff(gs2)
	if diffSet.Size() > 0 {
		var (
			temIds  []uint
			temKeys []string
		)
		diffSet.Iterator(func(v interface{}) bool {
			temIds = append(temIds, v.(uint))
			return true
		})
		if temKeys, err = hashid.BatchEncode(temIds, salt); err != nil {
			return err
		}
		return gerror.NewCode(biz.WithCodef(biz.TableKeyInvalid, strings.Join(temKeys, ",")))
	}

	return nil
}

// hashFc 条件处理函数
func hashFc(m *gdb.Model, field string, value string, mapTableName string, mapfieldName string) *gdb.Model {
	fields, ok := global.HashidSaltMap[mapTableName]
	if !ok {
		log.Fatalf("table: %s not find in hash map", mapTableName)
		// panic("table key not find in hash map")
	}
	salt, ok := fields[mapfieldName]
	if !ok {
		log.Fatalf("table field: %s not find in hash map", mapfieldName)

		// panic("table field %s not find in hash map")
	}
	keyId, _ := hashid.Decode(value, salt)
	return m.Where(field, keyId)
}
