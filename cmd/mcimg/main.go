package main

import (
	"image/png"
	"os"

	"github.com/smikulcik/mcview"

	"github.com/Tnze/go-mc/save/region"
)

func main() {
	infilename := "examples/r.0.0.mca"
	outfilename := "examples/r.0.0.0.png"
	r, err := region.Open(infilename)
	if err != nil {
		panic(err)
	}

	// convert the regionfile to an image
	img, err := view.RegionToImage(r)
	if err != nil {
		panic(err)
	}

	// write the image out
	outfile, err := os.Create(outfilename)
	if err != nil {
		panic(err)
	}
	err = png.Encode(outfile, img)
	if err != nil {
		panic(err)
	}
}
