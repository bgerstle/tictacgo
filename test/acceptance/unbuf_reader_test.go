package tictacgo_test

import (
	"io"
	"strings"
)

type UnbufferedReader struct {
	reader io.Reader
}

func NewUnbufferedReader(reader io.Reader) UnbufferedReader {
	return UnbufferedReader{reader: reader}
}

func (r UnbufferedReader) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		readCount, readErr := r.reader.Read(buf[i : i+1])
		if readErr != nil {
			return buf, readErr
		}
		if readCount == 0 {
			break
		}
	}
	return buf, nil
}

func (r UnbufferedReader) ReadByte() (byte, error) {
	bs, error := r.ReadBytes(1)
	return bs[0], error
}

func (r UnbufferedReader) ReadBytesUntil(delim byte) (s string, e error) {
	builder := strings.Builder{}
	for {
		b, e := r.ReadByte()
		if e != nil {
			break
		}
		if b == delim {
			break
		}
		builder.WriteByte(b)
	}
	s = builder.String()
	return
}

func (r UnbufferedReader) ReadLinesUntil(sentinel string, maxLines int) (lines []string, error error) {
	for i := 0; i < maxLines; i++ {
		line, error := r.ReadBytesUntil(byte('\n'))
		if error != nil {
			break
		}
		lines = append(lines, line)
		if line == sentinel {
			break
		}
	}
	return
}
