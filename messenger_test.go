package messenger

import (
	"context"
	"testing"
)

type MockHandler struct{}

type Request struct {
	Name string
}

type Response struct {
	Name string
}

type NonExistentRequest struct{}

func (m *MockHandler) Handle(ctx context.Context, request Request) (Response, error) {
	return Response{Name: "test response"}, nil
}

func Test(t *testing.T) {
	handler := &MockHandler{}

	err := Register[Request, Response](handler)
	if err != nil {
		t.Errorf("Error registering handler: %v", err)
	}

	// Try registering the same handler again (should fail)
	err = Register[Request, Response](handler)
	if err == nil || err.Error() != "handler already registered" {
		t.Errorf("Expected 'handler already registered' error, got: %v", err)
	}

	// Send a request
	var req = Request{Name: "test request"}
	response, err := Send[Request, Response](context.Background(), req)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Check response
	expectedResponse := Response{Name: "test response"}
	if response != expectedResponse {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}

	// Test sending a request with no registered handler
	_, err = Send[NonExistentRequest, Response](context.Background(), NonExistentRequest{})
	if err == nil || err.Error() != "could not found response for this handler" {
		t.Errorf("Expected 'could not found response for this handler' error, got: %v", err)
	}
}
