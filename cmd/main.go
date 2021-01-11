package main

import (
	"infrastructure/cmd/app"
	"infrastructure/cmd/init"
)

func main() {
	init.Init()
	app.Start()
}
