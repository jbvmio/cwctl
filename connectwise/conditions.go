package connectwise

import (
	"fmt"
	"strings"
)

type Conditional string

type Conditionals struct {
	ands []Conditional
	ors  []Conditional
}

func (C *Conditionals) String() string {
	var condition, andStr, orStr string
	switch len(C.ands) {
	case 0:
	case 1:
		andStr = string(C.ands[0])
	default:
		andStr = string(C.ands[0])
		for _, c := range C.ands[1:] {
			andStr += string(AND) + string(c)
		}
	}
	switch len(C.ors) {
	case 0:
	case 1:
		orStr = string(C.ors[0])
	default:
		orStr = string(C.ors[0])
		for _, c := range C.ors[1:] {
			orStr += string(OR) + string(c)
		}
	}
	switch {
	case andStr != "" && orStr != "":
		condition = `(` + andStr + `) and (` + orStr + `)`
	default:
		condition = andStr + orStr
	}
	return condition
}

func (C *Conditionals) AndIN(key string, values ...interface{}) *Conditionals {
	return C.AND(BuildConditional(ConOpIN, NONE, key, values...))
}

func (C *Conditionals) OrIN(key string, values ...interface{}) *Conditionals {
	return C.OR(BuildConditional(ConOpIN, NONE, key, values...))
}

func (C *Conditionals) AndContains(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.AND(BuildConditional(ConOpCONTAINS, logic, key, values...))
}

func (C *Conditionals) OrContains(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.OR(BuildConditional(ConOpCONTAINS, logic, key, values...))
}

func (C *Conditionals) AndEquals(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.AND(BuildConditional(ConOpEQUAL, logic, key, values...))
}

func (C *Conditionals) OrEquals(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.OR(BuildConditional(ConOpEQUAL, logic, key, values...))
}

func (C *Conditionals) AndLike(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.AND(BuildConditional(ConOpLIKE, logic, key, values...))
}

func (C *Conditionals) OrLike(logic LogicOp, key string, values ...interface{}) *Conditionals {
	return C.OR(BuildConditional(ConOpLIKE, logic, key, values...))
}

func (C *Conditionals) AND(c Conditional) *Conditionals {
	if c != "" {
		C.ands = append(C.ands, c)
	}
	return C
}

func (C *Conditionals) OR(c Conditional) *Conditionals {
	if c != "" {
		C.ors = append(C.ors, c)
	}
	return C
}

type LogicOp string

const (
	NONE LogicOp = `, `
	AND  LogicOp = ` and `
	OR   LogicOp = ` or `
)

// ConOp IDs Condition Operators.
type ConOp int

const (
	ConOpEQUAL ConOp = iota
	ConOpLIKE
	ConOpCONTAINS
	ConOpIN
)

var ConOpStrings = [...]string{
	`eq`,
	`like`,
	`contains`,
	`in`,
}

func (op ConOp) String() string {
	return ConOpStrings[op]
}

func BuildConditional(op ConOp, logic LogicOp, key string, values ...interface{}) Conditional {
	if len(values) < 1 {
		return ""
	}
	var cond string
	switch op {
	case ConOpIN:
		cond = `(` + key + ` ` + op.String() + `(`
		cond += buildStringList("", "", string(logic), values)
		cond += `))`
	case ConOpCONTAINS:
		cond = `(`
		cond += buildStringList(key+` `+op.String()+`(`, ")", string(logic), values)
		cond += `)`
	default:
		cond = `(`
		cond += buildStringList(key+` `+op.String()+` `, "", string(logic), values)
		cond += `)`
	}
	return Conditional(cond)
}

func buildStringList(prefix, suffix, separator string, values []interface{}) string {
	var strVal string
	switch len(values) {
	case 0:
	case 1:
		strVal = prefix + stringValue(values[0]) + suffix
	default:
		strVal = prefix + stringValue(values[0]) + suffix
		for _, v := range values[1:] {
			strVal += separator + prefix + stringValue(v) + suffix
		}
	}
	return strVal
}

func stringValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		switch strings.ToLower(v) {
		case `true`, `false`:
			return v
		default:
			return fmt.Sprintf("'%s'", v)
		}
	case bool:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("'%v'", v)
	}
}
