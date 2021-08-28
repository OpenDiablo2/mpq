package pkg

import (
	"io"
)

// DataStream represents a stream for MPQ data.
type DataStream struct {
	stream *Stream
}

// Read reads data from the data stream
func (m *DataStream) Read(p []byte) (n int, err error) {
	totalRead, err := m.stream.Read(p, 0, uint32(len(p)))
	return int(totalRead), err
}

// Seek sets the position of the data stream
func (m *DataStream) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		m.stream.Position = uint32(offset)
	case io.SeekEnd:
		m.stream.Position = uint32(m.stream.Size) - uint32(offset)
	case io.SeekCurrent:
		m.stream.Position += uint32(offset)
	}

	return int64(m.stream.Position), nil
}

// Close closes the data stream
func (m *DataStream) Close() error {
	m.stream = nil
	return nil
}
