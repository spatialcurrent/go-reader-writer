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

// CloseReaderAndWriter closes an input reader and output writer,
// and does not flush the writer if the pipe is broken.
// CloseReaderAndWriter will close the writer even if closing the reader returned an error.
//
// The returns values are the errors returned when closing the reader and closing the writer, respectively.
func CloseReaderAndWriter(inputReader Reader, outputWriter Writer, brokenPipe bool) (error, error) {

	errorReader := Close(inputReader)
	if errorReader != nil {
		errorReader = fmt.Errorf("error closing input resource: %w", errorReader)
	}

	if brokenPipe {
		// close without flushing to output writer
		errorWriter := Close(outputWriter)
		if errorWriter != nil {
			errorWriter = fmt.Errorf("error closing output resource: %w", errorWriter)
		}
		return errorReader, errorWriter
	}

	errorFlusher := Flush(outputWriter)
	if errorFlusher != nil {
		errorFlusher = fmt.Errorf("error flushing output resource: %w", errorFlusher)
	}

	errorWriter := Close(outputWriter)
	if errorWriter != nil {
		errorWriter = fmt.Errorf("error closing output resource: %w", errorWriter)
	}

	if errorFlusher != nil && errorWriter == nil {
		return errorReader, errorFlusher
	}

	return errorReader, errorWriter
}
