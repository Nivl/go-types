package date_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Nivl/go-types/date"
)

func TestToday(t *testing.T) {
	today := date.Today()
	year, month, day := time.Now().UTC().Date()

	assert.Equal(t, year, today.Year(), "Un expected year")
	assert.Equal(t, month, today.Month(), "Un expected month")
	assert.Equal(t, day, today.Day(), "Un expected day")
}

func TestNew(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description   string
		input         string
		shouldFail    bool
		expectedYear  int
		expectedMonth time.Month
		expectedDay   int
	}{
		{
			"2017-09-08 should work",
			"2017-09-08",
			!shouldFail,
			2017, time.September, 8,
		},
		{
			"2013-01 should work",
			"2013-01",
			!shouldFail,
			2013, time.January, 1,
		},
		{
			"03-25-1989 should fail",
			"03-25-1989",
			shouldFail,
			0, time.January, 0,
		},
		{
			"nothing should fail",
			"",
			shouldFail,
			0, time.January, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			d, err := date.New(tc.input)

			if tc.shouldFail {
				assert.Error(t, err, "New() should have fail")
				assert.True(t, d.IsZero(), "d should be zero value")
			} else {
				assert.NoError(t, err, "New() should have work")
				assert.Equal(t, tc.expectedYear, d.Year(), "invalid year")
				assert.Equal(t, tc.expectedMonth, d.Month(), "invalid month")
				assert.Equal(t, tc.expectedDay, d.Day(), "invalid day")
			}
		})
	}
}

func TestValue(t *testing.T) {
	t.Run("valid date should work", func(t *testing.T) {
		t.Parallel()

		d, err := date.New("2017-09-09")
		require.NoError(t, err, "New() should have work")

		v, err := d.Value()
		require.NoError(t, err, "d.Value() should not have fail")
		assert.Equal(t, d.Format(date.DATE), v.(string), "d.Value() returned an unexpected value")
	})

	t.Run("nil date should work", func(t *testing.T) {
		t.Parallel()

		var d *date.Date
		v, err := d.Value()
		require.NoError(t, err, "d.Value() should not have fail")
		assert.Nil(t, v, "d.Value() should have returned nil")
	})
}

func TestScan(t *testing.T) {
	expectedDate, err := date.New("2017-09-09")
	require.NoError(t, err, "New() should have work")

	d := date.Date{}
	err = d.Scan(expectedDate.Time)
	require.NoError(t, err, "dt.Scan() should not have fail")
	assert.Equal(t, expectedDate.String(), d.String(), "dt.Value() should not have fail")
}

func TestScanString(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description   string
		input         string
		shouldFail    bool
		expectedYear  int
		expectedMonth time.Month
		expectedDay   int
	}{
		{
			"2017-09-08 should work",
			"2017-09-08",
			!shouldFail,
			2017, time.September, 8,
		},
		{
			"2013-01 should work",
			"2013-01",
			!shouldFail,
			2013, time.January, 1,
		},
		{
			"03-25-1989 should fail",
			"03-25-1989",
			shouldFail,
			0, time.January, 0,
		},
		{
			"nothing should fail",
			"",
			shouldFail,
			0, time.January, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			d := &date.Date{}
			err := d.ScanString(tc.input)

			if tc.shouldFail {
				assert.Error(t, err, "ScanString() should have fail")
			} else {
				assert.NoError(t, err, "ScanString() should have work")
				assert.Equal(t, tc.expectedYear, d.Year(), "invalid year")
				assert.Equal(t, tc.expectedMonth, d.Month(), "invalid month")
				assert.Equal(t, tc.expectedDay, d.Day(), "invalid day")
			}
		})
	}
}

func TestIsBefore(t *testing.T) {
	// sugar
	shouldBeBefore := true

	testCases := []struct {
		description    string
		source         string
		target         string
		shouldBeBefore bool
	}{
		{
			"2011-01-01 should be before 2017-09-09",
			"2011-01-01", "2017-09-09",
			shouldBeBefore,
		},
		{
			"2017-09-08 should not be before 2017-08-08",
			"2017-09-08", "2017-08-08",
			!shouldBeBefore,
		},
		{
			"2017-09-09 should not be before 2017-09-08",
			"2017-09-09", "2017-09-08",
			!shouldBeBefore,
		},
		{
			"2017-09-09 should not be before 2017-09-09",
			"2017-09-09", "2017-09-09",
			!shouldBeBefore,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			source, err := date.New(tc.source)
			require.NoError(t, err, "New(source) should have not fail")

			target, err := date.New(tc.target)
			require.NoError(t, err, "New(target) should have not fail")

			assert.Equal(t, tc.shouldBeBefore, source.IsBefore(target), "IsBefore() did not return the expected value")
		})
	}
}

