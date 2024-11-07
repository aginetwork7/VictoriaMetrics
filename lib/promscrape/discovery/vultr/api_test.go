package vultr

import (
	"testing"

	"github.com/aginetwork7/VictoriaMetrics/lib/promauth"
)

func TestNewAPIConfig_Failure(t *testing.T) {
	sdc := &SDConfig{}
	baseDir := "."
	_, err := newAPIConfig(sdc, baseDir)
	if err == nil {
		t.Fatalf("expecting non-nil error")
	}
}

func TestNewAPIConfig_Success(t *testing.T) {
	sdc := &SDConfig{
		HTTPClientConfig: promauth.HTTPClientConfig{
			BearerToken: &promauth.Secret{
				S: "foobar",
			},
		},
	}
	baseDir := "."
	_, err := newAPIConfig(sdc, baseDir)
	if err != nil {
		t.Fatalf("newAPIConfig failed with, err: %v", err)
	}
}
