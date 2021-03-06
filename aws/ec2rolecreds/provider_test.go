package ec2rolecreds_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"github.com/aws/aws-sdk-go-v2/aws/ec2rolecreds"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/internal/sdk"
)

const credsRespTmpl = `{
  "Code": "Success",
  "Type": "AWS-HMAC",
  "AccessKeyId" : "accessKey",
  "SecretAccessKey" : "secret",
  "Token" : "token",
  "Expiration" : "%s",
  "LastUpdated" : "2009-11-23T0:00:00Z"
}`

const credsFailRespTmpl = `{
  "Code": "ErrorCode",
  "Message": "ErrorMsg",
  "LastUpdated": "2009-11-23T0:00:00Z"
}`

func initTestServer(expireOn string, failAssume bool) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/latest/meta-data/iam/security-credentials/" {
			fmt.Fprintln(w, "RoleName")
		} else if r.URL.Path == "/latest/meta-data/iam/security-credentials/RoleName" {
			if failAssume {
				fmt.Fprintf(w, credsFailRespTmpl)
			} else {
				fmt.Fprintf(w, credsRespTmpl, expireOn)
			}
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}))
	return server
}

func TestProvider(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()

	server := initTestServer("2014-12-16T01:51:37Z", false)
	defer server.Close()

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL + "/latest")

	p := ec2rolecreds.New(ec2metadata.New(cfg))

	creds, err := p.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := "accessKey", creds.AccessKeyID; e != a {
		t.Errorf("Expect access key ID to match")
	}
	if e, a := "secret", creds.SecretAccessKey; e != a {
		t.Errorf("Expect secret access key to match")
	}
	if e, a := "token", creds.SessionToken; e != a {
		t.Errorf("Expect session token to match")
	}

	sdk.NowTime = func() time.Time {
		return time.Date(2014, 12, 16, 0, 55, 37, 0, time.UTC)
	}

	if creds.Expired() {
		t.Errorf("Expect not expired")
	}
}

func TestProvider_FailAssume(t *testing.T) {
	server := initTestServer("2014-12-16T01:51:37Z", true)
	defer server.Close()

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL + "/latest")

	p := ec2rolecreds.New(ec2metadata.New(cfg))

	creds, err := p.Retrieve(context.Background())
	if err == nil {
		t.Fatalf("expect error, got none")
	}

	var aerr awserr.Error
	if !errors.As(err, &aerr) {
		t.Fatalf("expect %T error, got %v", err, aerr)
	}
	if e, a := "ErrorCode", aerr.Code(); e != a {
		t.Errorf("expect %v code, got %v", e, a)
	}
	if e, a := "ErrorMsg", aerr.Message(); e != a {
		t.Errorf("expect %v message, got %v", e, a)
	}

	nestedErr := errors.Unwrap(aerr)
	if nestedErr != nil {
		t.Fatalf("expect no nested error, got %v", err)
	}

	if e, a := "", creds.AccessKeyID; e != a {
		t.Errorf("Expect access key ID to match")
	}
	if e, a := "", creds.SecretAccessKey; e != a {
		t.Errorf("Expect secret access key to match")
	}
	if e, a := "", creds.SessionToken; e != a {
		t.Errorf("Expect session token to match")
	}
}

func TestProvider_IsExpired(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()

	server := initTestServer("2014-12-16T01:51:37Z", false)
	defer server.Close()

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL + "/latest")

	p := ec2rolecreds.New(ec2metadata.New(cfg))

	sdk.NowTime = func() time.Time {
		return time.Date(2014, 12, 16, 0, 55, 37, 0, time.UTC)
	}

	creds, err := p.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if creds.Expired() {
		t.Errorf("expect not to be expired")
	}

	sdk.NowTime = func() time.Time {
		return time.Date(2014, 12, 16, 1, 55, 37, 0, time.UTC)
	}

	if !creds.Expired() {
		t.Errorf("expect to be expired")
	}
}

func TestProvider_ExpiryWindowIsExpired(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()

	server := initTestServer("2014-12-16T01:51:37Z", false)
	defer server.Close()

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL + "/latest")

	p := ec2rolecreds.New(ec2metadata.New(cfg), func(options *ec2rolecreds.ProviderOptions) {
		options.ExpiryWindow = time.Hour
	})

	sdk.NowTime = func() time.Time {
		return time.Date(2014, 12, 16, 0, 40, 37, 0, time.UTC)
	}

	creds, err := p.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if creds.Expired() {
		t.Errorf("expect not to be expired")
	}

	sdk.NowTime = func() time.Time {
		return time.Date(2014, 12, 16, 1, 30, 37, 0, time.UTC)
	}

	if !creds.Expired() {
		t.Errorf("expect to be expired")
	}
}

func BenchmarkProvider(b *testing.B) {
	server := initTestServer("2014-12-16T01:51:37Z", false)
	defer server.Close()

	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(server.URL + "/latest")

	p := ec2rolecreds.New(ec2metadata.New(cfg))

	if _, err := p.Retrieve(context.Background()); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Retrieve(context.Background()); err != nil {
			b.Fatal(err)
		}
	}
}
