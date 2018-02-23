package a // import "github.com/zevdg/vgo-a"

import b "github.com/zevdg/vgo-b"

func A() string {
	return "A1" + b.B()
}
