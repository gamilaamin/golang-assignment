package usecase

import (
	"golang-assignment/entity"
	"golang-assignment/repository"
)

type PetUsecase struct {
	PetRepo repository.PetsRepository
}

func NewPetUsecase() *PetUsecase {
	return &PetUsecase{
		PetRepo: *repository.NewPetsRepository(),
	}
}

func (petUsecase *PetUsecase) Create(pet entity.Pet) (*entity.Pet, error) {
	return petUsecase.PetRepo.Create(pet)
}

func (petUsecase *PetUsecase) Update(pet entity.Pet) (*entity.Pet, error) {
	return petUsecase.PetRepo.Update(pet)
}

func (petUsecase *PetUsecase) Delete(id uint) error {
	return petUsecase.PetRepo.Delete(id)
}

func (petUsecase *PetUsecase) Get(id uint) (entity.Pet, error) {
	return petUsecase.PetRepo.Get(id)
}
func (petUsecase *PetUsecase) GetAll() ([]entity.Pet, error) {
	return petUsecase.PetRepo.GetAll()
}
