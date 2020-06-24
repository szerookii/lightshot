package database

import (
    "database/sql"
    "time"
    "os"
    "io/ioutil"
    _ "github.com/mattn/go-sqlite3"
)

func Init() error {
    if _, err := os.Stat("database.db"); os.IsNotExist(err) {
        if err = ioutil.WriteFile("database.db", []byte(""), 0777); err != nil {
            return err
        }
    }
    
    db, err := GetDB()
    if err != nil {
        return err
    }
    
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS images(id TEXT, data TEXT, createdAt TEXT)")
    if err != nil {
        return err
    }
    
    tick := time.Tick(30 * time.Second)
        
    go func() {
        for {
            select {
                case <-tick:
                CheckDeletable()
                break
            }
        }
    }()
    
    return nil
}

func GetDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "database.db")
    if err != nil {
        return nil, err
    }
    
    return db, nil
}