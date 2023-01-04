package structfield

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {

	tests := []struct {
		name       string
		dst        interface{}
		src        interface{}
		wantErr    bool
		wantString string
	}{
		{
			name: "Testcase B only",
			dst: &struct {
				A string
				B string
			}{"-", "-"},
			src: struct {
				B string
				C string
			}{"B", "C"},
			wantErr:    false,
			wantString: "&{- B}"},
		{
			name: "Testcase no pointer",
			dst: struct {
				A string
				B string
			}{"-", "-"},
			src: struct {
				B string
				C string
			}{"B", "C"},
			wantErr:    true,
			wantString: "{- -}"},
		{
			name: "Different field type",
			dst: &struct {
				A string
				B int
			}{"-", 5},
			src: struct {
				B string
				C string
			}{"B", "C"},
			wantErr:    false,
			wantString: "&{- 5}"},
		{
			name: "Tags",
			dst: &struct {
				A string
				B string
				C string
			}{"-", "-", "-"},
			src: struct {
				A string `structfield:"nocopy"`
				B string
				C string
			}{"A", "B", "C"},
			wantErr:    false,
			wantString: "&{- B C}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.dst, tt.src); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if fmt.Sprintf("%v", tt.dst) != tt.wantString {
				t.Errorf("Got: \n%v\nWant:\n%s", tt.dst, tt.wantString)
			}
		})
	}
}
