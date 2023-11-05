package controller

import (
	"errors"
	"golang-assignment/adapter/handlers"
	"golang-assignment/adapter/presenter"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type PetController struct {
	petHandler handlers.PetHandler
}

func NewPetController() *PetController {
	return &PetController{
		petHandler: *handlers.NewPetHandler(),
	}
}

func (petController PetController) Create(cxt echo.Context) error {
	var pet presenter.PetDTO
	err := cxt.Bind(&pet)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request", "message": err.Error()})
	}
	createdPet, err := petController.petHandler.Create(pet)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, map[string]string{"error": "Create failed", "message": err.Error()})
	} else {
		return cxt.JSON(http.StatusCreated, createdPet)
	}
}

func (petController PetController) Update(cxt echo.Context) error {
	var pet presenter.PetDTO
	idStr := cxt.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID", "message": err.Error()})
	}
	err = cxt.Bind(&pet)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request", "message": err.Error()})
	}
	pet.ID = uint(id)

	updatedpet, err := petController.petHandler.Update(pet)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, map[string]string{"error": "Update failed", "message": err.Error()})
	}
	return cxt.JSON(http.StatusOK, updatedpet)
}

func (petController PetController) Delete(cxt echo.Context) error {
	idStr := cxt.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID", "message": err.Error()})
	}
	err = petController.petHandler.Delete(uint(id))
	if err != nil {
		return cxt.String(http.StatusInternalServerError, "Failed to delete pet: "+err.Error())
	}
	return cxt.JSON(http.StatusOK, "Pet deleted successfully")
}

func (petController PetController) Get(cxt echo.Context) error {
	idStr := cxt.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID", "message": err.Error()})
	}
	pet, err := petController.petHandler.Get(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cxt.JSON(http.StatusNotFound, map[string]string{"error": "Pet not found", "message": err.Error()})
		}
		return cxt.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error", "message": err.Error()})
	}
	return cxt.JSON(http.StatusOK, pet)
}

func (petController PetController) GetAll(cxt echo.Context) error {
	if pets, err := petController.petHandler.GetAll(); err != nil {
		return cxt.String(http.StatusInternalServerError, "Failed to Retrive pets List: "+err.Error())
	} else {
		return cxt.JSON(http.StatusOK, pets)
	}
}
