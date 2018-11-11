package pkg

import (
	"testing"
)

func Test_buildDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"basic time func test", "November 11, 2018"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildDate(); got != tt.want {
				t.Errorf("buildDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleJSONBody(t *testing.T) {
	var j = `{
"size":"medium",
"colour":"yellow",
"style":"nice",
"friends": {"first":"alice","second":"bob"}
}`

	o := handleJSONBody(j)
	if (*o)["size"] != "medium" {
		t.Errorf("JSONBody handler failed %+v", o)
	}

}
