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
		ABLY_KEY    = "xVLyHw.ZCOcRw:AANYGugTk7v7tQpuT7-hzjotlo9dmB6nys_TSXZmNrU"
		channelName = "test"
	)

	client, err := ably.NewRealtime(
		ably.WithKey(ABLY_KEY),
		ably.WithLogLevel(ably.LogVerbose),
	)
	if err != nil {
		log.Fatal(err)
	}

	channel := client.Channels.Get(channelName)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	messages := []string{}
	u, err := channel.SubscribeAll(ctx, func(m *ably.Message) {
		messages = append(messages, m.Data.(string))
	})
	if err != nil {
		log.Fatal(err)
	}

	defer u()

	<-ctx.Done()

	serializedString, err := json.MarshalIndent(messages, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join("../generated", "subscribed.json")
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
