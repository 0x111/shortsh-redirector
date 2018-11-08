package utils

import (
	"github.com/0x111/shortsh-backend/models"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

func WriteVisitorsData(engine *xorm.Engine, c echo.Context, urlMeta *models.Url) error {
	ip := c.RealIP()

	_, err := engine.Insert(models.Visitors{Url: urlMeta.Id, Ip: ip})

	if err != nil {
		return err
	}

	return nil
}
