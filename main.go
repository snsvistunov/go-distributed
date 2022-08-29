package main

import (
	"fmt"
	"log"

	"github.com/snsvistunov/go-distributed/config"
	"github.com/snsvistunov/go-distributed/pkg/file"
	"github.com/snsvistunov/go-distributed/pkg/mergesort"
)

const configPath string = "config/envs"

func main() {
	var data file.Content

	err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	data.NewContent(config.Cfg.SourceFile)
	mergesort.ParralelSort(data)
	data.WriteContent(config.Cfg.OutputFile)
	fmt.Printf("Sorted data %v\n", data)
}
