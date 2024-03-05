package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseBody struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Origin string `json:"origin"`
	Data   any    `json:"data"`
}

func Success(c *gin.Context, data ...interface{}) {
	resp := &responseBody{
		Code:   SUCCESS.Code,
		Msg:    SUCCESS.Message,
		Origin: SUCCESS.Origin,
		Data:   data,
	}
	c.JSON(http.StatusOK, resp)
}

func Fail(c *gin.Context, err error) {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		e = SERVER_ERROR.WithOrigin(err)
	}

	var resp responseBody
	resp.Code = e.Code
	resp.Msg = e.Message
	resp.Origin = e.Origin

	c.JSON(int(e.Code/100), resp)
	c.Abort()
}
