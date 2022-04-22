package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ResponseJson(ctx *gin.Context, status int, data interface{}) {
	res := map[string]any{
		"code":    status,
		"message": messageError(status),
		"result":  data,
	}
	ctx.JSON(status, res)
}

func messageError(code int) string {
	if code == http.StatusOK {
		return "Success"
	} else {
		return "Failed"
	}
}

func GetRawBody(c *gin.Context) (map[string]interface{}, interface{}, string) {
	var data map[string]interface{}
	var thisError interface{}
	raw, _ := c.GetRawData()
	err := json.Unmarshal(raw, &data)
	if err != nil {
		if os.Getenv("APP_ENV") != "development" {
			thisError = err.Error()
		} else {
			thisError = "Body is not valid, please check your request"
		}
		return nil, thisError, err.Error()
	}
	return data, nil, ""
}

func GetErrorString(err error) interface{} {
	var thisError interface{}
	if err != nil {
		if os.Getenv("APP_ENV") != "development" {
			thisError = err.Error()
		} else {
			thisError = "Internal server error"
		}
		return thisError
	}
	return nil
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
