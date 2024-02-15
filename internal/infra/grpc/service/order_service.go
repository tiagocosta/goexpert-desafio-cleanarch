package service

import (
	"context"

	"github.com/tiagocosta/goexpert-desafio-cleanarch/internal/infra/grpc/pb"
	"github.com/tiagocosta/goexpert-desafio-cleanarch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrdersList, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order

	for _, orderOutputDto := range output {
		order := &pb.Order{
			Id:         orderOutputDto.ID,
			Price:      float32(orderOutputDto.Price),
			Tax:        float32(orderOutputDto.Tax),
			FinalPrice: float32(orderOutputDto.FinalPrice),
		}

		orders = append(orders, order)
	}

	return &pb.OrdersList{Orders: orders}, nil
}
