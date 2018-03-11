package main

import (
	"fmt"
)

/*

Problem description:

This program finds the number of (valid) phone numbers pressed on a phone keypad
with the constraint that a (valid) phone number sequence:
    is 7 digits in length,
    cannot start with 0 or 1,
    *must* be a sequence of knight movements

A knight movement parallels the eponymously named chess piece's valid movement
(that is it is an L shape).

1 2 3
4 5 6
7 8 9
_ 0 _


Problem Solution:

Consider a valid 2 digit sequence.  The first number is the starting number
(which can't be 0 or 1) and the next number is its possibilities.  For example,
starting at 6, we can go to 3 separate positions, namely 0, 7, and 1.
Continue this process for the rest of the digits (0 through 9 inclusively) and we
can imagine a row of counts, summing them would give us all the possible valid
sequences of length 2.  In this case, (+ 0, 0, 2, 2, 3, 0, 3, 2, 2, 2) = 16.
Hence, when the sequence length = 2, there are 16 valid knight movement sequences
we can generate on a phone pad.

Consider a valid 3 digit sequence.  This could be calculated using the prior
computed 2 digit sequences counts.  For example, if the first number is 3,
the next number can be either a 2 digit sequence starting with 4, or a 2 digit
sequence starting with 8.  Fortunately,  these  2 digit sequence counts were already computed.
We can simply visit the respective indices, summing them together.
Hence a 3 digit sequence starting at 3, is followed by 2 (2 digit) sequences
starting at 4 and 8, which gives us a count of 3+2 = 5.  We can similarly get a count
for *all* 3 digit sequence by following the above process for numbers (0 - 9 inclusively).

Hence, to compute a valid 7 digit sequence of knight movement phone numbers, we work
up from the base case  (1 digit sequences) until we've found a count for each number in
the 7th row, and sum them producing our result.

To keep track of the prior rows, we create a 7x10 matrix
where row (r) represents the number of r+1 valid knight movement sequences
which begin with column (c).

Our base case first row has only *1* valid move, the first row's values should all = 1,
with that added caveat that a sequence cannot start with 0 or 1.

*/

// compute valid phone numbers according to the problem description above
// should only need 2 rows of digits in memory simultaneously
func validPhoneNumbers(sequenceLength int) int {

	// handle edge case for incorrect sequence lengths
	if sequenceLength < 1 {
		return 0
	}

	// create the previous and next rows
	// populating with the base case
	previousRow := []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	nextRow := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// create the hashtable representing the possible digits from a starting digit
	possibleMoves := map[int][]int{
		0: {4, 6},
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {0, 3, 9},
		5: {},
		6: {0, 1, 7},
		7: {2, 6},
		8: {1, 3},
		9: {2, 4},
	}

	// for sequence length times
	for n := 1; n < sequenceLength; n++ {

		// zero next row (to accumulate the next iteration correctly)
		for i, _ := range nextRow {
			nextRow[i] = 0
		}

		// for each count in the next row
		for i := 0; i < 10; i++ {
			// accumulate the possible moves from that starting number
			for _, v := range possibleMoves[i] {
				nextRow[i] += previousRow[v]
			}

		}

		// copy next row into previous row
		copy(previousRow, nextRow)

	}

	// sum the final counts (previous row holds the last result)
	sum := 0
	for _, v := range previousRow {
		sum += v
	}

	return sum

}

func main() {
	formatString := "number of valid phone numbers of length: %d is %d\n"
	for i := 1; i <= 10; i++ {
		fmt.Printf(formatString, i, validPhoneNumbers(i))
	}
}
