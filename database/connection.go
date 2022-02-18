package database

import (
	"errors"
	"fmt"
	"github.com/defintly/backend/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"reflect"
	"time"
)

const DefaultTimeout = 3 * time.Second

type QueryResult struct {
	Results interface{}
	Error   error
}

var connection *sqlx.DB

var (
	NotAPointer      = errors.New("given type is not a pointer")
	NotAStruct       = errors.New("given type is not a struct")
	NoMatchingStruct = errors.New("no matching struct type found")
)

func OpenConnection(hostname string, port int, username string, password string, database string, sslMode string) {
	connection = sqlx.MustOpen("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", username, password, hostname, port, database, sslMode))
}

func PrepareStatement(query string, values ...interface{}) error {
	statement, err := connection.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(values...)

	return err
}

func Query(structType reflect.Type, query string, values ...interface{}) *QueryResult {
	if structType.String() != "pq.Int64Array" {
		if structType.Kind() != reflect.Ptr {
			return &QueryResult{nil, NotAPointer}
		}
		if structType.Elem().Kind() != reflect.Struct {
			return &QueryResult{nil, NotAStruct}
		}
	}

	switch structType.String() {
	case "*types.Category":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.Category)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.Collection":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.Collection)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.Criteria":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.Criteria)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.Concept":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.Concept)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.User":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.User)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.UserLoginInformation":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.UserLoginInformation)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.UsernameInformation":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.UsernameInformation)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.MailInformation":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.MailInformation)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.Role":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.Role)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.RolePermission":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.RolePermission)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	case "*types.IdInformation":
		results := reflect.MakeSlice(reflect.SliceOf(structType), 0, 0).Interface().([]*types.IdInformation)
		err := connection.Select(&results, query, values...)
		return &QueryResult{results, err}
	default:
		return &QueryResult{nil, NoMatchingStruct}
	}
}

func NamedQuery(object interface{}, query string, values interface{}) error {
	rows, err := connection.NamedQuery(query, values)
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.StructScan(object)
		if err != nil {
			return err
		}
	}

	return err
}

func NamedPrepareStatement(query string, values interface{}) error {
	statement, err := connection.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(values)

	return err
}

func MustExec(query string, values ...interface{}) {
	connection.MustExec(query, values...)
}
