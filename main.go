package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"runtime/pprof"

	"github.com/alberthier/pathtracer"
)

func main() {
	width := flag.Int("width", 400, "Rendered image width")
	height := flag.Int("height", 200, "Rendered image height")
	samples := flag.Int("samples", 100, "Samples per pixel")
	startframe := flag.Int("startframe", 1, "Animation start frame")
	length := flag.Int("length", 1, "Animation length (frames)")
	cpuprofile := flag.String("cpuprofile", "", "CPU profile file")
	prefix := flag.String("prefix", "", "Output file prefix")

	flag.Parse()

	worldFile := flag.Arg(0)

	if len(worldFile) == 0 {
		fmt.Println("No world file provided")
		os.Exit(1)
	}

	world := pathtracer.NewWorld()
	err := world.Load(worldFile, float64(*width)/float64(*height))
	if err != nil {
		fmt.Println(err)
	}

	if len(*cpuprofile) != 0 {
		f, _ := os.Create("cpuprofile")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	renderer := pathtracer.NewRenderer(*width, *height, *samples)
	for t := *startframe; t < (*startframe + *length); t++ {
		logprefix := fmt.Sprintf("Frame %d/%d - ", t+1, *length)
		img := renderer.Render(world, float64(t), logprefix)
		output, _ := os.Create(fmt.Sprintf("%s%03d.png", *prefix, t))
		png.Encode(output, img)
		output.Close()
	}
}
