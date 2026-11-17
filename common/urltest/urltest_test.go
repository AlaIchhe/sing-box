package urltest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	N "github.com/sagernet/sing/common/network"
)

func TestURLTestExpectedStatus(t *testing.T) {
	noContentServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNoContent)
	}))
	defer noContentServer.Close()
	okServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer okServer.Close()
	ctx := context.Background()
	if _, err := URLTest(ctx, noContentServer.URL, 0, N.SystemDialer); err != nil {
		t.Fatal("any status code should pass the test when expected_status is not set: ", err)
	}
	if _, err := URLTest(ctx, noContentServer.URL, http.StatusNoContent, N.SystemDialer); err != nil {
		t.Fatal("status code 204 should pass the test: ", err)
	}
	if _, err := URLTest(ctx, noContentServer.URL, http.StatusOK, N.SystemDialer); err == nil {
		t.Fatal("status code 204 should fail the test when 200 is expected")
	}
	if _, err := URLTest(ctx, okServer.URL, http.StatusOK, N.SystemDialer); err != nil {
		t.Fatal("status code 200 should pass the test: ", err)
	}
	if _, err := URLTest(ctx, okServer.URL, http.StatusNoContent, N.SystemDialer); err == nil {
		t.Fatal("status code 200 should fail the test when 204 is expected")
	}
}
