package levenshtein_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/levenshtein"
)

func ExampleDistance() {
	fmt.Println(levenshtein.Distance("kitten", "sitting"))
	fmt.Println(levenshtein.Distance("kitten", "sitting", levenshtein.Recursion))
	fmt.Println(levenshtein.Distance("kitten", "sitting", levenshtein.DynamicProgramming))

	// Output:
	// 3
	// 3
	// 3
}

func ExampleDP() {
	d, dp := levenshtein.DP("kitten", "sitting")
	fmt.Println(d)

	for _, r := range dp {
		fmt.Println(r)
	}

	// Output:
	// 3
	// [0 1 2 3 4 5 6 7]
	// [1 1 2 3 4 5 6 7]
	// [2 2 1 2 3 4 5 6]
	// [3 3 2 1 2 3 4 5]
	// [4 4 3 2 1 2 3 4]
	// [5 5 4 3 2 2 3 4]
	// [6 6 5 4 3 3 2 3]
}

func ExampleRec() {
	d := levenshtein.Rec("kitten", "sitting")
	fmt.Println(d)

	// Output:
	// 3
}

func TestDP(t *testing.T) {
	cases := []struct {
		s0, s1 string
		d      int
	}{
		{"kitten", "sitting", 3},
		{"すいぎょうまつ", "うんらいまつ", 5},
	}

	for _, c := range cases {
		got, dp := levenshtein.DP(c.s0, c.s1)
		if got != c.d {
			for _, row := range dp {
				t.Log(row)
			}

			t.Errorf("got=%v, want=%v", got, c.d)
		}
	}
}

func TestRec(t *testing.T) {
	cases := []struct {
		s0, s1 string
		d      int
	}{
		{"kitten", "sitting", 3},
		{"すいぎょうまつ", "うんらいまつ", 5},
	}

	for _, c := range cases {
		got := levenshtein.Rec(c.s0, c.s1)
		if got != c.d {
			t.Errorf("got=%v, want=%v", got, c.d)
		}
	}
}
