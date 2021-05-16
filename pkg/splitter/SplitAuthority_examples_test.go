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

func ExampleSplitAuthority_host() {
	userinfo, host, port := SplitAuthority("foo.bar")
	fmt.Printf("userinfo=%q host=%q port=%q\n", userinfo, host, port)
	// Output: userinfo="" host="foo.bar" port=""
}

func ExampleSplitAuthority_hostPort() {
	userinfo, host, port := SplitAuthority("foo.bar:80")
	fmt.Printf("userinfo=%q host=%q port=%q\n", userinfo, host, port)
	// Output: userinfo="" host="foo.bar" port="80"
}

func ExampleSplitAuthority_hostPortUser() {
	userinfo, host, port := SplitAuthority("joe@foo.bar:80")
	fmt.Printf("userinfo=%q host=%q port=%q\n", userinfo, host, port)
	// Output: userinfo="joe" host="foo.bar" port="80"
}

func ExampleSplitAuthority_hostPortUserPassword() {
	userinfo, host, port := SplitAuthority("joe:joey@foo.bar:80")
	fmt.Printf("userinfo=%q host=%q port=%q\n", userinfo, host, port)
	// Output: userinfo="joe:joey" host="foo.bar" port="80"
}
