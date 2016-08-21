package util

import (
	"testing"

	"github.com/bmizerany/assert"
)

// 选择 32[小写]
// http://tool.chinaz.com/tools/md5.aspx
func TestGenerateID(t *testing.T) {
	var cases = []struct {
		name   string
		date   string
		expect string
	}{
		{"Abc", "20160820", "3a11d432c75a9f8b9c50dd280e424a40"},
		{"张三", "20160821", "b8580455e406b89f6758381a3127970f"},
	}

	for _, c := range cases {
		//call
		got := GenerateID(c.name, c.date)
		//assert
		assert.Equal(t, c.expect, got)
	}
}
