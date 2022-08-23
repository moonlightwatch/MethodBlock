package methodblock_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moonlightwatch/methodblock"
)

func TestBlock(t *testing.T) {
	ctx := context.Background()
	nextCall := false
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) { nextCall = true })
	cfg := methodblock.CreateConfig()
	cfg.Message = "Method mot Allowed"
	cfg.Methods = []string{"DELETE", "HEAD"}
	handler, err := methodblock.New(ctx, next, cfg, "pluginName")

	if err != nil {
		t.Fatalf("New MethodBlock error: %+v\n", err)
	}

	{
		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodDelete, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(recorder, req)

		if recorder.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("Status Code want %d but got %d.\n", http.StatusMethodNotAllowed, recorder.Result().StatusCode)
		}
		if nextCall {
			t.Fatal("next was called\n")
		}
	}
	{
		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(recorder, req)

		if recorder.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("Status Code want %d but got %d.\n", http.StatusMethodNotAllowed, recorder.Result().StatusCode)
		}
		if nextCall {
			t.Fatal("next was called\n")
		}
	}
}

func TestPass(t *testing.T) {
	ctx := context.Background()
	nextCall := false
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) { nextCall = true })

	cfg := methodblock.CreateConfig()
	cfg.Message = "Method mot Allowed"
	cfg.Methods = []string{"DELETE", "HEAD"}
	handler, err := methodblock.New(ctx, next, cfg, "pluginName")

	if err != nil {
		t.Fatalf("New MethodBlock error: %+v\n", err)
	}

	{
		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(recorder, req)

		if recorder.Result().StatusCode != http.StatusOK {
			t.Fatalf("Status Code want 200 but got %d.\n", recorder.Result().StatusCode)
		}
		if !nextCall {
			t.Fatal("next was not called\n")
		}
	}
	{
		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(recorder, req)

		if recorder.Result().StatusCode != http.StatusOK {
			t.Fatalf("Status Code want 200 but got %d.\n", recorder.Result().StatusCode)
		}
		if !nextCall {
			t.Fatal("next was not called\n")
		}
	}
}
