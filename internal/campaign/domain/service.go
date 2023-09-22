package campaign

import (
	"fmt"
	"marketplace/internal/contract"
	internalerror "marketplace/internal/internal_error"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(createCampaignDTO contract.CreateCampaignDTO) (string,error) {
	campaign,err := NewCampaign(createCampaignDTO.Name,createCampaignDTO.Content,createCampaignDTO.Emails)
	if err != nil {
		return "",err
	}
	
	err = s.Repository.Save(campaign)
	if err != nil {
		fmt.Println(err)
		return "",internalerror.RepositoryErr
	}

	return campaign.ID,nil
}

func (s *Service) GetAll() []Campaign {
	campaigns,_ := s.Repository.GetAll()
	
	return campaigns
}