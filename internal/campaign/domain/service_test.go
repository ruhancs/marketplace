package campaign

import (
	"errors"
	"marketplace/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) GetAll() ([]Campaign, error) {
	//args := r.Called()
	return nil, nil
}

func (r *repositoryMock) GetOne(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	CREATEDTO = contract.CreateCampaignDTO{
		Name:      "test",
		Content:   "testing",
		Emails:    []string{"test1@email.com"},
		CreatedBy: "admin@email.com",
	}
	) 

func Test_CreateCampaignService(t *testing.T) {
	assert := assert.New(t)

	createCampaignDTO := contract.CreateCampaignDTO{
		Name:      "test",
		Content:   "testing",
		Emails:    []string{"test1@email.com"},
		CreatedBy: "admin@email.com",
	}

	repository := new(repositoryMock)
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
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

	id, err := service.Create(CREATEDTO)

	repository.AssertExpectations(t)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_GetCampaignByID(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	campaign, _ := NewCampaign("test", "testing", []string{"test@email.com"}, CREATEDTO.CreatedBy)
	repository.On("GetOne", mock.Anything).Return(campaign, nil)

	service := Service{
		Repository: repository,
	}

	output, err := service.GetOne("iquwerbj236")

	assert.Nil(err)
	assert.Equal(campaign.Name, output.Name)
	assert.Equal(campaign.Content, output.Content)
	assert.Equal(campaign.CreatedBy, output.CreatedBy)
	assert.Equal(campaign.Status, output.Status)
}

func Test_GetCampaignByID_NotFound(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	//campaign,_ := NewCampaign("test","testing",[]string{"test@email.com"})
	repository.On("GetOne", mock.Anything).Return(nil, errors.New("camapign not found"))

	service := Service{
		Repository: repository,
	}

	_, err := service.GetOne("iquwerbj236")

	assert.NotNil(err)
}

func Test_DeleteCampaign(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	campaign, _ := NewCampaign("test", "testing", []string{"test@email.com"}, CREATEDTO.CreatedBy)
	repository.On("GetOne", mock.Anything).Return(campaign, nil)
	repository.On("Delete", mock.Anything).Return(nil)

	service := Service{
		Repository: repository,
	}

	err := service.Delete("quwhegyu143")

	assert.Nil(err)
}

func Test_CampaignDoesNotBeDeleted(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	campaign := &Campaign{ID: "1", Status: Started, Name: "C1"}
	repository.On("GetOne", mock.Anything).Return(campaign, nil)

	service := Service{
		Repository: repository,
	}

	err := service.Delete("iquwerbj236")

	assert.NotNil(err)
	assert.Equal("campaign does not be canceled", err.Error())
}

func Test_Start_RecorNotFoundError(t *testing.T) {
	assert := assert.New(t)

	repository := new(repositoryMock)
	//campaign,_ := NewCampaign("test","testing",[]string{"test@email.com"})
	repository.On("GetOne", mock.Anything).Return(nil, errors.New("camapign not found"))

	service := Service{
		Repository: repository,
	}

	err := service.Start("iquwerbj236")

	assert.NotNil(err)
}

func Test_CampaignRepository(t *testing.T) {

	createCampaignDTO := contract.CreateCampaignDTO{
		Name:    "test",
		Content: "testing",
		Emails:  []string{"test1@email.com"},
		CreatedBy: "admin@email.com",
	}

	repository := new(repositoryMock)
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
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
