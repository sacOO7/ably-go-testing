package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/ably/ably-go/ably"
)

func main() {

	var (
		ABLY_KEY    = "xVLyHw.RnzM8g:iV7CnSzWDKG7HRFFOieUCYyDG6z1QYXMiqlk-RF0A6U"
		channelName = "test"
	)

	client, err := ably.NewRealtime(
		ably.WithKey(ABLY_KEY),
		ably.WithLogLevel(ably.LogWarning),
	)
	if err != nil {
		log.Fatal(err)
	}

	channel := client.Channels.Get(channelName)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	messages := []string{}
	var counter = 1
	u, err := channel.SubscribeAll(ctx, func(m *ably.Message) {
		messages = append(messages, m.Data.(string))
		fmt.Printf("message received %v\n", counter)
		counter++
	})
	if err != nil {
		log.Fatal(err)
	}

	defer u()

	fmt.Println("Start sending messages, once sending compleye, terminate the process to save received messages")
	<-ctx.Done()

	serializedString, err := json.MarshalIndent(messages, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join("../generated", "go_subscribed.json")
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	_, er := f.WriteString(string(serializedString))
	if er != nil {
		log.Fatal(er)
	}
	defer f.Close()

	fmt.Println("Total messages received ", len(messages))
}
