package main

import (
	"fmt"
	"log"
	"os"

	_ "rest-api/docs"
	"rest-api/internal/db"
	"rest-api/internal/env"
	"rest-api/internal/store"
)


// @title      Rest api
// @version    1.0
// @description Rest Api for test
// @BasePath   /api/
func main() {

	conf := &config{
		addr: env.GetEnv("PORT", ":4000"),
		db: dbConfig{
			addr:         env.GetEnv("MYSQL_DNS", "root:admin@(127.0.0.1:3306)/js_api?parseTime=true"),
			maxOpenConns: env.GetEnvInt("DB_MAX_OPEN_CONNS", 20),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE_CONNS", 10),
			maxIdleTime:  env.GetEnv("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		conf.db.addr,
		conf.db.maxOpenConns,
		conf.db.maxIdleConns,
		conf.db.maxIdleTime,
	)
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database connection pool established")
	fmt.Println("Server is running in http://localhost:4000/")

	store := store.NewStorage(db)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	
	app := &application{
		config:  *conf,
		store:   store,
		infoLog: infoLog,
	}

	mux := app.routes()

	err = app.run(mux)
	if err != nil {
		fmt.Errorf(err.Error())
	}

}
