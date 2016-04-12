package web

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
	p "patches"
)

var patches = p.Patches{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func PatchesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	patches, _ := p.GetPatches()

	fmt.Println(patches)

	if err := json.NewEncoder(w).Encode(patches); err != nil {
		panic(err)
	}
}

func PatchesShow(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")

	id := r.URL.Query().Get(":id")
	fmt.Println(id)

	patches, err := patches.GetPatch(id)

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(patches); err != nil {
		panic(err)
	}
}

func ServeAsset(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	http.ServeFile(w, r, r.URL.Path[1:])
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func Handlers() *pat.PatternServeMux {

	m := pat.New()
	//m.Get("/", http.FileServer(http.Dir("./")))
	m.Get("/", http.FileServer(http.Dir("./interface/build")))
	//m.Get("/build", http.HandlerFunc(ServeAsset))

	//(http.FileServer(http.Dir("./interface/build/")))

	m.Get("/hello/:name", http.HandlerFunc(HelloServer))
	m.Get("/patches", http.HandlerFunc(PatchesIndex))
	m.Get("/patches/:id", http.HandlerFunc(PatchesShow))

	//m.Get("/static", http.FileServer(http.Dir("./interface/build")))

	serveSingle("/javascripts/all.js", "./interface/build/javascripts/all.js")
	serveSingle("/stylesheets/all.css", "./interface/build/stylesheets/all.css")

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)

	return m

}

func Example() {

	//http.Handle("/static/", http.HandlerFunc(ServeAsset))
	Handlers()
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
