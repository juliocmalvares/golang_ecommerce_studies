package utils

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func BuildDefaultResponse(httpCode int, data interface{}, message string, err error) gin.H {
	v := reflect.ValueOf(data)
	size := 0
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		size = v.Len()
	case reflect.Struct:
		size = 1
	default:
		size = 0
	}
	return gin.H{
		"status":  httpCode,
		"message": message,
		"error":   err,
		"data":    data,
		"size":    size,
	}
}
