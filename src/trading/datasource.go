package trading

import "strings"
import "strconv"

/**
* Get Symbol of text
* @param string text
* @return string symbol
**/
func GetSymbol(text string) (symbol string) {
	var (
		symbolStruct Symbol
	)

	//split text by whitespace and if text is equal with data symbol
	//combine all text
	for _, splitText := range strings.Split(text, " ") {
		for _, data := range symbolStruct.DataSymbol() {
			if splitText == data.Text {
				symbol += data.Symbol
			}
		}
	}

	return symbol
}

/**
* Get Value of Romanian symbol
* @param string combination of symbol
* @return integer value
**/
func GetValue(text string) (value int) {
	var (
		valueStruct Value
	)

	//loop data value
	for _, data := range valueStruct.DataValue() {
		//if data symbol is from index 0 than
		if strings.Index(text, data.Symbol) == 0 {
			//set value and iterate GetValue until text and data symbol not equal from index 0
			value = data.Value + GetValue(strings.Replace(text, data.Symbol, "", 1))
			text = strconv.Itoa(value) //replace text with new data value
		}
	}

	return value
}
