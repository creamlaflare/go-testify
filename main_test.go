package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmptyf(t, responseRecorder.Body, "Response must not be empty")

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Wrong status code")
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Wrong status code")

	require.NotEmptyf(t, responseRecorder.Body, "Response must not be empty")

	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "wrong response")
}

func TestMainHandlerWhenLargeCount(t *testing.T) {
	totalCount := len(cafeList["moscow"])
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "wrong status code")

	require.NotEmptyf(t, responseRecorder.Body, "Response must not be empty")

	body := responseRecorder.Body.String()

	assert.Equal(t, totalCount, len(strings.Split(body, ",")), "wrong expected cafe count")
}
