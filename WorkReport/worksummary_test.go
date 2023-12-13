package WorkReport

import (
	"io"
	"testing"
)

type MockReader struct {
	contents []byte
	offset   int
}

func NewMockReader(contents string) *MockReader {
	mockReader := new(MockReader)
	mockReader.contents = []byte(contents)

	return mockReader
}

func (mock *MockReader) Read(buf []byte) (n int, err error) {
	if mock.offset < len(mock.contents) {
		count := len(mock.contents) - mock.offset
		if count > len(buf) {
			count = len(buf)
		}
		copy(buf, mock.contents[mock.offset:mock.offset+count])
		mock.offset += count

		return count, nil
	} else {
		return 0, io.EOF
	}
}

func TestWorkReport_Load(t *testing.T) {

}
