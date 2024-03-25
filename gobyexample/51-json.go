package main 

import (
	"fmt"
	"encoding/json"
	// "reflect"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
	ev bool
}

type response2 struct {
	Page int		`json:"page"`
	Fruits []string	`json:"fruits"`	
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.21)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("what's good?")
	fmt.Println(string(strB))
	
	slc := map[int]int{1: 12, 2:21, 3:45}
	slcB, _ := json.Marshal(slc)
	fmt.Println(string(slcB))


	r1 := &response1{
		Page: 20,
		Fruits: []string{"banana", "grapes", "kiwi"},
		ev: true,
	}
	r1B, _ := json.Marshal(r1)
	fmt.Println(string(r1B))
	
	r2 := &response2{
		Page: 20,
		Fruits: []string{"banana", "grapes", "kiwi"},
	}
	r2B, _ := json.Marshal(r2)
	fmt.Println(string(r2B))


	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{} // using interface since we don't know types of valye from key:value pair. so it can hold any arbitrary type

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	strs := data["strs"].([]interface{})
	str0 := strs[0].(string)
	fmt.Println(str0)

	str := `{"page": "int", "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Fruits[0], reflect.TypeOf(res.Page), res.Page)

	d1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(d1)
}