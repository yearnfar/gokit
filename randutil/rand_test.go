package randutil

import "testing"

func TestString(t *testing.T) {

	s := String(10)
	t.Log(s)

}

func TestNumber(t *testing.T) {
	s := Number(10)
	t.Log(s)
}
