package deveuigenerator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success case-check if there are 100 ids with unique short-form code",
		},
	}
	for _, tt := range tests {
		uniqueCheckMap := make(map[string]bool)
		t.Run(tt.name, func(t *testing.T) {
			got := Generate()
			if len(got) < 100 {
				t.Errorf("failed to generate 100 ids, %v", got)
			}
			for _, val := range got {
				if uniqueCheckMap[val[len(val)-5:]] {
					t.Errorf("failed to generate unique ids, %v", got)
				}
				uniqueCheckMap[val[len(val)-5:]] = true
			}
		})
	}
}
