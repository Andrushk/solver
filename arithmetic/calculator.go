package arithmetic

import (
	"github.com/andrushk/solver/string-browser"
	"strconv"
	"fmt"
	"github.com/andrushk/solver/common"
	"log"
)

type Calculator struct  {
	//вычисляемое выражение
	expr string

	//навигация по строке-выражению
	b string_browser.Browser
}

func (c *Calculator) init(expression string){
	c.b = string_browser.Browser{}
	c.b.Init(expression)
}

func (c *Calculator) Do(expression string) (interface{}, error) {
	c.init(expression)

	result, err := c.getNumber()
	if err != nil {
		return nil, err
	}

	for !c.b.IsEof() && IsSign(c.b.GetCurrentSymbol()){
		p2, p2err :=c.getNumber()

		if p2err != nil {
			return nil, p2err
		}

		switch c.b.GetCurrentSymbol() {
		case 50:
			log.Println("do+")

			result = Add(result, p2)
		case '-':
			log.Println("do-")

			result = Sub(result, p2)
		}
	}

	return result, nil
}

// Выделение из строки подстроки, соответствующей
// определению вещественного числа.
// Синтаксис вещественного числа
// <Number> ::= [<Sign>] <Digit> {<Digit>} [<Separator> <Digit> {<Digit>}] [<Exponent> [<Sign>] <Digit> {<Digit>}]
// <Digit> ::= '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
// <Sign> ::= '+' | '-'
// <Separator> ::= '.'
// <Exponent> ::= 'E' | 'e'
func (c *Calculator) getNumber() (interface{}, error) {
	if c.b.IsEof() {
		return nil, fmt.Errorf(common.ErrorInPosition, c.b.GetCurrentIdx(), common.EmptyStringInNotNumber)
	}

	idxF, idxL := c.b.MoveWhile(IsDigit)

	//var str string
	//for{
	//	if !IsDigit(c.lastSymbol){
	//		panic("ожидалась цифра")
	//	}
	//}

	f, err := strconv.Atoi(c.b.Get(idxF, idxL))//strconv.ParseFloat(c.b.Get(idxF, idxL), 64)
	if (err != nil) {
		return nil, fmt.Errorf(common.ErrorInPosition, c.b.GetCurrentIdx(), err)
	}
	return f, nil
}