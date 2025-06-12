package main

import (
	_ "embed"
	"log"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/toast.v1"
)

//go:embed titles.txt
var Titles string

//go:embed messages.txt
var Messages string

//go:embed morning_messages.txt
var MorningMessages string

//go:embed night_messages.txt
var GoodNightMessages string

// func main() {
//
// 	ticker := time.NewTicker(10 * time.Minute)
// 	defer ticker.Stop()
//
// 	for {
// 		now := time.Now()
// 		titles, messages := strings.Split(strings.TrimSpace(Title), "\n"), strings.Split(strings.TrimSpace(Messages), "\n")
// 		morning_messages, night_messages := strings.Split(strings.TrimSpace(MorningMessages), "\n"), strings.Split(strings.TrimSpace(GoodNightMessages), "\n")
//
// 		select {
// 		case <-ticker.C:
// 			title, message := randomMessage(titles), randomMessage(messages)
// 			sendNotification(title, message)
// 		default:
// 			// Check for 11 AM message
// 			if now.Hour() == 11 && now.Minute() == 0 {
// 				title, morningMessage := randomMessage(titles), randomMessage(morning_messages)
// 				sendNotification(title, morningMessage)
// 				continue
// 			}
//
// 			// Check for 11 PM message
// 			if now.Hour() == 23 && now.Minute() == 0 {
// 				title, nightMessage := randomMessage(titles), randomMessage(night_messages)
// 				sendNotification(title, nightMessage)
// 				continue
// 			}
//
// 		}
// 	}
// }

func main() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		now := time.Now()

		// Check for 11 AM
		if now.Hour() == 11 && now.Minute() == 0 {
			title, morningMessage := randomMessage(splitLines(Titles)), randomMessage(splitLines(MorningMessages))
			sendNotification(title, morningMessage)
			time.Sleep(61 * time.Second) // Prevent multiple triggers in the same minute
			continue
		}

		// Check for 11 PM
		if now.Hour() == 23 && now.Minute() == 0 {
			title, nightMessage := randomMessage(splitLines(Titles)), randomMessage(splitLines(GoodNightMessages))
			sendNotification(title, nightMessage)
			time.Sleep(61 * time.Second)
			continue
		}

		// Wait for the 10 minute ticker
		<-ticker.C
		title, message := randomMessage(splitLines(Titles)), randomMessage(splitLines(Messages))
		sendNotification(title, message)
	}
}

func randomMessage(messages []string) string {
	return messages[rand.Intn(len(messages))]
}

func splitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func sendNotification(title, message string) {
	notification := toast.Notification{
		AppID:   "Heartbeats",
		Title:   title,
		Message: message,
	}

	if err := notification.Push(); err != nil {
		log.Printf("Error sending notification: %v\n", err)
	}
}
