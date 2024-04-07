package service

import (
	"CatalogCar/pkg/model"
	"context"
)

func (s *Service) InsertCars(ctx context.Context, regNums []string) (int64, error) {
	return s.repository.InsertCars(ctx, regNums)
}

func (s *Service) UpdateInfo(ctx context.Context, car *model.CarModel) error {
	return s.repository.UpdateInfo(ctx, car)
}

func (s *Service) UpdateOwner(ctx context.Context, owner *model.PeopleModel) error {
	return s.repository.UpdateOwner(ctx, owner)
}

func (s *Service) GetCars(ctx context.Context, filters model.CarFilter) ([]*model.CarModel, error) {
	return s.repository.GetCars(ctx, filters)
}

func (s *Service) DeleteCarByID(ctx context.Context, id int) (int64, error) {
	return s.repository.DeleteCarByID(ctx, id)
}
