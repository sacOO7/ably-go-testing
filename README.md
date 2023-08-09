## Ably-go-testing
Goal of this repo is to test missing messages while subscribing to channels

## Requirements
Node.js, Golang

## Getting things into action
- Clone the repo. (No need to setup ABLY_KEY)
- Contains two folders

1. `js-pubsub` - contains scripts for generating dummy message data, publishing to `test` channel and subscribing to the same.
```js
cd js-pubsub
npm install
npn run generate // Generates 1000 messages in a file under `generated/messages.json` at root.
npm run publish // Publishes ~100 msg/sec from `generated/messages.json` and saves them at `generated/js_published.json`
npm run subscribe // starts listening to messages and saves them in `generated/js_subscribed.json` at root.
```

2. `go-pubsub` - scripts responsible for listening and sending messages to `test` channel.

```go
cd go-pubsub
go install
go run subscriber.go // starts listening to messages and saves them in `generated/go_subscribed.json` at root.
go run publisher-rest.go //  Publishes 10-15 msg/sec from `generated/messages.json` and saves them at `generated/go_published_rest.json`
go run publisher-realtime.go //  Publishes 45 msg/sec from `generated/messages.json` and saves them at `generated/go_published_realtime.json`
```
