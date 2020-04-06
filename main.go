package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/pkg/errors"
)


func main() {
	// Display help if all args not provided.
	if len(os.Args) < 4 {
		fmt.Print("Usage: sign-cloudfront-url CLOUDFRONT_URL KEY_PAIR_ID PRIVATE_KEY_PATH\n")
		return
	}

	// Otherwise get inputs.
	url := os.Args[1]
	keypairID := os.Args[2]
	privateKeyPath := os.Args[3]

	// Read private key.
	fmt.Printf("reading private key from %s ...\n", privateKeyPath)
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Printf("cannot read private key: %s", privateKeyPath)
		os.Exit(-1)
	}
	privateKeyText := strings.TrimSpace(string(privateKeyData))

	// Sign url.
	fmt.Printf("signing url %s with keypair %s and private key ...\n", url, keypairID)
	signed, err := getSignedURL(url, keypairID, privateKeyText)
	fmt.Printf("\n%s\n", signed)
}

// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-creating-signed-url-custom-policy.html#private-content-custom-policy-statement
// and https://docs.aws.amazon.com/sdk-for-go/api/service/cloudfront/sign/ for details on how this works.

func getCustomPolicy(url string, beginTime time.Time, endTime time.Time) *sign.Policy {
	return &sign.Policy{
		Statements: []sign.Statement{
			{
				Resource: url,
				Condition: sign.Condition{
					DateGreaterThan: &sign.AWSEpochTime{Time: beginTime.UTC()},
					DateLessThan:    &sign.AWSEpochTime{Time: endTime.UTC()},
				},
			},
		},
	}
}

func getPrivateKey(privateKeyText string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyText))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func getSignedURL(url string, keyID string, privateKeyText string) (string, error) {
	var signedURL string

	// Load private key from text.
	privateKey, err := getPrivateKey(privateKeyText)
	if err != nil {
		return "", errors.Wrap(err, "get private key failed")
	}

	// Get url signer for key pair with private key.
	signer := sign.NewURLSigner(keyID, privateKey)

	// Sign url.
	now := time.Now()
	beginTime := now
	endTime := now.Add(24 * time.Hour)
	policy := getCustomPolicy(url, beginTime, endTime)
	signedURL, err = signer.SignWithPolicy(url, policy)
	if err != nil {
		return "", errors.Wrap(err, "sign url failed")
	}

	// Success!
	return signedURL, nil
}
