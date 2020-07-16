package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)
	if err != nil {	
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	user, err2 := json.Marshal(u1)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(string(user))
	}
}