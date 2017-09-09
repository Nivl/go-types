package datetime_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Nivl/go-types/datetime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {
	loc, err := time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err, "timezone not found")

	tm := time.Date(2017, time.September, 7, 23, 18, 42, 0, loc)
	dt := &datetime.DateTime{Time: tm}

	v, err := dt.Value()
	require.NoError(t, err, "dt.Value() should not have fail")
	assert.Equal(t, dt.Format(datetime.ISO8601), v.(string), "dt.Value() should not have fail")
}

func TestValueNil(t *testing.T) {
	var dt *datetime.DateTime

	v, err := dt.Value()
	require.NoError(t, err, "dt.Value() should not have fail")
	assert.Nil(t, v, "dt.Value() should return a nil value")
}

func TestScan(t *testing.T) {
	loc, err := time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err, "timezone not found")

	tm := time.Date(2017, time.September, 7, 23, 18, 42, 0, loc)

	dt := &datetime.DateTime{}
	err = dt.Scan(tm)
	require.NoError(t, err, "dt.Scan() should not have fail")
	assert.Equal(t, tm.String(), dt.String(), "dt.Value() should not have fail")
}

func TestEqual(t *testing.T) {
	loc, err := time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err, "timezone not found")

	tm := time.Date(2017, time.September, 7, 23, 18, 42, 0, loc)
	dt := &datetime.DateTime{Time: tm}

	isEqual := dt.Equal(&datetime.DateTime{Time: tm})
	assert.True(t, isEqual, "Equal should have returned true")

	isEqual = dt.Equal(datetime.Now())
	assert.False(t, isEqual, "Equal should have returned false")
}

func TestAddDate(t *testing.T) {
	loc, err := time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err, "timezone not found")

	tm := time.Date(2017, time.September, 7, 23, 18, 42, 0, loc)
	dt := &datetime.DateTime{Time: tm}

	newDate := dt.AddDate(1, -2, 7)
	assert.Equal(t, 2018, newDate.Year())
	assert.Equal(t, time.July, newDate.Month())
	assert.Equal(t, 14, newDate.Day())
	assert.Equal(t, 23, newDate.Hour())
	assert.Equal(t, 18, newDate.Minute())
	assert.Equal(t, 42, newDate.Second())
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("raw function", func(t *testing.T) {
		dt := &datetime.DateTime{}
		err := dt.UnmarshalJSON([]byte(`"2017-09-07T23:18:42-0700"`))
		assert.NoError(t, err, "UnmarshalJSON() should have work")

		assert.Equal(t, 2017, dt.Year())
		assert.Equal(t, time.September, dt.Month())
		assert.Equal(t, 7, dt.Day())
		assert.Equal(t, 23, dt.Hour())
		assert.Equal(t, 18, dt.Minute())
		assert.Equal(t, 42, dt.Second())
	})

	t.Run("json.Decode", func(t *testing.T) {
		body := strings.NewReader(`{"date":"2017-09-07T23:18:42-0700"}`)
		var pld struct {
			Datetime *datetime.DateTime `json:"date"`
		}

		err := json.NewDecoder(body).Decode(&pld)
		assert.NoError(t, err, "json.NewDecoder() should have work")
		if assert.NotNil(t, pld, "pld should not be nil") {
			assert.Equal(t, 2017, pld.Datetime.Year())
			assert.Equal(t, time.September, pld.Datetime.Month())
			assert.Equal(t, 7, pld.Datetime.Day())
			assert.Equal(t, 23, pld.Datetime.Hour())
			assert.Equal(t, 18, pld.Datetime.Minute())
			assert.Equal(t, 42, pld.Datetime.Second())
		}
	})
}

func TestMarshalJSON(t *testing.T) {
	loc, err := time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err, "timezone not found")
	tm := time.Date(2017, time.September, 7, 23, 18, 42, 0, loc)
	dt := &datetime.DateTime{Time: tm}
	expectedOutput := `"2017-09-07T23:18:42-0700"`

	t.Run("raw function", func(t *testing.T) {
		data, err := dt.MarshalJSON()
		assert.NoError(t, err, "MarshalJSON() should have work")
		assert.Equal(t, expectedOutput, string(data), "Wrong data returned by MarshalJSON()")
	})

	t.Run("json.Marshal", func(t *testing.T) {
		testStruct := struct {
			Datetime *datetime.DateTime `json:"date"`
		}{Datetime: dt}
		expected := fmt.Sprintf(`{"date":%s}`, expectedOutput)

		output, err := json.Marshal(&testStruct)
		assert.NoError(t, err, "json.Marshal() should have work")
		assert.Equal(t, expected, string(output), "json.Marshal() did not return the expected output")
	})
}
