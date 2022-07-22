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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate JWT token with custom claims",
	Long:  `Generate JWT token with custom claims. Private key support: RSA private key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vip := vipers["create"]
		claimFile := vip.GetString("claims")
		privateKeyFile := vip.GetString("privkey")

		payload, err := os.ReadFile(claimFile)
		if err != nil {
			log.Fatal(err)
			return err
		}

		privateKey, err := os.ReadFile(privateKeyFile)
		if err != nil {
			log.Fatal(err)
			return err
		}

		token, err := JwtToken(payload, privateKey)
		if err != nil {
			fmt.Printf("failed to create token")
			return err
		} else {
			fmt.Printf("%s\n", token)
		}
		return nil
	},
}

func JwtToken(payload, privKey []byte) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return "", err
	}

	// Map the claims from the payload.
	claims := new(jwt.MapClaims)
	if err = json.Unmarshal(payload, &claims); err != nil {
		return "", fmt.Errorf("error unmarshaling jwt claims: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return signed, nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("claims", "c", "", "A path to customs claims file")
	createCmd.PersistentFlags().StringP("privkey", "k", "", "A path to RSA private key (PEM format)")

	vip := viper.New()
	vipers["create"] = vip

	if err := vip.BindPFlags(createCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}
