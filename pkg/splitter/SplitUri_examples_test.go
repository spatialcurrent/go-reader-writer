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

func ExampleSplitUri_none() {
	scheme, remainder := SplitUri("foo.bar")
	fmt.Println(fmt.Sprintf("scheme=%q remainder=%q", scheme, remainder))
	// Output: scheme="" remainder="foo.bar"
}

func ExampleSplitUri_fTP() {
	scheme, remainder := SplitUri("ftp://foo.bar")
	fmt.Println(fmt.Sprintf("scheme=%q remainder=%q", scheme, remainder))
	// Output: scheme="ftp" remainder="foo.bar"
}

func ExampleSplitUri_hTTPS() {
	scheme, remainder := SplitUri("https://foo.bar")
	fmt.Println(fmt.Sprintf("scheme=%q remainder=%q", scheme, remainder))
	// Output: scheme="https" remainder="foo.bar"
}

func ExampleSplitUri_file() {
	scheme, remainder := SplitUri("file:///path/to/file")
	fmt.Println(fmt.Sprintf("scheme=%q remainder=%q", scheme, remainder))
	// Output: scheme="file" remainder="/path/to/file"
}
