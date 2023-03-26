package levenshtein

type Algorithm int

const (
	DynamicProgramming = iota
	Recursion
)

// Distance returns the Levenshtein distance between s0 and s1.
func Distance(s0, s1 string, alg ...Algorithm) int {
	if len(alg) == 0 {
		alg = append(alg, DynamicProgramming)
	}

	if alg[0] == Recursion {
		return Rec(s0, s1)
	}

	d, _ := DP(s0, s1)
	return d
}

// DP returns the Levenshtein distance between s0 and s1 and the dynamic programming matrix.
func DP(s0, s1 string) (int, [][]int) {
	r0, r1 := []rune(s0), []rune(s1)
	n, m := len(r0), len(r1)

	// matrix
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, m+1)
	}

	// init
	for i := 0; i < n+1; i++ {
		d[i][0] = i
	}

	for j := 0; j < m+1; j++ {
		d[0][j] = j
	}

	// calc
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			var cost int
			if r0[i-1] != r1[j-1] {
				cost = 1
			}

			d[i][j] = min(d[i-1][j-1]+cost, d[i][j-1]+1, d[i-1][j]+1)
		}
	}

	return d[n][m], d
}

// Rec returns the Levenshtein distance between s0 and s1 using recursion.
func Rec(s0, s1 string) int {
	r0, r1 := []rune(s0), []rune(s1)

	if len(r0) == 0 {
		return len(r1)
	}

	if len(r1) == 0 {
		return len(r0)
	}

	var cost int
	if r0[0] != r1[0] {
		cost = 1
	}

	return min(
		Rec(string(r0[1:]), string(r1[1:]))+cost,
		Rec(string(r0[1:]), s1)+1,
		Rec(s0, string(r1[1:]))+1,
	)
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}
