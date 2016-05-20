package arithmetic

import (
	"testing"
	"github.com/andrushk/solver/common"
	"log"
)

func TestCalculateSimpleSum(t *testing.T) {
	const correctValue int = 3
	c := Calculator{}
	answer, err := c.Do("1+2")
	if err != nil {
		t.Fatal(err)
	}

	if v, ok := answer.(int); !ok || v != correctValue {
		log.Println("ok = ", ok)
		t.Fatalf(common.WrongResponse, correctValue, answer)
	}
}

func TestGetNumber(t *testing.T){
	testGetNumber(t, "125", 125)
	testGetNumber(t, "125+88", 125)
	testGetNumber(t, "-125", -125)
}

func testGetNumber(t *testing.T, src string, expected interface{}){
	c := Calculator{}
	c.init(src)
	result, err := c.getNumber()
	if err != nil {
		t.Fatalf(common.ParsingError, src, err)
	}

	if expected != result{
		t.Fatalf(common.WrongResponse, expected, result)
	}
}