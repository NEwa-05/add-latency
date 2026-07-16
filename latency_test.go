package latency

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLatency(t *testing.T) {
	startTime := time.Now().Second()
	cfg := CreateConfig()
	cfg.AddedLatency = 10

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "latency")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	endTime := time.Now().Second()

	assertLatency(t, cfg.AddedLatency, startTime, endTime)
}

func assertLatency(t *testing.T, latency int, sTime int, eTime int) {
	t.Helper()
	if (eTime - sTime) < latency {
		t.Errorf("proper latency not apply")
	}
}
