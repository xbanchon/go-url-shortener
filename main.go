package main

import "log"

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := NewStorage(cfg.dbCfg.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		cfg:     cfg,
		storage: db,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
