package repository

import (
	models "crudecho/model"

	"gorm.io/gorm"
)

type SiteManagerRepository struct {
	DB *gorm.DB
}

// Create a new site manager
func (repo *SiteManagerRepository) CreateSiteManager(manager *models.Manager) error {
	return repo.DB.Create(manager).Error
}

// Get all site managers
func (repo *SiteManagerRepository) GetAllSiteManagers() ([]models.Manager, error) {
	var managers []models.Manager
	err := repo.DB.Find(&managers).Error
	return managers, err
}

// Get site manager by ID
func (repo *SiteManagerRepository) GetSiteManagerByID(id uint) (models.Manager, error) {
	var manager models.Manager
	err := repo.DB.First(&manager, id).Error
	return manager, err
}

// Update site manager by ID
func (repo *SiteManagerRepository) UpdateSiteManager(id uint, update models.Manager) error {
	var manager models.Manager
	err := repo.DB.First(&manager, id).Error
	if err != nil {
		return err
	}
	return repo.DB.Model(&manager).Updates(update).Error
}

// Delete site manager by ID
func (repo *SiteManagerRepository) DeleteSiteManager(id uint) error {
	var manager models.Manager
	err := repo.DB.First(&manager, id).Error
	if err != nil {
		return err
	}
	return repo.DB.Delete(&manager).Error
}
