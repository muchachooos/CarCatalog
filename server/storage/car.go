package storage

import (
	"CarCatalog/model"
	"database/sql"
	"errors"
	"strconv"
)

func (s *Storage) AddCars(cars model.AddCarsReq) error {
	for i := range cars.RegNums {
		_, err := s.DB.Exec("INSERT INTO cars (reg_num) VALUES ($1)", cars.RegNums[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) ModifyCars(car model.AddParamInCar) error {
	var checkCar string

	err := s.DB.Get(&checkCar, "SELECT reg_num FROM cars WHERE reg_num = $1", car.RegNum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrCarNotFound
		}
		return err
	}

	if car.Mark != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET mark = :mark WHERE reg_num = :reg_num", car)
		if err != nil {
			return err
		}
	}

	if car.Model != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET model = :model WHERE reg_num = :reg_num", car)
		if err != nil {
			return err
		}
	}

	if car.Year != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET year = :year WHERE reg_num = :reg_num", car)
		if err != nil {
			return err
		}
	}

	if car.Owner != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET owner = :owner WHERE reg_num = :reg_num", car)
		if err != nil {
			return err
		}
	}

	if car.NewNum != nil {
		_, err := s.DB.NamedExec("UPDATE cars SET reg_num = :newNum WHERE reg_num = :reg_num", car)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) DeleteCar(regNum string) error {
	res, err := s.DB.Exec("DELETE FROM cars WHERE reg_num = $1", regNum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrCarNotFound
		}
	}

	countOfModifiedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		return model.ErrCarNotFound
	}

	return nil
}

func (s *Storage) GetCars(filter model.CarsFilter) ([]model.Car, error) {
	query := "SELECT * FROM cars WHERE TRUE "

	if filter.Mark != nil {
		query += " AND mark = " + "'" + *filter.Mark + "'"
	}

	if filter.Model != nil {
		query += " AND model = " + "'" + *filter.Model + "'"
	}

	if filter.Year != nil {
		query += " AND year = " + "'" + strconv.Itoa(*filter.Year) + "'"
	}

	if filter.Owner != nil {
		query += " AND owner = " + "'" + *filter.Owner + "'"
	}

	var res []model.Car

	err := s.DB.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
