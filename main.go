package main

import (
	"context"
	"fmt"
	"log"
	"moody/api"
	"moody/communication"
	"moody/models"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	retries = 5
)

func main() {
	// Explicit logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.Println("Starting up")

	// Set up a safe exit mechanism
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Load conf file
	conf := mustInitConf()

	// Db connection
	attempt := 0
	for attempt < retries {
		dbFileName := fmt.Sprintf("%s/%s", conf.ConfPath, conf.DbName)
		db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
		if err == nil {
			models.DB = db
			break
		}
		attempt += 1
		time.Sleep(15 * time.Second)
	}

	if attempt == retries {
		log.Fatal("Couldn't connect to the database")
	}

	if err := communication.StartCommInterface(conf.ProtocolConf); err != nil {
		log.Println("an error occurred while starting the communication interface")
		log.Fatal(err)
	}

	if err := communication.CommConnect(); err != nil {
		log.Println("an error occurred while connecting through the communication interface")
		log.Fatal(err)
	}

	// TODO set static folder path
	server := api.HttpListenAndServer(conf.ServerPort, "")
	go func() { log.Fatal(server.ListenAndServe()) }()

	<-quit
	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatal(err)
	}

	communication.CommClose()
	log.Println("Gateway - Shutting Down")
}
