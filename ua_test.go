package ua

import (
	"testing"
)

func TestRandomUserAgent(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomUserAgent())
	}
}

func TestRandomMobileUserAgent(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomMobileUserAgent())
	}
}
