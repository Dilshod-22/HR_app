package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=caboose.proxy.rlwy.net user=postgres password=AtubJXGofjbxXOktxxmfwhEuSiFJOANF dbname=railway port=50084 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rows, err := db.Raw("SELECT column_name, data_type, is_nullable FROM information_schema.columns WHERE table_name='employees' ORDER BY ordinal_position").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var col, typ, nullable string
		rows.Scan(&col, &typ, &nullable)
		fmt.Printf("%-25s %-20s %s\n", col, typ, nullable)
	}
}
