// Package datasource @author uangi 2023-05
package datasource

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/real-uangi/cockc/common/character"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/config"
	"time"
)

// database/sql 级别的DB,Stmt都是并发安全的

var dbs = make(map[string]*sql.DB)

var logger = plog.New("datasource")

func InitDataSource() {
	conf := config.GetPropertiesRO().Server.Datasource
	logger.Info("Initializing pu55y datasource...")
	var err error = nil
	// init multiple datasource
	for _, c := range conf {
		var cs string
		var db *sql.DB
		cs = character.AppendAll(
			"host=", c.Host,
			" port=", c.Port,
			" user=", c.User,
			" password=", c.Password,
			" dbname=", c.Database,
			" sslmode=disable",
		)
		db, err = sql.Open("postgres", cs)
		if err != nil {
			logger.Error(err.Error())
		} else {
			db.SetConnMaxLifetime(time.Hour)
			db.SetConnMaxIdleTime(5 * time.Minute)
			db.SetMaxIdleConns(2)
			db.SetMaxOpenConns(8)
		}
		dbs[c.Name] = db
		//check connection
		rows, err := dbs[c.Name].Query("select 1 as ans")
		if err != nil {
			logger.Error(err.Error())
			logger.Error("Datasource [" + c.Name + "] failed to initialize")
		} else {
			var a string
			rows.Next()
			e := rows.Scan(&a)
			if e != nil {
				logger.Error(e.Error())
			}
			logger.Info("Datasource [" + c.Name + "] initialized completed, test responded as : " + a)
		}

	}

}

func Get(name string) *sql.DB {
	return dbs[name]
}
