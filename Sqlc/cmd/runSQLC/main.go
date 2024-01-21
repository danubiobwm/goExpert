package main

import (
	"context"
	"database/sql"

	"github.com/danubiobwm/goExpert/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConne, err := sql.Open("mysql", "admin:admin@admin@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConne.Close()

	queries := db.New(dbConne)
	// CreateCategory

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Curso de Go",
	// 	Description: sql.NullString{String: "Curso de Go desde cero", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// UpdateCategory
	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "9c003df3-8abf-4433-a587-0b1ca1c1b866",
	// 	Name:        "Curso de Go Expert",
	// 	Description: sql.NullString{String: "Curso de Go desde Zero", Valid: true},
	// })
	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// DeleteCategory

	err = queries.DeleteCategory(ctx, "9c003df3-8abf-4433-a587-0b1ca1c1b866")
	if err != nil {
		panic(err)
	}
}
