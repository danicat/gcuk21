package main

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var ttfRoboto font.Face
var ttfRobotoLarge font.Face

func init() {
	f, err := os.Open("assets/Roboto-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(bytes)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	ttfRoboto, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	ttfRobotoLarge, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
