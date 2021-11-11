package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
	"webhook-better/helpers"
)

var isTrx = map[string]bool{
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
	"PATCH":  true,
	"GET":    false,
}

func GORM(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			ctx := req.Context()
			if !helpers.ContextHelper().IsValidToken(ctx) {
				log.Warn("Invalid token")
				return next(c)
			}
			if isTrx[req.Method] {
				log.Debug(">> trx", c.Request())
				tx := begin(db)
				defer func() {
					if r := recover(); r != nil {
						rollback(tx)
					}
				}()

				if err = tx.Error; err != nil {
					return
				}
				if err = next(c); err != nil {
					rollback(tx)
					return
				}

				if c.Response().Status >= 500 {
					return rollback(tx)
				}
				c.SetRequest(req.WithContext(helpers.ContextHelper().SetDB(ctx, tx)))
				return commit(tx)
			}

			log.Debug(">> none trx", c.Request())
			c.SetRequest(req.WithContext(helpers.ContextHelper().SetDB(ctx, db)))
			return next(c)
		}
	}
}

func rollback(db *gorm.DB) (err error) {
	log.Debug(">> rollback")
	if err = db.Rollback().Error; err != nil {
		log.Error("database rollback error", err.Error())
	}
	return
}

func commit(db *gorm.DB) (err error) {
	log.Debug(">> commit")
	if err = db.Commit().Error; err != nil {
		log.Error("database commit error", err.Error())
		rollback(db)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return
}

func begin(db *gorm.DB) *gorm.DB {
	log.Debug(">> begin")
	return db.Begin()
}
