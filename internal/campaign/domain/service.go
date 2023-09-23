package campaign

import (
	"errors"
	"fmt"
	"marketplace/internal/contract"
	internalerror "marketplace/internal/internal_error"
)

type ServiceInterface interface {
	Create(createCampaignDTO contract.CreateCampaignDTO) (string, error)
	GetAll() []Campaign
	GetOne(id string) (*contract.OutGetCampaignByID, error)
	Cancel(id string) error
	Delete(id string) error
}

type Service struct {
	Repository Repository
}

func (s *Service) Create(createCampaignDTO contract.CreateCampaignDTO) (string, error) {
	campaign, err := NewCampaign(createCampaignDTO.Name, createCampaignDTO.Content, createCampaignDTO.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		fmt.Println(err)
		return "", internalerror.RepositoryErr
	}

	return campaign.ID, nil
}

func (s *Service) GetAll() []Campaign {
	campaigns, _ := s.Repository.GetAll()

	return campaigns
}

func (s *Service) GetOne(id string) (*contract.OutGetCampaignByID, error) {
	campaign, err := s.Repository.GetOne(id)
	if err != nil {
		fmt.Println(err)
		return &contract.OutGetCampaignByID{}, err
	}

	output := contract.OutGetCampaignByID{
		Name:             campaign.Name,
		Content:          campaign.Content,
		Status:           campaign.Status,
		QuantityContacts: len(campaign.Contacts),
	}

	return &output, nil
}

func (s *Service) Cancel(id string) error {
	campaign, err := s.Repository.GetOne(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if campaign.Status != Pending {
		return errors.New("campaign does not be canceled")
	}

	campaign.Cancel()
	err = s.Repository.Update(campaign)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Service) Delete(id string) error {
	campaign, err := s.Repository.GetOne(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if campaign.Status != Pending {
		return errors.New("campaign does not be canceled")
	}

	err = s.Repository.Delete(campaign)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
