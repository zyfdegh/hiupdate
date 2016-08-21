package util

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestGetDate(t *testing.T) {
	var cases = []struct {
		t      time.Time
		expect string
	}{
		{time.Now(), "20160821"},
	}

	for _, c := range cases {
		//call
		got := GetDate(c.t)
		//assert
		assert.Equal(t, c.expect, got)
	}
}
