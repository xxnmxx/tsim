package tsim

import "fmt"

func Eval(node Node, env *Enviroment) Object {
	switch node := node.(type) {
	// Statements
	case *Program:
		return evalProgram(node, env)
	case *NewStatement:
		//val := Eval(node.Value, env)
		//if isError(val) {
		//	return val
		//}
		env.Set(node.Name.Value, node.Value)
		return node.Value
	// Expressions
	case *CorpLiteral:
		return NewCorp()
	case *Identifier:
		return evalIdentifier(node, env)
	}
	return &Error{Message: "out of switch"}
}

// FixMe:( This must be wrong.
func evalProgram(program *Program, env *Enviroment) Object {
	var result Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		//switch result.(type) {
		//case *ReturnValue:
		//	return result.Value
		//case *Error:
		//	return result
		//}
	}
	return result
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func evalIdentifier(node *Identifier, env *Enviroment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError("identifier not found: " + node.Value)
}
