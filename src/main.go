package main

import (
	"main/api"
)

func main() {
	r := api.InitRouters()
	r.Run()
}
