package initialize

import (
	"golang-assignment/router"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func InitServer(address string) {
	var petRouter router.PetRouter

	e := echo.New()
	// add routes
	petRouter.InitPetRouter(e)

	s := http.Server{
		Addr:        address,
		Handler:     e,
		ReadTimeout: 20 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
