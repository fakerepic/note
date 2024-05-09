package mw

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
)

func GetIDFromPath(path string) string {
	return strings.Split(strings.TrimPrefix(path, "/userdb-"), "/")[0]
}

func CouchAuthHandler(c echo.Context) error {
	info := apis.RequestInfo(c)
	auth_record := info.AuthRecord
	req := c.Request()
	if auth_record == nil || !auth_record.Verified() {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	db_owner_id := GetIDFromPath(req.URL.Path)
	if auth_record.Id != db_owner_id {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	req.Header.Set("X-Auth-CouchDB-UserName", auth_record.Id)
	req.Header.Set("X-Auth-CouchDB-Roles", "user")
	return nil
}

func AIServiceAuthHandler(c echo.Context) error {
	info := apis.RequestInfo(c)
	auth_record := info.AuthRecord
	if auth_record == nil || !auth_record.Verified() {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	_userid := c.QueryParam("user_id")

	if _userid == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing user_id")
	}

	if _userid != auth_record.Id {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	return nil
}
