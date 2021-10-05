package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
)

func One(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Printf("head type is %T", header)

	for k, v := range header {
		w.Header().Set(k, strings.Join(v, ","))
	}
	w.Header().Set("test", "test response header")

	fmt.Fprintf(w, "get is success")

}

func Two(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	for _, value := range env {
		tmp := strings.Split(value, "=")
		w.Header().Set(tmp[0], tmp[1])
	}
	fmt.Fprintf(w, "get is success: %s", env)
}

func Third(w http.ResponseWriter, r *http.Request) {
	result := r.Header.Get("X-FORWARDED-FOR")
	if result != "" {
		fmt.Printf("ip is : %s", result)
	} else {
		fmt.Printf("ip is: %s", r.RemoteAddr)
	}

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Healthz (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "check is success")
}

func main() {
	http.HandleFunc("/one", One)
	http.HandleFunc("/two", Two)
	http.HandleFunc("/third", Third)
	http.HandleFunc("/healthz", Healthz)

	fmt.Println("web service start is success!!!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}