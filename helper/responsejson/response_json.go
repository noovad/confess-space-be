package responsejson

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success mengirim response sukses dengan message default atau custom
func Success(ctx *gin.Context, data interface{}, message ...string) {
	code := http.StatusOK
	status := "OK"
	msg := "Success"

	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	ctx.JSON(code, Response{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	})
}

func Created(ctx *gin.Context, data interface{}, message ...string) {
	code := http.StatusCreated
	status := "Created"
	msg := "Resource created successfully"

	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}

	ctx.JSON(code, Response{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	})
}

func Unauthorized(ctx *gin.Context, message ...string) {
	msg := "Unauthorized"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	ctx.JSON(http.StatusUnauthorized, Response{
		Code:    http.StatusUnauthorized,
		Status:  "Unauthorized",
		Message: msg,
		Data:    nil,
	})
}

func InternalServerError(ctx *gin.Context, err error, message ...string) {
	msg := "Internal Server Error"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: msg,
		Data:    err.Error(),
	})
}

func BadRequest(ctx *gin.Context, err error, message ...string) {
	msg := "Bad Request"
	fmt.Println("Bad Request:", err.Error())
	fmt.Println("Bad Request message slice:", message)
	if len(message) > 0 && message[0] != "" {
		fmt.Println("Bad Request custom message:", message[0])
		msg = message[0]
	}
	fmt.Println("Bad Request final message:", msg)
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: msg,
		Data:    err.Error(),
	})
}

func Forbidden(ctx *gin.Context, message ...string) {
	msg := "Forbidden"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	ctx.JSON(http.StatusForbidden, Response{
		Code:    http.StatusForbidden,
		Status:  "Forbidden",
		Message: msg,
		Data:    nil,
	})
}

func NotFound(ctx *gin.Context, message ...string) {
	msg := "Not Found"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	ctx.JSON(http.StatusNotFound, Response{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: msg,
		Data:    nil,
	})
}

func Conflict(ctx *gin.Context, message ...string) {
	msg := "Conflict"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	ctx.JSON(http.StatusConflict, Response{
		Code:    http.StatusConflict,
		Status:  "Conflict",
		Message: msg,
		Data:    nil,
	})
}
