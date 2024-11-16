package random_user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset     string `json:"offset"`
	Description string `json:"description"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Location struct {
	Street    Street    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Postcode  interface{} `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone  Timezone  `json:"timezone"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Dob struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Login struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Gender   string   `json:"gender"`
	Name     Name     `json:"name"`
	Location Location `json:"location"`
	Email    string   `json:"email"`
	Login    Login    `json:"login"`
	Dob      Dob      `json:"dob"`
	Phone    string   `json:"phone"`
	Cell     string   `json:"cell"`
	Id       struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"id"`
	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	Nat string `json:"nat"`
}

type RandomUserResponse struct {
	Results []User `json:"results"`
	Info    struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}

func GetRandomUser() (string, string, string) {

	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		fmt.Println("Error:", err)
		return "","",""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "","",""
	}

	var result RandomUserResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "","", ""
	}
	location := result.Results[0].Location
	return location.Country, location.City, location.Street.Name
}