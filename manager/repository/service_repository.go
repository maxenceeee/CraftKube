package repository

import "xamence.eu/craftkube/internal"

type ServiceRepository struct {
	services map[string]internal.Service
}

func NewServiceRepository() *ServiceRepository {
	return &ServiceRepository{
		services: make(map[string]internal.Service),
	}
}
