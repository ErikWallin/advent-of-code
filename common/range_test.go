package common

import (
	"testing"
)

func TestAddToRange(t *testing.T) {
	ranges := []Range{}
	ranges = AddToRange(ranges, Range{1, 4})
	if len(ranges) != 1 {
		t.Errorf("AddToRange expected len 1, got %d", len(ranges))
	}
	ranges = AddToRange(ranges, Range{8, 10})
	if len(ranges) != 2 {
		t.Errorf("AddToRange expected len 2, got %d", len(ranges))
	}
	ranges = AddToRange(ranges, Range{-5, -1})
	if len(ranges) != 3 {
		t.Errorf("AddToRange expected len 3, got %d", len(ranges))
	}
	ranges = AddToRange(ranges, Range{-7, -2})
	if len(ranges) != 3 {
		t.Errorf("AddToRange expected len 3, got %d", len(ranges))
	}
	ranges = AddToRange(ranges, Range{11, 13})
	if len(ranges) != 3 {
		t.Errorf("AddToRange expected len 3, got %d", len(ranges))
	}
	ranges = AddToRange(ranges, Range{-100, 100})
	if len(ranges) != 1 {
		t.Errorf("AddToRange expected len 1, got %d", len(ranges))
	}
}
