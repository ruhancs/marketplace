package campaign

import "marketplace/internal/contract"

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
		return "",err
	}

	return campaign.ID,nil
}