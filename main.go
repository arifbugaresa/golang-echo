package main

import (
	"golang-echo/config"
	"golang-echo/route"
)

func main() {

	config.InitDB()

	e := route.New()
	e.Start(":8080")

}
