package interfaces

import "io"

type countingWriter struct {
	count  *int64
	writer io.Writer
}

func (c countingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err == nil {
		*c.count += int64(n)
	}
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cntr int64
	retval := countingWriter{&cntr, w}
	return retval, retval.count
}

// Altro esempio

type limitReader struct {
	nRem   int
	reader io.Reader
}

func (l *limitReader) Read(p []byte) (int, error) {
	if l.nRem <= 0 {
		return 0, io.EOF
	}
	if len(p) > l.nRem {
		p = p[0:l.nRem]
	}
	n, err := l.reader.Read(p)
	l.nRem -= n
	return n, err
}

// LimitReader returns a Reader that reads from r
// but reports an EOF condition after n bytes.
func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{n, r}
}
