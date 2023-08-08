package gosubscriber

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/ably/ably-go/ably"
)

func main() {

	var (
		ABLY_KEY    = ""
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

	u, err := channel.SubscribeAll(ctx, func(m *ably.Message) {
		log.Println("%v", m)
	})
	if err != nil {
		log.Fatal(err)
	}

	defer u()

	<-ctx.Done()
}
