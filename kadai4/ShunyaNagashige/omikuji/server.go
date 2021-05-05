package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func pickResult()string{
	_, month, day := time.Now().Date()

	if month == 1 && day >= 1 && day <= 3 {
		return "大吉"
	}
	
	return "小吉"
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//_, month, day := time.Now().Date()
	omikuji := make(map[string]string, 1)

	// if month == 1 && day >= 1 && day <= 3 {
	// 	omikuji["result"] = "大吉"
	// } else {
	// 	omikuji["result"] = "吉"
	// }
	omikuji["result"]=pickResult()

	enc := json.NewEncoder(w)
	if err := enc.Encode(omikuji); err != nil {
		http.Error(w, "failed to Encode", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/omikuji", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
