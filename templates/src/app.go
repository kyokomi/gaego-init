package app

import "net/http"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World {{ .AppName }}"))
	})
}
