package main

import (
	commands "app/bot"
	"app/web"
)

func main() {
	go commands.StartBot()
	web.StartServer()
}
