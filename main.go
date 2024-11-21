package main

import "log"

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := NewPostgresConn(cfg.dbCfg.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storage := NewStorage(db)

	app := &application{
		cfg:     cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
