package data

var dataInterface DataAccessInterface = nil

func InitializeDatabaseService(dbType string) {
	switch dbType {
	case "IN_MEM":
		dataInterface = getInMemDBConnection()
	case "SQL":
		dataInterface = GetSQLDBConnection()
	default:
		//do nothing
		break
	}
}

func GetDataSource() DataAccessInterface {
	return dataInterface
}
