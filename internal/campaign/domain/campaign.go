package campaign

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `valid:"email" gorm:"size:150"`
	CampaignID string `gorm:"size:50"`
}

const (
	Pending  string = "pending"
	Started  string = "started"
	Done     string = "done"
	Canceled string = "canceled"
)

type Campaign struct {
	ID        string    `valid:"required" gorm:"size:50"`
	Name      string    `valid:"required,stringlength(4|15)" gorm:"size:100"`
	CreatedBy string    `gorm:"size:100"`
	CreatedAt time.Time `valid:"required"`
	Content   string    `valid:"required" gorm:"size:1000"`
	Contacts  []Contact `valid:"required"`
	Status    string    `valid:"required" gorm:"size:20"`
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedBy: createdBy,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := campaign.Validate()
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Start() {
	c.Status = Started
}

func (c *Campaign) Done() {
	c.Status = Done
}

func (c *Campaign) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
