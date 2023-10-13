package cmd

import "testing"

func TestIsConfig(t *testing.T) {

	got := isConfig("examples/tracking_example")

	if got != true {
		t.Errorf("got %t, wanted %t", got, true)
	}
}