func TestIsAfter(t *testing.T) {
	// sugar
	shouldBeAfter := true

	testCases := []struct {
		description   string
		source        string
		target        string
		shouldBeAfter bool
	}{
		{
			"2011-01-01 should not be after 2017-09-09",
			"2011-01-01", "2017-09-09",
			!shouldBeAfter,
		},
		{
			"2017-09-08 should be after 2017-08-08",
			"2017-09-08", "2017-08-08",
			shouldBeAfter,
		},
		{
			"2017-09-09 should be after 2017-09-08",
			"2017-09-09", "2017-09-08",
			shouldBeAfter,
		},
		{
			"2017-09-09 should not be after 2017-09-09",
			"2017-09-09", "2017-09-09",
			!shouldBeAfter,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			source, err := date.New(tc.source)
			require.NoError(t, err, "New(source) should have not fail")

			target, err := date.New(tc.target)
			require.NoError(t, err, "New(target) should have not fail")

			assert.Equal(t, tc.shouldBeAfter, source.IsAfter(target), "IsAfter() did not return the expected value")
		})
	}
}
func TestEqual(t *testing.T) {
	// sugar
	shouldBeEqual := true

	testCases := []struct {
		description   string
		source        string
		target        string
		shouldBeEqual bool
	}{
		{
			"2017-09 should be equal 2017-09-01",
			"2017-09", "2017-09-01",
			shouldBeEqual,
		},
		{
			"2016-09-01 should not be equal 2017-09-01",
			"2016-09-01", "2017-09-01",
			!shouldBeEqual,
		},
		{
			"2017-09-01 should not be equal 2017-10-01",
			"2017-09-01", "2017-10-01",
			!shouldBeEqual,
		},
		{
			"2017-09-02 should not be equal 2017-09-01",
			"2017-09-02", "2017-09-01",
			!shouldBeEqual,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			source, err := date.New(tc.source)
			require.NoError(t, err, "New(source) should have not fail")

			target, err := date.New(tc.target)
			require.NoError(t, err, "New(target) should have not fail")

			assert.Equal(t, tc.shouldBeEqual, source.Equal(target), "Equal() did not return the expected value")
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	rawDate := "2017-09-07"
	dt, err := date.New(rawDate)
	require.NoError(t, err, "New() shoud have not failed")

	t.Run("raw function", func(t *testing.T) {
		data, err := dt.MarshalJSON()
		assert.NoError(t, err, "MarshalJSON() should have work")
		assert.Equal(t, `"`+rawDate+`"`, string(data), "Wrong data returned by MarshalJSON()")
	})

	t.Run("json.Marshal", func(t *testing.T) {
		testStruct := struct {
			Date date.Date `json:"date"`
		}{Date: dt}
		expected := fmt.Sprintf(`{"date":"%s"}`, rawDate)

		output, err := json.Marshal(&testStruct)
		assert.NoError(t, err, "json.Marshal() should have work")
		assert.Equal(t, expected, string(output), "json.Marshal() did not return the expected output")
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("raw function", func(t *testing.T) {
		d := date.Date{}
		err := d.UnmarshalJSON([]byte(`"2017-09-09"`))
		assert.NoError(t, err, "UnmarshalJSON() should have work")

		assert.Equal(t, 2017, d.Year())
		assert.Equal(t, time.September, d.Month())
		assert.Equal(t, 9, d.Day())
		assert.Equal(t, 0, d.Hour())
		assert.Equal(t, 0, d.Minute())
		assert.Equal(t, 0, d.Second())
	})

	t.Run("json.Decode", func(t *testing.T) {
		body := strings.NewReader(`{"date":"2017-09-08"}`)
		var pld struct {
			Date *date.Date `json:"date"`
		}

		err := json.NewDecoder(body).Decode(&pld)
		assert.NoError(t, err, "json.NewDecoder() should have work")
		if assert.NotNil(t, pld, "pld should not be nil") {
			assert.Equal(t, 2017, pld.Date.Year())
			assert.Equal(t, time.September, pld.Date.Month())
			assert.Equal(t, 8, pld.Date.Day())
			assert.Equal(t, 0, pld.Date.Hour())
			assert.Equal(t, 0, pld.Date.Minute())
			assert.Equal(t, 0, pld.Date.Second())
		}
	})

	t.Run("null", func(t *testing.T) {
		d := date.Date{}
		err := d.UnmarshalJSON([]byte(`null`))
		assert.NoError(t, err, "UnmarshalJSON() should have work")

		assert.True(t, d.IsZero())
	})
}
