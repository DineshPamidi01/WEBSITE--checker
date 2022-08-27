package main

import (
	"JOSH/checker"
	"JOSH/helper"
	"fmt"
	"net/http"
	"time"
)

var websiteList = make(map[string]bool)

func main() {

	http.HandleFunc("/GetList", getListHandler)

	http.HandleFunc("/CheckStatus", checkStatusHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	websiteList, err = helper.GetList(r)

	if err != nil {
		fmt.Println("\t400 BAD REQUEST\n...........Failed to retrieve the WebsiteList :(...........")
	} else {
		fmt.Fprint(w, websiteList)

		fmt.Println("\t200 OK\n...........Successfully retrieved the WebsiteList :)...........")

	}
}

func checkStatusHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	optionalQuery := r.Form.Get("name")
	if optionalQuery == "" {
		for {
			fmt.Println(fmt.Sprintf("\n\n" + time.Now().String() + "\n\n"))
			helper.GetStatusHelper(websiteList)
			time.Sleep(15 * time.Second)
		}
	} else {
		fmt.Println(fmt.Sprintf("\nObtained Optional Query parameter is " + optionalQuery))
		_, ok := websiteList[optionalQuery]
		if ok {
			checkerObj := checker.StatusChecker{}
			res, _ := checkerObj.CheckStatus(optionalQuery)
			if res {
				fmt.Println(fmt.Sprintf("The Website " + optionalQuery + " is UP"))
			} else {
				fmt.Println(fmt.Sprintf("The Website " + optionalQuery + " is DOWN"))
			}
		} else {
			fmt.Println("Sorry!... No such Website Provided... :(")
		}
	}
}
