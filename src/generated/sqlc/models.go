// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"database/sql"
)

type User struct {
	ID   int32
	Name string
	Bio  sql.NullString
}
