package main

import (
	"context"
	"encoding/base64"
	"log"

	"crypto/rand"
	"crypto/sha256"

	"github.com/jackc/pgx/v5"
	"github.com/pablu23/blinkscan/backend/database"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=db user=postgres dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	accounts, err := queries.GetAccounts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(accounts)

	pwd := "test"
	salt := make([]byte, 24)
	rand.Read(salt)
	hash := sha256.Sum256(append([]byte(pwd), salt...))
	b64hash := base64.StdEncoding.EncodeToString(hash[:])
	b64salt := base64.StdEncoding.EncodeToString(salt)

	insertedAccount, err := queries.CreateAccount(ctx, database.CreateAccountParams{
		Name:          "pablu",
		Base64PwdHash: b64hash,
		Base64PwdSalt: b64salt,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(insertedAccount)
}
