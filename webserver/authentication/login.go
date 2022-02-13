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

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()

		loginData := &types.LoginData{}

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		err = json.Unmarshal(body, loginData)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		if strings.TrimSpace(loginData.UsernameOrMail) == "" || strings.TrimSpace(loginData.Password) == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidLoginData)
			return
		}

		authInfo, err := users.Login(loginData.UsernameOrMail, loginData.Password, ctx.GetHeader("User-Agent"))
		if err != nil {
			if err == users.UserNotFound || err == users.IncorrectPassword {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidLoginData)
			} else if err == users.InvalidMailAddressOrUsername {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidMailAddressOrUsername)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
				general.Log.Error("Failed to login user: ", err)
			}
			return
		}

		ctx.JSON(http.StatusOK, authInfo)
	}
}
