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

	world := pathtracer.NewWorld()
	err := world.Load(worldFile)
	if err != nil {
		fmt.Println(err)
	}

	renderer := pathtracer.NewRenderer(320, 200, 100)
	img := renderer.Render(world)
	output, _ := os.Create("output.png")
	png.Encode(output, img)
	output.Close()
}
