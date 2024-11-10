package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitMain(_ *testing.T) {
	initMain()
}

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
	assert.False(t, finished)
	assert.Nilf(t, r, "panic not expected, got: %v", r)
}
