var app = require('app');  // Module to control application life.
var BrowserWindow = require('browser-window');  // Module to create native browser window.


require('electron-debug')({
    showDevTools: true
});

/*
var WebSocketServer = require('ws').Server
  , wss = new WebSocketServer({ port: 8081 });
*/

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
var mainWindow = null;

var path = require('path');
var spawn = require('child_process').spawn;

//spawn go app
//var child = spawn(path.join(__dirname, '..', 'main'), ['game.config', '--debug']);
var child = spawn(path.join(__dirname, 'oscflow') );

child.on('close', function(code) {
    console.log("close go server");
        //return result;
});


    child.stdout.on('data', function(data) {
        //result += data.toString();
        console.log(data.toString());
        //ws.send(data.toString());
        //console.log("sadd");
        //ws.send('something from osc');
    });

/*
//websocket
wss.on('connection', function connection(ws) {
    ws.on('message', function incoming(message) {
      console.log('ws server rcv: %s', message);
    });

    child.stdout.on('data', function(data) {
        //result += data.toString();
        //console.log(data.toString());
        ws.send(data.toString());
        //console.log("sadd");
        //ws.send('something from osc');
    });

  //ws.send('something from server');
});
*/

// attach events, etc.

// Quit when all windows are closed.
app.on('window-all-closed', function() {
    // On OS X it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform != 'darwin') {
        app.quit();
    }
});

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
app.on('ready', function() {

    // Create the browser window.
    mainWindow = new BrowserWindow({width: 320, height: 600});

    // and load the index.html of the app.
    mainWindow.loadURL('file://' + __dirname + '/index.html');

    // Emitted when the window is closed.
    mainWindow.on('closed', function() {
        // Dereference the window object, usually you would store windows
        // in an array if your app supports multi windows, this is the time
        // when you should delete the corresponding element.
        mainWindow = null;
    });
});