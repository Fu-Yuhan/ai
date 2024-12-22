package sqlx

// import "github.com/go-sql-driver/mysql"
import "github.com/lib/pq"

const (
	mysqlDriverName           = "postgres"
	duplicateEntryCode ="23505"
)

// NewMysql returns a mysql connection.
func NewPostgres(datasource string, opts ...SqlOption) SqlConn {
	opts = append(opts, withMysqlAcceptable())
	return NewSqlConn(mysqlDriverName, datasource, opts...)
}

func postgresAcceptable(err error) bool {
	if err == nil {
		return true
	}

	pqerr, ok := err.(*pq.Error)
	if !ok {
		return false
	}

	switch pqerr.Code {
	case duplicateEntryCode:
		return true
	default:
		return false
	}
}

func withMysqlAcceptable() SqlOption {
	return func(conn *commonSqlConn) {
		conn.accept = postgresAcceptable
	}
}
