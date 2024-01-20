package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/danubiobwm/goExpert/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %f", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCouse(ctx, db.CreateCouseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	dbConne, err := sql.Open("mysql", "admin:admin@admin@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConne.Close()

	queries := db.New(dbConne)
	courses, err := queries.ListCourses(ctx)

	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s, course ID: %s, course name: %s, course description: %s, course price: %f", course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)

	}

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Curso de Go Backend",
	// 	Description: sql.NullString{String: "Curso de Go Backend", Valid: true},
	// 	Price:       12.5,
	// }
	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConne)

	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)

	// if err != nil {
	// 	panic(err)
	// }

}
