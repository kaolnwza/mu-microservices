package service

import (
	"github.com/gin-gonic/gin"
)

func UpdateHoroServiceImage(c *gin.Context) {
	// if err := repository.NewHoroServiceImage(database.NewPostgresDB(), &*upload, upload.Path, upload.Bucket, userUUID); err != nil {
	// 	log.Error(err)
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(201, map[string]string{
	// 	"upload_uuid": upload.UUID.String(),
	// })

}

// func (s GrpcStorageServer) GetProfileImage(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
// 	upload := entity.Upload{}
// 	userUUID, _ := uuid.Parse(req.UserUuid)

// 	err := repository.FetchImageProfile(database.NewPostgresDB(), &upload, userUUID)
// 	if err != nil && err != sql.ErrNoRows {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	url := ""
// 	if err != sql.ErrNoRows {
// 		url, err = helper.GenerateImageURI(ctx, upload.Bucket, upload.Path)
// 		if err != nil {
// 			log.Error(err)
// 			return nil, err
// 		}
// 	}

// 	return &pb.ProfileResponse{Url: url}, nil
// }
