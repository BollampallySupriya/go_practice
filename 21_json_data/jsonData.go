package main

import (
	"encoding/json"
	"fmt"
)

// json:"name to be reflected in json format"
// json:"-" this will hide the field from the json result
//json:"omitempty"  will not show null values in the json result.

type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Encoding JSON!!!")
	DecodeJson()
}

func EncodeJson() {
	lco_courses := []course{
		{"reactjs", 499, "Youtube", "*1234*", []string{"WebDevelopment", "Software"}},
		{"golang", 299, "LearnCodeOnline", "368528735", []string{"FullStack", "Engineering"}},
		{"Django", 399, "GeeksForGeeks", "34678936", nil},
	}
	fmt.Printf("Type of lco_courses is %T\n  value is %v\n ",lco_courses, lco_courses)

	// finalCourses, err := json.Marshal(lco_courses)

	finalCourses, err := json.MarshalIndent(lco_courses, "", "\t")
	HandleNilError(err)
	fmt.Printf("Type of final is %T\n  value is %s\n ",finalCourses, finalCourses)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "golang",
                "Price": 299,
                "website": "LearnCodeOnline",
                "tags": ["FullStack", "Engineering"]
    	}
	`)
	var lcocourse course 
	isJsonValid := json.Valid(jsonDataFromWeb)
	if isJsonValid {
		fmt.Println("JSON is valid.")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)

	} else {
		fmt.Println("JSON is not valid!!!")
	}


	// some cases where you just want to add the key value pairs.

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData) // copies json data to myonline data
	// fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}
}

func HandleNilError(err error) {
	if err != nil {
		panic(err)
	}
}