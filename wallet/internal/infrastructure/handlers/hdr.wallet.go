package handler

import (
	"net/http"

	port "github.com/kaolnwza/muniverse/wallet/internal/ports"
	log "github.com/kaolnwza/muniverse/wallet/lib/logs"
)

type walHdr struct {
	svc port.WalletService
}

func NewWalletHandler(s port.WalletService) walHdr {
	return walHdr{svc: s}
}

func (h *walHdr) GetFundByUserUUIDHandler(c port.Context) {
	userUUID := c.AccessUserUUID()

	wallet, err := h.svc.GetWalletByUserUUID(c.Ctx(), userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"fund": wallet.Fund,
	})
}
