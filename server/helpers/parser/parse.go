package parser

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
)

func ParseBody(c *gin.Context, out interface{}) error {
	bodyInBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyInBytes, out)
	if err != nil {
		return err
	}

	return nil
}
