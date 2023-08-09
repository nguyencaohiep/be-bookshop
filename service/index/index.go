package index

import (
	"net/http"
	"soc1-service/pkg/db"
	"soc1-service/pkg/router"
	"soc1-service/pkg/server"
	"strings"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "200", "Venture NIKA is running")
}

// GetHealth Function to Show Health Check Status
func GetHealth(w http.ResponseWriter, r *http.Request) {
	// Check Database Connections
	if len(server.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := db.MySQL.Ping()
			if err != nil {
				router.ResponseInternalError(w, "mysql-health-check", err)
				return
			}
		case "postgres":
			err := db.PSQL.Ping()
			if err != nil {
				router.ResponseInternalError(w, "postgres-health-check", err)
				return
			}
		case "mongo":
			err := db.MongoSession.Ping()
			if err != nil {
				router.ResponseInternalError(w, "mongo-health-check", err)
				return
			}
		}
	}

	// Return Success response
	router.ResponseSuccess(w, "200", "Health is ok")
}
