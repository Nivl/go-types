package slices_test

import (
	"fmt"
	"testing"

	"github.com/Nivl/go-types/slices"
	"github.com/stretchr/testify/assert"
)

func TestHasInt(t *testing.T) {
	data := []int{1, 0, 5, 8}

	t.Run("Parallel", func(t *testing.T) {
		testCases := []struct {
			toFind        int
			shouldBeFound bool
		}{
			{toFind: 1, shouldBeFound: true},
			{toFind: 2, shouldBeFound: false},
			{toFind: 0, shouldBeFound: true},
			{toFind: 8, shouldBeFound: true},
			{toFind: 3, shouldBeFound: false},
		}
		for _, tc := range testCases {
			tc := tc
			description := fmt.Sprintf("Looking for %d", tc.toFind)
			t.Run(description, func(t *testing.T) {
				t.Parallel()
				found := slices.HasInt(tc.toFind, data)
				assert.Equal(t, tc.shouldBeFound, found)
			})
		}
	})
}
