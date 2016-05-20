package string_browser

import (
	"testing"
	"unicode/utf8"
	"github.com/andrushk/solver/common"
)

const testString = "test string"

func TestInit(t *testing.T) {
	br := Browser{}
	br.Init(testString)

	if br.src != testString || br.currIndex != 0 || br.currSymbol != 0 || br.IsEof() {
		t.Fatal("Init failed")
	}

	if br.r == nil || br.r.Size() != int64(utf8.RuneCountInString(testString)) {
		t.Fatal("Reader not initialized")
	}
}

func TestFirstStep(t *testing.T) {
	br := Browser{}
	br.Init(testString)
	result := br.Next()

	if br.src != testString {
		t.Fatal("Неправильно инициализирована строка")
	}

	if !result{
		t.Fatal("На первом шаге получен признак 'конец строки'")
	}

	if br.currIndex != 1 {
		t.Fatalf("Ошибка извлечения первого символа, currIndex ожидался %v, фактически %v", 1,
			br.GetCurrentIdx())
	}

	if br.currSymbol != 't' {
		t.Fatalf("Ошибка извлечения первого символа, currSymbol ожидался '%v', фактически '%v'", "t",
			br.GetCurrentSymbol())
	}
}

func TestReadAll(t *testing.T) {
	br := Browser{}
	br.Init(testString)
	length := utf8.RuneCountInString(testString)

	var c int
	for br.Next() {
		c++
	}

	if c != length {
		t.Fatalf("Должно было быть сделано '%v' шагов, а фактически сделано '%v'", length, c)
	}

	if br.currSymbol != 'g' {
		t.Fatalf("После прохода по строке ожидался currSymbol = '%v', а фактически '%v'", "g",
			br.GetCurrentSymbol())
	}

	if br.currIndex != length {
		t.Fatalf("После прохода по строке ожидался currIndex = '%v', а фактически '%v'", length,
			br.GetCurrentIdx())
	}
}

func TestGet(t *testing.T) {
	br := Browser{}
	br.Init(testString)

	exp := testString[0:4]
	act := br.Get(0, 4)

	if act != exp {
		t.Fatalf("Ошибка извлечения подстроки, ожидалось '%v', фактически '%v'", exp, act)
	}
}

func TestMoveWhile(t *testing.T) {
	br := Browser{}
	br.Init(testString)

	//первый символ пропускаем
	br.Next()

	//двигаться пока символ <> 't'
	s, f := br.MoveWhile(func(symbol rune) bool {
		return symbol != 't'
	})

	if s != 1 {
		t.Fatalf(common.WrongResponse, 1, s)
	}

	if f != 3 {
		t.Fatalf(common.WrongResponse, 3, f)
	}
}