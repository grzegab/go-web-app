package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ListenForRequest(port string) {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello 1: "+req.Host)
	})
	http.ListenAndServe(":"+port, serverMux)
}

func makeJson(req string) []byte {
	out, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		log.Println("Starting logger failed:", err)
	}

	return out
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}
