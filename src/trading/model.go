package trading

type Symbol struct {
	Text   string
	Symbol string
}

type Value struct {
	Symbol string
	Value  int
}

func (Value) DataValue() []Value {
	return []Value{
		Value{Symbol: "M", Value: 1000},
		Value{Symbol: "CM", Value: 900},
		Value{Symbol: "D", Value: 500},
		Value{Symbol: "CD", Value: 400},
		Value{Symbol: "C", Value: 100},
		Value{Symbol: "XC", Value: 90},
		Value{Symbol: "L", Value: 50},
		Value{Symbol: "XL", Value: 40},
		Value{Symbol: "X", Value: 10},
		Value{Symbol: "IX", Value: 9},
		Value{Symbol: "V", Value: 5},
		Value{Symbol: "IV", Value: 4},
		Value{Symbol: "I", Value: 1},
	}
}

func (Symbol) DataSymbol() []Symbol {
	return []Symbol{
		Symbol{Text: "tegj", Symbol: "L"},
		Symbol{Text: "pishtegj", Symbol: "XL"},
		Symbol{Text: "pish", Symbol: "X"},
		Symbol{Text: "globpish", Symbol: "IX"},
		Symbol{Text: "prok", Symbol: "V"},
		Symbol{Text: "globprok", Symbol: "IV"},
		Symbol{Text: "glob", Symbol: "I"},
	}
}
