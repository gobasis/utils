package utils

import (
	"testing"
	"fmt"
)

func TestUUID(t *testing.T) {
	fmt.Println(UUID.Get())
	fmt.Println(UUID.Get())
}