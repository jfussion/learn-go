package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, return url of the fastest", func(t *testing.T) {
		slowServer := delayedServer(20 * time.Millisecond)
		fastServer := delayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if want != got {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if server doesn't respond within 10 seconds", func(t *testing.T) {
		server := delayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't got one")
		}
	})
}

func delayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
