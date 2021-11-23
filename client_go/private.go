package main

import (
	"log"
	"main/util"

	"golang.org/x/net/websocket"
)

var doorState = util.DoorState{}
var config *Configuration

func readMsg(ws *websocket.Conn) string {
	var msg = make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		config.printError("Could not read message")
		log.Fatalln("Failed reading message")
	}
	return string(msg[:n])
}

func processMessage(msg string) {
	switch msg {
	case "open":
		doorState.SetDoorOpen(true)
		config.runOnDoorOpen()
	case "closed":
		doorState.SetDoorOpen(false)
		config.runOnDoorClose()
	default:
		log.Printf("[Warning] Unsupported message: %s", msg)
	}
}

func main() {
	config = loadConfiguration()
	origin := "http://localhost/"
	ws, err := websocket.Dial(config.Esp8266Url, "", origin)

	if err != nil {
		config.printError("Connection failed")
		log.Fatalln("Could not connect to " + config.Esp8266Url)
	}

	msg := readMsg(ws)
	config.runOnStart(doorState.IsDoorOpen())
	processMessage(msg)

	go func() {
		for {
			msg := readMsg(ws)
			log.Printf("Received a message: %s", msg)
			processMessage(msg)
		}
	}()

	config.waitUntilClose()
}
