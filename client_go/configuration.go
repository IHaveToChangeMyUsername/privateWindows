package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/util"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Configuration struct {
	Esp8266Url            string
	DisplayErrorCommand   string
	DoorOpenedCommand     string
	DoorClosedCommand     string
	OnStartOpenCommand    string
	OnStartClosedCommand  string
	WaitUntilCloseCommand string
}

var win = util.Windows{}

func isDefined(s string) bool {
	return len(s) != 0
}

func openFile(fileName string, defaultContent []byte) *os.File {
	file, err := os.Open(fileName)
	if err != nil { // file does probably not exists
		file.Close()
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatal("Failed creating file " + fileName)
		}
		file.Write(defaultContent)
		file.Close()
		file, err = os.Open(fileName)
		if err != nil {
			log.Fatal("Failed to open file " + fileName)
		}
		return file
	}
	return file
}

func loadConfiguration() *Configuration {
	file := openFile("config.json", windows().toBytes())

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}

func runCommand(command string) {
	if !isDefined(command) {
		return
	}

	if runtime.GOOS == "windows" {
		win.Run(command)
	} else {
		cmd := exec.Command(command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		log.Println("RUN!")
		err := cmd.Run()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func (configuration Configuration) printError(msg string) {
	command := configuration.DisplayErrorCommand
	if !isDefined(command) {
		fmt.Print(msg)
	} else {
		go runCommand(strings.ReplaceAll(command, "${ERROR_MESSAGE}", msg))
	}
}

func (configuration Configuration) runOnDoorOpen() {
	go runCommand(configuration.DoorOpenedCommand)
}

func (configuration Configuration) runOnDoorClose() {
	go runCommand(configuration.DoorClosedCommand)
}

func (configuration Configuration) runOnStart(isOpen bool) {
	if win.CreateSession() != nil {
		log.Fatalln("Could not create powershell session")
	}
	if isOpen {
		go runCommand(configuration.OnStartOpenCommand)
	} else {
		go runCommand(configuration.OnStartClosedCommand)
	}
}

func (configuration Configuration) waitUntilClose() {
	if isDefined(configuration.WaitUntilCloseCommand) {
		runCommand(configuration.WaitUntilCloseCommand)
	} else {
		time.Sleep(time.Duration(1<<63 - 1))
	}
	win.DestroySession()
}

func (configuration Configuration) toBytes() []byte {
	data, err := json.MarshalIndent(configuration, "", "  ")
	if err != nil {
		data = []byte("Json Encode error: " + err.Error())
	}
	return data
}
