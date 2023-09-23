package db

import (
	campaign "marketplace/internal/campaign/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) Save(campaign *campaign.Campaign) error {
	tx := repo.DB.Create(campaign)//utilizado para criar e atualizar
	return tx.Error
}

func (repo *Repository) Update(campaign *campaign.Campaign) error {
	tx := repo.DB.Save(campaign)//utilizado para criar e atualizar
	return tx.Error
}

func (repo *Repository) GetAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := repo.DB.Find(&campaigns)
	return campaigns, tx.Error
}

func (repo *Repository) GetOne(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := repo.DB.Preload("Contacts").First(&campaign, "id=?", id)
	return &campaign, tx.Error
}

func (repo *Repository) Delete(campaign *campaign.Campaign) error {
	
	tx := repo.DB.Select("Contacts").Delete(campaign)
	return tx.Error
}