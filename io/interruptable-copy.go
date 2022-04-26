package io

/*
golang标准库里的Copy，是基于dst来判断复制是否继续，InerruptableCopy可以通过context，从外部控制流复制。当调用cancelF()后，应尽快让Read返回，才能达到中断目的。
例如对于net.Conn,使用 Conn.SetDeadline来触发read timeout。
*/
import (
	"context"
	"io"
)

type readerFunc func(p []byte) (n int, err error)

func (rdf readerFunc) Read(p []byte) (n int, err error) {
	return rdf(p)
}

func InerruptableCopy(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	n, err := io.Copy(dst, readerFunc(func(p []byte) (int, error) {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			return src.Read(p)
		}
	}))
	return n, err
}
