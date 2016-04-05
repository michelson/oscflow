// Let us open a web socket
var ws = new WebSocket("ws://localhost:8081/echo");
ws.onopen = function()
{
  // Web Socket is connected, send data using send()
  ws.send("Message to send");
  console.log("Message is sent...");
  ws.send('something from client');

};

ws.onmessage = function (evt) 
{ 
  var received_msg = evt.data;
  console.log("Message is received...");
  console.log(evt.data)
  document.getElementById("stats").innerHTML = received_msg;

};

ws.onclose = function()
{ 
  // websocket is closed.
  console.log("Connection is closed..."); 
};
