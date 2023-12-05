package logs

import "io"

var WriterAdapter = make(map[string]InitLogWriterFunc, 0)

type InitLogWriterFunc func() LogWriter

type LogWriter interface {
	Flush() //关闭连接
	io.Writer
}

func RegisterInitWriterFunc(adapterName string, writerFunc InitLogWriterFunc) {
	WriterAdapter[adapterName] = writerFunc
}
