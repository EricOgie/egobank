package konstants

// Common Konstants strings constants
//
// Unless otherwise noted, these are defined for use
const (
	PORT          = "localhost:8000"
	ContentType   = "Content-Type"
	TypeXML       = "application/xml"
	TypeJSON      = "application/json"
	ServerError   = "Unexpected Server Error"
	AppStart      = "Starting Appliction ...."
	LogType       = "Logger-Type"
	Time          = "timestamp"
	NoClient      = "Customer Not Found"
	QueryError    = "Query Error Noticed! Error Msg: "
	DBCredentials = "root@tcp(localhost)/banking"
)

// GetClientStatus read Client status
func GetClientStatus(id string) string {
	if id == "0" {
		return "inactive"
	} else {
		return "active"
	}
}
