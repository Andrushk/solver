package arithmetic

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	if !(IsDigit('0')&&IsDigit('1')&&IsDigit('2')&&IsDigit('3')&&IsDigit('4')&&IsDigit('5')&&
	IsDigit('6')&&IsDigit('7')&&IsDigit('8')&&IsDigit('9')) {
		t.Fatal("IsDigit пропустил цыфру")
	}

	if IsDigit('a')||IsDigit(' ')||IsDigit('+')||IsDigit('/') {
		t.Fatal("IsDigit посчитал не цыфру за цыфру")
	}
}

func TestIsSeparator(t *testing.T)  {
	if !IsSeparator('.'){
		t.Fatal("IsSeparator пропустил символ-разделитель")
	}

	if IsSeparator(',')||IsSeparator(':')||IsSeparator('+')||IsSeparator('/')||IsSeparator('-')||IsSeparator(' ') {
		t.Fatal("IsSeparator поситал разделителем символ не разделитель")
	}
}

func TestIsSign(t *testing.T)  {
	if !(IsSign('+')||IsSign('-')) {
		t.Fatal("IsSign пропустил символ-знак числа")
	}

	if IsSign(',')||IsSign(':')||IsSign('.')||IsSign('/')||IsSign('*')||IsSign(' ') {
		t.Fatal("IsSign поситал знаком символ не знак")
	}
}
