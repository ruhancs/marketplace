package db

import campaign "marketplace/internal/campaign/domain"

type Repository struct {
	campaigns []campaign.Campaign
}

func (repo *Repository) Save(campaign *campaign.Campaign) error {
	repo.campaigns = append(repo.campaigns, *campaign)
	return nil
}

func (repo *Repository) GetAll() ([]campaign.Campaign,error) {
	
	return repo.campaigns,nil
}