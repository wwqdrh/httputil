package httputil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDoReq(t *testing.T) {
	type NameT struct {
		Name string `form:"name" validate:"required"`
	}
	engine := gin.Default()
	engine.GET("/test", func(ctx *gin.Context) {
		handler := new(Handler)
		var req NameT
		if err := handler.DoReq(ctx, Query, &req); err != nil {
			handler.DoRes(ctx, ParamInvalid, gin.H{
				"err": err.Error(),
			})
		} else {
			handler.DoRes(ctx, ServerOK, nil)
		}
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test?name=123", nil)
	engine.ServeHTTP(w, r)
	fmt.Println(w.Body.String())

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/test", nil)
	engine.ServeHTTP(w, r)
	fmt.Println(w.Body.String())
}
