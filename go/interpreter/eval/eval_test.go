package eval

import (
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"2.3", 2.3},
		{"0.1", 0.1},
		{".4", 0.4},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestEvalNullExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"fn (a) { a; }"},
		{"ahoi"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNullObject(t, evaluated)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestMinusOperator(t *testing.T) {
	t.Run("on Integer", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int64
		}{
			{"-1", -1},
			{"--1", 1},
		}
		for _, tt := range tests {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		}
	})
	t.Run("on Float", func(t *testing.T) {
		tests := []struct {
			input    string
			expected float64
		}{
			{"-.2", -0.2},
			{"--3.34", 3.34},
		}
		for _, tt := range tests {
			evaluated := testEval(tt.input)
			testFloatObject(t, evaluated, tt.expected)
		}
	})
	t.Run("on object", func(t *testing.T) {
		tests := []struct {
			input string
		}{
			{"-test"},
			{"--hello"},
		}
		for _, tt := range tests {
			evaluated := testEval(tt.input)
			testNullObject(t, evaluated)
		}
	})
}

// HELPER

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f",
			result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	_, ok := obj.(*object.Null)
	if !ok {
		t.Errorf("object is not Null. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
