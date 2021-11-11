package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
	"strings"
	"webhook-better/helpers"
)

func JwtToken() echo.MiddlewareFunc {
	//jwtAuthentication = security.JwtAuthentication{}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("Call JwtToken")
			var accessToken string
			defer func() {
				fmt.Println("End JwtToken")
			}()

			req := c.Request()
			accessToken = req.Header.Get("Authorization")
			if len(accessToken) == 0 {
				return next(c)
			}

			index := strings.Index(accessToken, "Bearer")
			if index < 0 {
				index = strings.Index(accessToken, "Bearer")
			}
			if index >= 0 {
				accessToken = accessToken[index+len("Bearer"):]
				accessToken = strings.Trim(accessToken, " ")
			}

			//userClaim, err := jwtAuthentication.ConvertTokenUserClaim(accessToken)
			//if err != nil {
			//	return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			//}

			if accessToken != ""{
				req = req.WithContext(helpers.ContextHelper().SetValidToken(req.Context()))
			}
			//req = req.WithContext(helpers.ContextHelper().SetUserClaim(req.Context(), userClaim))
			c.SetRequest(req)

			return next(c)
		}
	}
}