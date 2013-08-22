package api

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, Response{"success": true, "message": "Hello!"})
}

func RegisterHandlers() {
	http.HandleFunc("/api", testHandler)
}
