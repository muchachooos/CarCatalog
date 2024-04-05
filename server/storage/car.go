package storage

import "CarCatalog/model"

func (s *Storage) AddCars(cars model.AddCarsReq) error {
	for i := range cars.RegNums {
		_, err := s.DB.Exec("INSERT INTO cars (regNum) VALUES ($1)", cars.RegNums[i])
		if err != nil {
			return err
		}
	}
	return nil
}
