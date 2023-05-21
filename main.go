package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Car struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Transmission string `json:"transmission"`
	Price        string `json:"price"`
}

func main() {
	e := echo.New()

	cars := []Car{
		{"Ford", "Fiesta", "Manual", "165000000"},
		{"Ford", "Fiesta", "Manual", "175000000"},
		{"Ford", "Fiesta", "Automatic", "18000000"},
		{"Ford", "Fiesta", "Manual", "155000000"},
		{"VW", "Polo", "Manual", "170000000"},
		{"VW", "Beetle", "Manual", "265000000"},
		{"VW", "Polo", "Automatic", "165000000"},
	}

	e.GET("/cars", func(c echo.Context) error {
		// Get query parameters
		brand := c.QueryParam("brand")
		model := c.QueryParam("model")
		transmission := c.QueryParam("transmission")

		// Filter cars based on query parameters
		filteredCars := make([]Car, 0)
		for _, car := range cars {
			if (brand == "" || strings.EqualFold(car.Brand, brand)) &&
				(model == "" || strings.EqualFold(car.Model, model)) &&
				(transmission == "" || strings.EqualFold(car.Transmission, transmission)) {
				filteredCars = append(filteredCars, car)
			}
		}

		filteredCars1 := make([]Car, 0)
		filteredCars1 = append(filteredCars1, filteredCars...)
		// filteredCars1 := filteredCars
		for i := 1; i < len(filteredCars); i++ {
			if filteredCars[i].Brand == filteredCars1[i-1].Brand {
				filteredCars[i].Brand = ""
			}
			if filteredCars[i].Model == filteredCars1[i-1].Model {
				filteredCars[i].Model = ""
			}
			if filteredCars[i].Transmission == filteredCars1[i-1].Transmission {
				filteredCars[i].Transmission = ""
			}
		}

		return c.JSON(http.StatusOK, filteredCars)
	})

	e.Start(":8000")
}
