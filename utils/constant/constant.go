package constant

type Dialect string

func (d Dialect) String() string {
	return string(d)
}

const (
	PostgresDialect Dialect = "postgres"
	MysqlDialect    Dialect = "mysql"
)

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	EmployeeTableName       TableName = "employees"
	UserTableName           TableName = "users"
	RoleTableName           TableName = "roles"
	PermissionTableName     TableName = "permissions"
	RolePermissionTableName TableName = "role_permissions"
)

type DateTimeFormat string

func (d DateTimeFormat) String() string {
	return string(d)
}

const (
	DateFormat DateTimeFormat = "2006-02-01"
)

type RegexFormat string

func (d RegexFormat) String() string {
	return string(d)
}
