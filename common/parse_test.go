package common

import (
	"testing"
)

func TestParseIntList(t *testing.T) {
	var input = `1
23
45`
	var ints = ParseIntList(input, "\n")
	if len(ints) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(ints))
	}
	if ints[1] != 23 {
		t.Errorf("ParseIntList expected item 23, got %d", ints[1])
	}
}

func TestParseIntCsvList(t *testing.T) {
	var input = `1,23,45`
	var ints = ParseIntList(input, ",")
	if len(ints) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(ints))
	}
	if ints[1] != 23 {
		t.Errorf("ParseIntList expected item 23, got %d", ints[1])
	}
}

func TestParseStringList(t *testing.T) {
	var input = `hej
på
dig`
	var rows = ParseStringList(input, "\n")
	if len(rows) != 3 {
		t.Errorf("ParseStringList expected len 3, got %d", len(rows))
	}
	if rows[1] != "på" {
		t.Errorf("ParseStringList expected item <på>, got %s", rows[1])
	}
}

func TestParseStringList_sentence(t *testing.T) {
	var input = `hej på dig`
	var ints = ParseStringList(input, " ")
	if len(ints) != 3 {
		t.Errorf("ParseStringList expected len 3, got %d", len(ints))
	}
	if ints[1] != "på" {
		t.Errorf("ParseStringList expected item <på>, got %s", ints[1])
	}
}

func TestParseStringStringList(t *testing.T) {
	var input = `hej

Kapitel 1
Blabla

Modo är bäst
`
	var groups = ParseStringStringList(input, "\n\n", "\n")
	if len(groups) != 3 {
		t.Errorf("ParseStringStringList expected len 3, got %d", len(groups))
	}
	if groups[1][1] != "Blabla" {
		t.Errorf("ParseStringStringList expected item <Blabla>, got %s", groups[1][1])
	}
}

func TestParseRuneListList(t *testing.T) {
	var input = `abc
def
ghi`
	var runes = ParseRuneListList(input)
	if len(runes) != 3 {
		t.Errorf("ParseRuneListList expected len 3, got %d", len(runes))
	}
	if len(runes[0]) != 3 {
		t.Errorf("ParseRuneListList expected len 3, got %d", len(runes[0]))
	}
	if runes[2][1] != 'h' {
		t.Errorf("ParseRuneListList expected rune x, got %q", runes[2][1])
	}
	if runes[1][1] != 'e' {
		t.Errorf("ParseRuneListList expected rune e, got %q", runes[1][1])
	}
}

func TestParseRuneListListList(t *testing.T) {
	var input = `abc
def
ghi

123
456
789`
	var runes = ParseRuneListListList(input)
	if len(runes) != 2 {
		t.Errorf("ParseRuneListListList expected len 2, got %d", len(runes))
	}
	if len(runes[0]) != 3 {
		t.Errorf("ParseRuneListListList expected len 3, got %d", len(runes[0]))
	}
	if len(runes[0][1]) != 3 {
		t.Errorf("ParseRuneListListList expected len 3, got %d", len(runes[0][1]))
	}
	if runes[0][2][1] != 'h' {
		t.Errorf("ParseRuneListListList expected rune h, got %q", runes[0][2][1])
	}
	if runes[1][1][0] != '4' {
		t.Errorf("ParseRuneListList expected rune 4, got %q", runes[1][1][0])
	}
}
