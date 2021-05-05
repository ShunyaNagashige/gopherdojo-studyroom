package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//年始のテストもできるようにする
//pickresultのテストは不要？
//→必要かも．TestOmikujiHandler=実際にjsonで値を取得できるかどうかのテスト
//TestPickResult=ある日付に対応した結果を，正しく得られているかどうかのテスト
func TestOmikujiHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/omikuji", nil)
	omikujiHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	expected := make(map[string]string, 1)
	
	expected["result"]=pickResult()

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	output := make(map[string]string, 1)
	if err := json.Unmarshal(b, &output); err != nil {
		t.Fatal("cannot parse the data.")
	}

	if output["result"] != expected["result"] {
		t.Fatalf("unexpected response: want %#v, got %#v", expected, output)
	}
}
