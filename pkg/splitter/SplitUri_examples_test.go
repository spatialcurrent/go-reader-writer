// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package splitter

import (
	"fmt"
)

func ExampleSplitURI_none() {
	scheme, remainder := SplitURI("foo.bar")
	fmt.Printf("scheme=%q remainder=%q\n", scheme, remainder)
	// Output: scheme="" remainder="foo.bar"
}

func ExampleSplitURI_fTP() {
	scheme, remainder := SplitURI("ftp://foo.bar")
	fmt.Printf("scheme=%q remainder=%q\n", scheme, remainder)
	// Output: scheme="ftp" remainder="foo.bar"
}

func ExampleSplitURI_hTTPS() {
	scheme, remainder := SplitURI("https://foo.bar")
	fmt.Printf("scheme=%q remainder=%q\n", scheme, remainder)
	// Output: scheme="https" remainder="foo.bar"
}

func ExampleSplitURI_file() {
	scheme, remainder := SplitURI("file:///path/to/file")
	fmt.Printf("scheme=%q remainder=%q\n", scheme, remainder)
	// Output: scheme="file" remainder="/path/to/file"
}
