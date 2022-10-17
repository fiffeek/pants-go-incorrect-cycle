package example_test

import (
	"github.com/fiffeek/pants-go-incorrect-cycle/src/example"
	"github.com/fiffeek/pants-go-incorrect-cycle/src/example/testutils"
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		if example.ExampleValue != testutils.ExampleValueFromTestutils {
			t.Error("Not equal")
		}
	})
}
