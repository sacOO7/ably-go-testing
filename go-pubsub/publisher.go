package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ably/ably-go/ably"
)

func main() {

	var (
		ABLY_KEY    = "xVLyHw.RnzM8g:iV7CnSzWDKG7HRFFOieUCYyDG6z1QYXMiqlk-RF0A6U"
		channelName = "test"
	)

	client, err := ably.NewREST(
		ably.WithKey(ABLY_KEY),
		ably.WithLogLevel(ably.LogWarning),
	)
	if err != nil {
		log.Fatal(err)
	}

	channel := client.Channels.Get(channelName)
	// channel.Attach(context.Background())
	// time.Sleep(time.Second)

	data, err := os.ReadFile(filepath.Join("../generated", "messages.json"))
	if err != nil {
		log.Fatal(err)
	}
	var messages []map[string]interface{}
	err = json.Unmarshal(data, &messages)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	sentMessages := make([]map[string]interface{}, len(messages))
	for index := 0; index < len(messages); index++ {
		message := messages[index]
		err := channel.Publish(context.Background(), "greeting", message)
		if err != nil {
			fmt.Printf("Error publishing message %v\n", err)
		} else {
			fmt.Printf("published message %v\n", index)
			sentMessages[index] = message
		}
		if index == len(messages)-1 {
			cancel()
		}
	}

	<-ctx.Done()

	fmt.Printf("total messages published %v", len(sentMessages))
	serializedString, err := json.MarshalIndent(sentMessages, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join("../generated", "go_published.json")
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	_, er := f.WriteString(string(serializedString))
	if er != nil {
		log.Fatal(er)
	}
	defer f.Close()

}
