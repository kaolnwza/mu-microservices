package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type horoOdrSvc struct {
	tx      port.Transactor
	repo    port.HoroOrderRepository
	vochSvc port.VoucherService
	walSvc  port.WalletService
	seerSvc port.SeerService
	horoSvc port.HoroService
	chatSvc port.ChatService
}

func NewHoroOrderService(tx port.Transactor, repo port.HoroOrderRepository, vochSvc port.VoucherService, walSvc port.WalletService, horoSvc port.HoroService, seerSvc port.SeerService, chatSvc port.ChatService) port.HoroOrderService {
	return &horoOdrSvc{
		tx:      tx,
		repo:    repo,
		vochSvc: vochSvc,
		walSvc:  walSvc,
		horoSvc: horoSvc,
		seerSvc: seerSvc,
		chatSvc: chatSvc,
	}
}

func (s *horoOdrSvc) CreateHoroOrder(ctx context.Context, userUUID uuid.UUID, horoUUID uuid.UUID, voucherUUID uuid.UUID, price int, startTime string, endTime string) error {
	newVoucherUUID := "NULL"
	if voucherUUID != uuid.Nil {
		newVoucherUUID = "'" + voucherUUID.String() + "'"
	}

	//paymentUUID

	if err := s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		err := s.repo.Create(tx, userUUID, horoUUID, uuid.Nil, newVoucherUUID, price, startTime, endTime)
		if err != nil {
			return err
		}

		horo, err := s.horoSvc.GetHoroByHoroUUID(tx, horoUUID)
		if err != nil {
			return err
		}

		seerUsrUUID, err := s.seerSvc.GetUserUUIDBySeerUUID(tx, horo.SeerUUID)
		if err != nil {
			return err
		}

		if userWallet, err := s.walSvc.DecreaseUserWallet(tx, userUUID, int64(price)); err != nil {
			fmt.Println(userWallet)
			return err
		}

		if _, err := s.walSvc.IncreaseUserWallet(tx, *seerUsrUUID, int64(price)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *horoOdrSvc) GetUpcomingCustomerOrder(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.OrderCustomer, error) {
	order := []*entity.OrderCustomer{}

	if err := s.repo.GetUpcomingCustomerOrder(ctx, &order, seerUUID); err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *horoOdrSvc) GetCustomerOrderHistory(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.OrderCustomer, error) {
	order := []*entity.OrderCustomer{}

	if err := s.repo.GetCustomerOrderHistory(ctx, &order, seerUUID); err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *horoOdrSvc) GetOrderByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.OrderCustomer, error) {
	order := []*entity.OrderCustomer{}

	if err := s.repo.GetOrderByUserUUID(ctx, &order, userUUID); err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *horoOdrSvc) GetOrderHistoryByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.OrderCustomer, error) {
	order := []*entity.OrderCustomer{}

	if err := s.repo.GetOrderHistoryByUserUUID(ctx, &order, userUUID); err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *horoOdrSvc) UpdateOrderStatusConfirmedByUUID(ctx context.Context, status entity.HoroOrderStatus, orderUUID uuid.UUID, seerUUID uuid.UUID) error {

	return s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		order := &entity.Order{}
		if err := s.repo.GetOrderByUUID(tx, order, orderUUID); err != nil {
			return err
		}

		if err := s.repo.UpdateOrderStatusByUUID(tx, status, orderUUID); err != nil {
			return err
		}

		return s.chatSvc.CreateChatRoom(
			tx,
			orderUUID,
			order.UserUUID,
			seerUUID,
			order.StartTime,
			order.EndTime,
		)
	})
}

func (s *horoOdrSvc) UpdateOrderStatusSuccessByUUID(ctx context.Context, status entity.HoroOrderStatus, orderUUID uuid.UUID) error {

	return s.repo.UpdateOrderStatusByUUID(ctx, status, orderUUID)
}
