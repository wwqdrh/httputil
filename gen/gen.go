package gen

import (
	"fmt"
	"strings"
)

type SwagInfo struct {
	Url         string
	Method      string
	Tag         string // 标签
	Summary     string // 概扩
	Description string // 描述
	Request     string // 只需要指定变量的路径名字
	Response    string // 只需要指定变量的路径名字
	SlotName    string // 函数方法的slotname
}

// example:
// API = []struct {
// 	method  string
// 	url     string
// 	handler []gin.HandlerFunc
// }{
// 	{"GET", "/api/etcd/conf/verify", etcd.ConfVerify{}},
// 	{"GET", "/api/etcd/conf/status", etcd.ConfStatus{}},
// 	{"POST", "/api/etcd/conf/update", etcd.ConfUpdate{}},
// 	{"GET", "/api/etcd/key/list", etcd.KeyList{}},
// 	{"POST", "/api/etcd/key/put", etcd.KeyPut{}},
// 	{"POST", "/api/etcd/key/delete", etcd.KeyDelete{}},
// }
// output:
// @Tags {手动添加}
// @Summary {手动添加}
// @Description {手动添加}
// @accept application/json {etcd.ConfVerify.Request}
// @Produce application/json {etcd.ConfVerify.Response}
// @Param data body object true "请求值"
// @Success      200  {object}  object
// @Faliure 400 {object} object
// @Router /路由 [方法]
func NewSwagInfoList(apis []string) []*SwagInfo {
	ans := make([]*SwagInfo, len(apis))
	for i, api := range apis {
		ans[i] = &SwagInfo{}
		parts := strings.Split(api[1:len(api)-3], ",")
		ans[i].Method = strings.Trim(strings.TrimSpace(parts[0]), `"`)
		ans[i].Url = strings.Trim(strings.TrimSpace(parts[1]), `"`)

		end := strings.Index(parts[2], "{")
		if end == -1 {
			end = len(parts[2])
		}
		if reqCtx := strings.TrimSpace(parts[2][:end]); reqCtx != "" {
			ans[i].Request = reqCtx + ".Request"
			ans[i].Response = reqCtx + ".Response"
		}

		ans[i].SlotName = fmt.Sprintf("slotFn%d", i)
	}
	return ans
}

func (s *SwagInfo) String() string {
	tpl := fmt.Sprintf(`// @Tags %s
// @Summary %s
// @Description %s
// @Accept application/json
// @Produce application/json
// @Param data body object true "请求值"
// @Success      200  {object}  object
// @Faliure 400 {object} object
// @Router %s [%s]
func %s() {}

`, s.Tag, s.Summary, s.Description, s.Url, s.Method, s.SlotName)
	return tpl
}
