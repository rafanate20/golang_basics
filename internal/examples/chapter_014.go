package examples

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatingAnEchoServer() {

	fmt.Println("Chapter 014 - Instance of echo Web server:")
	e := echo.New()

	//route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GoLang Web Server (Echo v4)")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
