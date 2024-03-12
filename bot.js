#!/usr/bin/env node

const tmi = require("tmi.js");

const client = new tmi.Client({
  identity: {
    username: "sergio_sneto",
    password: "xytqxy3pffs0xwoaktldkvvnudobb7",
  },
  channels: ["ayellol"],
});

client.connect();

client.on("connected", onConnectedHandler);

client.on("message", (channel, tags, message, self) => {
  console.log(`message:${tags["display-name"]}:${message}`);
});

client.on("cheer", (channel, userstate, message) => {
  console.log(`bits:${userstate["display-name"]}:${userstate.bits}:${message}`);
});

function onConnectedHandler(addr, port) {
  console.log(`* Connected to ${addr}:${port}`);
}
