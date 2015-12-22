package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"reflect"
	"strconv"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
        "google.golang.org/appengine/user"
)

type Person struct {
	FirstName string
	LastName  string
}

func init() {

	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}


	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// receive form submission

	         c := appengine.NewContext(req)
                 u := user.Current(c)
	         key := datastore.NewKey(ctx, "Name", fName, 0, nil)
	         var name Name
	         err = datastore.Get(c, key, &name)
	         if err != nil {
		 profile.Email = u.Email
	         }
		fName := req.FormValue("first")
		lName := req.FormValue("last")
		fmt.Println("fName: ", fName)
		fmt.Println("[]byte(fName): ", []byte(fName))
		fmt.Println("typeOf: ", reflect.TypeOf(fName))

		// execute template
		err = tpl.Execute(res, Person{fName, lName})
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
