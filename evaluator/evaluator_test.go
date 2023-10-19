package evaluator

import (
	"testing"

	"github.com/Krishanu230/Flipbook-Language/ast"
	"github.com/Krishanu230/Flipbook-Language/object"
)

func TestSwirlEffect(t *testing.T) {
	pages := []object.PageProperty{
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
	}
	SwirlEffect(pages, 10)
	if pages[0].ImagesProps[0].Image.Rotation != 0 || pages[1].ImagesProps[0].Image.Rotation != 10 {
		t.Errorf("SwirlEffect did not correctly update the rotation of the images")
	}
}

func TestEvalSwirlEffect(t *testing.T) {
	env := object.NewEnvironment()
	env.Set("bookone", &object.Book{Pages: []object.PageProperty{
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
	}})
	node := &ast.SwirlEffectStatement{Book: &ast.Identifier{Value: "bookone"}, RotationRate: &ast.IntegerLiteral{Value: 10}}
	evalSwirlEffect(node, env)
	book, _ := env.Get("bookone")
	if book.(*object.Book).Pages[0].ImagesProps[0].Image.Rotation != 0 || book.(*object.Book).Pages[1].ImagesProps[0].Image.Rotation != 10 {
		t.Errorf("evalSwirlEffect did not correctly update the rotation of the images")
	}
}

func TestEvalSave(t *testing.T) {
	env := object.NewEnvironment()
	env.Set("bookone", &object.Book{Pages: []object.PageProperty{
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
		{ImagesProps: []object.ImageProperty{{Image: object.Image{Rotation: 0}}}},
	}})
	node := &ast.SaveStatement{Book: &ast.Identifier{Value: "bookone"}, OutputName: &ast.StringLiteral{Value: "out.pdf"}}
	evalSave(node, env)
	book, _ := env.Get("bookone")
	if book.(*object.Book).Pages[0].ImagesProps[0].Image.Rotation != 0 || book.(*object.Book).Pages[1].ImagesProps[0].Image.Rotation != 10 {
		t.Errorf("evalSave did not correctly update the rotation of the images")
	}
}
