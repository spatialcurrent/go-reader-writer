// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flate

import (
	"compress/flate"
	"fmt"
	"io"
)

type Writer struct {
	*flate.Writer
	underlying io.WriteCloser
}

type flusher interface {
	Flush() error
}

// Close closes the flate.Writer, and then flushes and closes the underlying WriteCloser.
func (w *Writer) Close() error {
	err := w.Writer.Close()
	if err != nil {
		return fmt.Errorf("error closing flate.Writer: %w", err)
	}
	// When the flate writer is closed is flushes one last time.
	// Therefore, we need to flush the underlying writer one last time before we close it.
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	err = w.underlying.Close()
	if err != nil {
		return fmt.Errorf("error closing underlying writer: %w", err)
	}
	return nil
}

// Flush flushes any pending data to the underlying writer.
// It is useful mainly in compressed network protocols, to ensure that
// a remote reader has enough data to reconstruct a packet.
// Flush does not return until the data has been written.
// Calling Flush when there is no pending data still causes the Writer
// to emit a sync marker of at least 4 bytes.
// If the underlying writer returns an error, Flush returns that error.
//
// In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
func (w *Writer) Flush() error {
	err := w.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing flate.Writer: %w", err)
	}
	if f, ok := w.underlying.(flusher); ok {
		err = f.Flush()
		if err != nil {
			return fmt.Errorf("error flushing underlying writer: %w", err)
		}
	}
	return nil
}

// NewWriter returns a new Writer compressing data at the default level.
func NewWriter(w io.WriteCloser) (*Writer, error) {
	fw, err := flate.NewWriter(w, DefaultCompression)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: fw, underlying: w}, nil
}

// NewWriterLevel returns a new Writer compressing data at the given level.
// Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
// higher levels typically run slower but compress more. Level 0
// (NoCompression) does not attempt any compression; it only adds the
// necessary DEFLATE framing.
// Level -1 (DefaultCompression) uses the default compression level.
// Level -2 (HuffmanOnly) will use Huffman compression only, giving
// a very fast compression for all types of input, but sacrificing considerable
// compression efficiency.
//
// If level is in the range [-2, 9] then the error returned will be nil.
// Otherwise the error returned will be non-nil.
func NewWriterLevel(w io.WriteCloser, level int) (*Writer, error) {
	fw, err := flate.NewWriter(w, level)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: fw, underlying: w}, nil
}

// NewWriterDict is like NewWriter but initializes the new
// Writer with a preset dictionary. The returned Writer behaves
// as if the dictionary had been written to it without producing
// any compressed output. The compressed data written to w
// can only be decompressed by a Reader initialized with the
// same dictionary.
func NewWriterDict(w io.WriteCloser, level int, dict []byte) (*Writer, error) {
	fw, err := flate.NewWriterDict(w, level, dict)
	if err != nil {
		return nil, err
	}
	return &Writer{Writer: fw, underlying: w}, nil
}
