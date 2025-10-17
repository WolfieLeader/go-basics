package main

import (
	"fmt"
	"slices"
	"time"
	"unique"
)

// "slices" package provides many utility functions for slices
func slicesPackageExample() {
	src := make([]uint8, 10, 12) // len 10, cap 12
	copy(src, []uint8{0, 0, 2, 3, 5, 5, 5, 8, 12, 13})

	// Returns a cloned copy, instead of using `copy()` and pre-allocating
	cloned := slices.Clone(src)
	fmt.Printf("\n- Cloned: %v, Src: %v, Are Equal? %t\n", cloned, src, slices.Equal(src, cloned))

	// Returns a clipped copy with cap == len
	clipped := slices.Clip(src)
	fmt.Printf("- Clipped capacity: %d, Src capacity: %d\n", cap(clipped), cap(src))

	compacted := slices.Compact(src)
	fmt.Printf("\n- Compacted: %v, Src: %v, Are Equal? %t\n", compacted, src, slices.Equal(src, compacted))

	bigger := slices.Grow(src, 5) // cap 24 because if 5 elements are added, cap will exceed current cap of 12
	fmt.Printf("- Bigger Capacity: %d, Src Capacity: %d\n", cap(bigger), cap(src))

	fmt.Printf("- Contains 5? %t, Contains 100? %t\n", slices.Contains(src, 5), slices.Contains(src, 100))
	fmt.Printf("- Index of 8: %d, Index of 100: %d\n", slices.Index(src, 8), slices.Index(src, 100))
	fmt.Printf("- Smallest: %d, Largest: %d\n", slices.Min(src), slices.Max(src))
	fmt.Printf("- IsSorted? %t\n", slices.IsSorted(src))

	indexOf12, found12 := slices.BinarySearch(src, 12)
	fmt.Printf("- Binary search for 12: index %d, found? %t\n", indexOf12, found12)

	slices.Reverse(src)
	fmt.Printf("- Reversed: %v, Is Sorted? %t\n", src, slices.IsSorted(src))

	slices.Sort(src)
	fmt.Printf("- Sorted: %v, and is Sorted? %t\n", src, slices.IsSorted(src))

	fmt.Printf("- New slice delete elements at index 3 to 5: %v\n", slices.Delete(src, 3, 6))
	fmt.Printf("- New slice insert 99 and 88 at index 4: %v\n", slices.Insert(src, 4, 99, 88))
	fmt.Printf("- Get repeat of a slice A, B, C: %v\n", slices.Repeat([]byte{'A', 'B', 'C'}, 3))
}

// "unique" interns comparable values: equal values share one canonical handle.
// Comparing handles is cheaper than comparing full values and can save memory
// when the same values repeat many times.
func uniquePackageExample() {
	longWords := []string{
		"Supercalifragilisticexpialidocious",
		"Hippopotomonstrosesquippedaliophobia",
		"Floccinaucinihilipilification",
	}
	repeated := slices.Repeat(longWords, 1000) // 3,000 entries

	// Prepare handles
	handles := make([]unique.Handle[string], 0, len(repeated))
	for _, word := range repeated {
		handles = append(handles, unique.Make(word))
	}

	creationTime := getTime(func() {
		fake := make([]unique.Handle[string], 0, len(repeated))
		for _, word := range repeated {
			fake = append(fake, unique.Make(word))
		}
	})

	var diffStr []string
	strDiffTime := getTime(func() {
		diffs := make([]string, 0, len(repeated))
		for _, word := range repeated {
			if !slices.Contains(diffs, word) {
				diffs = append(diffs, word)
			}
		}
		diffStr = diffs
	})

	var diffHandles []unique.Handle[string]
	handlesDiffTime := getTime(func() {
		diffs := make([]unique.Handle[string], 0, len(handles))
		for _, handle := range handles {
			if !slices.Contains(diffs, handle) {
				diffs = append(diffs, handle)
			}
		}
		diffHandles = diffs
	})

	fmt.Printf("- Total Entries: %d\n", len(repeated))
	fmt.Printf("- Strings (1k times): Distinct=%d, Total = %d Milliseconds\n", len(diffStr), strDiffTime.Milliseconds())
	fmt.Printf("- Handles (1k times):  Distinct=%d, Total = %d Milliseconds (Creation = %d Milliseconds)\n", len(diffHandles), handlesDiffTime.Milliseconds(), creationTime.Milliseconds())

	fmt.Print("- Values:")
	for _, handle := range diffHandles {
		fmt.Printf(" %q", handle.Value()) // `Value()` retrieves the original value
	}
	fmt.Println()
}

func getTime(fn func()) time.Duration {
	start := time.Now()
	for range 1000 {
		fn()
	}
	return time.Since(start)
}
