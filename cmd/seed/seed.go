package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/yannis94/kagami/handlers"
	"github.com/yannis94/kagami/store"
	"golang.org/x/crypto/ssh"
)

func init() {
    err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
    fmt.Println("Kagami - seed script")

    username := flag.String("username", "", "Create admin username")
    pwd := flag.String("password", "", "Create admin password")
    filePath := flag.String("key", "", "SSH public key file path")

    flag.Parse()
    db := store.NewSQLiteStorage()
    db.Init()

    if err := isValidPubKey(*filePath); err != nil {
        panic(err)
    }

    adminHandler := handlers.NewAdminHandler(db)
    if err := adminHandler.CreateAdmin(*username, *pwd); err != nil {
        panic(err)
    }

    fmt.Println("Database seeded!")
}

func isValidPubKey(filePath string) error {
    keyBytes, err := ioutil.ReadFile(filePath)

    if err != nil {
        return err
    }

    pubKey, _, _, _, err := ssh.ParseAuthorizedKey(keyBytes)

    if err != nil {
        return err
    }

    fmt.Printf("Public key authorized.\nKey type: %s\nKey: %s", pubKey.Type(), string(keyBytes))

    return nil
}
