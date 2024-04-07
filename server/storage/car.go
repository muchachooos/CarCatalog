package storage

import (
	"CarCatalog/model"
	"database/sql"
	"errors"
)

func (s *Storage) AddCars(cars model.AddCarsReq) error {
	for i := range cars.RegNums {
		_, err := s.DB.Exec("INSERT INTO cars (regNum) VALUES ($1)", cars.RegNums[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) ModifyCars(car model.Car) error {
	if car.Mark != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET mark = :mark WHERE regNum = :regNum", car)
		if err != nil {
			return err
		}
	}

	if car.Model != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET model = :model WHERE regNum = :regNum", car)
		if err != nil {
			return err
		}
	}

	if car.Year != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET year = :year WHERE regNum = :regNum", car)
		if err != nil {
			return err
		}
	}

	if car.Owner != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET owner = :owner WHERE regNum = :regNum", car)
		if err != nil {
			return err
		}
	}

	if car.NewNum != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET regNum = :newNum WHERE regNum = :regNum", car)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) DeleteCar(regNum string) error {
	var car string

	err := s.DB.Get(&car, "SELECT regNum FROM cars WHERE regNum = $1", regNum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrCarNotFound
		}
		return err
	}

	_, err = s.DB.Exec(`DELETE FROM cars WHERE regNum = $1`, regNum)
	if err != nil {
		return err
	}

	return nil
}
