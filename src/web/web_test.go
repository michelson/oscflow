package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server = httptest.NewServer(Handlers()) //Creating new server with the user handlers

func TestHome(t *testing.T) {

	usersUrl := fmt.Sprintf("%s/patches", server.URL)

	request, err := http.NewRequest("GET", usersUrl, nil)

	t.Logf("req %v", usersUrl)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode > 290 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

}

func TestGetPatch(t *testing.T) {

	usersUrl := fmt.Sprintf("%s/patches/hello_world.pd", server.URL)

	request, err := http.NewRequest("GET", usersUrl, nil)

	t.Logf("req %v", usersUrl)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode > 290 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

}
