package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string, deposit int64) *Account {
	acc, err := NewAccount(fname, lname, pw, deposit)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account =>", acc.AccountNumber)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Henry", "Cavil", "1233", 1_000)
}

func main() {
	seed := flag.Bool("seed", false, "seed db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding database")
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.Run()

}
