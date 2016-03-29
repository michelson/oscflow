var WebSocket = require('ws');
var ws = new WebSocket('ws://localhost:8080');

console.log("loading app");

ws.on('open', function open() {
  var array = new Float32Array(5);

  for (var i = 0; i < array.length; ++i) {
    array[i] = i / 2;
  }

  ws.send(array, { binary: true, mask: true });
});


ws.on('open', function open() {
  ws.send('something from client');
});

ws.on('message', function(data, flags) {
  console.log(data);
  document.getElementById("stats").innerHTML = data;
  // flags.binary will be set if a binary data is received.
  // flags.masked will be set if the data was masked.
});