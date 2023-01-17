package types

type ExampleRequest struct {
	Query string `json:"query" protobuf:"bytes,1,opt,name=query"`
}

type ExampleResponse struct {
	Content string `json:"content" protobuf:"bytes,1,opt,name=content"`
}
