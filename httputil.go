package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type reqType uint8

const (
	JSON reqType = iota
	XML
	Form
	Query
	FormPost
	FormMultipart
	ProtoBuf
	MsgPack
	YAML
	Uri
	Header
	TOML
)

type Base interface {
	DoReq(ctx *gin.Context, typ_ reqType, req interface{}) error
	DoRes(ctx *gin.Context, code ResponseCode, ext map[string]interface{})
}

var DefaultHandler Base = new(Handler)

type Handler struct{}

// 序列化请求以及进行校验
func (b *Handler) DoReq(ctx *gin.Context, typ_ reqType, req interface{}) error {
	switch typ_ {
	case JSON:
		if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case XML:
		if err := ctx.ShouldBindWith(req, binding.XML); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case Form:
		if err := ctx.ShouldBindWith(req, binding.Form); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case Query:
		if err := ctx.ShouldBindWith(req, binding.Query); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case FormPost:
		if err := ctx.ShouldBindWith(req, binding.FormPost); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case FormMultipart:
		if err := ctx.ShouldBindWith(req, binding.FormMultipart); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case ProtoBuf:
		if err := ctx.ShouldBindWith(req, binding.ProtoBuf); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case MsgPack:
		if err := ctx.ShouldBindWith(req, binding.MsgPack); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case YAML:
		if err := ctx.ShouldBindWith(req, binding.YAML); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case Uri:
		if err := ctx.ShouldBindUri(req); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case Header:
		if err := ctx.ShouldBindWith(req, binding.Header); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	case TOML:
		if err := ctx.ShouldBindWith(req, binding.TOML); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	default:
		if err := ctx.ShouldBind(req); err != nil {
			return fmt.Errorf("参数绑定失败: %w", err)
		}
	}
	if err := ctx.BindQuery(req); err != nil {
		return fmt.Errorf("参数绑定失败: %w", err)
	}

	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("参数校验失败: %w", err)
	}
	return nil
}

func (b *Handler) DoRes(ctx *gin.Context, code ResponseCode, ext map[string]interface{}) {
	body := gin.H{
		"code": code,
		"msg":  code.String(),
	}
	for key, val := range ext {
		body[key] = val
	}

	ctx.JSON(http.StatusOK, body)
}
