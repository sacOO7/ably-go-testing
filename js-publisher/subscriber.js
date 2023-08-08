var Ably = require('ably');
const fs = require('fs');
const path = require('path');
const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

const ABLY_KEY = "xVLyHw.ZCOcRw:AANYGugTk7v7tQpuT7-hzjotlo9dmB6nys_TSXZmNrU";
const receivedMessagesPath = '../generated' + path.sep + 'js_subscribed.json';

var ablyClient = new Ably.Realtime({ key: ABLY_KEY, logLevel: 'warn'});
var channel = ablyClient.channels.get('test');

var receivedMessages = []
var counter = 0;
channel.subscribe(message => {
    console.log('message received '+ ++counter)
    receivedMessages.push(message.data)
}, attachErr => {
    if (attachErr) {
        console.error('\nerror attaching to test channel')
    } else {
        console.log('\nattached to test channel')
    }
 })

 rl.question('\nStart sending messages, press enter to save received messages', _ => {
    console.log("total messages received " + receivedMessages.length)
    fs.writeFileSync(receivedMessagesPath, JSON.stringify(receivedMessages, null, 2));
    rl.close();
});
