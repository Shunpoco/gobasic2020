package application

import (
	"flag"
	"fmt"
	"os"

	"ch4/ex11/github"
)

func Application() {
	// define subcommand
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	closeCmd := flag.NewFlagSet("close", flag.ExitOnError)

	// set options
	createRepo := createCmd.String("repo", "", "repository")
	searchRepo := searchCmd.String("repo", "", "repository")
	updateRepo := updateCmd.String("repo", "", "repository")
	closeRepo := closeCmd.String("repo", "", "repository")

	searchNum := searchCmd.Int("issue", 0, "issue number")
	updateNum := updateCmd.Int("issue", 0, "issue number")
	closeNum := closeCmd.Int("issue", 0, "issue number")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands: 'create', 'search', 'update' and 'close'.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		// Issue作成時の挙動を書く
		createCmd.Parse(os.Args[2:])
		github.Create(*createRepo)
	case "search":
		// Issue検索時の挙動を書く
		searchCmd.Parse(os.Args[2:])
		github.Search(*searchRepo, *searchNum)
	case "update":
		// Issue更新時の挙動を書く
		updateCmd.Parse(os.Args[2:])
		github.Update(*updateRepo, *updateNum)
	case "close":
		// Issueクローズ時の挙動を書く
		closeCmd.Parse(os.Args[2:])
		github.Close(*closeRepo, *closeNum)
	case "default":
		fmt.Println("expected subcommands: 'create', 'search', 'update' and 'close'.")
		os.Exit(1)
	}
}
