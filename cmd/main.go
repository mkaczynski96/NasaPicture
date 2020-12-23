package main

import (
	"gogoapps-go/pkg/api"
	client "gogoapps-go/pkg/client"
	"gogoapps-go/pkg/config"
	"log"
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	var cfg config.Config
	config.ReadFile(&cfg)
	config.ReadEnv(&cfg)
	log.Printf("Loaded config: %+v", cfg)

	newClient, err := client.NewClient(cfg)
	if err != nil {
		return err
	}

	newApi, err := api.New(cfg, newClient)
	if err != nil {
		return err
	}

	if err = newApi.SetServer(); err != nil {
		return err
	}
	return nil
}
