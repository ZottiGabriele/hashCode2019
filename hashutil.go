package main

import (
	"fmt"
	"os"
)

type SlideShow struct {
	n_of_slides int
	slides      []Slide
}

type Slide struct {
	photos []Photo
	tags   []string
}

type Photo struct {
	ID        int
	orient    byte
	n_of_tags int
	tags      []string
}

type PhotoCollection struct {
	n_of_photos int
	photos      []Photo
}

func readInput(filePath string) PhotoCollection {
	file, err := os.Open(filePath)
	check(err)

	N := 0
	_, err = fmt.Fscanf(file, "%d", &N)
	check(err)
	out := PhotoCollection{n_of_photos: N}
	out.photos = make([]Photo, N)

	for i := 0; i < N; i++ {
		current := Photo{ID: i}

		_, err = fmt.Fscanf(file, "%c", &current.orient)
		check(err)

		_, err = fmt.Fscanf(file, "%d", &current.n_of_tags)
		check(err)

		for j := 0; j < current.n_of_tags; j++ {
			currentTag := ""
			_, err = fmt.Fscanf(file, "%s", &currentTag)
			check(err)
			current.tags = append(current.tags, currentTag)
		}
		out.photos[i] = current
	}

	return out
}

func writeOutput(slideShow SlideShow, outPath string) {
	out, err := os.Create(outPath)
	check(err)
	out.WriteString(fmt.Sprintf("%d\n", slideShow.n_of_slides))
	for _, slide := range slideShow.slides {
		for _, photo := range slide.photos {
			out.WriteString(fmt.Sprintf("%d ", photo.ID))
		}
		out.WriteString(fmt.Sprintf("\n"))
	}
	err = out.Close()
	check(err)
}

func toInt(b byte) int {
	return int(b) - int('0')
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
