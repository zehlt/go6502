package asrt

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatalf("Received: %v (%v), expected: %v (%v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func True(t *testing.T, a bool) {
	Equal(t, a, true)
}

func False(t *testing.T, a bool) {
	Equal(t, a, false)
}
