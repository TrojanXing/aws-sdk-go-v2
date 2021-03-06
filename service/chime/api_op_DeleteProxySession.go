// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteProxySessionInput struct {
	_ struct{} `type:"structure"`

	// The proxy session ID.
	//
	// ProxySessionId is a required field
	ProxySessionId *string `location:"uri" locationName:"proxySessionId" min:"1" type:"string" required:"true"`

	// The Amazon Chime voice connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProxySessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProxySessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProxySessionInput"}

	if s.ProxySessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProxySessionId"))
	}
	if s.ProxySessionId != nil && len(*s.ProxySessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProxySessionId", 1))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.VoiceConnectorId != nil && len(*s.VoiceConnectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VoiceConnectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProxySessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProxySessionId != nil {
		v := *s.ProxySessionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "proxySessionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteProxySessionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProxySessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProxySessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteProxySession = "DeleteProxySession"

// DeleteProxySessionRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the specified proxy session from the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using DeleteProxySessionRequest.
//    req := client.DeleteProxySessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteProxySession
func (c *Client) DeleteProxySessionRequest(input *DeleteProxySessionInput) DeleteProxySessionRequest {
	op := &aws.Operation{
		Name:       opDeleteProxySession,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/proxy-sessions/{proxySessionId}",
	}

	if input == nil {
		input = &DeleteProxySessionInput{}
	}

	req := c.newRequest(op, input, &DeleteProxySessionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteProxySessionRequest{Request: req, Input: input, Copy: c.DeleteProxySessionRequest}
}

// DeleteProxySessionRequest is the request type for the
// DeleteProxySession API operation.
type DeleteProxySessionRequest struct {
	*aws.Request
	Input *DeleteProxySessionInput
	Copy  func(*DeleteProxySessionInput) DeleteProxySessionRequest
}

// Send marshals and sends the DeleteProxySession API request.
func (r DeleteProxySessionRequest) Send(ctx context.Context) (*DeleteProxySessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProxySessionResponse{
		DeleteProxySessionOutput: r.Request.Data.(*DeleteProxySessionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProxySessionResponse is the response type for the
// DeleteProxySession API operation.
type DeleteProxySessionResponse struct {
	*DeleteProxySessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProxySession request.
func (r *DeleteProxySessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
