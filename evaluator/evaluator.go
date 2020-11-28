package evaluator

import (
	"fmt"
	"strconv"

	"github.com/Krishanu230/Flipbook-Language/ast"
	"github.com/Krishanu230/Flipbook-Language/object"

	"github.com/signintech/gopdf"
)

var (
	NULL = &object.Null{}
)

//eval a ast node. recursive nature
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.NewStatement:
		val := evalNew(node, env)
		if isError(val) {
			return val
		}

	case *ast.InsertStatement:
		val := evalInsert(node, env)
		if isError(val) {
			return val
		}

	case *ast.KeyframeStatement:
		val := evalKeyframe(node, env)
		if isError(val) {
			return val
		}

	case *ast.SaveStatement:
		val := evalSave(node, env)
		if isError(val) {
			return val
		}
	} //end of switch
	return nil
}

//eval a bunch of statements by passing them to Eval()
func evalStatements(sts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range sts {
		result = Eval(statement, env)
	}

	return result
}

//resolve an identifier
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if ok {
		return val
	}
	return newError("identifier not found: " + node.Value)
}

//evaluate New type of statements
func evalNew(inp *ast.NewStatement, env *object.Environment) object.Object {
	st := *inp
	if st.DType.Value == "image" {
		imgObj := object.Image{
			Filename: st.Attribute.Value,
			DimX:     st.DimX.Value,
			DimY:     st.DimY.Value,
		}
		env.Set(st.Name.Value, &imgObj)
	}
	if st.DType.Value == "book" {
		b := object.NewBook(st.DimX.Value, st.DimY.Value, st.Attribute.Value)
		env.Set(st.Name.Value, b)
	}
	return &object.Null{}
}

//evaluate Insert statements
func evalInsert(inp *ast.InsertStatement, env *object.Environment) object.Object {
	//env.Print()
	r := evalIdentifier(inp.Image, env)
	img, ok := r.(*object.Image)
	if !ok {
		return r //if not okay the img contains error object
	}
	r2 := evalIdentifier(inp.Book, env)
	book, ok := r2.(*object.Book)
	if !ok {
		return r2 //if not okay the img contains error object
	}

	pgs := inp.StartPage.Value
	pge := inp.EndPage.Value
	posx := inp.PositionX.Value
	posy := inp.PositionY.Value

	//for all pages in range
	for i := pgs - 1; i <= pge-1; i++ {
		//add the image to the image array of the page
		imgprop := object.ImageProperty{
			Image: *img,
			Scale: 100,
			PosX:  posx,
			PosY:  posy,
		}
		orgPage := book.Pages[i].ImagesProps
		modPage := append(orgPage, imgprop)
		book.Pages[i] = object.PageProperty{ImagesProps: modPage}
	}
	//book.PrintPagesMetadata()
	return &object.Null{}
}

//evaluate keyframe statements
func evalKeyframe(inp *ast.KeyframeStatement, env *object.Environment) object.Object {
	r := evalIdentifier(inp.Image, env)
	img, ok := r.(*object.Image)
	if !ok {
		return r //if not okay the img contains error object
	}
	r2 := evalIdentifier(inp.Book, env)
	book, ok := r2.(*object.Book)
	if !ok {
		return r2 //if not okay the img contains error object
	}

	prop := inp.Property.Value
	pgs := inp.StartPage.Value
	pge := inp.EndPage.Value
	sProp := inp.StartProperty.Value
	eProp := inp.EndProperty.Value

	//get the linear rate of change of property
	delProp := eProp - sProp
	delPgCount := pge - pgs + 1
	slope := int(delProp / delPgCount)

	//for all pages in range
	for i := pgs - 1; i <= pge-1; i++ {
		//check all the images in the image array of a page
		for imgNo, imgitr := range book.Pages[i].ImagesProps {
			//if the image exists on the page update the property
			if imgitr.Image.Filename == img.Filename {
				switch prop {
				case "scale":
					orgImg := imgitr
					orgImg.Scale = sProp + slope*(i)
					book.Pages[i].ImagesProps[imgNo] = orgImg
				case "positionX":
					orgImg := imgitr
					orgImg.PosX = sProp + slope*(i)
					book.Pages[i].ImagesProps[imgNo] = orgImg
				case "positionY":
					orgImg := imgitr
					orgImg.PosY = sProp + slope*(i)
					book.Pages[i].ImagesProps[imgNo] = orgImg
				}
			}
		}
	}

	return &object.Null{}
}

func evalSave(inp *ast.SaveStatement, env *object.Environment) object.Object {
	r := evalIdentifier(inp.Book, env)
	book, ok := r.(*object.Book)
	if !ok {
		return r
	}
	fname := inp.OutputName.Value
	pw := book.DimX
	ph := book.DimY

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: float64(pw), H: float64(ph)}})

	for pno, page := range book.Pages {
		pdf.AddPage()
		for _, img := range page.ImagesProps {
			imgpath := img.Image.Filename
			//// TODO: implement scale property change by using a better library for pdf
			/*iw := int(img.Image.DimX * (0))
			ih := int(img.Image.DimX * (0))

			if (iw > pw) || (ih > ph) {
				println("ERROR1")
				return newError("Image " + fname + " size larger than the page " + strconv.Itoa(pno))
			}*/
			ix := img.PosX
			iy := img.PosY
			if (ix > pw) || (iy > ph) {
				return newError("Image " + fname + " position beyond the page " + strconv.Itoa(pno))
			}
			pdf.Image(imgpath, float64(ix), float64(iy), nil)
		}

	}

	pdf.WritePdf(fname)
	return &object.Null{}
}

//Helper Functions
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
