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
	"bytes"
	"fmt"
	"net"

	"github.com/pkg/sftp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

// sftpCmd represents the sftp command
var sftpCmd = &cobra.Command{
	Use:   "sftp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sftpUser := viper.GetString("SFTP_USER")
		sftpPass := viper.GetString("SFTP_PASS")
		sftpAddr := viper.GetString("SFTP_ADDR")
		sftpPath := viper.GetString("SFTP_PATH")
		// addr := "167.99.226.2:22"
		config := &ssh.ClientConfig{
			User: sftpUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(sftpPass),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			//Ciphers: []string{"3des-cbc", "aes256-cbc", "aes192-cbc", "aes128-cbc"},
		}
		conn, err := ssh.Dial("tcp", sftpAddr+":22", config)
		if err != nil {
			panic("Failed to dial: " + err.Error())
		}

		client, err := sftp.NewClient(conn)
		if err != nil {
			panic("Failed to create client: " + err.Error())
		}
		// Close connection
		defer client.Close()
		wd, err := client.Getwd()
		if err != nil {
			panic("Failed to get working directory: " + err.Error())
		}
		fmt.Printf("Current working directory: %+v\n", wd)

		file, err := client.Create(sftpPath + "/file.txt")
		if err != nil {
			panic("Failed to make file:" + err.Error())
		}

		buf := []byte("hello\ngo\nshit")

		_, err = file.Write(buf)
		if err != nil {
			panic("Failed to write file:" + err.Error())
		}

		// fmt.Println("n", n)

		b := make([]byte, len(buf))
		w := bytes.NewBuffer(b)
		f, err := client.Open(sftpPath + "/file.txt")
		if err != nil {
			panic("Failed to open file: " + err.Error())
		}

		_, err = f.WriteTo(w)
		if err != nil {
			panic("Failed to read: " + err.Error())
		}
		// fmt.Println("n1", n1)
		fmt.Printf("%+v\n", w)
	},
}

func init() {
	rootCmd.AddCommand(sftpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sftpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sftpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
