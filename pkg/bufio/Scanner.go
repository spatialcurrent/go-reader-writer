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

// A scanner that propagates the close method.
type Scanner struct {
	*bufio.Scanner
	underlying io.Closer
}

func (s *Scanner) Close() error {
	err := s.underlying.Close()
	if err != nil {
		return fmt.Errorf("error closing underlying scanner: %w", err)
	}
	return nil
}

// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.ReadCloser) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(r), underlying: r}
}
