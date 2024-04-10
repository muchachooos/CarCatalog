package handler

import (
	"CarCatalog/helpers/parser"
	"CarCatalog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) AddCarsHandler(c *gin.Context) {
	var carsReq model.AddCarsReq

	err := parser.ParseBody(c, &carsReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Parse body error: " + err.Error()})
		return
	}

	for i := range carsReq.RegNums {
		if carsReq.RegNums[i] == "" {
			c.JSON(http.StatusBadRequest, model.Err{Error: "Empty string in the body"})
			return
		}
	}

	err = s.Storage.AddCars(carsReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) ModifyCarsHandler(c *gin.Context) {
	regNum, ok := c.GetQuery("reg_num")
	if regNum == "" || !ok {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Number is missing"})
		return
	}

	var car model.AddParamInCar

	err := parser.ParseBody(c, &car)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Parse body error: " + err.Error()})
		return
	}

	car.RegNum = regNum

	err = s.Storage.ModifyCars(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) GetCarsHandler(c *gin.Context) {
	var filter model.CarsFilter

	regNum := c.Query("reg_num")
	if regNum != "" {
		filter.RegNum = &regNum
	}

	mark := c.Query("mark")
	if mark != "" {
		filter.Mark = &mark
	}

	modelCar := c.Query("model")
	if modelCar != "" {
		filter.Model = &modelCar
	}

	yearInStr := c.Query("year")
	if yearInStr != "" {
		year, _ := strconv.Atoi(yearInStr)
		filter.Year = &year
	}

	owner := c.Query("owner")
	if owner != "" {
		filter.Owner = &owner
	}

	res, err := s.Storage.GetCars(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteCarHandler(c *gin.Context) {
	regNum, ok := c.GetQuery("regNum")
	if regNum == "" || !ok {
		c.JSON(http.StatusBadRequest, model.Err{Error: "Number is missing"})
		return
	}

	err := s.Storage.DeleteCar(regNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
