# Use instead
https://github.com/awslabs/aws-lambda-go-api-proxy

# echolam
echolam is a library that bind [AWS Lambda](https://aws.amazon.com/lambda/), [AWS API Gateway](https://aws.amazon.com/api-gateway/) to [`Echo`](https://echo.labstack.com/) framework

# Install

```shell
go get -u github.com/aws-serverless-go/echolam
```

# Example

```go
package main

import (
	"github.com/aws-serverless-go/echolam"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	if echolam.IsLambdaRuntime() {
		echolam.StartLambdaWithAPIGateway(e)
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
```

# License
[MIT License](LICENSE)
