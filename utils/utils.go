package utils

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/short-sh/shortsh-backend/models"
)

func WriteVisitorsData(engine *xorm.Engine, c echo.Context, urlMeta *models.Url) error {
	ip := c.RealIP()

	_, err := engine.Insert(models.Visitors{Url: urlMeta.Id, Ip: ip, Referrer: c.Request().Referer()})

	if err != nil {
		return err
	}

	return nil
}
