package log

type Log struct {
	Mode string
	Type string
}

func CreateLogs() []Log {
	logs := []Log{
		{
			Mode: "INFO",
			Type: "user_login",
		},

		{
			Mode: "ERROR",
			Type: "database_connection",
		},

		{
			Mode: "INFO",
			Type: "page_open",
		},

		{
			Mode: "WARNING",
			Type: "high_memory",
		},

		{
			Mode: "ERROR",
			Type: "database_connection",
		},

		{
			Mode: "INFO",
			Type: "user_logout",
		},
	}

	return logs
}
