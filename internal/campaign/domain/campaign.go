package campaign

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `valid:"email"`
}

type Campaign struct {
	ID        string    `valid:"required"`
	Name      string    `valid:"required,alpha,stringlength(4|15)"`
	CreatedAt time.Time `valid:"required"`
	Content   string    `valid:"required"`
	Contacts  []Contact `valid:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}
	err := campaign.Validate()
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (c *Campaign) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
