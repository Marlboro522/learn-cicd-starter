package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey hellomothefather"},
			},
			wantAPIKey: "hellomothefather",
			wantErr:    nil,
		},
		{
			name:       "missing authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name: "wrong authorization scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer hellomothefather"},
			},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
		{
			name: "missing API key value",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	//run the fucking tests.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)

			//handle the ouputs.
			// tt is the test case.
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf(
						"expected error %q, got nil",
						tt.wantErr,
					)
				}

				if err.Error() != tt.wantErr.Error() {
					t.Fatalf(
						"expected error %q, got %q",
						tt.wantErr,
						err,
					)
				}
			} else if err != nil {
				t.Fatalf(
					"expected no error, got %v",
					err,
				)
			}

			if gotAPIKey != tt.wantAPIKey {
				t.Errorf(
					"expected API key %q, got %q",
					tt.wantAPIKey,
					gotAPIKey,
				)
			}
		})
	}

}
