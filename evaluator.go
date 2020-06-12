package tsim

func Eval(node Node, env *Enviroment) Object {
	switch node := node.(type) {
	// Statements
	case *Program:
		return evalProgram(node, env)
	case *NewStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	// Expressions
	case *CorpLiteral:
		return NewCorp()
	}
	return &Error{Message:"out of switch"}
}

func evalProgram(program *Program,env *Enviroment) Object {
	var result Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		//switch result.(type) {
		//case *ReturnValue:
		//	return result.Value
		case *Error:
			return result
		}
	}
	return result
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}
