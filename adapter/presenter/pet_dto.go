package presenter

import "golang-assignment/entity"

type PetDTO struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Breed string `json:"breed" form:"breed"`
}

func FromEntity(pet entity.Pet) *PetDTO {
	return &PetDTO{
		ID:    pet.ID,
		Name:  pet.Name,
		Age:   pet.Age,
		Breed: pet.Breed,
	}
}

func ToEntity(pet PetDTO) *entity.Pet {
	return &entity.Pet{
		GMODEL: entity.GMODEL{
			ID: pet.ID,
		},
		Name:  pet.Name,
		Age:   pet.Age,
		Breed: pet.Breed,
	}
}

func FromEntityList(pets []entity.Pet) (petDTOs []PetDTO) {
	for _, pet := range pets {
		petDTO := FromEntity(pet)
		petDTOs = append(petDTOs, *petDTO)
	}
	return petDTOs
}
