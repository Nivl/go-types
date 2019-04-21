//go:generate mockgen -destination mocks_test.go -package filetype_test io ReadSeeker,Reader

package filetype_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"os"
	"path"

	"github.com/Nivl/go-types/filetype"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA256Sum(t *testing.T) {
	testCases := []struct {
		content  string
		expected string
	}{
		{"this is a test", "2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.content, func(t *testing.T) {
			t.Parallel()

			r := bytes.NewReader([]byte(tc.content))
			sum, err := filetype.SHA256Sum(r)
			assert.NoError(t, err, "SHA256Sum() should have succeed")
			assert.Equal(t, tc.expected, sum, "invalid sum")
		})
	}
}

func TestSHA256SumSeekFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Seek(int64(0), io.SeekCurrent).Return(int64(0), errors.New("seek failed"))

	mime, err := filetype.SHA256Sum(reader)
	assert.Error(t, err, "SHA256Sum() should have failed")
	assert.Empty(t, mime, "SHA256Sum() should have not returned a value")
}

func TestSHA256SumCopyFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Seek(int64(0), io.SeekCurrent).Return(int64(1), nil)
	reader.EXPECT().Read(gomock.Any()).Return(0, errors.New("read failed"))
	reader.EXPECT().Seek(int64(1), io.SeekStart).Return(int64(0), nil)

	mime, err := filetype.SHA256Sum(reader)
	assert.Error(t, err, "SHA256Sum() should have failed")
	assert.Empty(t, mime, "SHA256Sum() should have not returned a value")
}

func TestMimeType(t *testing.T) {
	testCases := []struct {
		description string
		filename    string
		expected    string
	}{
		{"png", "black_pixel.png", "image/png"},
		{"jpg", "black_pixel.jpg", "image/jpeg"},
		{"pdf", "black_pixel.pdf", "application/pdf"},
		{"text file with no ext", "LICENSE", "text/plain; charset=utf-8"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			filePath := path.Join("fixtures", tc.filename)
			f, err := os.Open(filePath)
			require.NoError(t, err, "os.Open() should have succeed")
			defer func() {
				assert.NoError(t, f.Close(), "Close() should have worked")
			}()
			mime, err := filetype.MimeType(f)
			assert.NoError(t, err, "MimeType() should have succeed")
			assert.Equal(t, tc.expected, mime, "invalid mimetype")
		})
	}
}

func TestMimeTypeReaderFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Read(gomock.Any()).Return(0, errors.New("read failed"))

	mime, err := filetype.MimeType(reader)
	assert.Error(t, err, "MimeType() should have failed")
	assert.Empty(t, mime, "MimeType() should have not returned a value")
}

func TestMimeTypeSeekFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reader := NewMockReadSeeker(mockCtrl)
	reader.EXPECT().Read(gomock.Any()).Return(50, nil)
	reader.EXPECT().Seek(int64(-50), io.SeekCurrent).Return(int64(0), errors.New("seek failed"))

	mime, err := filetype.MimeType(reader)
	assert.Error(t, err, "MimeType() should have failed")
	assert.Empty(t, mime, "MimeType() should have not returned a value")
}
