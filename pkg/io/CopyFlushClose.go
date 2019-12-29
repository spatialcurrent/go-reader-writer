// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package io

import (
	"fmt"
)

// CopyFlushClose copies all data from the reader to the writer,
// flushes the writer, and then
// closes the reader and writer.
// If there is an error closing the reader,
// will still attempt to close the writer.
func CopyFlushClose(w Writer, r Reader) error {

	_, err := Copy(w, r)
	if err != nil {
		return fmt.Errorf("error copying from reader to writer: %w", err)
	}

	err = Flush(w)
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	errReader, errWriter := CloseReaderAndWriter(r, w, true)
	if errReader != nil {
		return fmt.Errorf("error closing reader: %w", errReader)
	}

	if errWriter != nil {
		return fmt.Errorf("error closing writer: %w", errWriter)
	}

	return nil
}
