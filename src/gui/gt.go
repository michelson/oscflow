package gui

import (
	"fmt"
	commands "github.com/miketheprogrammer/go-thrust/lib/commands"
	"github.com/miketheprogrammer/go-thrust/thrust"
)

func Example() {
	//flag.Parse()
	thrust.InitLogger()
	thrust.Start()

	size := commands.SizeHW{Width: 320, Height: 240}

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: fmt.Sprintf("http://127.0.0.1:%d", 8181),
		//HasFrame: true,
		Size: size,
	})

	thrustWindow.Show()
	thrustWindow.Focus()
	//thrustWindow.Fullscreen(true)
	thrustWindow.OpenDevtools()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	thrust.LockThread()
}

/*

/*
import (
  "fmt"
  "net/http"

  "github.com/miketheprogrammer/go-thrust/thrust"
  "github.com/miketheprogrammer/go-thrust/tutorials/provisioner"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func Example() {
	http.HandleFunc("/", handler)
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	mysession := thrust.NewSession(false, false, "cache")

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://localhost:8080/",
		Session: mysession,
	})
	thrustWindow.Show()
	//thrustWindow.Maximize()
	thrustWindow.Focus()

	// See, we dont use thrust.LockThread() because we now have something holding the process open
	http.ListenAndServe(":8080", nil)
}

var htmlIndex string = `
<html>
  <body>
    <h2> Welcome to Go-Thrust <h3>
    <img height="50px" width="120px" src="http://i.imgur.com/DwFKI0J.png"/>
    <script>
      window.onload = function() {
        setTimeout(function() {
          var webview = document.createElement("webview");
          document.body.appendChild(webview);
          webview.src = "http://www.google.com";
          //webview.classList.add("active");
        }, 0);
      }
    </script>
  </body>
</html>
`
*/
