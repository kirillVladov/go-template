package di

import someClient "go-template/pkg/some_client"

type DI struct {
	someClient *someClient.SomeClient
}

func New() *DI {
	return &DI{}
}
