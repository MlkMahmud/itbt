package main

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

const (
	bucketNameMinLen = 3
	bucketNameMaxLen = 63
)

var (
	errBucketNameAdjacentPeriods       = errors.New("bucket names must not contain two adjacent periods (.)")
	errBucketNameFormat                = errors.New("bucket names must not be formatted as an IP address (for example, 192.168.5.4)")
	errBucketNameInvalidFirstCharacter = errors.New("bucket names must begin with a letter or number")
	errBucketNameInvalidCharacter      = errors.New("bucket names can consist only of lowercase letters, numbers, periods (.), and hyphens (-)")
	errBucketNameInvalidLastCharacter  = errors.New("bucket names must end with a letter or number")
	errBucketNameInvalidPrefix         = errors.New("bucket names must not start with the prefix")
	errBucketNameInvalidSuffix         = errors.New("bucket names must not start with the suffix")
	errBucketNameLen                   = fmt.Errorf("bucket names must be between %d (min) and %d (max) characters long", bucketNameMinLen, bucketNameMaxLen)
)

func validateBucketName(bucketName string) error {
	bucketNameLen := len(bucketName)

	if bucketNameLen < bucketNameMinLen || bucketNameLen > bucketNameMaxLen {
		return errBucketNameLen
	}

	if ip := net.ParseIP(bucketName); ip != nil {
		return errBucketNameFormat
	}

	if match, _ := regexp.MatchString("\\.{2,}", bucketName); match {
		return errBucketNameAdjacentPeriods
	}

	if match, _ := regexp.MatchString("[^-.a-z0-9]", bucketName); match {
		return errBucketNameInvalidCharacter
	}

	if match, _ := regexp.MatchString("^[a-z0-9]", bucketName); !match {
		return errBucketNameInvalidFirstCharacter
	}

	if match, _ := regexp.MatchString("[a-z0-9]$", bucketName); !match {
		return errBucketNameInvalidLastCharacter
	}

	for _, prefix := range []string{
		"xn--",
		"sthree-",
		"amzn-s3-demo-",
	} {
		if strings.HasPrefix(bucketName, prefix) {
			return fmt.Errorf("%w \"%s\"", errBucketNameInvalidPrefix, prefix)
		}
	}

	for _, suffix := range []string{
		"-s3alias",
		"--ol-s3",
		".mrap",
		"--x-s3",
		"--table-s3",
	} {
		if strings.HasSuffix(bucketName, suffix) {
			return fmt.Errorf("%w \"%s\"", errBucketNameInvalidSuffix, suffix)
		}
	}

	return nil
}
