// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"github.com/pkg/errors"
)

//CloseReaderAndWriter closes an input ByteReadCloser and output ByteWriteCloser,
// and does not flush the writer if the pipe is broken.
// CloseReaderAndWriter will continue and close the writer even if closing the reader returned an error.
//
// The returns values are the errors returned when closing the reader and closing the writer, respectively.
func CloseReaderAndWriter(inputReader ByteReadCloser, outputWriter ByteWriteCloser, brokenPipe bool) (error, error) {
	errorReader := inputReader.Close()
	if errorReader != nil {
		errorReader = errors.Wrap(errorReader, "error closing input resource")
	}

	if brokenPipe {
		// close without flushing to output writer
		errorWriter := outputWriter.Close()
		if errorWriter != nil {
			errorWriter = errors.Wrap(errorWriter, "error closing output resource")
		}
		return errorReader, errorWriter
	}

	errorFlusher := outputWriter.Flush()
	if errorFlusher != nil {
		errorFlusher = errors.Wrap(errorFlusher, "error flushing output resource")
	}

	errorWriter := outputWriter.Close()
	if errorWriter != nil {
		errorWriter = errors.Wrap(errorWriter, "error closing output resource")
	}

	if errorFlusher != nil && errorWriter == nil {
		return errorReader, errorFlusher
	}

	return errorReader, errorWriter
}
