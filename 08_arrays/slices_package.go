package main

import (
	"fmt"
	"slices"
)

// "slices" package provides many utility functions for slices
// "sort" package simply calls "slices" package functions
func slicesPackageExample() {
	src := make([]uint8, 10, 12) // len 10, cap 12
	copy(src, []uint8{0, 0, 2, 3, 5, 5, 5, 8, 12, 13})

	// returns a cloned copy, instead of using `copy()` to existing slice
	cloned := slices.Clone(src)
	fmt.Printf("\n- Cloned: %v, Src: %v, Are Equal? %t\n", cloned, src, slices.Equal(src, cloned))

	// returns a clipped copy with cap == len
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
