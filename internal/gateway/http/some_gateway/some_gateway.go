package some_gateway

import (
	"context"
	"go-template/internal/app/dto"
	someClient "go-template/pkg/some_client"
)

type SomeClient interface {
	Get(_ context.Context) (someClient.GetSomeItemsResponse, error)
}

type SomeGateway struct {
	someClient SomeClient
}

func New(client SomeClient) *SomeGateway {
	return &SomeGateway{
		someClient: client,
	}
}

func (g *SomeGateway) Get(ctx context.Context) (dto.SomeDto, error) {
	response, err := g.someClient.Get(ctx)
	if err != nil {
		return dto.SomeDto{}, err
	}

	return convertToApp(response), nil
}
