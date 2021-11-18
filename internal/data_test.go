package internal_test

import (
	"errors"
	"testing"

	"github.com/SaWLeaDeR/key-value-store/internal"
)

func TestData_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Data
		withErr bool
	}{
		{
			"OK",
			internal.Data{
				Key:   "dummyKey",
				Value: "dummyValue",
			},
			false,
		},
		{
			"ERR: Data",
			internal.Data{
				Key: "dummyData",
			},
			true,
		},
		{
			"ERR: Data",
			internal.Data{
				Key: "Key",
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.Validate()
			if (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}

			var ierr *internal.Error
			if tt.withErr && !errors.As(actualErr, &ierr) {
				t.Fatalf("expected %T error, got %T", ierr, actualErr)
			}
		})
	}
}
