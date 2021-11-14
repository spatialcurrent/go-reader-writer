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

func ExampleSplitUserInfo_none() {
	user, password, err := SplitUserInfo("")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user=%q password=%q\n", user, password)
	// Output: user="" password=""
}

func ExampleSplitUserInfo_user() {
	user, password, err := SplitUserInfo("foo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user=%q password=%q\n", user, password)
	// Output: user="foo" password=""
}

func ExampleSplitUserInfo_userColon() {
	user, password, err := SplitUserInfo("foo:")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user=%q password=%q\n", user, password)
	// Output: user="foo" password=""
}

func ExampleSplitUserInfo_userPassword() {
	user, password, err := SplitUserInfo("foo:bar")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user=%q password=%q\n", user, password)
	// Output: user="foo" password="bar"
}
