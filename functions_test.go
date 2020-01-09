package mysqltools

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGetQueryString(t *testing.T) {

	query := fmt.Sprintf("INSERT INTO @uuid user_device (uuid, appVersion, osVersion, os, deviceModel, timezoneStr, languaje, timezone, onesignal_id, status, user_id, created_at, updated_at) VALUES ( @uuid, @appVersion, @osVersion, @os, @deviceModel, @timezone, @languaje, @timezone, @onesignalID, @status, @userID, NOW(), NOW())")

	query2, err := GetQueryString(
		query,
		sql.Named("uuid", "hola"),
		sql.Named("appVersion", 1),
		sql.Named("osVersion", 1),
		sql.Named("os", 1),
		sql.Named("deviceModel", 1),
		sql.Named("timezone", 1),
		sql.Named("languaje", 1),
		sql.Named("onesignalID", 1),
		sql.Named("status", "active"),
		sql.Named("userID", 1),
	)

	if err != nil {
		t.Log(query2)

		t.Error(err)
		return
	}

	t.Log(query2)
}
