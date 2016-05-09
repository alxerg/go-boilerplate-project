// hello_world_test.go
package helloworld

import "testing"

func TestHello(t *testing.T) {
	exp := "hello go world"
	got := Hello("Go")

	if exp != got {
		t.Errorf("\nExp: %s\nGot: %s", exp, got)
	}
}
