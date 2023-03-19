package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
	log "github.com/kaolnwza/muniverse/order/lib/logs"
)

type orderHdr struct {
	svc port.HoroOrderService
}

func NewHoroOrderHandler(svc port.HoroOrderService) orderHdr {
	return orderHdr{svc: svc}
}

// voucherUUID uuid.UUID, price int, startTime string, endTime string) error
func (h *orderHdr) CreateNewHoroOrderHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//empty --> uuid.Nil
	voucherUUID, _ := uuid.Parse(c.FormValue("voucher_uuid"))

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	startTime := c.FormValue("start_time")
	endTime := c.FormValue("end_time")

	if err := h.svc.CreateHoroOrder(c.Ctx(), userUUID, horoUUID, voucherUUID, price, startTime, endTime); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *orderHdr) GetUpcomingCustomerOrderHandler(c port.Context) {
	seerUUID := c.AccessSeerUUID()

	order, err := h.svc.GetUpcomingCustomerOrder(c.Ctx(), seerUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *orderHdr) GetCustomerOrderHistoryHandler(c port.Context) {
	seerUUID := c.AccessSeerUUID()

	order, err := h.svc.GetCustomerOrderHistory(c.Ctx(), seerUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *orderHdr) GetOrderByUserUUIDHandler(c port.Context) {
	userUUID := c.AccessUserUUID()

	order, err := h.svc.GetOrderByUserUUID(c.Ctx(), userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
func (h *orderHdr) GetOrderHistoryByUserUUIDHandler(c port.Context) {
	seerUUID := c.AccessUserUUID()

	order, err := h.svc.GetOrderHistoryByUserUUID(c.Ctx(), seerUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *orderHdr) UpdateOrderStatusConfirmedByUUIDHandler(c port.Context) {
	seerUUID := c.AccessSeerUUID()
	if seerUUID == uuid.Nil {
		err := errors.New("role has no authorized")
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	seerUserUUID := c.AccessUserUUID()

	orderUUID, err := uuid.Parse(c.Param("order_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	status := entity.HORO_ORDER_STATUS_CONFIRMED
	if err := h.svc.UpdateOrderStatusConfirmedByUUID(c.Ctx(), status, orderUUID, seerUserUUID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)

}

func (h *orderHdr) UpdateOrderStatusSuccessByUUIDHandler(c port.Context) {
	seerUUID := c.AccessSeerUUID()
	if seerUUID == uuid.Nil {
		err := errors.New("role has no authorized")
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	orderUUID, err := uuid.Parse(c.Param("order_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	status := entity.HORO_ORDER_STATUS_SUCCESS
	if err := h.svc.UpdateOrderStatusSuccessByUUID(c.Ctx(), status, orderUUID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)

}
