package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	json := []byte(`
	{
    	"token": "testtest",
    	"challenge": "testchallenge",
    	"type": "url_verification"
	}`)
	byteReader := bytes.NewReader(json)
	r := httptest.NewRequest("POST", "/", byteReader)
	Handle(w, r)
	rw := w.Result()
	if rw.StatusCode != http.StatusOK {
		t.Errorf("want %d, but result %d \n", http.StatusOK, rw.StatusCode)
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != "testchallenge" {
		t.Errorf("want %s, but acutaul %s \n", string(b), "testchallenge")
	}
}
