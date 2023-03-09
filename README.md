# levenshtein

```golang
d := levenshtein.Distance("kitten", "sitting")
fmt.Println(d)

// Output:
// 3
```

```golang
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
```
