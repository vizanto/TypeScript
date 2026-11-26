//go:build !linux

package batch

// reserveHugePageHeap only applies to Linux transparent huge pages.
func reserveHugePageHeap(memory uint64) {}
