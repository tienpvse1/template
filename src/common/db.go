package common

import (
	"database/sql"
	"fmt"
	"template/src/generated/sqlc"
)

func InitConnectionAndGetQueries(env Env) *sqlc.Queries {
	db, err := sql.Open("postgres", fmt.Sprintf(`
		user=%s 
		dbname=%s 
		host=%s 
		password=%s 
		sslmode=disable`,
		env.Get("POSTGRES_USER"), env.Get("POSTGRES_DB"), env.Get("POSTGRES_HOST"), env.Get("POSTGRES_PASSWORD")))
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(db)
	return queries
}
