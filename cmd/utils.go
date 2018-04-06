package cmd

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"strings"
)

func buildSession() (*session.Session) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if (err != nil) {
		panic("Unable to describe parameters")
	}
	return sess
}

//func streamline(s string) string {
//	var re = regexp.MustCompile(`(^|[^_])\bproducts\b([^_]|$)`)
//	return re.ReplaceAllString(s, `$1.$2`)
//}

func buildModuleStageName(prefix string, module string, stage string) string {
	s := fmt.Sprintf("%s-%s-%s", prefix, module, stage)
	return strings.Replace(s, "--", "-", -1)
}

func buildParamName(prefix string, module string, stage string, param string) string {
	s := fmt.Sprintf("%s-%s", buildModuleStageName(prefix, module, stage), param)
	return strings.Replace(s, "--", "-", -1)
}
