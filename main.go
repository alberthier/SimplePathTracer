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

	/*
		f, _ := os.Create("cpuprofile")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	*/

	renderer := pathtracer.NewRenderer(width, height, samples)
	for t := 0; t < 100; t++ {
		img := renderer.Render(world, float64(t))
		output, _ := os.Create(fmt.Sprintf("output%03d.png", t))
		png.Encode(output, img)
		output.Close()
	}
}
