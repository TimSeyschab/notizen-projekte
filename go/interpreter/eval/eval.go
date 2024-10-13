package eval

import (
	"interpreter/ast"
	"interpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return getNativeBooleanObject(node.Value)

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	}

	return NULL
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case isNumber(left) && isNumber(right):
		return evalNumberInfixExpression(operator, left, right)
	case operator == "==":
		return getNativeBooleanObject(left == right)
	case operator == "!=":
		return getNativeBooleanObject(left != right)
	default:
		return NULL
	}
}

func isNumber(obj object.Object) bool {
	switch obj.Type() {
	case object.INTEGER_OBJ, object.FLOAT_OBJ:
		return true
	default:
		return false
	}
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal, isLeftFloat := getValueAndType(left)
	rightVal, isRightFloat := getValueAndType(right)
	if leftVal == nil || rightVal == nil {
		return NULL
	}

	switch operator {
	case "<":
		return getNativeBooleanObject(leftVal.(float64) < rightVal.(float64))
	case ">":
		return getNativeBooleanObject(leftVal.(float64) > rightVal.(float64))
	case "==":
		return getNativeBooleanObject(leftVal.(float64) == rightVal.(float64))
	case "!=":
		return getNativeBooleanObject(leftVal.(float64) != rightVal.(float64))
	}

	if isLeftFloat || isRightFloat {
		return evalFloatInfix(operator, leftVal.(float64), rightVal.(float64))
	}
	return evalIntInfix(operator, int64(leftVal.(float64)), int64(rightVal.(float64)))
}

func getValueAndType(obj object.Object) (interface{}, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return float64(v.Value), false
	case *object.Float:
		return v.Value, true
	default:
		return nil, false
	}
}

func evalFloatInfix(operator string, leftVal, rightVal float64) object.Object {
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	default:
		return NULL
	}
}

func evalIntInfix(operator string, leftVal, rightVal int64) object.Object {
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	default:
		return NULL
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Float:
		right.Value = right.Value * -1
		return right
	case *object.Integer:
		right.Value = right.Value * -1
		return right
	}

	return NULL
}

func getNativeBooleanObject(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
