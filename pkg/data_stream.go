package pkg

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
	m.stream.Position = uint32(offset + int64(whence))
	return int64(m.stream.Position), nil
}

// Close closes the data stream
func (m *DataStream) Close() error {
	m.stream = nil
	return nil
}
