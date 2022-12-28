package chessboard

// File - Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Chessboard - Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given File.
func CountInFile(cb Chessboard, file string) int {
	var result int
	for _, v := range cb[file] {
		if v {
			result++
		}
	}
	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	result := 0

	for _, v := range cb {
		if v[rank-1] {
			result++
		}
	}

	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	result := 0

	for _, v := range cb {
		result += len(v)
	}

	return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	result := 0

	for _, v := range cb {
		for _, v2 := range v {
			if v2 {
				result++
			}
		}
	}

	return result
}
