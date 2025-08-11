package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)
func main() {
	s := gocron.NewScheduler(time.UTC)

	// Runs every 5 seconds
	s.Every(5).Seconds().Do(func ()  {
		fmt.Println("Task running at", time.Now())
	})

	//Start schedular
	s.

}