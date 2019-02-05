package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    server := echo.New()

    server.Debug = true

    // Handler
    server.GET("/", func(context echo.Context) error {

        address := context.RealIP()
        fmt.Println(address)
        return context.String(http.StatusOK, "Hello, " + address + "!")
    })

    // Start server
    server.Logger.Fatal(server.Start(":5000"))
}