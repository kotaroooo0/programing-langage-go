package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kotaroooo0/programing-language-go/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	lessThanMonth := make([]*github.Issue, 0)
	lessThanYear := make([]*github.Issue, 0)
	moreThanYear := make([]*github.Issue, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(time.Now().AddDate(0, -1, 0)):
			lessThanMonth = append(lessThanMonth, item)
		case item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)):
			lessThanYear = append(lessThanYear, item)
		default:
			moreThanYear = append(moreThanYear, item)
		}
	}

	fmt.Println("less than a month")
	for _, issue := range lessThanMonth {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("less than a year")
	for _, issue := range lessThanYear {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Println("more than a year:")
	for _, issue := range moreThanYear {
		fmt.Printf("  #%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
