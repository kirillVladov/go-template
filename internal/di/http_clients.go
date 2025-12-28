package di

import (
	someClient "go-template/pkg/some_client"
)

func (di *DI) SomeClient() *someClient.SomeClient {
	if di.someClient != nil {
		return di.someClient
	}

	di.someClient = someClient.New()

	return di.someClient
}
