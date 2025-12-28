package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_validateBucketName(t *testing.T) {
	inputs := map[string]error{
		// length checks
		"ab": errInvalidLen,
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": errInvalidLen, // 64 chars

		// IP address format
		"192.168.0.1": errInvalidFormat,

		// adjacent periods
		"my..bucket": errAdjacentPeriods,

		// invalid characters (uppercase, underscore, spaces, etc.)
		"my_Bucket":       errInvalidCharacter,
		"Mybucket":        errInvalidCharacter,
		"name with space": errInvalidCharacter,

		// invalid first character
		"-bucket": errInvalidFirstCharacter,
		".bucket": errInvalidFirstCharacter,

		// invalid last character
		"bucket-": errInvalidLastCharacter,
		"bucket.": errInvalidLastCharacter,

		// prohibited prefixes
		"xn--example":      errInvalidPrefix,
		"sthree-test":      errInvalidPrefix,
		"amzn-s3-demo-foo": errInvalidPrefix,

		// prohibited suffixes
		"good-suffix-s3alias": errInvalidSuffix,
		"name--ol-s3":         errInvalidSuffix,
		"name.mrap":           errInvalidSuffix,
		"name--x-s3":          errInvalidSuffix,
		"name--table-s3":      errInvalidSuffix,

		// valid names
		"abc":       nil,
		"a1b":       nil,
		"my-bucket": nil,
		"a.b-c1":    nil,
		// 63 characters (max allowed)
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": nil,
	}

	for bucketName, expectedErr := range inputs {
		t.Run(fmt.Sprintf("validate bucket name: %s", bucketName), func(t *testing.T) {
			receivedErr := validateBucketName(bucketName)

			if !errors.Is(receivedErr, expectedErr) {
				t.Errorf("Expected %v got %v", expectedErr, receivedErr)
			}
		})
	}
}
