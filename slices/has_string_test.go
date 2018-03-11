package slices_test

import (
	"fmt"
	"testing"

	"github.com/Nivl/go-types/slices"
	"github.com/stretchr/testify/assert"
)

func TestHasString(t *testing.T) {
	data := []string{"ONE", "TWO", "FOUR"}

	t.Run("Parallel", func(t *testing.T) {
		testCases := []struct {
			toFind        string
			shouldBeFound bool
		}{
			{toFind: "ONE", shouldBeFound: true},
			{toFind: "TWO", shouldBeFound: true},
			{toFind: "THREE", shouldBeFound: false},
			{toFind: "FOUR", shouldBeFound: true},
		}
		for _, tc := range testCases {
			tc := tc
			description := fmt.Sprintf("Looking for %s", tc.toFind)
			t.Run(description, func(t *testing.T) {
				t.Parallel()
				found := slices.HasString(tc.toFind, data)
				assert.Equal(t, tc.shouldBeFound, found)
			})
		}
	})
}
