package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	greeting := greetPerson("Nama Ta")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, greeting+" is alive")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func greetPerson(nama string) string {
	return fmt.Sprintf("Hello %s", nama)
}
