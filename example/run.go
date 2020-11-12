package main

import (
	"errors"
	"log"

	"github.com/rfyiamcool/go-retry"
)

func main() {
	r := retry.New()
	var running = false
	err := r.Ensure(func() error {
		log.Println("enter")
		if !running {
			log.Println("111")
			running = true
			return retry.Retriable(errors.New("diy"))
		}

		log.Println("222")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
