package helper

import (
	"postTrainingProject/checker"
	"sync"

	"github.com/rodaine/table"
)

var waitGroup sync.WaitGroup

func GetStatusHelper(websiteList map[string]bool) {
	tbl := table.New("\t\tWebsite", "\tStatus")
	tbl.WithPadding(15)
	for website, _ := range websiteList {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			checkerObj := &checker.StatusChecker{}

			result, _ := checkerObj.CheckStatus(website)
			if result {
				tbl.AddRow("\t"+website, "UP")
			} else {
				tbl.AddRow("\t"+website, "DOWN")
			}
		}()
		waitGroup.Wait()
	}
	tbl.Print()
}
