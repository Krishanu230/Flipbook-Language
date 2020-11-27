package object

import "Flipbook/object"

type ObjectType string

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"
	IMAGE_OBJ = "IMAGE"
	BOOK_OBJ  = "BOOK"
)

var (
	NULL = &object.Null{}
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
	return "Image: " + e.Filename + " dim: " + string(e.DimX) + ", " + string(e.DimY)
}

type Book struct {
	Bookname   string
	DimX       int
	DimY       int
	CountPages int
	Pages      map[int]PageProperty
}

func (e *Book) Type() ObjectType { return BOOK_OBJ }
func (e *Book) Inspect() string {
	return "Book: " + e.Bookname + " dim: " + string(e.DimX) + ", " + string(e.DimY) + " PagesCount: " + string(e.CountPages)
}

type PageProperty struct {
	ImagesProps []ImageProperty
}

type ImageProperty struct {
	Image Image
	Scale int
	PosX  int
	PosY  int
}
