package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/dylanreid7/web-servers/internal/database"
)

type DB struct {
	path string
	mu   *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type Chirp struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mu:   &sync.RWMutex{},
	}
	err := db.ensureDB()
	return db, err
}

func (db *DB) CreateChirp(body string) (Chirp, error) {
	// read the db (load db)
	dbStructure, err := db.loadDB()
	chirp := Chirp{}
	if err != nil {
		return chirp, err
	}
	// create a chirp with body, id = len of chirps + 1
	id := len(dbStructure)
	chirp = {
		ID: id,
		Body: body,
	}
	dbStructure[id] = chirp
	// write that to the db file
	err := db.writeDB(dbStructure)
	if err != nil {
		return chirp, err
	}
	return chirp, nil
}

func (db *DB) GetChirps() ([]Chirp, error) {
	// 
}

func (db *DB) createDB() error {
	// initiate a new db
	dbStructure, err := DBStructure{
		Chirps: map[int]Chirp{}
	}
	return db.writeDB(dbStructure)
}

func (db *DB) ensureDB() error {
	// check if database file exists
	// if yes, return nil
	// err = create the database file, empty
	// return err
	_, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}
	return err
}

func (db *DB) loadDB() (DBStructure, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	dbStructure := DBStructure{}
	// use os readfile to read the file
	dat, err := os.readFile(db.path)
	// check for err
	if err != nil {
		return dbStructure, err
	}
	// unmarshal json (turn into dbstructure)
	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return dbStructure, err
	}
	return dbStructure, nil
}

func (db *DB) writeDB(dbStructure DBStructure) error {
	// mu lock
	db.mu.Lock()
	defer db.mu.Unlock()
	// turn dbStructure into json
	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}
	// err = take the db structure and write it into the database.json file
	err = os.writeFile(db.path, dat, 0600)
	// return err
	return err
}
