package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {

	server := httptest.NewServer(Handlers()) //Creating new server with the user handlers

	//req, _ := http.NewRequest("GET", "http://example.com/patches", nil)
	//w := httptest.NewRecorder()

	usersUrl := fmt.Sprintf("%s/patches", server.URL)

	request, err := http.NewRequest("GET", usersUrl, nil)

	//t.Logf("w %v", w)
	//t.Logf("req %v", req)

	t.Logf("req %v", usersUrl)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode > 290 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

}
