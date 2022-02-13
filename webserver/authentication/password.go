package authentication

import (
	"encoding/json"
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ChangePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)

		defer ctx.Request.Body.Close()

		var body map[string]interface{}

		bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
			return
		}

		if value, ok := body["old_password"]; ok {
			if oldPassword, ok := value.(string); ok {
				if value, ok := body["new_password"]; ok {
					if newPassword, ok := value.(string); ok {

						err = users.ChangePassword(user, oldPassword, newPassword)
						if err != nil {
							if err == users.IncorrectPassword {
								ctx.AbortWithStatusJSON(http.StatusForbidden, errors.InvalidLoginData)
							} else {
								ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
							}
							return
						}

						ctx.JSON(http.StatusOK, gin.H{"status": "changed"})
						return
					}
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
		return
	}
}
