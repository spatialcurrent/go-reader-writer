// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grw

import (
	"bytes"
	"github.com/pkg/errors"
	"testing"
)

type testCase struct {
	Uri       string
	Algorithm string
	Content   []byte
}

func TestReadFromResource(t *testing.T) {

	testCases := []testCase{
		testCase{Uri: "../test/doc.txt", Algorithm: "", Content: []byte("hello world")},
		testCase{Uri: "../test/doc.txt.bz2", Algorithm: "bzip2", Content: []byte("hello world")},
		testCase{Uri: "../test/doc.txt.gz", Algorithm: "gzip", Content: []byte("hello world")},
		testCase{Uri: "../test/doc.txt.sz", Algorithm: "snappy", Content: []byte("hello world")},
		testCase{Uri: "../test/doc.txt.zip", Algorithm: "zip", Content: []byte("hello world")},
	}

	for _, testCase := range testCases {

		brc, _, err := ReadFromResource(testCase.Uri, testCase.Algorithm, 4096, false, nil)
		if err != nil {
			t.Errorf(errors.Wrap(err, "cannot opening resource for testing at uri "+testCase.Uri).Error())
		}
		got, err := brc.ReadAllAndClose()
		if err != nil {
			t.Errorf(errors.Wrap(err, "cannot reading resource for testing at uri"+testCase.Uri).Error())
		}
		if !bytes.Equal(got, testCase.Content) {
			t.Errorf("content does match for uri %v, got %v, want %v", testCase.Uri, string(got), string(testCase.Content))
		}
	}

}
