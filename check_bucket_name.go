package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errFailedCheck = errors.New("bucket check failed")
)

func checkBucketName(bucketName string) (bool, error) {
	bucketURL := fmt.Sprintf("https://%s.s3.amazonaws.com", bucketName)
	resp, err := http.Get(bucketURL)

	if err != nil {
		return false, fmt.Errorf("%w: %w", errFailedCheck, err)
	}

	defer resp.Body.Close()

	// In S3, both StatusOK (bucket exists and is public) and StatusForbidden (bucket exists but is private) mean the bucket name is already taken.
	switch resp.StatusCode {
	case http.StatusOK, http.StatusForbidden:
		return false, nil

	case http.StatusNotFound:
		return true, nil

	default:
		return false, errFailedCheck
	}
}
