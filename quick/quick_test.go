package quick

import (
	"fmt"
	"testing"
	"testing/quick"
)

var N = 100000

func TestAddWithSystem(t *testing.T) {
	conditions := func(a, b uint16) bool {
		return Add(a, b) == (a + b)
	}
	err := quick.Check(conditions, &quick.Config{MaxCount: N})
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddWithItSelf(t *testing.T) {
	conditions := func(a, b uint16) bool {
		return Add(a, b) == Add(b, a)
	}
	err := quick.Check(conditions, &quick.Config{MaxCount: N})
	if err != nil {
		fmt.Println("error: ", err)
	}
}
