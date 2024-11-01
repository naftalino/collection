package main

import "app/web"

func main() {
	go StartBot()
	web.StartServer()
}
