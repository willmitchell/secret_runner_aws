package cmd

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"strings"
	"regexp"
)

func buildSession() (*session.Session) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if (err != nil) {
		panic("Unable to build Session.  Check your AWS credentials.")
	}
	return sess
}

func streamline(s string) string {
	var re = regexp.MustCompile(`/+`)
	return re.ReplaceAllString(s,"/")
}

func buildModuleStageName(prefix string, module string, stage string) string {
	s := fmt.Sprintf("/%s/%s/%s", prefix, module, stage)
	return streamline(s)
}

func getCurrentParamPath() string {
	//return fmt.Sprintf("%v/*",buildModuleStageName(prefix, module, stage))
	s:= buildModuleStageName(prefix, module, stage)
	return streamline(s)
}

func getCurrentParamPathPlus() string {
	return fmt.Sprintf("%v/",getCurrentParamPath())
}

func removeCurrentParamPath(s string) string {
	return strings.Replace(s, getCurrentParamPathPlus(), "", -1)
}

func buildParamName(param string) string {
	s := fmt.Sprintf("%s/%s", buildModuleStageName(prefix, module, stage), param)
	return streamline(s)
}
