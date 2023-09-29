package admin

type cAuthRoleAccess struct{}

var AuthRoleAccess = cAuthRoleAccess{}

// // 路由选择展柜
// func (c *cAuthRoleAccess) GetPage(ctx context.Context, req *v1.AuthRoleAccessGetPageReq) (res *v1.AuthRoleAccessGetPageRes, err error) {
// 	var (
// 		in  *model.AuthRoleAccessPageInput
// 		out *model.AuthRoleAccessPageOutput
// 	)
// 	if err = gconv.Struct(req, &in); err != nil {
// 		return
// 	}
// 	if out, err = service.Auth().GetRoleAccessPage(ctx, in); err != nil {
// 		return
// 	}
// 	err = gconv.Struct(out, &res)
// 	return
// }
