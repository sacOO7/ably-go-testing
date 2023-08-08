var Ably = require('ably');
const fs = require('fs');
const path = require('path');

const ABLY_KEY = "xVLyHw.ZCOcRw:AANYGugTk7v7tQpuT7-hzjotlo9dmB6nys_TSXZmNrU";
const publishedMessagesPath = '../generated' + path.sep + 'js_published.json';

var ablyClient = new Ably.Rest({ key: ABLY_KEY, logLevel: 'warn'});
var channel = ablyClient.channels.get('test');

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

const data = fs.readFileSync('../generated' + path.sep + 'messages.json', 'utf8');
const messagesToBePublished = JSON.parse(data)

publishMessages(messagesToBePublished).then(messagesPublished => {
    console.log("total messages sent "+ messagesPublished.length)
    fs.writeFileSync(publishedMessagesPath, JSON.stringify(messagesPublished, null, 2));
});

async function publishMessages(messagesToBePublished) {  
    var messagesPublished = [];
    var resolve = () => {};
    const promise = new Promise(res => resolve = res);
    for (let index = 0; index < messagesToBePublished.length; index++) {
        const message = messagesToBePublished[index];
        await sleep(10);
        channel.publish('greeting', message, function(err) {
            if(err) {
                console.log('publish failed with error ' + err);
            } else {
                messagesPublished.push(message);
                console.log('publish succeeded');
                if(index == messagesToBePublished.length - 1) {
                    resolve(messagesPublished);
                }
            }
        });
    }
    return promise;
}
  
