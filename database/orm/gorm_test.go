package orm

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
		name: "1",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init()
		})
	}
}
