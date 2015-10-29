package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/doctordesh/rest/models"
	"github.com/doctordesh/rest/repositories"
	_ "github.com/go-sql-driver/mysql"
)

const (
	AUTH_HEADER = "X-Auth"
)

func main() {

	db, err := sql.Open("mysql", "root@/golang")

	if err != nil {
		panic(err)
	}

	repo := repositories.NewInviteRepository(db)

	invite := new(models.Invite)
	invite.ID = 5
	invite.CreatedAt = time.Now()
	invite.UpdatedAt = time.Now()
	invite.Code = "toto"
	invite.Token = "aabbcc-ddeeff-gghhii-jjkkll-mmnnoo"

	repo.Save(invite)

	//_, err := newDb()

	// if err != nil {
	// 	panic(err)
	// }

	port := os.Getenv("PORT")
	router := NewRouter()

	if port == "" {
		port = "8000"
	}

	fmt.Printf("Starting server on port %v\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
