package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign,error) {
	contacts := make([]Contact, len(emails))

	if name == "" {
		return nil, errors.New("name is required")
	}

	for index,email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID: xid.New().String(),
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts,
	}, nil
}