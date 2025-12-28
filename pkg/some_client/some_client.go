package some_client

import (
	"context"
)

type SomeClient struct{}

func New() *SomeClient {
	return &SomeClient{}
}

func Get(_ context.Context) (GetSomeItemsResponse, error) {
	responseAnswer := GetSomeItemsResponse{
		ID: 1,
	}

	return responseAnswer, nil
}
