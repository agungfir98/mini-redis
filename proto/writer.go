package proto

import (
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(wr io.Writer) *Writer {
	w := &Writer{writer: wr}

	return w
}

func (w *Writer) Write(v RespMessage) error {
	buf := v.Marshal()

	_, err := w.writer.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
