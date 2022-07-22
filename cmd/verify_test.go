package cmd

import "testing"

func TestVerify(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.QrwpfrOGmresmFzSv6fJVCxsNZKcUraV10oLhL_75KM4JNfsysmljB04WaOdqJ4Kr-uphZTyjpCub1b18iULpl50669qms1vev80yEH5EbYOMBrqZqDEyOfZiC4dcBhK9NrrlZBjPrBkBK2LQ-oex0VYWaewLtAAn_-SCXHhznbGVuRIrzAxLU0FnYWQx49tE9I6xCSqc5wchWj-vBA8JH6ezCXEJs_kppW5_cKHvSWrdM2Di09op0IiHU9Ezk7htJIaoEmSVerYWAzDni2bFX-uk2vbxELltNLt5-ezloc2Fug-rMOgmLAhedq6nfe3ZTIS2aWJHBjt8v269NtIFg"
	args := []string{
		"verify",
		"-t",
		token,
		"-k",
		"../default/public-key.pem",
	}

	rootCmd.AddCommand(verifyCmd)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()

	if err != nil {
		t.Fatal("verify command failed", err)
	}
}
