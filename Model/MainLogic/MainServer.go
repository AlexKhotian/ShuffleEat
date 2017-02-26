package main

import "ShuffleEat/Model/ServerHandler"

func main() {
	server := ServerHandler.ServerRoutineFactory()
	server.RunServer()
}
