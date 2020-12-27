package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

func setupNotifier() func() {
	note := gosxnotifier.NewNotification("Regroup in 10mins")
	note.Title = "It's break time 🤧"
	note.Subtitle = "You've been working for 1hr"
	note.Sound = gosxnotifier.Glass
	note.Group = "com.unique.yourapp.identifier"
	note.Sender = "com.apple.Safari"
	note.ContentImage = "gopher.png"
	return func() {
		if err := note.Push(); err != nil {
			log.Println("Uh oh!")
		}
	}
}

func main() {
	notifyUser := setupNotifier()
	ticker := time.NewTicker(1 * time.Hour)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				notifyUser()
			}
		}
	}()

	<-quit
	log.Println("shutting down ... ")

}
