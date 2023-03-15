package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
	"github.com/kaolnwza/muniverse/order/lib/helper"
)

type horoSvc struct {
	tx      port.Transactor
	repo    port.HoroRepository
	seerSvc port.SeerService
	imgSvc  port.ImageStorageService
}

func NewHoroService(tx port.Transactor, repo port.HoroRepository, seerSvc port.SeerService, imgSvc port.ImageStorageService) port.HoroService {
	return &horoSvc{tx: tx, repo: repo, seerSvc: seerSvc, imgSvc: imgSvc}
}

func (s *horoSvc) CreateHoroService(
	ctx context.Context,
	userUUID uuid.UUID,
	title string,
	desc string,
	price int,
	horoLocation string,
	isMeet bool,
	longitude string,
	latitude string,
	isChat bool,
	isVoiceCall bool,
	isVideoCall bool,
	imagesJSON string,
) error {
	seer, err := s.seerSvc.GetSeerByUserUUID(ctx, userUUID)
	if err != nil {
		return err
	}

	newHoroLocation := entity.HoroLocationMap[horoLocation]
	if newHoroLocation == "" {
		return fmt.Errorf("horo_location is empty")
	}

	meetLocation := ""
	if longitude != "" && latitude != "" {
		meetLocation = helper.LatLongToPoint(latitude, longitude)
	}

	horo := entity.Horo{}
	if err := s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		if err := s.repo.Create(tx, seer.UUID, title, desc, price, newHoroLocation, isMeet, meetLocation, isChat, isVoiceCall, isVideoCall, &horo); err != nil {
			return err
		}

		if imagesJSON != "" {
			images := []*entity.HoroImagesRequest{}

			if err := helper.StringToJSON(&images, imagesJSON); err != nil {
				return err
			}

			if err := s.imgSvc.NewHoroServiceImages(tx, horo.UUID, images); err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

func (s *horoSvc) UpdateHoroService(
	ctx context.Context,
	horoUUID uuid.UUID,
	title string,
	desc string,
	price int,
	horoLocation string,
	isMeet bool,
	longitude string,
	latitude string,
	isChat bool,
	isVoiceCall bool,
	isVideoCall bool,
) error {
	newHoroLocation := entity.HoroLocationMap[horoLocation]
	if newHoroLocation == "" {
		return fmt.Errorf("horo_location is empty")
	}

	meetLocation := ""
	if longitude != "" && latitude != "" {
		meetLocation = helper.LatLongToPoint(latitude, longitude)
	}

	return s.repo.Update(ctx, horoUUID, title, desc, price, newHoroLocation, isMeet, meetLocation, isChat, isVoiceCall, isVideoCall)
}

func (s *horoSvc) UpdateHoroServiceOnEvent(ctx context.Context, horoUUID uuid.UUID, title string, desc string, horoLocation string, eventTime string) error {
	newHoroLocation := entity.HoroLocationMap[horoLocation]
	if newHoroLocation == "" {
		return fmt.Errorf("horo_location is empty")
	}

	if err := s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		if err := s.repo.UpdateOnEvent(tx, horoUUID, title, desc, newHoroLocation); err != nil {
			return err
		}
		if eventTime != "" {
			sche := []entity.HoroSchedule{}
			if err := helper.StringToJSON(&sche, eventTime); err != nil {
				return err
			}

			if err := s.repo.CreateSchedule(tx, sche, horoUUID); err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

func (s *horoSvc) UpdateHoroServiceStatus(ctx context.Context, horoUUID uuid.UUID, available bool) error {
	return s.repo.UpdateStatus(ctx, horoUUID, available)
}

func (s *horoSvc) GetHoroAvailableEventByDate(ctx context.Context, horoUUID uuid.UUID, date string) (*entity.HoroAvailableEvent, error) {
	sche := []*entity.HoroScheduleTime{}

	if err := s.repo.GetAvailableEventByDate(ctx, &sche, horoUUID, date); err != nil {
		return nil, err
	}

	horoInfo := entity.Horo{}
	if err := s.repo.GetHoroByHoroUUID(ctx, &horoInfo, horoUUID); err != nil {
		return nil, err
	}

	horo := entity.HoroAvailableEvent{
		HoroLocation:    horoInfo.HoroLocation,
		MeetingStatus:   horoInfo.MeetingStatus,
		MeetingLocation: horoInfo.MeetingLocation.String,
		ChatStatus:      horoInfo.ChatStatus,
		VoiceCallStatus: horoInfo.VoiceCallStatus,
		VideoCallStatus: horoInfo.VideoCallStatus,
	}

	for _, item := range sche {
		horo.Schedule = append(horo.Schedule, &entity.HoroScheduleTime{
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
		})
	}

	horo.Schedule = sche

	return &horo, nil
}

func (s *horoSvc) GetHoroScheduleEventByDate(ctx context.Context, horoUUID uuid.UUID, date string) (*[]*entity.HoroScheduleTime, error) {
	sche := []*entity.HoroScheduleTime{}

	if err := s.repo.GetScheduleEventByDate(ctx, &sche, horoUUID, date); err != nil {
		return nil, err
	}

	return &sche, nil
}

func (s *horoSvc) GetHoroByHoroUUID(ctx context.Context, horoUUID uuid.UUID) (*entity.HoroResponse, error) {
	horo := entity.Horo{}
	if err := s.repo.GetHoroByHoroUUID(ctx, &horo, horoUUID); err != nil {
		return nil, err
	}

	var lat, long = "", ""
	if horo.MeetingLocation.Valid {
		lat, long = helper.PointToLatLong(horo.MeetingLocation.String)
	}

	resp := entity.HoroResponse{
		UUID:            horo.UUID,
		SeerUUID:        horo.SeerUUID,
		HoroLocation:    horo.HoroLocation,
		HoroType:        horo.HoroType,
		Title:           horo.Title.String,
		Description:     horo.Description.String,
		Price:           horo.Price,
		Available:       horo.Available,
		MeetingStatus:   horo.MeetingStatus,
		MeetingLocation: entity.LineString{Latitude: lat, Longtitude: long},
		ChatStatus:      horo.ChatStatus,
		VoiceCallStatus: horo.VoiceCallStatus,
		VideoCallStatus: horo.VideoCallStatus,
	}

	img, err := s.imgSvc.GetHoroServiceImages(ctx, horoUUID)
	if err != nil {
		return nil, err
	}

	resp.Images = *img

	return &resp, nil
}

func (s *horoSvc) GetAllHoroService(ctx context.Context) (*[]*entity.HoroResponse, error) {
	horo := []*entity.Horo{}

	if err := s.repo.GetAllHoro(ctx, &horo); err != nil {
		return nil, err
	}

	resp := make([]*entity.HoroResponse, len(horo))

	imgSvcErr := make(chan error)

	for idx, item := range horo {
		go func(idx int, item *entity.Horo) {
			var lat, long = "", ""
			if item.MeetingLocation.Valid {
				lat, long = helper.PointToLatLong(item.MeetingLocation.String)
			}

			resp[idx] = &entity.HoroResponse{
				UUID:            item.UUID,
				SeerUUID:        item.SeerUUID,
				HoroLocation:    item.HoroLocation,
				HoroType:        item.HoroType,
				Title:           item.Title.String,
				Description:     item.Description.String,
				Price:           item.Price,
				Available:       item.Available,
				MeetingStatus:   item.MeetingStatus,
				ChatStatus:      item.ChatStatus,
				VoiceCallStatus: item.VoiceCallStatus,
				VideoCallStatus: item.VideoCallStatus,
				MeetingLocation: entity.LineString{Latitude: lat, Longtitude: long},
			}

			img, err := s.imgSvc.GetHoroServiceImages(ctx, resp[idx].UUID)
			imgSvcErr <- err

			resp[idx].Images = *img

		}(idx, item)
	}

	for i := 0; i < len(horo); i++ {
		if err := <-imgSvcErr; err != nil {
			return nil, err
		}
	}

	return &resp, nil
}

func (s *horoSvc) GetAllHoroServiceBySeerUUID(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.HoroResponse, error) {
	horo := []*entity.Horo{}

	if err := s.repo.GetAllHoroBySeerUUID(ctx, &horo, seerUUID); err != nil {
		return nil, err
	}

	resp := make([]*entity.HoroResponse, len(horo))

	imgSvcErr := make(chan error)

	for idx, item := range horo {
		go func(idx int, item *entity.Horo) {
			var lat, long = "", ""
			if item.MeetingLocation.Valid {
				lat, long = helper.PointToLatLong(item.MeetingLocation.String)
			}

			resp[idx] = &entity.HoroResponse{
				UUID:            item.UUID,
				SeerUUID:        item.SeerUUID,
				HoroLocation:    item.HoroLocation,
				HoroType:        item.HoroType,
				Title:           item.Title.String,
				Description:     item.Description.String,
				Price:           item.Price,
				Available:       item.Available,
				MeetingStatus:   item.MeetingStatus,
				ChatStatus:      item.ChatStatus,
				VoiceCallStatus: item.VoiceCallStatus,
				VideoCallStatus: item.VideoCallStatus,
				MeetingLocation: entity.LineString{Latitude: lat, Longtitude: long},
			}

			img, err := s.imgSvc.GetHoroServiceImages(ctx, resp[idx].UUID)
			imgSvcErr <- err

			resp[idx].Images = *img

		}(idx, item)
	}

	for i := 0; i < len(horo); i++ {
		if err := <-imgSvcErr; err != nil {
			return nil, err
		}
	}

	return &resp, nil
}
