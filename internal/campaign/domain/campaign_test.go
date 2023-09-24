package campaign_test

import (
	campaign "marketplace/internal/campaign/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name      = "test"
	content   = "testando"
	contacts  = []string{"test1@email.com", "test2@email.com", "test3@email.com"}
	createdBy = "admin@email.com"
)

func Test_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)

	newCampaign, _ := campaign.NewCampaign(name, content, contacts, createdBy)

	assert.NotEmpty(newCampaign.ID)
	assert.Equal(newCampaign.Name, "test")
	assert.Equal(newCampaign.Content, "testando")
	assert.Equal(len(newCampaign.Contacts), len(contacts))
	assert.Equal(newCampaign.CreatedBy, createdBy)
	assert.NotEmpty(newCampaign.CreatedAt)
	assert.Equal(newCampaign.Status, campaign.Pending)
}

func Test_ValidateNameCampaign(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"test1@email.com", "test2@email.com", "test3@email.com"}

	_, err := campaign.NewCampaign("", "testando", contacts, createdBy)

	assert.NotNil(err)
}
