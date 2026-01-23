package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	date    = "unknown"
	version = "dev"
)

func main() {
	app := &cli.App{
		Name: "Is This Bucket Taken",
		Action: func(ctx *cli.Context) error {
			bucketName := ctx.Args().First()

			if err := validateBucketName(bucketName); err != nil {
				return err
			}

			bucketIsAvailable, err := checkBucketName(bucketName)

			if err != nil {
				return err
			}

			if bucketIsAvailable {
				fmt.Println("bucket name is not taken")
			} else {
				fmt.Println("bucket name is taken")
			}

			return nil
		},
		Authors: []*cli.Author{
			{
				Name:  "Malik Mahmud",
				Email: "almalikmahmud@gmail.com",
			},
		},
		Args:        true,
		Description: "Is This Bucket Taken (itbt) is a small command-line utility for validating AWS S3 bucket names and checking whether a given bucket name is already registered. When a name is taken, the tool can optionally generate suggested available alternatives to help you quickly find an acceptable name. Designed for interactive use, scripting, and pre-deployment checks.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "generate-suggestions",
				Aliases:  []string{"g"},
				Required: false,
				Usage:    "generate a fixed number of available bucket names if the checked bucket name is already taken",
				Value:    false,
			},
			&cli.IntFlag{
				Name:     "num-of-suggestions",
				Aliases:  []string{"n"},
				Required: false,
				Usage:    "number of bucket names to suggest (integer between 1 and 5, inclusive)",
				Value:    1,
				Base:     10,
			},
		},
		HelpName: "itbt",
		Usage:    "A command line utility for checking if an AWS S3 bucket name is taken.",
		UsageText: `itbt bucket_name [global options]

bucket_name:	the name of the AWS S3 bucket to check (must be a valid S3 bucket name: 3-63 characters, lowercase letters, numbers, dots and hyphens; must start and end with a letter or number)`,
		Version: version,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
