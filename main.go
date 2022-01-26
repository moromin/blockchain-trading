package main

import (
	"blockchain-trading/config"
	"blockchain-trading/di"
	"blockchain-trading/infrastructure"
	"database/sql"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	driverName, dataSourceName := config.SetDBConfig()
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	container, err := di.NewOHLC(conn)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	infrastructure.OHLCRouting(e, container)
	e.Logger.Fatal(e.Start(":8000"))
}
