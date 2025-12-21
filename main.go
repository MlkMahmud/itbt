package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Is This Bucket Taken",
		Authors: []*cli.Author{
			{
				Name:  "Malik Mahmud",
				Email: "almalikmahmud@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "check",
				Action: func(ctx *cli.Context) error {
					bucketName := ctx.String("bucket-name")

					if err := validateBucketName(bucketName); err != nil {
						return err
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "bucket-name",
						Aliases:  []string{"b"},
						Required: true,
						Usage:    "the name of the AWS S3 bucket to check (must be a valid S3 bucket name: 3-63 characters, lowercase letters, numbers, dots and hyphens; must start and end with a letter or number)",
					},
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
				Usage: "check if an AWS S3 bucket name is taken",
			},
		},
		Description: "A command line utility for checking if an AWS S3 bucket name is taken.",
		HelpName:    "itbt",
		Usage:       "check if an AWS S3 bucket name is taken",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
