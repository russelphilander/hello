package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name    string `bson: "name"`
	Surname string `bson: "surname"`
	Age     int    `bson: "age"`
}

type Test struct {
	school string
	course string
	mark   string
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydb").C("Person")

	/*
		err = c.Insert(&Person{Name: "Pienut", Surname: "Butter"})

		if err != nil {
			panic(err)
		}
	*/

	result := new(Person) //bson.M{}
	iter := c.Find(bson.M{}).Iter()

	for iter.Next(result) {
		fmt.Printf("Result: %v\n", result)

		if iter.Err() != nil {
			panic(iter.Err())
		}
	}

	fmt.Println(result)

	testing := new(Test)
	testing.school = "Laerskool Darling"
	testing.course = "Maths"
	testing.mark = "87"

	fmt.Println(testing)
}
