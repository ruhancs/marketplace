package campaign_test

import (
	campaign "marketplace/internal/campaign/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"test1@email.com", "test2@email.com", "test3@email.com"}

	newCampaign,_ := campaign.NewCampaign("test", "testando", contacts)

	assert.NotEmpty(newCampaign.ID)
	assert.Equal(newCampaign.Name, "test")
	assert.Equal(newCampaign.Content, "testando")
	assert.Equal(len(newCampaign.Contacts), len(contacts))
	assert.NotEmpty(newCampaign.CreatedAt)
	assert.Equal(newCampaign.Status, campaign.Pending)
}

func Test_ValidateNameCampaign(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"test1@email.com", "test2@email.com", "test3@email.com"}

	_, err := campaign.NewCampaign("", "testando", contacts)

	assert.NotNil(err)	
}
