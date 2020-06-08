package tsim

func Eval(node Node, env *Enviroment) Object {
	switch node.(type) {
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
	case CorpLiteral:
		return NewCorp()
	}
	return nil
}
