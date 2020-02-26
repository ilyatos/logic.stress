package main

import (
	"fmt"
	"github.com/ilyatos/logic.stress/pkg/client"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

//TODO: pass as opt
const logicHost = "http://logic.hacktory.ai"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	u, err := url.Parse(logicHost)
	if err != nil {
		log.Fatalln(err)
	}

	c := client.NewClient(u, os.Getenv("LOGIC_AUTH_TOKEN"), nil)

	message, err := c.Index()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message.Message)
}
