// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMonitoringExecutionsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only jobs created after a specified time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only jobs created before a specified time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// Name of a specific endpoint to fetch jobs for.
	EndpointName *string `type:"string"`

	// A filter that returns only jobs modified before a specified time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only jobs modified after a specified time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of jobs to return in the response. The default value is
	// 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// Name of a specific schedule to fetch jobs for.
	MonitoringScheduleName *string `min:"1" type:"string"`

	// The token returned if the response is truncated. To retrieve the next set
	// of job executions, use it in the next request.
	NextToken *string `type:"string"`

	// Filter for jobs scheduled after a specified time.
	ScheduledTimeAfter *time.Time `type:"timestamp"`

	// Filter for jobs scheduled before a specified time.
	ScheduledTimeBefore *time.Time `type:"timestamp"`

	// Whether to sort results by Status, CreationTime, ScheduledTime field. The
	// default is CreationTime.
	SortBy MonitoringExecutionSortKey `type:"string" enum:"true"`

	// Whether to sort the results in Ascending or Descending order. The default
	// is Descending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that retrieves only jobs with a specific status.
	StatusEquals ExecutionStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListMonitoringExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMonitoringExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMonitoringExecutionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MonitoringScheduleName != nil && len(*s.MonitoringScheduleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MonitoringScheduleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMonitoringExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// A JSON array in which each element is a summary for a monitoring execution.
	//
	// MonitoringExecutionSummaries is a required field
	MonitoringExecutionSummaries []MonitoringExecutionSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of jobs, use it in the subsequent reques
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMonitoringExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMonitoringExecutions = "ListMonitoringExecutions"

// ListMonitoringExecutionsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns list of all monitoring job executions.
//
//    // Example sending a request using ListMonitoringExecutionsRequest.
//    req := client.ListMonitoringExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListMonitoringExecutions
func (c *Client) ListMonitoringExecutionsRequest(input *ListMonitoringExecutionsInput) ListMonitoringExecutionsRequest {
	op := &aws.Operation{
		Name:       opListMonitoringExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMonitoringExecutionsInput{}
	}

	req := c.newRequest(op, input, &ListMonitoringExecutionsOutput{})
	return ListMonitoringExecutionsRequest{Request: req, Input: input, Copy: c.ListMonitoringExecutionsRequest}
}

// ListMonitoringExecutionsRequest is the request type for the
// ListMonitoringExecutions API operation.
type ListMonitoringExecutionsRequest struct {
	*aws.Request
	Input *ListMonitoringExecutionsInput
	Copy  func(*ListMonitoringExecutionsInput) ListMonitoringExecutionsRequest
}

// Send marshals and sends the ListMonitoringExecutions API request.
func (r ListMonitoringExecutionsRequest) Send(ctx context.Context) (*ListMonitoringExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMonitoringExecutionsResponse{
		ListMonitoringExecutionsOutput: r.Request.Data.(*ListMonitoringExecutionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMonitoringExecutionsRequestPaginator returns a paginator for ListMonitoringExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMonitoringExecutionsRequest(input)
//   p := sagemaker.NewListMonitoringExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMonitoringExecutionsPaginator(req ListMonitoringExecutionsRequest) ListMonitoringExecutionsPaginator {
	return ListMonitoringExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMonitoringExecutionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListMonitoringExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMonitoringExecutionsPaginator struct {
	aws.Pager
}

func (p *ListMonitoringExecutionsPaginator) CurrentPage() *ListMonitoringExecutionsOutput {
	return p.Pager.CurrentPage().(*ListMonitoringExecutionsOutput)
}

// ListMonitoringExecutionsResponse is the response type for the
// ListMonitoringExecutions API operation.
type ListMonitoringExecutionsResponse struct {
	*ListMonitoringExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMonitoringExecutions request.
func (r *ListMonitoringExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
