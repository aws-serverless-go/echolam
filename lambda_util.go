package echolam

import "github.com/aws-serverless-go/httplam"

func IsLambdaRuntime() bool {
	return httplam.IsLambdaRuntime()
}
