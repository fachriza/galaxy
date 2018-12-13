package trading

import "testing"

func TestGetSymbol(t *testing.T) {
	text := "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?"
	if GetSymbol(text) == "" {
		return
	}
}

func TestGetValue(t *testing.T) {
	symbol := "XL"
	if GetValue(symbol) == 40 {
		return
	}
}

func TestZeroValue(t *testing.T) {
	symbol := "A"
	if GetValue(symbol) == 0 {
		return
	}
}
