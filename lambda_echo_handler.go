package echolam

import (
	"context"
	"encoding/json"
	"github.com/aws-serverless-go/httplam"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
)

var _ lambda.Handler = (*defaultHandler)(nil)

type defaultHandler struct {
	e *echo.Echo
}

func (h *defaultHandler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	var request events.APIGatewayV2HTTPRequest
	h.e.Logger.Info("json.Unmarshal, events.APIGatewayV2HTTPRequest")
	err := json.Unmarshal(payload, &request)
	if err != nil {
		h.e.Logger.Error("failed to json.Unmarshal events.APIGatewayV2HTTPRequest, weird payload")
		h.e.Logger.Error("payload data ", string(payload))
		h.e.Logger.Error(err)
		return nil, err
	}

	h.e.Logger.Info("constructing *http.Request from *events.APIGatewayV2HTTPRequest")
	req, err := httplam.NewHTTPRequest(ctx, &request)
	if err != nil {
		h.e.Logger.Error("failed to httplam.NewHTTPRequest")
		h.e.Logger.Error(err)
		return nil, err
	}

	var response events.APIGatewayV2HTTPResponse
	rw := httplam.NewAPIGatewayV2HTTPResponseBuilder(&response)

	h.e.Logger.Info("call *echo.ServeHTTP")
	h.e.ServeHTTP(rw, req)

	h.e.Logger.Info("call httplam.APIGatewayV2HTTPResponseBuilder.Build")
	_, err = rw.Build()
	if err != nil {
		h.e.Logger.Error("failed to httplam.APIGatewayV2HTTPResponseBuilder.Build")
		h.e.Logger.Error(err)
		return nil, err
	}

	h.e.Logger.Info("call json.Marshal events.APIGatewayV2HTTPResponse")
	return json.Marshal(response)
}

func StartLambdaWithAPIGateway(e *echo.Echo) {
	lambda.Start(&defaultHandler{e: e})
}
