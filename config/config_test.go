package config

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	x := Add(1, 2)
	fmt.Printf("%d", x)
}
