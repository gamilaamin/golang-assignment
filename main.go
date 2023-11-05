package main

import (
	"fmt"
	"golang-assignment/global"
	"golang-assignment/initialize"
)

func main() {
	global.VP = initialize.InitViper("./config.yaml") // Initialize the Viper
	global.DB = initialize.InitDB()                   // gorm Connect to the database for writing

	if global.DB != nil {
		initialize.RegisterTables(global.DB) // Initialize the table
		// Close the database link before the end of program
		db, _ := global.DB.DB()
		defer db.Close()
	}
	fmt.Println("DB: ", global.CONFIG.Mysql)
	fmt.Println("=========================================================")
	fmt.Println("========== Welcom To Pets Backend System v1.0 ===========")
	fmt.Println("=========================================================")
	fmt.Println(global.CONFIG.System.Addr)
	initialize.InitServer(global.CONFIG.System.Addr)
}
