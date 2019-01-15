package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(firstURL, secondURL string) (winner string, error error) {
	return ConfigurableRacer(firstURL, secondURL, tenSecondTimeout)
}

func ConfigurableRacer(firstURL, secondURL string, timeout time.Duration) (string, error) {

	select {
	case <-ping(firstURL):
		return firstURL, nil
	case <-ping(secondURL):
		return secondURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", firstURL, secondURL)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
