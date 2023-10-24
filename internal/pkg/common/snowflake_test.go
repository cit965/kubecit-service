package common

import (
	"testing"
	"time"
)

func TestIdGenerator_GenID(t *testing.T) {
	idg, err := NewIdGenerator(time.Now(), 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(idg.GenID())
}
