package typing_test

import (
	"bytes"
	"testing"

	"github.com/ShunyaNagashige/golang-de/typing"
)

func TestInput(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBufferString("Good morning!")
	expected := "Good morning"
	actual := <-typing.Input(buf)
	
	if actual != expected {
		t.Errorf(`expected="%s",actual="%s"`, expected, actual)
	}
}
