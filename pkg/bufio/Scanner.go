// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package bufio

import (
	"bufio"
	"fmt"
	"io"
)

type Scanner struct {
	*bufio.Scanner
	underlying io.Reader
}

// Close, calls the Close method of the underlying reader, if it implements io.Closer.
func (s *Scanner) Close() error {
	if c, ok := s.underlying.(io.Closer); ok {
		err := c.Close()
		if err != nil {
			return fmt.Errorf("error closing underlying reader: %w", err)
		}
	}
	return nil
}

// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(r), underlying: r}
}
