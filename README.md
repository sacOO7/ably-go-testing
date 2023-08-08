## Ably-go-testing
Goal of this repo is to test missing messages while subscribing to channels

## Requirements
Node.js, Golang

## Getting things into action
- Clone the repo. (No need to setup ABLY_KEY)
- Contains two folders
1. `go-subscriber` - script responsible for listening to all messages published on `test` channel.

```go
cd go-subscriber
go mod install
go run main.go // starts listening to messages and saves them in `generated/go_subscribed.json` at root.
```

2. `js-publisher` - contains scripts for generating dummy message data, publishing to `test` channel and subscribing to the same.
```js
cd js-publisher
npm run install
npn run generate // Generates 1000 messages in a file under `generated/messages.json` at root.
npm run publish // Publishes messages from `generated/messages.json` and saves them at ``generated/js_published.json``
npm run subscribe // starts listening to messages and saves them in `generated/js_subscribed.json` at root.
```
