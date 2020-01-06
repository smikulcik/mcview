package view

import (
	"fmt"
	"image"

	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
)

// RegionToImage convert a mc region to a
func RegionToImage(r *region.Region) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))

	x, z := 0, 0
	// for x := 0; x < 32; x++ {
	// for z := 0; z < 32; z++ {
	data, err := r.ReadSector(x, z)
	if err != nil {
		return nil, err
	}
	var col save.Column
	err = col.Load(data)
	if err != nil {
		return nil, err
	}
	for k, v := range col.Level.Heightmaps {
		fmt.Printf("Key: %s Len: %d\n", k, len(v))
	}
	fmt.Printf("Len heightmaps: %d\n", len(col.Level.Heightmaps))
	// 	}
	// }

	fmt.Printf("Len col sections: %d\n", len(col.Level.Sections))
	for _, cnk := range col.Level.Sections {
		fmt.Printf("Len col sections states: %d y: %d len pallet: %d\n", len(cnk.BlockStates), cnk.Y, len(cnk.Palette))
		for _, blk := range cnk.Palette {
			fmt.Printf()
		}
	}

	return img, nil
}
