var Ably = require('ably');
const fs = require('fs');
const path = require('path');
const readline = require('readline');

const ABLY_KEY = "xVLyHw.ZCOcRw:AANYGugTk7v7tQpuT7-hzjotlo9dmB6nys_TSXZmNrU";
const receivedMessagesPath = '../generated' + path.sep + 'js_subscribed.json';

var ablyClient = new Ably.Realtime({ key: ABLY_KEY, logLevel: 'warn'});
var channel = ablyClient.channels.get('test');

var receivedMessages = []
var counter = 0;
channel.subscribe(message => {
    console.log('message received '+ ++counter)
    receivedMessages.push(message.data)
})
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question('Press enter to save received messages', _ => {
    console.log("total messages received " + receivedMessages.length)
    fs.writeFileSync(receivedMessagesPath, JSON.stringify(receivedMessages, null, 2));
    rl.close();
});

