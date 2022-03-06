package migration

import (
	"fmt"
	"time"

	"togo/models"
	"togo/models/dbcon"
	"togo/pkg/util"
)

func (a Migration) Migrate1Up() {
	sqlXDB := dbcon.GetSqlXDB()
	var columnName []string
	userModel := []models.Member{
		{ID: 1, Username: "test", Password: util.GetMD5Hash("123456"), MaxTask: 5, Status: true, CreatedAt: time.Now()},
	}
	columnName = []string{"ID", "Username", "Password", "MaxTask", "Status", "CreatedAt"}
	sqlXDB.NamedExec(fmt.Sprintf("insert into %s (%s) values (%s)",
		(models.Member{}).TableName(),
		models.ColumnsName((models.Member{}).TableName(), &models.Member{}, columnName),
		models.ColumnsNameValueList(&models.Member{}, columnName),
	), userModel)
}
