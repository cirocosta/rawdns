package lib

// masks is a map of bit masks in the form
// of a slice where each entry retrieves a
// mask that when `AND`ed gives only the
// last `n` bits (0-indexed).
//
// 7 6 5 4 3 2 1 0
//
// 0 0 0 0 0 0 0 1     1   = (1 << 1 ) - 1
// 0 0 0 0 0 0 1 1     3   = (1 << 2 ) - 1
// 0 0 0 0 0 1 1 1     7   = (1 << 3 ) - 1
// 0 0 0 0 1 1 1 1     15  = (1 << 4 ) - 1
// 0 0 0 1 1 1 1 1     31  = (1 << 5 ) - 1
// 0 0 1 1 1 1 1 1     63  = (1 << 6 ) - 1
// 0 1 1 1 1 1 1 1     127 = (1 << 7 ) - 1
// 1 1 1 1 1 1 1 1     255 = (1 << 8 ) - 1
var masks = []byte{
	(1 << 1) - 1,
	(1 << 2) - 1,
	(1 << 3) - 1,
	(1 << 4) - 1,
	(1 << 5) - 1,
	(1 << 6) - 1,
	(1 << 7) - 1,
	(1 << 8) - 1,
}
