package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinesProcessing(t *testing.T) {
	var linesProcessingTests = []struct {
		inputOptions Options
		inputLine    []string
		out          []string
	}{
		{Options{CFlagStated: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			[]string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik.",
			},
		},
		{Options{DFlagStated: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			[]string{
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		{Options{UFlagStated: true},
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			[]string{
				"",
				"Thanks.",
			},
		},
		{Options{FFlagStated: 1},
			[]string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{Options{SFlagStated: 1},
			[]string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},
		{Options{IFlagStated: true},
			[]string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love music of Kartik.",
			},
			[]string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
				"I love music of kartik.",
			},
		},
		{Options{FFlagStated: 2, SFlagStated: 2, IFlagStated: true},
			[]string{
				"I love 56Music.",
				"A hate 34muSIc.",
				"Listen to 12MusIc.",
				"",
				"I love 78music of Kartik.",
				"We hate 46music of karTIK.",
				"Thanks.",
			},
			[]string{
				"I love 56Music.",
				"",
				"I love 78music of Kartik.",
				"Thanks.",
			},
		},
		{Options{FFlagStated: 2, SFlagStated: 3, CFlagStated: true},
			[]string{
				"I love AAAmusic.",
				"A hate BBBmusic.",
				"Listen to CCCmusic.",
				"",
				"I love AAAmusic of Kartik.",
				"We hate BBBmusic of Kartik.",
				"Thanks.",
			},
			[]string{
				"3 I love AAAmusic.",
				"1 ",
				"2 I love AAAmusic of Kartik.",
				"1 Thanks.",
			},
		},
		{Options{FFlagStated: 2, SFlagStated: 1, DFlagStated: true},
			[]string{
				"I love Amusic.",
				"A hate Bmusic.",
				"Listen to Cmusic.",
				"",
				"I love Amusic of Kartik.",
				"We hate Bmusic of Kartik.",
				"Thanks.",
			},
			[]string{
				"I love Amusic.",
				"I love Amusic of Kartik.",
			},
		},
		{Options{FFlagStated: 2, SFlagStated: 5, UFlagStated: true},
			[]string{
				"I love AAAmusic.",
				"A hate BBBmusic.",
				"Listen to CCCmusic.",
				"",
				"I love AAAmusic of Kartik.",
				"We hate BBBmusic of Kartik.",
				"Thanks.",
			},
			[]string{
				"",
				"Thanks.",
			},
		},
	}
	for _, test := range linesProcessingTests {
		result := LinesProcessing(test.inputOptions, test.inputLine)
		assert.Equal(t, test.out, result)
	}
}
