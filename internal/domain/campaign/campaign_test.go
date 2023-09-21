package campaign_test

import (
	"marketplace/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"test1@email.com", "test2@email.com", "test3@email.com"}

	campaign,_ := campaign.NewCampaign("test", "testando", contacts)

	assert.NotEmpty(campaign.ID)
	assert.Equal(campaign.Name, "test")
	assert.Equal(campaign.Content, "testando")
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotEmpty(campaign.CreatedAt)
}

func Test_ValidateNameCampaign(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"test1@email.com", "test2@email.com", "test3@email.com"}

	_, err := campaign.NewCampaign("", "testando", contacts)

	assert.Equal("name is required", err.Error())	
}
