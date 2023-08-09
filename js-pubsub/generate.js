var fs = require('fs');
const path = require('path')
const directory = '../generated'
const filePath = directory + path.sep + 'messages.json'
const messageSize = 2048; // 2kb each
const noOfMessages = 1000;

const generateRandomString = (length) => {
    let result = '';
    const characters =
      'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
  };
  
if (!fs.existsSync(directory)){
    fs.mkdirSync(directory);
}

const messages = []

for (let index = 0; index < noOfMessages; index++) {
    messages.push({
        id : index,
        message : generateRandomString(messageSize)
    })
}

fs.writeFileSync(filePath, JSON.stringify(messages, null, 2));

