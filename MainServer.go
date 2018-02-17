package main

import (
	"ShuffleEat/Model/ServerHandler"
	"log"
)

func main() {
	server, err := ServerHandler.ServerRoutineFactory()
	if err != nil {
		log.Println(err)
		return
	}
	server.RunServer()
}
