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
	errAdjacentPeriods       = errors.New("bucket names must not contain two adjacent periods (.)")
	errInvalidCharacter      = errors.New("bucket names can consist only of lowercase letters, numbers, periods (.), and hyphens (-)")
	errInvalidFirstCharacter = errors.New("bucket names must begin with a letter or number")
	errInvalidFormat         = errors.New("bucket names must not be formatted as an IP address (for example, 192.168.5.4)")
	errInvalidLastCharacter  = errors.New("bucket names must end with a letter or number")
	errInvalidLen            = fmt.Errorf("bucket names must be between %d (min) and %d (max) characters long", bucketNameMinLen, bucketNameMaxLen)
	errInvalidPrefix         = errors.New("bucket names must not start with the prefix")
	errInvalidSuffix         = errors.New("bucket names must not start with the suffix")
)

func validateBucketName(bucketName string) error {
	bucketNameLen := len(bucketName)

	if bucketNameLen < bucketNameMinLen || bucketNameLen > bucketNameMaxLen {
		return errInvalidLen
	}

	if ip := net.ParseIP(bucketName); ip != nil {
		return errInvalidFormat
	}

	if match, _ := regexp.MatchString("\\.{2,}", bucketName); match {
		return errAdjacentPeriods
	}

	if match, _ := regexp.MatchString("[^-.a-z0-9]", bucketName); match {
		return errInvalidCharacter
	}

	if match, _ := regexp.MatchString("^[a-z0-9]", bucketName); !match {
		return errInvalidFirstCharacter
	}

	if match, _ := regexp.MatchString("[a-z0-9]$", bucketName); !match {
		return errInvalidLastCharacter
	}

	for _, prefix := range []string{
		"xn--",
		"sthree-",
		"amzn-s3-demo-",
	} {
		if strings.HasPrefix(bucketName, prefix) {
			return fmt.Errorf("%w \"%s\"", errInvalidPrefix, prefix)
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
			return fmt.Errorf("%w \"%s\"", errInvalidSuffix, suffix)
		}
	}

	return nil
}
