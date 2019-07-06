// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: cases/msg.gohtml

package cases

import (
	"github.com/sipin/gorazor/gorazor"
	"io"
	. "kp/models"
	"strings"
)

func Msg(u *User) string {
	var _b strings.Builder
	RenderMsg(&_b, u)
	return _b.String()
}

func RenderMsg(_buffer io.StringWriter, u *User) {

	getName := func(u *User) string {
		return "(" + u.Name + ")"
	}

	var username string
	if u.Email != "" {
		username = getName(u) + "(" + u.Email + ")"
	}

	_buffer.WriteString("\n<div class=\"welcome\">\n<h4>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(username))
	_buffer.WriteString("</h4>\n\n<div>")
	_buffer.WriteString((u.Intro))
	_buffer.WriteString("</div>\n</div>")

}
