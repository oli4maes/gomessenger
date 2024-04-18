package mediator

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

var (
	registeredHandlers sync.Map
)

func init() {
	registeredHandlers = sync.Map{}
}

type key[TRequest any, TResponse any] struct{}

func Register[TRequest any, TResponse any](handler RequestHandler[TRequest, TResponse]) error {
	k := key[TRequest, TResponse]{}

	_, exists := registeredHandlers.LoadOrStore(reflect.TypeOf(k), handler)
	if exists {
		return errors.New("handler already registered")
	}
	return nil
}

func Send[TRequest any, TResponse any](ctx context.Context, r TRequest) (TResponse, error) {
	var noResponse TResponse

	var k key[TRequest, TResponse]
	handler, found := registeredHandlers.Load(reflect.TypeOf(k))
	if !found {
		return noResponse, errors.New("could not found response for this handler")
	}

	switch handler := handler.(type) {
	case RequestHandler[TRequest, TResponse]:
		return handler.Handle(ctx, r)
	}

	return noResponse, errors.New("handler not valid")
}

type RequestHandler[TRequest any, TResponse any] interface {
	Handle(ctx context.Context, request TRequest) (TResponse, error)
}
