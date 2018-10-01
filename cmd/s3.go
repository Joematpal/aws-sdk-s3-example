// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("s3 called")

		awsArgs := []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"}

		for _, s := range awsArgs {
			key := viper.GetString(s)
			_ = os.Setenv(s, key)
		}

		fmt.Println(os.Getenv("AWS_ACCESS_KEY_ID"))
		fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))

		// item := "2018-10-01T05:00:10.30146400Z"
		// item := "MOCK_DATA.csv"
		// file, err := os.Create(item)
		// if err != nil {
		// 	panic("cant create file:" + err.Error())
		// }

		// path := "/out/handyman_matters/customer"
		// path := "customres/"

		// bucket := "connector-partner"
		bucket := "lambda-test-bucket-jmp"

		// file, err := os.Create("./" + item)
		// if err != nil {
		// 	panic("Unable to open file %q, %v" + err.Error())
		// }

		// defer file.Close()

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewEnvCredentials(),
		})

		if err != nil {
			panic("Failed to make new session" + err.Error())
		}

		svc := s3.New(sess)

		// delimiter := "/"

		out, err := svc.ListObjects(&s3.ListObjectsInput{
			Bucket: &bucket,
			// Prefix:    &path,
			// Delimiter: &delimiter,
		})

		if err != nil {
			fmt.Println("Failed to List Objects" + err.Error())
		}

		// fmt.Println("path: ", path)
		fmt.Printf("%+v\n", *out)

		// buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
		// if err != nil {
		// 	panic("Failed to list buckets: " + err.Error())
		// }
		// fmt.Printf("%+v\n", *buckets)

		// resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
		// if err != nil {
		// 	panic("Unable to list items in bucket: " + err.Error())
		// }

		// for _, item := range resp.Contents {
		// 	fmt.Println("Name:         ", *item.Key)
		// 	fmt.Println("Last modified:", *item.LastModified)
		// 	fmt.Println("Size:         ", *item.Size)
		// 	fmt.Println("Storage class:", *item.StorageClass)
		// 	fmt.Println("")
		// }

		// fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
		// fmt.Println("")

		// downloader := s3manager.NewDownloader(sess)

		// numBytes, err := downloader.Download(file,
		// 	&s3.GetObjectInput{
		// 		Bucket: aws.String(bucket),
		// 		Key:    aws.String(path + "/" + item),
		// 	})

		// fmt.Println("PATH: ", path+"/"+item)
		// fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
