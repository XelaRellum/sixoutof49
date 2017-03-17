package sixoutof49

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sort"
)

// errors
var (
	ErrMinTooLow     = errors.New("min is too low (minimum 1)")
	ErrMaxTooLow     = errors.New("max is too low (maximum 1)")
	ErrCountTooLow   = errors.New("count ist too low (minimum 1)")
	ErrMinGreaterMax = errors.New("min is greater than max")
)

func createRand(max int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	if err != nil {
		return -1, err
	}

	return int(r.Int64()), nil
}

// Create create count unique and random numbers in the range of min to max
// note: the result is sorted in ascending order
func Create(min int, max int, count int) ([]int, error) {
	if count < 1 {
		return nil, ErrCountTooLow
	}

	if min < 1 {
		return nil, ErrMinTooLow
	}

	if max < 1 {
		return nil, ErrMaxTooLow
	}

	if min > max {
		return nil, ErrMinGreaterMax
	}

	// create a bag full of 49 numbers
	bagSize := max - min + 1
	bag := make([]int, bagSize)
	for i := 0; i < bagSize; i++ {
		bag[i] = i + min
	}

	balls := []int{}

	// now make the throws
	for len(balls) < count {
		// now select one ball out of the bag randomly
		slot, err := createRand(len(bag) - 1)
		if err != nil {
			return nil, err
		}

		balls = append(balls, bag[slot])

		// and remove this ball
		bag = append(bag[:slot], bag[slot+1:]...)
	}

	// return the sorted array
	sort.Ints(balls)
	return balls, nil
}
