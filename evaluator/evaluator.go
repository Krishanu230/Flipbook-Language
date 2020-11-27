package evaluator

import (
	"Flipbook/ast"
	"Flipbook/object"
	"fmt"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.NewStatement:
		val := evalNew(node)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	} //end of switch
	return nil
}

func evalStatements(sts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range sts {
		result = Eval(statement)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	return newError("identifier not found: " + node.Value)
}

func evalNew(inp ast.Node) object.Object {
	fmt.Printf("gotta eval a new statement")
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
