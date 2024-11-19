package services

import (
	"crudecho/config"
	model "crudecho/model"

	"gorm.io/gorm"
)

type ManagerService struct {
	db *gorm.DB
}

func NewManagerService(config *config.Config) *ManagerService {
	return &ManagerService{db: config.PostgresDB}
}

func (s *ManagerService) CreateManager(manager model.Manager) error {
	return s.db.Create(&manager).Error
}

func (s *ManagerService) GetManagerByID(id uint) (model.Manager, error) {
	var manager model.Manager
	err := s.db.First(&manager, id).Error
	return manager, err
}

func (s *ManagerService) GetAllManagers() ([]model.Manager, error) {
	var managers []model.Manager
	err := s.db.Find(&managers).Error
	return managers, err
}

func (s *ManagerService) UpdateManager(id uint, update model.Manager) error {
	var manager model.Manager
	err := s.db.First(&manager, id).Error
	if err != nil {
		return err
	}
	return s.db.Model(&manager).Updates(update).Error
}

func (s *ManagerService) DeleteManager(id uint) error {
	var manager model.Manager
	err := s.db.First(&manager, id).Error
	if err != nil {
		return err
	}
	return s.db.Delete(&manager).Error
}
