package utilities

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func ReadAndLogBody(c *gin.Context) ([]byte, error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Received JSON: %s\n", string(body))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func BindJSON(c *gin.Context, body []byte, obj interface{}) error {
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	fmt.Printf("Parsed Data: %+v\n", obj)
	return nil
}
