/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// poCmd represents the po command
var poCmd = &cobra.Command{
	Use:   "po",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		arg1, err := hex.DecodeString(args[0])
		if err != nil {
			panic(err)
		}
		arg2, err := hex.DecodeString(args[1])
		if err != nil {
			panic(err)
		}

		po := Proximity(arg1, arg2)
		fmt.Println("proximity between addresses", po)

		fmt.Println("addr1", args[0])
		printBinary(arg1)

		fmt.Println("addr2", args[1])
		printBinary(arg2)
	},
}

func init() {
	rootCmd.AddCommand(poCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const MaxPO = 16

func Proximity(one, other []byte) (ret int) {
	b := (MaxPO-1)/8 + 1
	if b > len(one) {
		b = len(one)
	}
	m := 8
	for i := 0; i < b; i++ {
		oxo := one[i] ^ other[i]
		for j := 0; j < m; j++ {
			if (oxo>>uint8(7-j))&0x01 != 0 {
				return i*8 + j
			}
		}
	}
	return MaxPO
}

const n = 6

func printBinary(addr []byte) {
	endStr := ""
	for i, v := range addr {
		str := ""
		for ii := 0; ii < 8; ii++ {
			if (v>>ii)&0x1 == 1 {
				str = "1" + str
			} else {
				str = "0" + str
			}
		}
		endStr = endStr + str

		if i == n-1 {
			break
		}
	}
	fmt.Println(endStr)
}
