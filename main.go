package main

import "go_chat/router"

func main() {
	e := router.Router()
	e.Run("0.0.0.0:8080")
}
