package migration

import (
	"fmt"
	"time"

	"github.com/khoale193/togo/models"
	"github.com/khoale193/togo/models/dbcon"
	"github.com/khoale193/togo/pkg/util"
)

func (a Migration) Migrate1Up() {
	sqlXDB := dbcon.GetSqlXDB()
	var columnName []string
	userModel := []models.Member{
		{ID: 1, Username: "test1", Password: util.GetMD5Hash("123456"), MaxTask: 6, Status: true, CreatedAt: time.Now()},
		{ID: 2, Username: "test2", Password: util.GetMD5Hash("123456"), MaxTask: 9, Status: true, CreatedAt: time.Now()},
	}
	columnName = []string{"ID", "Username", "Password", "MaxTask", "Status", "CreatedAt"}
	sqlXDB.NamedExec(fmt.Sprintf("insert into %s (%s) values (%s)",
		(models.Member{}).TableName(),
		models.ColumnsName((models.Member{}).TableName(), &models.Member{}, columnName),
		models.ColumnsNameValueList(&models.Member{}, columnName),
	), userModel)
}
