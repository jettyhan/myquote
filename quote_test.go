// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package myquote

import "testing"

func TestHello(t *testing.T) {
	hello := "Hello, world."
	if out := Hello(); out != hello {
		t.Errorf("Hello() = %q, want %q", out, hello)
	}
}

func TestGlass(t *testing.T) {
	glass := "I can eat glass and it doesn't hurt me."
	if out := Glass(); out != glass {
		t.Errorf("Glass() = %q, want %q", out, glass)
	}
}
