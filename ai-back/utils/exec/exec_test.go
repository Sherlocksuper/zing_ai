package exece

import (
	"testing"
)

func TestExec(t *testing.T) {
	err := exectest()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestExec")
}
