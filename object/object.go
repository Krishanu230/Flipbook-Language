package object

import (
	"fmt"
	"strconv"
)

type ObjectType string

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"
	IMAGE_OBJ = "IMAGE"
	BOOK_OBJ  = "BOOK"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Image struct {
	Filename string
	DimX     int
	DimY     int
}

func (e *Image) Type() ObjectType { return IMAGE_OBJ }
func (e *Image) Inspect() string {
	dx := strconv.Itoa(e.DimX)
	dy := strconv.Itoa(e.DimY)
	return "Image: " + e.Filename + " dim: " + dx + ", " + dy
}

type Book struct {
	DimX       int
	DimY       int
	CountPages int
	Pages      []PageProperty
}

func (e *Book) Type() ObjectType { return BOOK_OBJ }
func (e *Book) Inspect() string {
	return "Bookdim: " + strconv.Itoa(e.DimX) + ", " + strconv.Itoa(e.DimY) + " PagesCount: " + strconv.Itoa(e.CountPages)

}

//Debugging Helper function
func (e *Book) PrintPagesMetadata() {
	for i, page := range e.Pages {
		fmt.Printf("\nPage no %d", i)
		for _, img := range page.ImagesProps {
			fmt.Printf("Image Name: %s, Scale: %d, PosX: %d, PosY: %d", img.Image.Filename, img.Scale, img.PosX, img.PosY)
		}
	}
}

//returns a blank new book
func NewBook(dimX int, dimY int, pCount string) *Book {
	pcnt, _ := strconv.Atoi(pCount)
	pages := []PageProperty{}
	i := 0
	for i < pcnt {
		blankPage := []ImageProperty{}
		pages = append(pages, PageProperty{ImagesProps: blankPage})
		i += 1
	}
	b := &Book{
		DimX:       dimX,
		DimY:       dimY,
		CountPages: pcnt,
		Pages:      pages,
	}
	return b
}

//every page is just a collection of images it contains.
type PageProperty struct {
	ImagesProps []ImageProperty
}

//every image obj in book's pages is a wrapper around image type
type ImageProperty struct {
	Image Image
	Scale int
	PosX  int
	PosY  int
}
