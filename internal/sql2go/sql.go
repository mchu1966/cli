package sql2go

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string

	Charset string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// (m *DBModel)  有两个东西，此函数是获取第一个东西	DBEngine *sql.DB
func (m *DBModel) Connect() error {
	var err error
	var dsn string
	switch m.DBInfo.DBType {
	case "mysql":
		s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
		dsn = fmt.Sprintf( // dsn dataSourceName
			s,
			m.DBInfo.Username,
			m.DBInfo.Password,
			m.DBInfo.Host,
			m.DBInfo.Charset,
		)
	case "postgres":
		dsn = fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			m.DBInfo.Username,
			m.DBInfo.Password,
			m.DBInfo.Host,
			"postgres")
	default:
	}
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	// 第一个参数为驱动名称，eg mysql;
	// 第二个参数为驱动连接数据库的连接信息；dsn dataSourceName
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	var query string
	switch m.DBInfo.DBType {
	case "mysql":
		query = "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
		rows, err := m.DBEngine.Query(query, dbName, tableName)
		if err != nil {
			return nil, err
		}
		if rows == nil {
			return nil, errors.New("data not found")
		}
		defer rows.Close()

		var columns []*TableColumn
		for rows.Next() {
			var column TableColumn
			err := rows.Scan(&column.ColumnName, &column.DataType,
				&column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
			if err != nil {
				return nil, err
			}

			columns = append(columns, &column)
		}
		return columns, nil
	case "postgres":
		fields := []string{
			"COLUMN_NAME",
			"DATA_TYPE",
			"IS_NULLABLE",
			"CHARACTER_MAXIMUM_LENGTH",
			"CHARACTER_OCTET_LENGTH",
			"NUMERIC_PRECISION",
			"NUMERIC_PRECISION_RADIX",
			"NUMERIC_SCALE",
			"DATETIME_PRECISION",
			"INTERVAL_TYPE",
		}
		query = fmt.Sprintf("SELECT %s FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", strings.Join(fields, ","))
		rows, err := m.DBEngine.Query(query, dbName, tableName)
		if err != nil {
			return nil, err
		}
		if rows == nil {
			return nil, errors.New("data not found")
		}
		defer rows.Close()

		// var columns []*TableColumn
		for rows.Next() {

		}
	}
	return nil, errors.New("invalid db type")
}
