package pipe

import (
	"v2ray.com/core/common/buf"
)

// Writer is a buf.Writer that writes data into a pipe.
type Writer struct {
	pipe *pipe
}

// WriteMultiBuffer implements buf.Writer.
func (w *Writer) WriteMultiBuffer(mb buf.MultiBuffer) error {
	return w.pipe.WriteMultiBuffer(mb)
}

// Close implements io.Closer. After the pipe is closed, writing to the pipe will return io.ErrClosedPipe, while reading will return io.EOF.
func (w *Writer) Close() error {
	return w.pipe.Close()
}

// CloseError sets the pipe to error state. Both reading and writing from/to the pipe will return io.ErrClosedPipe.
func (w *Writer) CloseError() {
	w.pipe.CloseError()
}
