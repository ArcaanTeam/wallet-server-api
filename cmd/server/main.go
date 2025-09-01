package main

import (
	"wallet-api/internal/wire"
)

// TODO: pass "ctx" all the way down from controller
// TODO: move "migration" to repositories
func main() {
	wire.NewApp().Run()
}
