package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)


func main(){
 mux := http.NewServeMux()
 mux.HandleFunc("/",GetAirtime)

 server := &http.Server{
	Handler: mux,
	Addr: ":8000",
 }
 
 if err := server.ListenAndServe(); err != nil{
	fmt.Print(err)
 }
}


type AirtimeResponse struct{
	TotalAmount string `json:"totalAmount" binding:"required"`
	ErrorMessage string `json:"errorMessage"`
	TotalDiscount string `json:"totalDiscount" binding:"required"`
	Responses  []Response `json:"responses"`

}
type Response struct{
	PhoneNumber string `json:"phoneNumber"`
	Status string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
 }

 
func GetAirtime(w http.ResponseWriter, r *http.Request){
	key := "32b810b22ebe6865d01bf597b2fd5028511122d880d0a1f4bf93486e0c0bb9a6"
	burl := "https://api.sandbox.africastalking.com/version1/airtime/send"

	type Recipient struct{
		PhoneNumber string `json:"phoneNumber"`
		Amount string `json:"amount"`
	}

	var users []*Recipient

	pp,_ := io.ReadAll(r.Body)

	if err := json.Unmarshal(pp, &users); err != nil{
		if err != nil{
			fmt.Print(err)
		}
	}

	jsonresp,err := json.Marshal(users)
	if err != nil{
		fmt.Print(err)
	}
		values := url.Values{}
		values.Add("username","sandbox")
		values.Add("recipients",string(jsonresp))


	req, err := http.NewRequest(http.MethodPost, burl, strings.NewReader(values.Encode()))
	if err != nil{
		fmt.Print(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept" ,"application/json")
	req.Header.Set("apiKey", key)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil{
		fmt.Print(err)
	}
	bodybyte, _ := ioutil.ReadAll(resp.Body)

	var tt *AirtimeResponse
	if err := json.Unmarshal(bodybyte, &tt); err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v", tt)
	w.Write(bodybyte)
}
