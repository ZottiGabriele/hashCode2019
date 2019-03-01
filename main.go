package main

import (
	"fmt"
	"os"
	"sort"
)

func swap(arr []Photo, a int, b int) {

	temp := arr[b]
	arr[b] = arr[a]
	arr[a] = temp
}

type ByTag []Photo

func (a ByTag) Len() int           { return len(a) }
func (a ByTag) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTag) Less(i, j int) bool { return a[i].n_of_tags < a[j].n_of_tags }

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("ERROR: Usage: arg1 arg2 (arg1 is input path, arg2 is output path")
		return
	}

	fmt.Println("Reading input...")
	collection := readInput(args[0])
	slideShow := SlideShow{}
	slides := []Slide{}
	var slidePhoto []Photo
	var count int

	for _, photo := range collection.photos {

		count = 0
		if photo.orient == 'H' {
			count++
			slides = append(slides, Slide{photos: []Photo{photo}})
			slidePhoto = append(slidePhoto, photo)
		}

	}

	sort.Sort(ByTag(slidePhoto))

	punteggio := 0
	var pos int

	fmt.Println("Doing magical stuff...")
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {

			slide1 := Slide{photos: []Photo{slidePhoto[i]}}
			slide2 := Slide{photos: []Photo{slidePhoto[j]}}

			temp := calcolaPunteggio(slide1, slide2)
			if punteggio < temp {
				punteggio = temp
				pos = j
			}
		}

		swap(slidePhoto, pos, i+1)
	}

	i := 0
	for _, photo := range slidePhoto {
		slideShow.slides = append(slideShow.slides, Slide{[]Photo{photo}, photo.tags})
		i++
	}

	slideShow.n_of_slides = i
	fmt.Println("Writing output at", args[1])
	writeOutput(slideShow, args[1])
	fmt.Println("Finished :D")
}

func calcolaPunteggio(photo1, photo2 Slide) int {

	aElem := photo1.photos[0].n_of_tags
	bElem := photo2.photos[0].n_of_tags
	intersElem := 0

	for i := 0; i < photo1.photos[0].n_of_tags; i++ {

		for j := i + 1; j < photo1.photos[0].n_of_tags; j++ {

			if photo1.tags[i] == photo1.tags[j] {
				aElem--
				bElem--
				intersElem++
			}
		}
	}

	var min int

	if aElem < bElem {
		min = aElem
	} else {
		min = bElem
	}

	if min > intersElem {

		min = intersElem
	}

	return min

}
