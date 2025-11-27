package utils

import "testing"

func TestTool(t *testing.T) {
	testAdd(t)
}

func testAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"case1", 2, 3, 5},
		{"case2", -1, -1, -2},
		{"case3", 0, 0, 999}, // 故意错误
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("失败")
			}
		})
	}
}
