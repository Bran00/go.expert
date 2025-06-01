package main

import (
	"context"
	"database/sql"

	"github.com/Bran00/go.expert/48/internal/db"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	/* err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	} */

	/* categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	} */

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "02280b4c-332b-4027-8855-8b6cc5c53924",
		Name:        "Backend Updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	})

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	
}
