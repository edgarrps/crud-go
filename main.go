package main

import "github.com/edgarrps/crud-go/configs"

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
}
