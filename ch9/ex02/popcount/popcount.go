package popcount

import "sync"

var pc [256]byte
var makeLookupTableOnce sync.Once

func PopCount(x uint64) int {
	makeLookupTableOnce.Do(makeLookupTable)

	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}

func makeLookupTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
