package frontend

import (
	"net/http"

	"goodjob/database"
)

func (f *Frontend) serveHomePage(w http.ResponseWriter, r *http.Request, db database.DB) error {
	return nil
}
