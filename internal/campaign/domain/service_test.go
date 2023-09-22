package campaign

import (
	"marketplace/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func(r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func(r *repositoryMock) GetAll() ([]Campaign,error ){
	//args := r.Called()
	return  nil,nil
}

func Test_CreateCampaignService(t *testing.T) {
	assert := assert.New(t)

	createCampaignDTO := contract.CreateCampaignDTO{
		Name: "test",
		Content: "testing",
		Emails: []string{"test1@email.com"},
	}
	
	repository := new(repositoryMock)
	repository.On("Save",mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != createCampaignDTO.Name ||
			campaign.Content != createCampaignDTO.Content ||
			len(campaign.Contacts) != len(createCampaignDTO.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{
		Repository: repository,
	}

	id,err := service.Create(createCampaignDTO)

	repository.AssertExpectations(t)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_CampaignRepository(t *testing.T) {

	createCampaignDTO := contract.CreateCampaignDTO{
		Name: "test",
		Content: "testing",
		Emails: []string{"test1@email.com"},
	}
	
	repository := new(repositoryMock)
	repository.On("Save",mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != createCampaignDTO.Name ||
			campaign.Content != createCampaignDTO.Content ||
			len(campaign.Contacts) != len(createCampaignDTO.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{
		Repository: repository,
	}

	service.Create(createCampaignDTO)

	repository.AssertExpectations(t)
}