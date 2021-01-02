// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cache

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

type Cache struct {
	Reader   io.ByteReadCloser
	Cursor   int
	Complete *bool
	Content  *[]byte
}

func NewCache(r io.ByteReadCloser) *Cache {
	complete := false
	content := make([]byte, 0)
	return &Cache{
		Reader:   r,
		Cursor:   0,
		Content:  &content,
		Complete: &complete,
	}
}

func NewCacheWithContent(r io.ByteReadCloser, c *[]byte, i int) *Cache {
	complete := false
	return &Cache{
		Reader:   r,
		Content:  c,
		Cursor:   i,
		Complete: &complete,
	}
}

func (c *Cache) ReadByte() (byte, error) {

	_, err := c.Reader.Read(make([]byte, 1))
	if err != nil {
		return (*c.Content)[len(*c.Content)-1], nil
	}

	return byte(0), nil

}

// Read reads a maximum len(p) bytes from the reader and returns an error, if any.
func (c *Cache) Read(p []byte) (n int, err error) {

	if *c.Complete {
		return 0, io.EOF
	}

	if len(p) == 0 {
		return 0, errors.New("error attempting to read zero bytes from file.")
	}

	n, err = c.Reader.Read(p)
	if n > 0 {
		*c.Content = append(*c.Content, p[0:n]...)
	}

	if err != nil || n == 0 {
		*c.Complete = true
	}

	return n, err
}

func (c *Cache) ReadBytes(delim byte) ([]byte, error) {

	if *c.Complete {
		return make([]byte, 0), io.EOF
	}

	b, err := c.Reader.ReadBytes(delim)
	if len(b) > 0 {
		*c.Content = append(*c.Content, b...)
	}

	if err != nil || len(b) == 0 {
		*c.Complete = true
	}

	return b, err
}

func (c *Cache) ReadString(delim byte) (string, error) {
	b, err := c.ReadBytes(delim)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// ReadFirst returns the first byte stating at the cursor.
func (c *Cache) ReadFirst() (byte, error) {
	//fmt.Println("ReadFirst()")

	if (c.Cursor) < len(*c.Content) {
		return (*c.Content)[c.Cursor], nil
	}

	if !*c.Complete {
		for c.Cursor >= len(*c.Content) {
			newContent := make([]byte, c.Cursor-len(*c.Content)+1)
			_, err := c.Read(newContent)
			//c.Content = append(c.Content, newContent[0:n]...)
			if err != nil {
				if err == io.EOF {
					//*c.Complete = true
					break
				} else {
					return byte(0), fmt.Errorf("error reading first byte starting at cursor %q: %w", fmt.Sprint(c.Cursor), err)
				}
			}
		}
	}

	return (*c.Content)[c.Cursor], nil
}

// ReadAt returns the byte at the index given by i
func (c *Cache) ReadAt(i int) (byte, error) {
	//fmt.Println("ReadAt("+fmt.Sprint(i)+")")

	if (c.Cursor + i) < len(*c.Content) {
		return (*c.Content)[c.Cursor+i], nil
	}

	if !*c.Complete {
		for (c.Cursor + i) >= len(*c.Content) {
			//newContent := make([]byte, c.Cursor + i - len(c.Content) + 1)
			newContent := make([]byte, c.Cursor+i-len(*c.Content)+1)
			//fmt.Println("Reading new Content: ", newContent)
			_, err := c.Read(newContent)
			//c.Content = append(9, newContent[0:n]...)
			if err != nil {
				if err == io.EOF {
					//*c.Complete = true
					return byte(0), err
				} else {
					return byte(0), fmt.Errorf("error reading bytes from index %q starting at cursor %q: %w", fmt.Sprint(i), fmt.Sprint(c.Cursor), err)
				}
			}
		}
	}

	if (c.Cursor + i) >= len(*c.Content) {
		return byte(0), io.EOF //errors.New("index "+fmt.Sprint(i)+" >= length of Content")
	}

	return (*c.Content)[c.Cursor+i], nil
}

// ReadAll reads all content from the underlying reader and returns the content
func (c *Cache) ReadAll() ([]byte, error) {
	//fmt.Println("ReadAll()")
	if !*c.Complete {
		newContent, err := ioutil.ReadAll(c.Reader)
		if err != nil {
			return newContent, fmt.Errorf("error reading all content from underlying reader: %w", err)
		}
		*c.Content = append(*c.Content, newContent...)
		*c.Complete = true
	}
	return *c.Content, nil
}

// ReadRange reads a range of btes from the cache.  End points to the index of the last byte to read, so [start:end+1]
func (c *Cache) ReadRange(start int, end int) ([]byte, error) {
	//fmt.Println("ReadRange("+fmt.Sprint(start)+","+fmt.Sprint(end)+") Content is",c.Content,"Complete is ", *c.Complete)
	if (c.Cursor + end) < len(*c.Content) {
		return (*c.Content)[c.Cursor+start : c.Cursor+end+1], nil
	}

	if !*c.Complete {
		for (c.Cursor + end + 1) > len(*c.Content) {
			//newContent := make([]byte, c.Cursor + end - len(c.Content) + 1)
			_, err := c.Read(make([]byte, c.Cursor+end-len(*c.Content)+1))
			//fmt.Println("Reading from range: n=", n, "err=", err, "complete=", *c.Complete)
			//c.Content = append(c.Content, newContent[0:n]...)
			if err != nil {
				if err == io.EOF {
					//fmt.Println("Setting complete to true in ReadRange.")
					//*c.Complete = true
					break
				} else {
					return make([]byte, 0), fmt.Errorf("error reading %d bytes for range %d-%d starting at cursor %d.  Current Content: %x: %w",
						c.Cursor+end-len(*c.Content)+1,
						start,
						end,
						c.Cursor,
						c.Content,
						err)
				}
			}
		}
	}

	if (c.Cursor + end + 1) > len(*c.Content) {
		return make([]byte, 0), errors.New(fmt.Sprintf(
			"Content is only %d bytes.  %x.  End is %d",
			len(*c.Content),
			fmt.Sprint(c.Content),
			c.Cursor+end+1))
	}

	//fmt.Println("End of ReadRange("+fmt.Sprint(start)+","+fmt.Sprint(end)+") has content ",c.Content)
	return (*c.Content)[c.Cursor+start : c.Cursor+end+1], nil
}

func (c *Cache) Close() error {
	return c.Reader.Close()
}

func (c *Cache) ReadAllAndClose() ([]byte, error) {
	b, err := c.Reader.ReadAllAndClose()

	if *c.Complete {
		return make([]byte, 0), io.EOF
	}

	if len(b) > 0 {
		*c.Content = append(*c.Content, b...)
	}

	if err != nil || len(b) == 0 {
		*c.Complete = true
	}

	return b, err
}
