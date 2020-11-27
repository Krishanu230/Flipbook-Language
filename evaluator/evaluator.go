package evaluator

import (
	"Flipbook/ast"
	"Flipbook/object"
	"fmt"
)

var (
	NULL = &object.Null{}
)

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
	} //end of switch
	return nil
}

func evalStatements(sts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range sts {
		result = Eval(statement, env)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if ok {
		println(val.Inspect())
		return val
	}
	return newError("identifier not found: " + node.Value)
}

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

	for i := pgs; i <= pge; i++ {
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
