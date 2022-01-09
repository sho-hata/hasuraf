package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sho-hata/hasura-fzf/lib/command"
)

func main() {
	result, err := command.Run()
	if err != nil {
		log.Fatal(err, result)
	}
	fmt.Println(result)
	os.Exit(0)
}
