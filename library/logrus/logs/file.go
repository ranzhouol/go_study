package logs

import "os"

const LOGPATH = "library/logrus/runtime/logs/log.log"

type fileWriter struct {
	*os.File
}

func (s *fileWriter) Flush() {
	s.Sync()
}

func newFileWriter() LogWriter {
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		file = os.Stderr
	}

	return &fileWriter{
		file,
	}
}

func init() {
	RegisterInitWriterFunc("file", newFileWriter)
}
