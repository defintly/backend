package users

import (
	"encoding/json"
	"github.com/defintly/backend/general"
	"github.com/defintly/backend/types"
	"github.com/defintly/backend/users"
	"github.com/defintly/backend/webserver/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ChangeUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*types.User)
		id := ctx.MustGet("userId").(int)

		if id != user.Id {
			ctx.AbortWithStatusJSON(http.StatusForbidden, errors.NoPermission)
			return
		}

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

		if value, ok := body["username"]; ok {
			if username, ok := value.(string); ok {

				err = users.ChangeUsername(user, username)
				if err != nil {
					if err == users.UserAlreadyExists {
						ctx.AbortWithStatusJSON(http.StatusConflict, errors.UserAlreadyExists)
					} else {
						general.Log.Error("Failed to change username: ", err)
						ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.InternalError)
					}
					return
				}

				ctx.JSON(http.StatusOK, gin.H{"status": "changed"})
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.InvalidRequest)
		return
	}
}
