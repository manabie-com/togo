//go:build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)
