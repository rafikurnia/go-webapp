package utils

const (
	// These constants standardize the definition of the condition checking to determine what
	// database is used by the application.
	DBMode     = "db"      // Use docker for DB
	DBModeMock = "db_mock" // Use mock data and ignore DB in docker
	DBModeTest = "db_test" // Use docker for DB during test
)

// To allow dynamic mapping of the resources
var ResourcesMap = map[string]map[string]string{
	"contacts": {
		"singleton":   "contact",
		"dbFieldName": "Contacts",
	},
}
