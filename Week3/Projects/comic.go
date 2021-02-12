package main

import "fmt"

//This program prints data about ficticious comic books with variable data

func main() {
	var writer, artist, title string
	var year, pageNumber int
	var grade float32

	title, writer, artist = "Mr. GoToSleep", "Tracey Hatchet", "Jewel Tampson"

	year = 1997

	pageNumber = 14

	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "in", year, "( pages:", pageNumber, "| grade:", grade, ")")

	//New Values
	title, writer, artist, year, pageNumber, grade = "Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", 2013, 160, 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "in", year, "( pages:", pageNumber, "| grade:", grade, ")")

}
