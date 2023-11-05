package handlers

import (
	"golang-assignment/adapter/presenter"
	"golang-assignment/usecase"
)

type PetHandler struct {
	PetUsecase usecase.PetUsecase
}

func NewPetHandler() *PetHandler {
	return &PetHandler{
		PetUsecase: *usecase.NewPetUsecase(),
	}
}

func (petHandler *PetHandler) Create(petDTO presenter.PetDTO) (*presenter.PetDTO, error) {
	entity := presenter.ToEntity(petDTO)
	pet, err := petHandler.PetUsecase.Create(*entity)
	if err != nil {
		return nil, err
	}
	petDTO = *presenter.FromEntity(*pet)
	return &petDTO, nil
}

func (petHandler *PetHandler) Update(petDTO presenter.PetDTO) (*presenter.PetDTO, error) {
	entity := presenter.ToEntity(petDTO)
	pet, err := petHandler.PetUsecase.Update(*entity)
	if err != nil {
		return nil, err
	}
	updatedPet, err := petHandler.PetUsecase.Get(pet.ID)
	if err != nil {
		return nil, err
	}
	petDTO = *presenter.FromEntity(updatedPet)
	return &petDTO, nil
}

func (petHandler *PetHandler) Delete(id uint) error {
	return petHandler.PetUsecase.Delete(id)
}

func (petHandler *PetHandler) Get(id uint) (petDTO presenter.PetDTO, err error) {

	pet, err := petHandler.PetUsecase.Get(id)
	if err != nil {
		return petDTO, err
	}
	petDTO = *presenter.FromEntity(pet)
	return petDTO, nil
}

func (petHandler *PetHandler) GetAll() ([]presenter.PetDTO, error) {
	pets, err := petHandler.PetUsecase.GetAll()
	if err != nil {
		return nil, err
	}
	return presenter.FromEntityList(pets), nil
}
