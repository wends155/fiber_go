package main

import (
	"github.com/wends155/fiber_go/src/config"
)

func main() {
	app := config.New()
	app.Listen(":3000")
}
