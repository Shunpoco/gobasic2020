package github

import (
	"ch4/ex11/editor"
	"fmt"
	"log"
	"os"
)

func Create(repo string) {
	b, err := editor.FetchInputFromEditor()
	if err != nil {
		log.Fatal("create: something went wrong...")
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func Search(repo string, issueNum int) {
	issue, err := search(repo, issueNum)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(issue)
}

func Update(repo string, issueNum int) {
	b, err := editor.FetchInputFromEditor()
	if err != nil {
		log.Fatal("update: something went wrong...")
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func Close(repo string, issueNum int) {
	fmt.Println(repo, issueNum)
}
