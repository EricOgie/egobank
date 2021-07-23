package helpers

import (
	"database/sql"
	"fmt"
	"strconv"
)

// SQLRowCount for row count
func SQLRowCount(rows *sql.Rows) int {
	count := 0
	for rows.Next() {
		count++
	}

	fmt.Println(strconv.Itoa(count))
	return count

}
