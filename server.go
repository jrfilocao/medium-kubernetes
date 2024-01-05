// A simple server. /healthz alternates its response every 60 seconds, alternating between HTTP 200 and HTTP 500 status codes

package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "time"
)

func main() {
    started := time.Now()
    healthy := true

    http.HandleFunc("/started", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        data := (time.Since(started)).String()
        w.Write([]byte(data))
    })
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        duration := time.Since(started)
        if duration.Seconds() > 60 {
            healthy = !healthy   // Toggle the health status
            started = time.Now() // Update the start time to measure next 60 seconds
        }
        // Alternating between healthy and unhealthy states
        if !healthy {
            w.WriteHeader(500)
            w.Write([]byte(fmt.Sprintf("ERROR: %v", duration.Seconds())))
        } else {
            w.WriteHeader(200)
            w.Write([]byte(fmt.Sprintf("OK: %v", duration.Seconds())))
        }
    })
    http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
        loc, err := url.QueryUnescape(r.URL.Query().Get("loc"))
        if err != nil {
            http.Error(w, fmt.Sprintf("invalid redirect: %q", r.URL.Query().Get("loc")), http.StatusBadRequest)
            return
        }
        http.Redirect(w, r, loc, http.StatusFound)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
