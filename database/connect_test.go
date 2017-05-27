package database

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestX(t *testing.T) {
	t.SkipNow()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT 2+2 FROM dual").WillReturnRows(sqlmock.NewRows([]string{"blah"}).FromCSVString("4"))
}
