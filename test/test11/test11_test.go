package test11

import (
	"testing"
)

func TestShow(t *testing.T) {
	Show()
}
func BenchmarkShow(b *testing.B) {
	Show()
}
