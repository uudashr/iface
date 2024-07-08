package compressor

import (
	"compress/gzip"
	"io"
)

func encoderGzip(w io.Writer, level int) io.Writer {
	gw, err := gzip.NewWriterLevel(w, level)
	if err != nil {
		return nil
	}
	return gw
}

type EncoderFunc func(w io.Writer, level int) io.Writer

type Compressor struct {
	level    int
	encoders map[string]EncoderFunc
}

func New(level int) *Compressor {
	c := &Compressor{
		level:    level,
		encoders: make(map[string]EncoderFunc),
	}

	c.SetEncoder("gzip", encoderGzip)

	return c
}

func (c *Compressor) SetEncoder(name string, fn EncoderFunc) {
	c.encoders[name] = fn
}
