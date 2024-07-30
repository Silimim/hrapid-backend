package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func InitDB() {

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	dsn := os.Getenv("MYSQL_USER") + ":" +
		os.Getenv("MYSQL_PASSWORD") + "@" +
		os.Getenv("MYSQL_PROTOCOL") + "(" +
		os.Getenv("MYSQL_HOST") + ":" +
		os.Getenv("MYSQL_PORT") + ")/hrapid?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error during db open: %s\n", err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Error during db migration: %s\n", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
	db.Close()
}

func INITORM() (db *gorm.DB, err error) {
	dsn := os.Getenv("MYSQL_USER") +
		":" +
		os.Getenv("MYSQL_PASSWORD") +
		"@" +
		os.Getenv("MYSQL_PROTOCOL") +
		"(" + os.Getenv("MYSQL_HOST") +
		":" + os.Getenv("MYSQL_PORT") +
		")/hrapid?charset=utf8mb4&parseTime=True&loc=Local"

	g := gen.NewGenerator(gen.Config{
		OutPath: "db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()

	return gormdb, err
}
