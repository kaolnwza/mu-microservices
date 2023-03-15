package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/auth/database/postgres"
	entity "github.com/kaolnwza/muniverse/auth/entities"
	grpcClient "github.com/kaolnwza/muniverse/auth/grpc/client"
	"github.com/kaolnwza/muniverse/auth/grpc/proto/pb"
	log "github.com/kaolnwza/muniverse/auth/logs"
	"github.com/kaolnwza/muniverse/auth/pkg"
	repository "github.com/kaolnwza/muniverse/auth/repositories"
)

func GoogleCallackByTokenHandler(c *gin.Context) {
	googleToken := c.Request.FormValue("access_token")
	displayName := c.Request.FormValue("display_name")
	dob := c.Request.FormValue("birthday")
	desc := c.Request.FormValue("description")
	telNumber := c.Request.FormValue("tel_number")

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + googleToken)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	oauth := entity.OAuth2{}
	if err := json.Unmarshal(contents, &oauth); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	db := postgres.NewPostgresDB()
	tx, _ := db.Beginx()
	user := entity.AuthGoogle{}

	if err := repository.FetchByEmail(tx, &user, oauth.Email); err != nil && err != sql.ErrNoRows {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	seerUUID := uuid.Nil
	userUUID := user.UserUUID
	existsUser := userUUID != uuid.Nil
	if !existsUser {

		rpcReq := &pb.CreateUserRequest{
			DisplayName: displayName,
			Description: desc,
			Birthday:    dob,
			TelNumber:   telNumber,
		}

		userConn := grpcClient.NewUserServiceClient()
		rpcResp, err := userConn.CreateUser(c.Request.Context(), rpcReq)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		userUUID, _ = uuid.Parse(rpcResp.Uuid)
		if err := repository.CreateAuthGoogle(db, userUUID, oauth.Email); err != nil {
			tx.Rollback()
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

	} else {
		seerConn := grpcClient.NewSeerServiceClient()
		seerRpc, _ := seerConn.GetSeerByUserUUID(c.Request.Context(), &pb.SeerRequest{UserUuid: userUUID.String()})

		seerUUID, err = uuid.Parse(seerRpc.Uuid)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	at, _, err := pkg.GenerateToken(userUUID, seerUUID)
	if err != nil {
		tx.Rollback()
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()

	c.JSON(200, *at)

}

func GoogleLoginTestHandler(c *gin.Context) {
	displayName := c.Request.FormValue("display_name")
	dob := c.Request.FormValue("birthday")
	desc := c.Request.FormValue("description")
	telNumber := c.Request.FormValue("tel_number")

	email := c.Request.FormValue("email")
	oauth := entity.OAuth2{
		Email: email,
	}

	db := postgres.NewPostgresDB()
	tx, _ := db.Beginx()
	user := entity.AuthGoogle{}

	err := repository.FetchByEmail(tx, &user, oauth.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	seerUUID := uuid.Nil
	userUUID := user.UserUUID
	existsUser := userUUID != uuid.Nil
	if !existsUser {
		rpcReq := &pb.CreateUserRequest{
			DisplayName: displayName,
			Description: desc,
			Birthday:    dob,
			TelNumber:   telNumber,
		}

		userConn := grpcClient.NewUserServiceClient()
		rpcResp, err := userConn.CreateUser(c.Request.Context(), rpcReq)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		userUUID, _ = uuid.Parse(rpcResp.Uuid)
		if err := repository.CreateAuthGoogle(db, userUUID, oauth.Email); err != nil {
			tx.Rollback()
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

	} else {
		seerConn := grpcClient.NewSeerServiceClient()
		seerRpc, err := seerConn.GetSeerByUserUUID(c.Request.Context(), &pb.SeerRequest{UserUuid: userUUID.String()})
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		seerUUID, err = uuid.Parse(seerRpc.Uuid)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		log.Info(seerRpc.Uuid)

	}

	at, _, err := pkg.GenerateToken(userUUID, seerUUID)
	if err != nil {
		tx.Rollback()
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()

	c.JSON(200, *at)

}
