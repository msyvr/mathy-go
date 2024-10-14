package main

import (
	"image"
	color "image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

type Config struct {
	xmin       int
	xmax       int
	ymin       int
	ymax       int
	rows       int
	columns    int
	iterations int
	scale      float64
}

func main() {
	const (
		xmin           = -2
		xmax           = 2
		ymin           = -2
		ymax           = 2
		rows       int = 1024
		iterations     = 5000
		scale          = 50
	)

	var columns = int(math.Floor(1024 * 1 / 1))

	cfg := Config{
		xmin,
		xmax,
		ymin,
		ymax,
		rows,
		columns,
		iterations,
		scale,
	}

	buffer := image.NewRGBA(image.Rect(0, 0, columns, rows))
	for point_y := range rows {
		y := ((float64(point_y) / float64(rows)) * (ymax - ymin)) + ymin
		for point_x := range columns {
			x := ((float64(point_x) / float64(columns)) * (xmax - xmin)) + xmin

			z := complex(x, y)
			buffer.Set(point_y, point_x, mandelbrot(cfg, point_x, point_y, z))

		}
	}

	f, err := os.Create("static/mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}

	if err = png.Encode(f, buffer); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

}

func mandelbrot(cfg Config, x int, y int, z complex128) color.RGBA {
	var value complex128
	for i := range cfg.iterations {
		value = value*value + z
		if cmplx.Abs(value) > 2.0 {
			return color.RGBA{
				uint8(x * 255.0 / cfg.rows),
				uint8(float64(i) * cfg.scale),
				uint8(y * 255.0 / cfg.columns),
				255}
		}
	}

	return color.RGBA{
		255 - uint8(y*255.0/(2.0*cfg.columns)),
		0,
		255 - uint8(x*255.0/(2.0*cfg.rows)),
		255}

}
