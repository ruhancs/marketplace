package campaign


type Repository interface {
	Save(campaign *Campaign) error
	GetAll() ([]Campaign,error)
	GetOne(id string) (*Campaign,error)
	Delete(campaign *Campaign) error
	Update(campaign *Campaign) error
}