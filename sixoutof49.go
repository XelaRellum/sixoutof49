package sixoutof49

import (
	"crypto/rand"
	"math/big"
	"sort"
)

func createRand(max int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return -1, err
	}

	return int(r.Int64()), nil
}

// Create create 4 unique and random numbers in the range of 1 to 49
// note: the result is sorted in ascending order
func Create() ([]int, error) {
	// create a bag full of 49 numbers
	bag := make([]int, 49)
	for i := 0; i < 49; i++ {
		bag[i] = i + 1
	}

	balls := []int{}

	// now make the throws
	for len(balls) < 6 {
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
