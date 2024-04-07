package parser

import (
	"CarCatalog/model"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

func ParseBody(c *gin.Context, out interface{}) error {
	bodyInBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Read body error: " + err.Error()})
		return err
	}

	err = c.Request.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Close body error: " + err.Error()})
		return err
	}

	err = json.Unmarshal(bodyInBytes, out)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
		return err
	}

	return nil
}
