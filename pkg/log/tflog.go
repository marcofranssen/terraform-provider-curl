package log

import (
	"context"
	"io"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TFLogWriter struct {
	ctx context.Context
}

func NewTFLogger(ctx context.Context) io.Writer {
	return &TFLogWriter{ctx}
}

// Write implements io.Writer
func (l *TFLogWriter) Write(p []byte) (n int, err error) {
	tflog.Info(l.ctx, string(p))
	return len(p), nil
}
