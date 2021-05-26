package main

import (
	"strings"
	"testing"
)

func TestWorker(t *testing.T) {
	in := make(chan string)
	out := make(chan string)

	go Worker(in, out)

	in <- ""
	if !strings.Contains(<-out, "unsupported protocol scheme") {
		t.FailNow()
	}

	close(in)
}
