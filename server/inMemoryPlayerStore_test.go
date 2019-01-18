package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Joseph"
	store := NewInMemoryPlayerStore()

	server := &PlayerServer{store}

	request := newPostWinRequest(player)

	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)

	response := httptest.NewRecorder()
	request = newGetScoreRequest(player)

	server.ServeHTTP(response, request)

	assert.Equal(t, "3", response.Body.String())
}
