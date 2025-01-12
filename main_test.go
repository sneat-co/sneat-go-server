package main

import (
	"testing"
	"time"
)

func TestMainFunc(t *testing.T) {
	finished := false
	var r any
	go func() {
		defer func() {
			r = recover()
		}()
		main()
		finished = true
	}()
	time.Sleep(9 * time.Millisecond)
	if finished {
		t.Fatal("main() finished unexpectedly")
	}
	if r != nil {
		t.Fatalf("panic not expected, got: %v", r)
	}
}
