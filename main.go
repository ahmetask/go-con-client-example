package main

import (
	"fmt"
	gocon "github.com/ahmetask/go-con"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	configServerPort := "8081"
	applicationPort := "8082"

	type customConfig struct {
		A string
		B bool
	}

	var MyConfig = customConfig{
		A: "go-con-client",
		B: true,
	}

	go func() {
		for {
			fmt.Println("go-con-client")
			fmt.Println(fmt.Sprintf("My Config %v", MyConfig))
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		// Pass Your Config variable address to Listener
		err := gocon.ListenConfigChange(configServerPort, &MyConfig, nil)
		panic(err)
	}()

	e.GET("/config", func(c echo.Context) error {
		return c.JSON(http.StatusOK, MyConfig)
	})

	e.Logger.Fatal(e.Start(":" + applicationPort))
}
