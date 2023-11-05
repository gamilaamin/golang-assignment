package entity

type Pet struct {
	GMODEL
	Name  string `gorm:"not null"`
	Age   int    `gorm:"not null"`
	Breed string `gorm:"not null"`
}

func (p *Pet) TableName() string {
	return "pets"
}
