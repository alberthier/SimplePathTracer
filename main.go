package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/alberthier/pathtracer"
)

func main() {
	flag.Parse()

	worldFile := flag.Arg(0)

	if len(worldFile) == 0 {
		fmt.Println("No world file provided")
		os.Exit(1)
	}

	width := 400
	height := 200
	samples := 100
	world := pathtracer.NewWorld()
	err := world.Load(worldFile, float64(width)/float64(height))
	if err != nil {
		fmt.Println(err)
	}

	renderer := pathtracer.NewRenderer(width, height, samples)
	img := renderer.Render(world)
	output, _ := os.Create("output.png")
	png.Encode(output, img)
	output.Close()
}
