package authentication

import (
	"encoding/json"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/defintly/backend/webserver/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()

		registrationData := &types.RegistrationData{}

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		err = json.Unmarshal(body, registrationData)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		if strings.TrimSpace(registrationData.Username) == "" || strings.TrimSpace(registrationData.Password) == "" ||
			strings.TrimSpace(registrationData.Mail) == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		authInfo, err := users.Register(registrationData.Username, registrationData.Mail, registrationData.Password,
			registrationData.FirstName, registrationData.LastName)
		if err != nil {
			if err == users.UserAlreadyExists {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.UserAlreadyExists)
			} else if err == users.InvalidMailAddressOrUsername {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidMailAddressOrUsername)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to register user: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, authInfo)
	}
}
