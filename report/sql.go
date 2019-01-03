package report

import (
	"database/sql"
	"fmt"

	"github.com/chennqqi/go-HoneyPot/config"

	_ "github.com/go-sql-driver/mysql"
)

type ormReport struct {
	db    *sql.DB
	table string
}

func NewSqlReporter(cfg *config.Database) (Reporter, error) {
	var rpt ormReport

	port := cfg.Port
	if port == "" {
		port = "3306"
	}
	connection := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, port, cfg.Name)

	db, err := sql.Open(cfg.Driver, connection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.Exec(`
		CREATE TABLE IF NOT EXISTS tbl_honeypot (
			id int AUTO_INCREMENT NOT NULL, 
			src varchar(16) not NULL, 
			dst varchar(16) not NULL,
			srcport int not null,
			dstport int not null,
			atime TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
			payload VARBINARY(4096),
			raw BLOB, 
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	rpt.db = db
	return &rpt, nil
}

func (r *ormReport) Pub(p *HoneypotRecord) error {
	db := r.db
	_, err := db.Exec(`INSERT INTO tbl_honeypot(src, dst, srcport, dstport, payload, raw)
		VALUES($1, $2, $3, $4, $5, $6)`,
		p.Src, p.Dst, p.Srcport, p.Dstport, p.Payload, sql.RawBytes(p.Raw))
	return err
}

func (r *ormReport) Close() error {
	if r.db != nil {
		r.db.Close()
		r.db = nil
	}
	return nil
}
