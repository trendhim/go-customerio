package customerio_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/customerio/go-customerio/v3"
)

// A failed collection call carries the status and Customer.io's answer: the first error's detail when the body has
// one, otherwise the body as it came.
func TestCollectionErrorKeepsTheAnswer(t *testing.T) {
	for name, tc := range map[string]struct {
		body string
		want string
	}{
		"an error detail":         {`{"errors":[{"detail":"the url's file could not be fetched","status":"400"}]}`, "the url's file could not be fetched"},
		"json in another shape":   {`{"meta":{"error":"request timed out"}}`, `{"meta":{"error":"request timed out"}}`},
		"an empty list of errors": {`{"errors":[]}`, `{"errors":[]}`},
		"text that is not json":   {"upstream request timeout", "upstream request timeout"},
		"no body":                 {"", ""},
	} {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusGatewayTimeout)
				w.Write([]byte(tc.body))
			}))
			defer srv.Close()

			api := customerio.NewAPIClient("key")
			api.URL = srv.URL

			request := &customerio.UpdateCollectionActionRequest{Name: "products", URL: "https://files.example/products.json"}
			calls := map[string]func() error{
				"create": func() error { _, err := api.CreateCollectionAction(context.Background(), request); return err },
				"list":   func() error { _, err := api.ListCollectionsAction(context.Background()); return err },
				"update": func() error { _, err := api.UpdateCollectionAction(context.Background(), "4", request); return err },
			}

			for call, do := range calls {
				var e *customerio.CollectionError
				if err := do(); !errors.As(err, &e) {
					t.Fatalf("%s: got %v, want a CollectionError", call, err)
				}
				if e.StatusCode != http.StatusGatewayTimeout || e.Error() != tc.want {
					t.Errorf("%s: got status %d and %q, want status 504 and %q", call, e.StatusCode, e.Error(), tc.want)
				}
			}
		})
	}
}
