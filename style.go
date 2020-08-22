// SPDX-License-Identifier: Apache-2.0

package main

import (
	"syscall/js"
)

const css = `
body * {
  font-family: sans-serif;
  font-size: 14px;
  line-height: 24px;

  padding: 0;
  margin: 0;
}
input {
  background-color: transparent;
  border-style: none;
  outline: none;
}
`

func init() {
	style := js.Global().Get("document").Call("createElement", "style")
	style.Set("textContent", css)

	head := js.Global().Get("document").Get("head")
	head.Call("appendChild", style)
}
