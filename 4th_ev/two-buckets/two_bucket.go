// Package twobucket solves two bucket problems
package twobucket

import (
	"errors"
)

type bucket struct {
	name    string
	size    int
	current int
}

// Solve solves the two bucket problems
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string,
	numSteps,
	otherBucketLevel int,
	e error) {

	if sizeBucketOne == 0 || sizeBucketTwo == 0 {
		return "", 0, 0, errors.New("invalid bucket size")
	}

	if goalAmount == 0 {
		return "", 0, 0, errors.New("invalid goal amount")
	}

	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, errors.New("invalid start bucket name")
	}

	var moves int
	var lastPoured bool

	bucketOne := bucket{
		name: "one",
		size: sizeBucketOne,
	}

	bucketTwo := bucket{
		name: "two",
		size: sizeBucketTwo,
	}

	if startBucket == "one" {
		fillBucket(&bucketOne, &moves, &lastPoured)
		if goalAmount == bucketTwo.size {
			fillBucket(&bucketTwo, &moves, &lastPoured)
			return bucketTwo.name, moves, bucketOne.current, nil
		}
	}

	if startBucket == "two" {
		fillBucket(&bucketTwo, &moves, &lastPoured)
		if goalAmount == bucketOne.size {
			fillBucket(&bucketOne, &moves, &lastPoured)
			return bucketOne.name, moves, bucketTwo.current, nil
		}
	}

	for goalAmount != bucketOne.current && goalAmount != bucketTwo.current {
		if bucketOne.current != 0 && bucketTwo.current != bucketTwo.size && !lastPoured {
			pourBucket(&bucketOne, &bucketTwo, &moves, &lastPoured)
			continue
		}
		if bucketTwo.current != 0 && bucketOne.current != bucketOne.size && !lastPoured {
			pourBucket(&bucketTwo, &bucketOne, &moves, &lastPoured)
			continue
		}

		if bucketOne.current == bucketOne.size && bucketTwo.current != 0 {
			emptyBucket(&bucketOne, &moves, &lastPoured)
			continue
		}

		if bucketTwo.current == bucketTwo.size && bucketOne.current != 0 {
			emptyBucket(&bucketTwo, &moves, &lastPoured)
			continue
		}

		if bucketOne.current == 0 && bucketTwo.size-bucketTwo.current != 0 {
			fillBucket(&bucketOne, &moves, &lastPoured)
			continue
		}

		if bucketTwo.current == 0 && bucketOne.size-bucketOne.current != 0 {
			fillBucket(&bucketTwo, &moves, &lastPoured)
			continue
		}

		return "", 0, 0, errors.New("no solution")

	}

	if goalAmount == bucketOne.current {
		return bucketOne.name, moves, bucketTwo.current, nil
	}

	if goalAmount == bucketTwo.current {
		return bucketTwo.name, moves, bucketOne.current, nil
	}

	return "", 0, 0, errors.New("no solution")

}

func pourBucket(bucketFrom *bucket, bucketTo *bucket, moves *int, lastPoured *bool) {
	*moves++
	if (bucketTo.size - bucketTo.current) > bucketFrom.current {
		bucketTo.current = bucketTo.current + bucketFrom.current
		bucketFrom.current = 0
		*lastPoured = true
		return
	}
	bucketFrom.current = bucketFrom.current - (bucketTo.size - bucketTo.current)
	bucketTo.current = bucketTo.size
	*lastPoured = true
}

func fillBucket(bucketFill *bucket, moves *int, lastPoured *bool) {
	*moves++
	*lastPoured = false
	bucketFill.current = bucketFill.size
}

func emptyBucket(bucketEmpty *bucket, moves *int, lastPoured *bool) {
	*moves++
	*lastPoured = false
	bucketEmpty.current = 0
}