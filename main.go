package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustTocken()
	fmt.Println(t)
	//token = flags.Get(token)

	//tgCLient = telegram.New(tocken)//8337556172:AAGIjLrGDmyElPTs4YC6sy520LE9zezIH6M

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)
}

func mustTocken() string {
	token := flag.String("t",
		"",
		"token for your bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is GOVNA KYSOK")
	}
	return *token
}
