// Package imaging does image generation based on various parameters
package imaging

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

const (
	height = 1000
	width  = 1000
)

// Generate generates a svg based on a name. See telegram for example pictures
func Generate(name string) string {
	r := rand.Intn(155)
	g := rand.Intn(155)
	b := rand.Intn(155)
	text := initials(name)
	fontSize := int(500 / math.Pow(float64(len(text)), 0.7))

	data := fmt.Sprintf(`<?xml version="1.0"?>
<svg width="1000" height="1000" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
	<circle cx="500" cy="500" r="500" style="fill:rgb(%d,%d,%d)" />
	<circle cx="500" cy="500" r="480" style="fill:rgb(%d,%d,%d)" />
	<text x="500" y="500" style="text-anchor:middle;font-size:%dpx;alignment-baseline:middle;font-family:Arial,Helvetica;fill:rgb(%d,%d,%d)">%s</text>
</svg>
`, r, g, b, r+100, g+100, b+100, fontSize, r, g, b, text)
	return data
}

func initials(name string) string {
	f := strings.Fields(name)
	val := ""
	for _, n := range f {
		val += strings.ToUpper(string(n[0]))
	}
	return val
}
