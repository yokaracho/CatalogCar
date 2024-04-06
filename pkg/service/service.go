package service

import (
	"NameService/pkg/model"
	"NameService/pkg/repository"
	"context"
)

type ImageImplementation interface {
	InsertCars(ctx context.Context, regNums []string) (int64, error)
	UpdateInfo(ctx context.Context, car *model.CarModel) error
	GetCars(ctx context.Context, filters model.CarFilter) ([]*model.CarModel, error)
	DeleteCarByID(ctx context.Context, id int) (int64, error)
}

type Service struct {
	repository repository.Implementation
}

type Implementation interface {
	ImageImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
