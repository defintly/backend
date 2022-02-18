package main

import (
	"github.com/alecthomas/kong"
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/importer"
	"github.com/defintly/backend/webserver"
)

var cli struct {
	DatabaseHostname string `help:"hostname of PostgreSQL instance" default:"127.0.0.1"`
	DatabasePort     int    `help:"port of the PostgreSQL instance" default:"5432"`
	DatabaseUser     string `help:"user of the PostgreSQL instance" default:"defintly"`
	DatabasePassword string `help:"password of the PostgreSQL instance" required:""`
	DatabaseName     string `help:"name of the PostgreSQL instance database" default:"defintly"`
	DatabaseSSLMode  string `help:"enable/disable SSL connection to the PostgreSQL instance (see PostgreSQL documentation of specific values to enable)" default:"disable"`

	Serve struct {
		WebserverHostname string `help:"ip to bind the webserver to" default:"127.0.0.1"`
		WebserverPort     int    `help:"port to bind the webserver to" default:"4269"`
	} `cmd:"" help:"Start the webserver."`

	Import struct {
		ExcelFile string `help:"Path to the excel file to import data from" type:"path" required:""`
	} `cmd:"" help:"Import data."`
}

func main() {
	kongCtx := kong.Parse(&cli)

	database.OpenConnection(cli.DatabaseHostname, cli.DatabasePort, cli.DatabaseUser, cli.DatabasePassword, cli.DatabaseName,
		cli.DatabaseSSLMode)

	initDDL()

	switch kongCtx.Command() {
	case "serve":
		webserver.Run(cli.Serve.WebserverHostname, cli.Serve.WebserverPort)
		break
	case "import":
		importer.ImportFromExcel(cli.Import.ExcelFile)
		break
	default:
		panic(kongCtx.PrintUsage(true))
	}
}

func initDDL() {
	database.MustExec("CREATE TABLE IF NOT EXISTS concept_comments(" +
		"id BIGSERIAL PRIMARY KEY," +
		"concept_id BIGINT NOT NULL," +
		"user_id BIGINT NOT NULL," +
		"text TEXT NOT NULL," +
		"parent_id BIGINT," +
		"allowed BOOLEAN NOT NULL" +
		")")
	database.MustExec("CREATE TABLE IF NOT EXISTS users(" +
		"id BIGSERIAL PRIMARY KEY," +
		"username TEXT NOT NULL," +
		"mail TEXT NOT NULL," +
		"first_name TEXT," +
		"last_name TEXT," +
		"password TEXT NOT NULL" +
		")")
	database.MustExec("CREATE TABLE IF NOT EXISTS user_sessions(" +
		"user_id BIGINT NOT NULL," +
		"session_key VARCHAR(36) NOT NULL," +
		"user_agent TEXT NOT NULL," +
		"creation_date DATE NOT NULL" +
		")")
	database.MustExec("CREATE TABLE IF NOT EXISTS roles(" +
		"id BIGSERIAL PRIMARY KEY," +
		"name TEXT NOT NULL," +
		"description TEXT" +
		")")
	database.MustExec("CREATE TABLE IF NOT EXISTS role_permissions(" +
		"role_id BIGINT NOT NULL," +
		"name TEXT NOT NULL" +
		")")
	database.MustExec("CREATE TABLE IF NOT EXISTS user_role_mapping(" +
		"user_id BIGINT NOT NULL," +
		"role_id BIGINT NOT NULL" +
		")")
}
