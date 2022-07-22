/*
Copyright © 2022 The Tran <tranthepq@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a JWT token",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		vip := vipers["verify"]
		token := vip.GetString("token")
		publicKeyFile := vip.GetString("pubkey")

		// verify token
		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// read pem file
			keyPem, err := os.ReadFile(publicKeyFile)
			if err != nil {
				log.Printf("%v", err)
				return nil, err
			}

			// convert to rsa key
			pub, err := jwt.ParseRSAPublicKeyFromPEM(keyPem)
			if err != nil {
				log.Printf("%v", err)
				return nil, err
			}

			return pub, nil
		})

		if parseToken.Valid {
			fmt.Println("token is valid")

			fmt.Println("===== token:")
			fmt.Println("header: ", parseToken.Header)
			fmt.Println("claims: ", parseToken.Claims)
			fmt.Println("method: ", parseToken.Method)
			fmt.Println("signature: ", parseToken.Signature)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				fmt.Println("invalid token format")
			case ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0:
				fmt.Println("token is expired")
			default:
				fmt.Println("cannot handle this token:", err)
			}
			return err
		} else {
			fmt.Println("cannot handle this token:", err)
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.PersistentFlags().StringP("token", "t", "", "JWT token string to verify")
	verifyCmd.PersistentFlags().StringP("pubkey", "k", "", "A path to RSA public key use to verify token")

	vip := viper.New()
	vipers["verify"] = vip

	if err := vip.BindPFlags(verifyCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}
