package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_validateBucketName(t *testing.T) {
	inputs := map[string]error{
		// length checks
		"ab": errBucketNameLen,
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": errBucketNameLen, // 64 chars

		// IP address format
		"192.168.0.1": errBucketNameFormat,

		// adjacent periods
		"my..bucket": errBucketNameAdjacentPeriods,

		// invalid characters (uppercase, underscore, spaces, etc.)
		"my_Bucket":       errBucketNameInvalidCharacter,
		"Mybucket":        errBucketNameInvalidCharacter,
		"name with space": errBucketNameInvalidCharacter,

		// invalid first character
		"-bucket": errBucketNameInvalidFirstCharacter,
		".bucket": errBucketNameInvalidFirstCharacter,

		// invalid last character
		"bucket-": errBucketNameInvalidLastCharacter,
		"bucket.": errBucketNameInvalidLastCharacter,

		// prohibited prefixes
		"xn--example":      errBucketNameInvalidPrefix,
		"sthree-test":      errBucketNameInvalidPrefix,
		"amzn-s3-demo-foo": errBucketNameInvalidPrefix,

		// prohibited suffixes
		"good-suffix-s3alias": errBucketNameInvalidSuffix,
		"name--ol-s3":         errBucketNameInvalidSuffix,
		"name.mrap":           errBucketNameInvalidSuffix,
		"name--x-s3":          errBucketNameInvalidSuffix,
		"name--table-s3":      errBucketNameInvalidSuffix,

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
