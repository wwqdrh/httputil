package gen

import (
	"fmt"
	"testing"
)

func TestGenSwag(t *testing.T) {
	for _, item := range NewSwagInfoList([]string{
		`{"GET", "/api/etcd/conf/verify", []gin.HandlerFunc{etcd.ConfVerify}},`,
		`{"GET", "/api/etcd/conf/status", []gin.HandlerFunc{etcd.ConfStatus}},`,
		`{"POST", "/api/etcd/conf/update", []gin.HandlerFunc{etcd.ConfUpdate}},`,
		`{"GET", "/api/etcd/key/list", []gin.HandlerFunc{etcd.KeyList}},`,
		`{"POST", "/api/etcd/key/put", []gin.HandlerFunc{etcd.KeyPut}},`,
		`{"POST", "/api/etcd/key/delete", []gin.HandlerFunc{etcd.KeyDelete}},`,
	}) {
		fmt.Println(item.String())
	}
}
