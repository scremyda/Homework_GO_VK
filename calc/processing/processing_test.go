package processing

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	var calcTestFloat = []struct {
		inputLine string
		answer    float64
	}{
		{"2+3*(5-2)/(1+2)-6/3+8*(4-2)", 19},
		{"(((12-3)*(4+6))/((5*1)+((8/2)-1)))+(((7*3)-15)/(2+((9/3)*2)))", 12},
		{"((500-250)*(80+20))/((150+50)-(10*4))", 156.25},
		{"((25+15)*(40-10))/((100+200)-(75*2))", 8},
		{"((8000+7000)*(60-15))/((100+500)-(40*4))*11", 16875},
		{"((3000+1500)*(120-30))/((750+250)-(50*5))", 540},
		{"((600-300)*(90+30))/((200+100)-(20*3))", 150},
		{"(((((1000+500)*2-750)/3+200)*4-1000)/5+300)*6-1500", 3660},
		{"1.1*1.2*1.7+0.001+0.002+1.1", 3.347},
		{"(1+2)-3", 0},
		{"(1+2)*3", 9},
		{"(((((((((1+2)-3)*4)+5)-6)*7)+8)-9)*10)", -80},
		{"((((((((((10+10))))))))))", 20},
		{"5*5", 25},
		{"2*2+3*3", 13},
		{"(2+2)*(3+3)", 24},
		{"1+2+3+4+5", 15},
		{"2-1*2-1*1*(1)", -1},
		{"0+0*0-0", 0},
		{"0", 0},
		{"100.45", 100.45},
	}

	var calcTestErrors = []struct {
		inputLine string
		err       error
	}{
		{"5/0", errors.New("expression is not written correctly")},
		{"0/0+100*100", errors.New("expression is not written correctly")},
		{"(545*11848+104*948+0.1)/0", errors.New("expression is not written correctly")},
		{"ffffffffffffff", errors.New("invalid symbols are given")},
		{"(1+1", errors.New("expression is not written correctly")},
		{"77*0+20)", errors.New("expression is not written correctly")},
		{"11*A+200/(134+0.1)", errors.New("invalid symbols are given")},
		{"A+B*2/C", errors.New("invalid symbols are given")},
		{"7*Q", errors.New("invalid symbols are given")},
	}
	for _, test := range calcTestFloat {
		result, _ := Calc(test.inputLine)
		assert.Equal(t, test.answer, result)
	}
	for _, test := range calcTestErrors {
		_, err := Calc(test.inputLine)
		assert.Equal(t, test.err, err)
	}
}
