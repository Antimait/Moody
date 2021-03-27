package models

import (
	"errors"
)

type StateValue uint64

type Service struct {
	Id   uint64 `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name;unique" validate:"nonzero" json:"name"`
}

type ServiceRequest struct {
	Service uint64 `json:"serviceId" validate:"nonzero"`
}

type ServiceByName struct {
	ServiceName string `json:"name" validate:"nonzero"`
}

type ServicesResponse struct {
	Services     []string `json:"services"`
	ServiceCount int      `json:"count"`
}

type ServiceTableResponse struct {
	Table        map[string]string `json:"services"`
	ServiceCount int               `json:"count"`
}

func (n *Service) TableName() string {
	return "service"
}

func GetAllServices() (services []*Service, totalRows int64, err error) {
	services = []*Service{}
	serviceOrm := DB.Model(&Service{})
	serviceOrm.Count(&totalRows)

	if err = serviceOrm.Find(&services).Error; err != nil {
		err := errors.New("not found")
		return nil, -1, err
	}

	return services, totalRows, nil
}

func GetService(id uint64) (*Service, error) {
	var service Service

	if err := DB.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func AddService(service *Service) error {
	if err := DB.Save(service).Error; err != nil {
		return InsertFailedError
	}
	return nil
}

func DeleteService(s *Service) error {
	service := &Service{}
	if err := DB.First(service, s.Id).Error; err != nil {
		return NotFound
	}

	if err := DB.Delete(service).Error; err != nil {
		return DeleteFailedError
	}
	return nil
}
