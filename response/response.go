package response

import (
	"net/http"

	"github.com/honmaple/forest"
)

type PageInfo struct {
	Page     int   `json:"page"  query:"page"`
	Limit    int   `json:"limit" query:"limit"`
	Total    int64 `json:"total"`
	NotLimit bool  `json:"-"`
}

func (s *PageInfo) GetLimit() (int, int) {
	if s.Page < 1 {
		s.Page = 1
	}
	if s.Limit < 1 {
		s.Limit = 10
	}
	offset := (s.Page - 1) * s.Limit
	if offset < 0 {
		offset = 0
	}
	return offset, s.Limit
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message"`
}

type ListResponse struct {
	PageInfo
	List interface{} `json:"list,omitempty"`
}

func New(code int, data ...interface{}) *Response {
	resp := &Response{
		Code: code,
	}

	if len(data) == 0 {
		resp.Message = http.StatusText(code)
	} else if len(data) == 1 {
		resp.Message = data[0]
	} else {
		resp.Message = data[0]
		resp.Data = data[1]
	}
	return resp
}

func Render(c forest.Context, code int, data ...interface{}) error {
	return c.JSON(code, New(code, data...))
}

func OK(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusOK, data...)
}

func BadRequest(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusBadRequest, data...)
}

func UnAuthorized(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusUnauthorized, data...)
}

func Forbidden(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusForbidden, data...)
}

func NotFound(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusNotFound, data...)
}

func ServerError(c forest.Context, data ...interface{}) error {
	return Render(c, http.StatusInternalServerError, data...)
}
