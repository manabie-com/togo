package read

import (
	"context"
	"net/url"
	"time"

	"github.com/go-pg/pg"
	"github.com/gofrs/uuid"

	"github.com/sirupsen/logrus"

	"github.com/manabie-com/togo/internal/pkg/util"

	read_side "github.com/manabie-com/togo/internal/services/read-side"
)

type taskList struct {
	ID        uuid.UUID `sql:"id"`
	Username  string    `sql:"username"`
	Content   string    `sql:"content"`
	CreatedAt time.Time `sql:"created_at"`
}

func (r *readRepo) GetTaskList(ctx context.Context, values url.Values) ([]*read_side.TaskList, string, error) {
	model := []taskList{}

	q := r.db.Model().Table("tasks").
		Column("tasks.id",
			"users.username",
			"tasks.content",
			"tasks.created_at").
		Join("LEFT JOIN users ON users.id = tasks.user_id")

	if v, ok := values["user_id"]; ok {
		if len(v) > 0 {
			q.Where("tasks.user_id = ? AND users.user_id = ?", v[0], v[0])
		}
	}

	// for next page
	if v, ok := values["cursor"]; ok {
		if len(v) > 0 {
			createdAt, id, err := util.DecodeCursor(v[0])
			if err != nil {
				logrus.WithError(err).Errorf("Decode cursor failed %s", v[0])
			} else {
				q.Where("tasks.created_at <= ? and (tasks.id < ? OR tasks.created_at < ?)", createdAt, id, createdAt)
			}
		}
	}

	limit := defaultLimit
	q.Limit(limit)

	q.Order("tasks.created_at DESC")

	err := q.Select(&model)
	if err != nil {
		if err == pg.ErrNoRows {
			return []*read_side.TaskList{}, "", err
		}

		return []*read_side.TaskList{}, "", nil
	}

	var cursor string
	list := []*read_side.TaskList{}
	for idx, m := range model {
		list = append(list, &read_side.TaskList{
			ID:        m.ID,
			Username:  m.Username,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})

		if idx == limit-1 {
			cursor = util.EncodeCursor(m.CreatedAt, m.ID.String())
		}
	}

	return list, cursor, nil
}
