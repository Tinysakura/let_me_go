package font

import (
	"bytes"
	"strconv"
)

var defaultFamily = "default"

type Font struct {
	family string
	size int
}

func New(family string, size int) *Font {
	return &Font{validFamily(family), size}
}

// 验证family字段
func validFamily(family string) string {
	if family == "" {
		return defaultFamily
	}

	if len(family) < 5 || len(family) > 144 {
		return defaultFamily
	}

	return family
}

func (font *Font) SetFamily(family string) {
	font.family = validFamily(family)
}

func (font *Font) SetSize(size int) {
	font.size = size
}

func (font *Font) String() string {
	buffer := bytes.NewBufferString("{")
	buffer.WriteString("front-family:")
	buffer.WriteString("\"")
	buffer.WriteString(font.family)
	buffer.WriteString("\"")
	buffer.WriteString(";")
	buffer.WriteString("font-size:")
	buffer.WriteString("\"")
	buffer.WriteString(strconv.Itoa(font.size))
	buffer.WriteString("\"")
	buffer.WriteString(";")

	return buffer.String()
}