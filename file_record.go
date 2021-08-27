package mpq

// FileRecord represents a file record in an MPQ
type FileRecord struct {
	MpqFile          string
	IsPatch          bool
	UnpatchedMpqFile string
}
