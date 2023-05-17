package helpers

import "testing"

func TestContains(t *testing.T) {
	items := []string{"test", "notAtest", "shouldMatch"}
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"notThere", "nope", false},
		{"there", "test", true},
		{"notThere-case-sensitive", "notatest", false},
		{"there-case-sensitive", "notAtest", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.value, items); got != tt.want {
				t.Errorf("Contains(%v) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
