package main

import (
	"fmt"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
    ID   bson.ObjectId `bson:"_id"`
    Name string        `bson:"name"`
    Age  int           `bson:"age"`
}

func main() {
	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/ok", okHandler)
	// http.ListenAndServe(":8080", nil)

	session, _ := mgo.Dial("mongo-db:27017")
    // defer session.Close()
	db := session.DB("test")

    /**
     * つくるところ
    **/
    mai := &Person{
        ID:   bson.NewObjectId(),
        Name: "まいまい",
        Age:  17,
    }
    col := db.C("people")
    if err := col.Insert(mai); err != nil {
        log.Fatalln(err)
    }
	db.C("people")

	/**
     * みつけるところ
    **/
    p := new(Person)
    query := db.C("people").Find(bson.M{})
    query.One(&p)

    /**
     * 結果
    **/
    fmt.Printf("%+v\n", p)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!\n")
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok!\n")
}