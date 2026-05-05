package main

import (
	"fmt"

	"github.com/kitaclysm/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cfg.SetUser("kitaclysm")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg2, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Print(cfg2)
	return
}