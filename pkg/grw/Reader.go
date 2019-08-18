// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"io"
	"io/ioutil"
)

// Reader is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Reader implements the ByteReader interface by wrapping around a subordinate ByteReader.
type Reader struct {
	Reader io.Reader // the instance of ByteReader used for reading bytes
	Closer *Closer   // the underlying closers
}

// Read reads a maximum len(p) bytes from the reader and returns an error, if any.
func (r *Reader) Read(p []byte) (n int, err error) {

	if r.Reader != nil {
		return r.Reader.Read(p)
	}

	return 0, nil
}

// ReadByte returns a single byte from the underlying reader.
func (r *Reader) ReadByte() (byte, error) {

	if r.Reader != nil {
		if br, ok := r.Reader.(ByteReader); ok {
			return br.ReadByte()
		}
		return byte(0), &ErrFunctionNotImplemented{Function: "ReadBytes", Object: "Reader"}
	}

	return 0, nil
}

// ReadBytes returns all bytes up to an including the first occurrence of the delimiter "delim" and an error, if any.
func (r *Reader) ReadBytes(delim byte) ([]byte, error) {
	if r.Reader != nil {
		if br, ok := r.Reader.(ByteReader); ok {
			return br.ReadBytes(delim)
		}
		return make([]byte, 0), &ErrFunctionNotImplemented{Function: "ReadBytes", Object: "Reader"}
	}
	return make([]byte, 0), nil
}

// ReadString returns a string of all the bytes to an including the first occurrence of the delimiter "delim" and an error, if any.
func (r *Reader) ReadString(delim byte) (string, error) {
	b, err := r.ReadBytes(delim)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// ReadFirst is not implemented by Reader
func (r *Reader) ReadFirst() (byte, error) {
	return byte(0), &ErrFunctionNotImplemented{Function: "ReadFirst", Object: "Reader"}
}

func (r *Reader) ReadAt(p []byte, off int64) (n int, err error) {
	if r.Reader != nil {
		if ra, ok := r.Reader.(io.ReaderAt); ok {
			return ra.ReadAt(p, off)
		}
		return 0, &ErrFunctionNotImplemented{Function: "ReadAt", Object: "Reader"}
	}
	return 0, nil
}

// ReadRange is not implemented by Reader
func (r *Reader) ReadRange(start int, end int) ([]byte, error) {
	return make([]byte, 0), &ErrFunctionNotImplemented{Function: "ReadRange", Object: "Reader"}
}

// ReadAll is not implemented by Reader
func (r *Reader) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(r.Reader)
}

func (r *Reader) Close() error {
	if r.Closer != nil {
		return r.Closer.Close()
	}
	return nil
}

// ReadAllAndClose reads all the bytes from the underlying reader and attempts to close the reader, even if there was an error reading.
func (r *Reader) ReadAllAndClose() ([]byte, error) {
	b, err := ioutil.ReadAll(r.Reader)
	if err != nil {
		// Attempt to close even if there is an error while reading
		r.Close() // #nosec
		return b, err
	}
	err = r.Close()
	return b, err
}
