package auth

import (
	"testing"
	// "errors"
)

func TestGetAPIKey(t *testing.T) {
	got, err := GetAPIKey(map[string][]string{"Authorization": {"ApiKey 123"}})
	if err != nil {
		t.Fatal(err)
	}
	want := "123"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestGetAuthHeaderMissingError(t *testing.T) {
	_, err := GetAPIKey(map[string][]string{})
	if err.Error() != "no authorization header included" {
		t.Fatalf("got %q, want %q", err, "no authorization header included")
	}
}

func TestGetAPIKeyL2Error(t *testing.T) {
	_, err := GetAPIKey(map[string][]string{"Authorization": {"Bearer 1"}})
	if err.Error() != "malformed authorization header" {
		t.Fatalf("got %q, want %q", err, "malformed authorization header")
	}
}

func TestGetAPIKeyNoKeyError(t *testing.T) {
	_, err := GetAPIKey(map[string][]string{"Authorization": {"1234"}})
	if err.Error() != "malformed authorization header" {
		t.Fatalf("got %q, want %q", err, "malformed authorization header")
	}
}
