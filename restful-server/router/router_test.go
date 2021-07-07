package router

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"foo", "bar"},
		{"manu", "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", tt.name, tt.password)))
			if encoded != "" {
				t.Logf("url=%s", encoded)
			}
		})
	}
}
