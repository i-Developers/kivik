package test

import (
	"net/http"

	"github.com/flimzy/kivik/test/kt"
)

func init() {
	RegisterSuite(SuiteCloudant, kt.SuiteConfig{
		"AllDBs.expected":         []string{"_replicator", "_users"},
		"AllDBs/NoAuth.status":    http.StatusUnauthorized,
		"AllDBs/RW/NoAuth.status": http.StatusUnauthorized,

		"Config/Admin/GetAll.status":             http.StatusForbidden,
		"Config/Admin/GetSection.sections":       []string{"chicken"},
		"Config/Admin/GetSection/chicken.status": http.StatusForbidden,
		"Config/NoAuth/GetAll.status":            http.StatusUnauthorized,
		"Config/NoAuth/GetSection.sections":      []string{"chicken"},
		"Config/NoAuth/GetItem.items":            []string{"foo.bar"},
		"Config/NoAuth/GetSection.status":        http.StatusUnauthorized,
		"Config/NoAuth/GetItem.status":           http.StatusUnauthorized,
		"Config/RW/NoAuth/Set.status":            http.StatusUnauthorized,
		"Config/RW/Admin/Set.status":             http.StatusForbidden,
		"Config/RW/NoAuth/Delete.status":         http.StatusUnauthorized,
		"Config/RW/Admin/Delete.status":          http.StatusForbidden,

		"CreateDB/NoAuth.status":         http.StatusUnauthorized,
		"CreateDB/Admin/Recreate.status": http.StatusPreconditionFailed,

		"AllDocs/Admin.databases":            []string{"_replicator", "chicken"},
		"AllDocs/Admin/_replicator.expected": []string{"_design/_replicator"},
		"AllDocs/Admin/_replicator.offset":   0,
		"AllDocs/Admin/chicken.status":       http.StatusNotFound,
		"AllDocs/NoAuth.databases":           []string{"_replicator", "chicken"},
		"AllDocs/NoAuth/_replicator.status":  http.StatusUnauthorized,
		"AllDocs/NoAuth/chicken.status":      http.StatusNotFound,
		"AllDocs/RW/NoAuth.status":           http.StatusUnauthorized,
	})
}
