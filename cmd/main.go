package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.Static("/", "ui/dist")
	//e.File("/", "ui/dist/index.html")
	e.Static("/lists", "ui/dist")
	//e.File("/lists", "ui/dist/index.html")
	e.Logger.Fatal(e.Start(":3000"))

}
