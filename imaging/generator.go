// Package imaging does image generation based on various parameters
package imaging

import (
	"fmt"
	"io"
	"math/rand"
	"strings"

	svg "github.com/ajstarks/svgo"
)

const (
	height = 1000
	width  = 1000
)

// Generate generates a svg based on a name. See telegram for example pictures
func Generate(name string, w io.Writer) *svg.SVG {
	canvas := svg.New(w)
	r := rand.Intn(155)
	g := rand.Intn(155)
	b := rand.Intn(155)
	color := canvas.RGB(r, g, b)
	lighterColor := canvas.RGB(r+100, g+100, b+100)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 500, color)
	canvas.Circle(width/2, height/2, 480, lighterColor)
	canvas.Text(width/2, height/2, initials(name), fmt.Sprintf("text-anchor:middle;font-size:250px;alignment-baseline:middle;font-family:Arial,Helvetica;%s", color))
	canvas.End()
	return canvas
}

func initials(name string) string {
	f := strings.Fields(name)
	val := ""
	for _, n := range f {
		val += strings.ToUpper(string(n[0]))
	}
	return val
}
