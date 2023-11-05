package router

import (
	"golang-assignment/adapter/controller"

	"github.com/labstack/echo/v4"
)

type PetRouter struct {
}

// InitAccountRouter initialization agent Route information
func (p *PetRouter) InitPetRouter(Router *echo.Echo) {
	PetRouter := Router.Group("/pets")
	var petController = controller.NewPetController()
	{
		PetRouter.POST("", petController.Create)
		PetRouter.GET("", petController.GetAll)
		PetRouter.GET("/:id", petController.Get)
		PetRouter.PUT("/:id", petController.Update)
		PetRouter.DELETE("/:id", petController.Delete)
	}
}
