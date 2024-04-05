package handler

import (
	"CarCatalog/helpers/parser"
	"CarCatalog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) AddCarsHandler(c *gin.Context) {
	var carsReq model.AddCarsReq

	err := parser.ParseBody(c, &carsReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Parse body error: " + err.Error()})
	}

	err = s.Storage.AddCars(carsReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
	}

	c.Status(http.StatusOK)
}
