package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	if err := MarshalConfig("config.json"); err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println(config)

	done := make(chan bool)
	go watchConfigChanges(watcher, done)

	err = watcher.Add("config.json")
	if err != nil {
		fmt.Println("Error watching config file:", err)
		return
	}

	<-done
}

func watchConfigChanges(watcher *fsnotify.Watcher, done chan<- bool) {
	defer close(done)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Println("Event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write && event.Name == "config.json" {
				if err := MarshalConfig("config.json"); err != nil {
					fmt.Println("Error reloading config:", err)
					continue
				}
				fmt.Println("Modified file:", event.Name)
				fmt.Println(config)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}

func MarshalConfig(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}
	return nil
}
