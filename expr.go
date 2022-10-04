package main

import "math/rand"

type Operator int

const (
	OpMul Operator = iota
	OpDiv
	OpAdd
	OpSub
)

var operatorMap = map[string]Operator{"+": OpAdd, "-": OpSub, "*": OpMul, "/": OpDiv}

func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

type Value struct {
	Dice          *Dice       `@@`
	Number        *float64    `|  @Int`
	Subexpression *Expression `| "(" @@ ")"`
}

type Dice struct {
	StandardDice *StandardDice `@@`
	SimpleDice   *SimpleDice   `| @@`
}

type StandardDice struct {
	Count *float64 `@Int`
	Value *float64 `"d" @Int`
}

type SimpleDice struct {
	Value *float64 `"d" @Int`
}

type OpFactor struct {
	Operator Operator `@("*" | "/")`
	Base     *Value   `@@`
}

type Term struct {
	Left  *Value      `@@`
	Right []*OpFactor `@@*`
}

type OpTerm struct {
	Operator Operator `@("+" | "-")`
	Term     *Term    `@@`
}

type Expression struct {
	Left  *Term     `@@`
	Right []*OpTerm `@@*`
}

// Evaluation

func (o Operator) Eval(l, r float64) float64 {
	switch o {
	case OpMul:
		return l * r
	case OpDiv:
		return l / r
	case OpAdd:
		return l + r
	case OpSub:
		return l - r
	}
	panic("unsupported operator")
}

func (v *Value) Eval() float64 {
	switch {
	case v.Number != nil:
		return *v.Number
	case v.Dice != nil:
		return v.Dice.Eval()
	default:
		return v.Subexpression.Eval()
	}
}

func (d *Dice) Eval() float64 {
	switch {
	case d.StandardDice != nil:
		return d.StandardDice.Eval()
	case d.SimpleDice != nil:
		return d.SimpleDice.Eval()
	}
	panic("unsupported dice")
}

func (d *StandardDice) Eval() float64 {
	result := 0
	for i := 1; i <= int(*d.Count); i++ {
		result += rand.Intn(int(*d.Value)) + 1
	}
	return float64(result)
}

func (d *SimpleDice) Eval() float64 {
	return float64(rand.Intn(int(*d.Value)) + 1)
}

func (t *Term) Eval() float64 {
	n := t.Left.Eval()
	for _, r := range t.Right {
		n = r.Operator.Eval(n, r.Base.Eval())
	}
	return n
}

func (e *Expression) Eval() float64 {
	l := e.Left.Eval()
	for _, r := range e.Right {
		l = r.Operator.Eval(l, r.Term.Eval())
	}
	return l
}
