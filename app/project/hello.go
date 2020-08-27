package main

import (
	"fmt"
	"net/http"
	"log"
	"reflect"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/graphql-go/graphql"
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
    defer session.Close()
	db := session.DB("test")

    /**
     * インサート
    **/
    newPerson := &Person{
        ID:   bson.NewObjectId(),
        Name: "て",
        Age:  17,
    }
    col := db.C("people")
    if err := col.Insert(newPerson); err != nil {
        log.Fatalln(err)
    }

	/**
     * 全部取ってくる
    **/
	var allPeople []Person
    query := db.C("people").Find(bson.M{})
    query.All(&allPeople)

	/**
	 * graphQLの型定義
	**/

	var personType = graphql.NewObject( // GraphQL用の型定義
		graphql.ObjectConfig{
		  Name: "Person",
		  Fields: graphql.Fields{
			"name": &graphql.Field{
			  Type: graphql.String,
			},
			"age": &graphql.Field{
			  Type: graphql.Int,
			},
		  },
		},
	  )

	  // ここからgraphqlの設定
	  // field定義
	  fields := graphql.Fields{ // フィールド(リクエストで使われるデータの扱い方に関する設定)
		"person": &graphql.Field{
		  Type: graphql.NewList(personType),
		  Description: "Fetch person by name",
		  Args: graphql.FieldConfigArgument{ // クエリに渡す引数についての設定
			"name": &graphql.ArgumentConfig{
			  Type: graphql.String,
			},
		  },
		  Resolve: func(param graphql.ResolveParams) (interface{}, error) { // 帰って来るデータの設定
			name, ok := param.Args["name"].(string)
			if ok {
			  for _, person := range allPeople { // monngoDBから取ってきた全Personデータ
				if person.Name == name { 
				  return person, nil
				}
			  }
			}
			return nil, nil
		  },
		},
		"list": &graphql.Field{
			Type: graphql.NewList(personType),
			Description: "Fetch person list",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			  return allPeople, nil // AllUsersはmonngoDBから取ってきた全てUserのデータ
			},
		  },
	  }

	  rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: fields,
	  }
	  schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	  }
	  schema, err := graphql.NewSchema(schemaConfig) // スキーマ
	  if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	  }
	
	gqlQuery := `
	{
		person(name: "て") {
		    age
		}
	  }
	`

	// gqlQuery := `
	// {
	// 	list {
	// 	    name
	// 	    age
	// 	}
	//   }
	// `

	params := graphql.Params{
		Schema: schema, 
		RequestString: gqlQuery, 
	}
	fmt.Println("%v",reflect.TypeOf(params))
	r := graphql.Do(params) // GraphQLの実行し、rに結果を格納
	fmt.Println("%v",reflect.TypeOf(r))
	if len(r.Errors) > 0 { // エラーがあれば
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r) // 構造体をJSONに変換して、
	fmt.Printf("%s \n", rJSON) // プリントする。
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!\n")
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok!\n")
}