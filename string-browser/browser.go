package string_browser

import (
	"strings"
	"io"
)

type Browser struct {
	src        string

	currIndex  int
	currSymbol rune
	isEof      bool
	r          *strings.Reader
}

func (b *Browser) Init(source string)  {
	b.src = source
	b.currIndex = 0
	b.currSymbol = 0
	b.isEof = false
	b.r = strings.NewReader(b.src)
}

//true, если удалось сдвинуться на символ
//false, если достигли конца строки
func (b *Browser) Next() bool {
	if b.isEof {
		return false
	}

	symbol, _, err := b.r.ReadRune()

	if err == io.EOF {
		b.isEof = true
		return false
	} else if err != nil {
		panic(err)
	}

	b.currSymbol = symbol
	b.currIndex++
	return true
}

//true, если достигнут конец строки
func (b *Browser) IsEof() bool {
	return b.isEof
}

//текущая позиция в просматриваемой строке
func (b *Browser) GetCurrentIdx() int{
	return b.currIndex
}

//возвращает символ, стоящий в текущае позиции просматриваемой строки
func (b *Browser) GetCurrentSymbol() rune{
	return b.currSymbol
}

//двигаться по строке пока выполняется переданное условие,
//вернуть индекс позиции начала движения и конца
func (b *Browser) MoveWhile(fn func(symbol rune) bool) (int, int){
	idxFrom := b.currIndex
	for b.Next() {
		if !fn(b.currSymbol) {
			return idxFrom, b.currIndex - 1
		}
	}
	return idxFrom, b.currIndex
}

func (b *Browser) Get(from int, to int) string {
	return b.src[from:to]
}