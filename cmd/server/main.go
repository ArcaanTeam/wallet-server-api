package main

import (
	"wallet-api/internal/wire"
)

func main() {
	wire.NewApp().Run()
}
