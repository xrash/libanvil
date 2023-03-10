package libanvil

import (
	"fmt"
	"testing"
)

func TestAnvil(t *testing.T) {
	a, err := RunAnvil(nil)
	if err != nil {
		fmt.Println("error running anvil", err)
		t.Fail()
	}

	if a == nil {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	defer a.Stop()
}
