package repository

import (
	"golang-assignment/entity"
	"golang-assignment/global"
)

type PetsRepository struct {
}

func NewPetsRepository() *PetsRepository {
	return &PetsRepository{}
}

func (petsRepository *PetsRepository) Create(pet entity.Pet) (*entity.Pet, error) {
	err := global.DB.Debug().Create(&pet).Error
	return &pet, err
}

func (petsRepository *PetsRepository) Update(pet entity.Pet) (*entity.Pet, error) {
	err := global.DB.Debug().Model(&entity.Pet{}).Where("id = ?", pet.ID).Updates(pet).Error
	return &pet, err
}

func (petsRepository *PetsRepository) Delete(id uint) error {
	return global.DB.Debug().Where("id = ?", id).Delete(&entity.Pet{}).Error
}

func (petsRepository *PetsRepository) Get(id uint) (pet entity.Pet, err error) {
	err = global.DB.Debug().Model(&entity.Pet{}).Where("id = ?", id).First(&pet).Error
	return pet, err
}

func (petsRepository *PetsRepository) GetAll() (pets []entity.Pet, err error) {
	err = global.DB.Model(&entity.Pet{}).Debug().Order("id DESC").Find(&pets).Error
	return pets, err
}
