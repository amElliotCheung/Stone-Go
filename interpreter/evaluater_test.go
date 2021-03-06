package interpreter

import (
	"strings"
	"testing"
)

func TestEvaluater_Eval(t *testing.T) {
	input := `
	adder := func (x, y) {
		x+y
	}
	a := -1
	b := 1
	adder(a, b)
	c := "one"
	d := " two"
	adder(c, d)`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval2(t *testing.T) {
	input := `
	adder := func (x, y) {
		x+y
	}
	n := 1
	a := 0
	b := 1
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}
func TestEvaluater_Eval3(t *testing.T) {
	input := `a := true
	b := !a
	b = !!!!!b
	a = b == (1 == 1)
	a = b != (1 == 1)
	1 != 2
	1+1 == 2
	3 > 5
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval4(t *testing.T) {
	input := `a := [1, 2]
	while len(a) <= 10 {
		a = append(a, a[len(a)-1] + a[len(a)-2])
	}
	a[9]
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}
func TestEvaluater_Eval5(t *testing.T) {
	input := `
	fib := func (n) {
		a := [1, 2]
		while len(a) <= n {
			a = append(a,  a[len(a)-1] + a[len(a)-2])
		}
		a[n-1]
	}
	add := func(x, y) {
		x+y
	}
	fib(50)
	add("one", "+two")`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval6(t *testing.T) {
	input := `
	fib := func (n) {
		if n <= 2 {
			n
		} else {
			fib(n-1)+fib(n-2)
		}
	}
	fib(1)`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval7(t *testing.T) {
	input := `
	getFibArray := func(size) {
		fib := func (n) {n <= 2 ? n : fib(n-1) + fib(n-2)}
		n := 1
		fibArray := []
		while n <= size {
			fibArray = append(fibArray, fib(n))
			n = n + 1
		}
		fibArray
	}
	getFibArray(20)
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval8(t *testing.T) {
	input := `
	makeAdder := func (x) {
		func(y) {
			x+y
		}
	}
	fiveAdder := makeAdder(5)
	fiveAdder(35)
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}

func TestEvaluater_Eval9(t *testing.T) {
	input := `
	makeAdder := func (x) {
		func(y) {
			x+y
		}
	}
	fiveAdder := makeAdder(5)
	fiveAdder(35)
	`
	in := strings.NewReader(input)
	lexer := NewLexer(in)
	parser := NewParser(lexer)
	c := make(chan Statement)
	go parser.Parse(c)
	e := NewEvaluater(c)
	e.Eval()
}
