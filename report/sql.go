package report

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mojachieee/go-HoneyPot/config"
	
	"database/sql"
)

type ormReport struct {
	db *gorm.DB
	table string
}

func NewSqlReporter(cfg config.Database) (Reporter, error) {
	var rpt ormReport

	port := cfg.Port
	if port == "" {
		port = "3306"
	}
	str := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, port, cfg.Name)
		
	sql.Open(cfg.Driver, "")
	
	db, err := gorm.Open("mysql", str)
	if err != nil {
		nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		nil, err
	}
	rpt.db = db
	rpt.table = cfg.Table
	return &rpt, nil
}

func (r *ormReport) Pub(p *HoneypotRecord) error {
	db := r.db
	sql := fmt.Sprintf(`INSERT INTO %v (Date, InIp, InPort, DestIP, DestPort, DataLength)
		VALUES ("%v", "%v", "%v", "%v", "%v", "%v")`,
		cfg.Table, time.Now().Format("20060102150405"), remHost, remPort, locHost, locPort, n)
	db.Exec(sql)
	db.Ex
}

