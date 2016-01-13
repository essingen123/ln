package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/fogleman/ln/ln"
)

func cube(x, y, z float64) ln.Shape {
	a := ln.Vector{x - 0.5, y - 0.5, z - 0.5}
	b := ln.Vector{x + 0.5, y + 0.5, z + 0.5}
	return ln.NewCube(a, b)
}

func load(path string) []ln.Vector {
	var result []ln.Vector
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	for _, record := range records {
		x, _ := strconv.ParseFloat(record[0], 64)
		y, _ := strconv.ParseFloat(record[1], 64)
		z, _ := strconv.ParseFloat(record[2], 64)
		result = append(result, ln.Vector{x, z, y})
	}
	return result
}

func main() {
	blocks := load("examples/mountain.csv")
	fmt.Println(len(blocks))
	scene := ln.Scene{}
	size := ln.Vector{0.5, 0.5, 0.5}
	for _, v := range blocks {
		scene.Add(ln.NewCube(v.Sub(size), v.Add(size)))
	}
	eye := ln.Vector{90, -90, 70}
	center := ln.Vector{0, 0, -15}
	up := ln.Vector{0, 0, 1}
	paths := scene.Render(eye, center, up, 50, 1, 0.1, 1000, 0.1)
	paths.Render("out.png", 1024)
	// paths.Print()
}
