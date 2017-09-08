package filetype_test

import (
	"errors"
	"image"
	"io"
	"os"
	"path"
	"testing"

	"github.com/Nivl/go-types/filetype"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIsImage(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description  string
		filename     string
		shouldFail   bool
		expectedMime string
	}{
		{"gif should work", "black_pixel.gif", !shouldFail, "image/gif"},
		{"png should work", "black_pixel.png", !shouldFail, "image/png"},
		{"jpg should work", "black_pixel.jpg", !shouldFail, "image/jpeg"},
		{"pdf should fail", "black_pixel.pdf", shouldFail, ""},
		{"LICENSE should fail", "LICENSE", shouldFail, ""},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			filePath := path.Join("fixtures", tc.filename)
			f, err := os.Open(filePath)
			if assert.NoError(t, err, "Open should not have failed") {
				isValid, mime, err := filetype.IsImage(f)
				if tc.shouldFail {
					assert.Error(t, err, "IsImage should have failed")
					assert.Empty(t, mime, "mime should be empty")
					assert.False(t, isValid, "isValid should be false")
				} else {
					assert.NoError(t, err, "IsImage should not have failed")
					assert.Equal(t, tc.expectedMime, mime, "ivalid mime")
					assert.True(t, isValid, "isValid should be true")
				}
			}
		})
	}
}

func TestIsImageMimeFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Read(gomock.Any()).Return(0, errors.New("read failed"))

	isValid, mime, err := filetype.IsImage(reader)
	assert.Error(t, err, "IsImage() should have failed")
	assert.Empty(t, mime, "IsImage() should have not returned a mime")
	assert.False(t, isValid, "isValid should have been false")
}

func TestValidateImageSeekFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Seek(int64(0), io.SeekCurrent).Return(int64(0), errors.New("seek failed"))

	isValid, err := filetype.ValidateImage(reader, nil)
	assert.Error(t, err, "ValidateImage() should have failed")
	assert.False(t, isValid, "isValid should have been false")
}

func TestValidateImageSeekStartFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Seek(int64(0), io.SeekCurrent).Return(int64(50), nil)
	reader.EXPECT().Seek(int64(50), io.SeekStart).Return(int64(0), errors.New("seek failed"))

	isValid, err := filetype.ValidateImage(reader, func(r io.Reader) (image.Image, error) {
		return nil, nil
	})
	assert.Error(t, err, "ValidateImage() should have failed")
	assert.False(t, isValid, "isValid should have been false")
}

func TestIsGIF(t *testing.T) {
	// sugar
	shouldBeValid := true

	testCases := []struct {
		description   string
		filename      string
		shouldBeValid bool
	}{
		{"gif should work", "black_pixel.gif", shouldBeValid},
		{"png should fail", "black_pixel.png", !shouldBeValid},
		{"jpg should fail", "black_pixel.jpg", !shouldBeValid},
		{"pdf should fail", "black_pixel.pdf", !shouldBeValid},
		{"LICENSE should fail", "LICENSE", !shouldBeValid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			filePath := path.Join("fixtures", tc.filename)
			f, err := os.Open(filePath)
			if assert.NoError(t, err, "Open should not have failed") {
				isValid, err := filetype.IsGIF(f)
				assert.NoError(t, err, "IsGIF should not have failed")
				assert.Equal(t, tc.shouldBeValid, isValid, "IsGIF did not return the expected value")
			}
		})
	}
}

func TestIsPNG(t *testing.T) {
	// sugar
	shouldBeValid := true

	testCases := []struct {
		description   string
		filename      string
		shouldBeValid bool
	}{
		{"png should fail", "black_pixel.png", shouldBeValid},
		{"gif should work", "black_pixel.gif", !shouldBeValid},
		{"jpg should fail", "black_pixel.jpg", !shouldBeValid},
		{"pdf should fail", "black_pixel.pdf", !shouldBeValid},
		{"LICENSE should fail", "LICENSE", !shouldBeValid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			filePath := path.Join("fixtures", tc.filename)
			f, err := os.Open(filePath)
			if assert.NoError(t, err, "Open should not have failed") {
				isValid, err := filetype.IsPNG(f)
				assert.NoError(t, err, "IsPNG should not have failed")
				assert.Equal(t, tc.shouldBeValid, isValid, "IsPNG did not return the expected value")
			}
		})
	}
}

func TestIsJPG(t *testing.T) {
	// sugar
	shouldBeValid := true

	testCases := []struct {
		description   string
		filename      string
		shouldBeValid bool
	}{
		{"jpg should fail", "black_pixel.jpg", shouldBeValid},
		{"gif should work", "black_pixel.gif", !shouldBeValid},
		{"png should fail", "black_pixel.png", !shouldBeValid},
		{"pdf should fail", "black_pixel.pdf", !shouldBeValid},
		{"LICENSE should fail", "LICENSE", !shouldBeValid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			filePath := path.Join("fixtures", tc.filename)
			f, err := os.Open(filePath)
			if assert.NoError(t, err, "Open should not have failed") {
				isValid, err := filetype.IsJPG(f)
				assert.NoError(t, err, "IsJPG should not have failed")
				assert.Equal(t, tc.shouldBeValid, isValid, "IsJPG did not return the expected value")
			}
		})
	}
}
