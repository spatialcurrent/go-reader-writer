// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"fmt"
	"io"
	"io/ioutil"
)

import (
	"github.com/pkg/errors"
)

type Cache struct {
	*Reader
	Cursor   int
	Complete *bool
	Content  *[]byte
}

func NewCache(r *Reader) *Cache {
	complete := false
	content := make([]byte, 0)
	return &Cache{
		Reader:   r,
		Cursor:   0,
		Content:  &content,
		Complete: &complete,
	}
}

func NewCacheWithContent(r *Reader, c *[]byte, i int) *Cache {
	complete := false
	return &Cache{
		Reader:   r,
		Content:  c,
		Cursor:   i,
		Complete: &complete,
	}
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
		//fmt.Println("Setting complete to true in Read.")
		//fmt.Println("Content is ", c.Content)
		//panic(nil)
		*c.Complete = true
	}

	//if n == 0 {
	//	c.Complete = true
	//}

	//fmt.Println("Read("+fmt.Sprint(p)+") returning ",p[0:n],".  Content = ", c.Content)

	return n, err
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
					return byte(0), errors.Wrap(err, "error reading first byte starting at cursor "+fmt.Sprint(c.Cursor))
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
					return byte(0), errors.Wrap(err, "error reading bytes from index "+fmt.Sprint(i)+" starting at cursor "+fmt.Sprint(c.Cursor))
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
			return newContent, errors.Wrap(err, "error reading all content from underlying reader")
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
					return make([]byte, 0), errors.Wrap(err, "error reading "+fmt.Sprint(c.Cursor+end-len(*c.Content)+1)+" bytes for range "+fmt.Sprint(start)+"-"+fmt.Sprint(end)+" starting at cursor "+fmt.Sprint(c.Cursor)+".  Current Content:"+fmt.Sprint(c.Content))
				}
			}
		}
	}

	if (c.Cursor + end + 1) > len(*c.Content) {
		return make([]byte, 0), errors.New("Content is only " + fmt.Sprint(len(*c.Content)) + " bytes.  " + fmt.Sprint(c.Content) + ".  End is " + fmt.Sprint(c.Cursor+end+1))
	}

	//fmt.Println("End of ReadRange("+fmt.Sprint(start)+","+fmt.Sprint(end)+") has content ",c.Content)
	return (*c.Content)[c.Cursor+start : c.Cursor+end+1], nil
}
