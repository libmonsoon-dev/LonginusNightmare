package main

import (
	"log"

	"github.com/libmonsoon-dev/LonginusNightmare/client"
)

func main() {
	c, err := client.New()
	if err != nil {
		log.Fatal("inti client:", err)
	}

	if err = c.Run(); err != nil {
		log.Fatal("runtime error:", err)
	}
}
