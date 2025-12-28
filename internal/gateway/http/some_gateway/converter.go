package some_gateway

import (
	"go-template/internal/app/dto"
	someClient "go-template/pkg/some_client"
)

func convertToApp(in someClient.GetSomeItemsResponse) dto.SomeDto {
	return dto.SomeDto{
		ID: in.ID,
	}
}
