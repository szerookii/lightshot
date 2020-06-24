package database

import (
    "encoding/base64"
    "math/rand"
    "time"
    "log"
    "github.com/Seyz123/lightshot/config"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Image struct {
    Id string
    Data []byte
    CreatedAt time.Time
}

func Save(data []byte) (string, error) {
    id := GenerateId(12)
    
    for Exists(id) {
        id = GenerateId(12) // bad but.. yeah ?
    }
    
    str := base64.StdEncoding.EncodeToString(data)
    
    db, err := GetDB()
    if err != nil {
        return "", err
    }
    defer db.Close()
    
    stmt, err := db.Prepare("INSERT INTO images(id, data, createdAt) VALUES (?, ?, ?)")
    if err != nil {
        return "", err
    }
    defer stmt.Close()
    
    _, err = stmt.Exec(id, str, time.Now().Format(time.RFC822))
    if err != nil {
        return "", err
    }
    
    return id, nil
}

func Exists(id string) bool {
    db, err := GetDB()
    if err != nil {
        return false
    }
    defer db.Close()
    
    stmt, err := db.Prepare("SELECT exists (SELECT * FROM images WHERE id=?)")
    if err != nil {
        return false
    }
    defer stmt.Close()
    
    row := stmt.QueryRow(id)
    
    var exists bool
    err = row.Scan(&exists)
    if err != nil {
        return false
    }
    
    return exists
}

func Get(id string) (*Image, error) {
    db, err := GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()
    
    stmt, err := db.Prepare("SELECT id, data, createdAt FROM images WHERE id=?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    
    row := stmt.QueryRow(id)
    if err != nil {
        return nil, err
    }
    
    var image Image
    var dataStr string
    var createdAt string
    
    err = row.Scan(&image.Id, &dataStr, &createdAt)
    if err != nil {
        return nil, err
    }
    
    bytes, err := base64.StdEncoding.DecodeString(dataStr)
    if err != nil {
        return nil, err
    }
    
    time, err := time.Parse(time.RFC822, createdAt)
    if err != nil {
        return nil, err
    }
    
    image.CreatedAt = time
    image.Data = bytes
    
    return &image, nil
}

func GetAll() ([]Image, error) {
    db, err := GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()
    
    stmt, err := db.Prepare("SELECT id, data, createdAt FROM images")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    
    rows, err := stmt.Query()
    if err != nil {
        return nil, err
    }
    
    var images []Image
    
    for rows.Next() {
        var image Image
        var dataStr string
        var createdAt string
        
        err = rows.Scan(&image.Id, &dataStr, &createdAt)
        if err != nil {
            return nil, err
        }
        
        bytes, err := base64.StdEncoding.DecodeString(dataStr)
        if err != nil {
            return nil, err
        }
        
        time, err := time.Parse(time.RFC822, createdAt)
        if err != nil {
            return nil, err
        }
        
        image.CreatedAt = time
        image.Data = bytes
        
        images = append(images, image)
    }
    
    return images, nil
}

func Delete(id string) error {
    db, err := GetDB()
    if err != nil {
        return err
    }
    defer db.Close()
    
    stmt, err := db.Prepare("DELETE FROM images WHERE id=?")
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    
    return nil
}

func CheckDeletable() {
    images, err := GetAll()
    if err != nil {
        return
    }
    
    config, err := config.GetConfig()
    if err != nil {
        return
    }
    
    for _, image := range images {
        diff := time.Now().Sub(image.CreatedAt)
        
        if diff.Hours() > 24 {
            err = Delete(image.Id)
            if err == nil {
                if config.Logs {
                    log.Println("Deleted screenshot with ID : " + image.Id)
                }
            }
        }
    }
}

func GenerateId(n int) string {
    b := make([]byte, n)
    
    for i := range b {
        b[i] = chars[rand.Intn(len(chars))]
    }
    
    return string(b)
}