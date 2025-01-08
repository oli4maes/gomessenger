# Messenger
In-process messaging library.

## Usage
An example of how to use the request/response pattern with the appropriate request handler.

### Request

```go
// Request - Create Product 
type Request struct {
	Name        string
	Description string
}
```

### Response

```go
// Response - Create Product
type Response struct {
	ProductID   uuid.UUID
	Name        string
	Description string
}
```

### Request Handler

```go
// Register createProductHandler
repo := NewProductRepo()

err := mediator.Register[Request, Response](handler{repo: repo})
if err != nil {
    panic (err)
}
```

```go
// handler is the request handler, all dependencies should be added here
type handler struct {
    repo ProductRepo
}

func (h handler) Handle(ctx context.Context, request Request) (Response, error) {
    product := Product {
        ProductID:   uuid.New(),
        Name:        request.Name,
        Description: request.Description,
    }
    
    createdProduct, err := h.repo.Create(ctx, product)
    if err != nil {
        return  Response{}, err
    }
    
    return Response {
            ProductID:   createdProduct.ProductID,
            Name:        createdProduct.Name,
            Description: createdProduct.Description,
    }, nil   
}
```

### Sending a request
```go
request := createproduct.Request {
	Name: "name",
	Description: "description",
}

response, err := messenger.Send[Request, Response](ctx, request)
```
