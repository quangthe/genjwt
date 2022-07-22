package cmd

import "testing"

func TestCreateCmd(t *testing.T) {
	args := []string{
		"create",
		"-c",
		"../default/claims.json",
		"-k",
		"../default/private-key.pem",
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal("create command failed", err)
	}
}
