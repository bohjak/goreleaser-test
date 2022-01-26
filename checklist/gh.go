package checklist

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func (app *appEnv) github() {
	var cmd *exec.Cmd
	var err error
	var out []byte

	fmt.Println("*** Create a release with for the latest tag ***")

	fmt.Println("Enter sprint number:")
	fmt.Scanln(&app.sprint)

	var branch string
	for branch != "master" {
		cmd = exec.Command("git", "switch", "master")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Switch to master branch and press <enter> to continue:")
			fmt.Scanln()
		}

		cmd = exec.Command("git", "branch", "--show-current")
		out, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s\n", out)
			fmt.Println(err)
		}
		branch = strings.TrimSpace(string(out))
	}

	var updated bool
	for !updated {
		cmd = exec.Command("git", "pull", "--ff-only")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Pull latest changes and press <enter> to continue:")
			fmt.Scanln()
		}

		cmd = exec.Command("git", "diff", "origin/master")
		out, err = cmd.Output()
		if err != nil {
			fmt.Printf("%s\n", out)
			log.Fatal(err)
		}
		updated = len(string(out)) == 0
	}

	cmd = exec.Command("git", "describe")
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		log.Fatal(err)
	}
	app.tag = string(out)

	// TODO: gh is being very uncooperative
	// cmd := exec.Command("gh", "--version")
	// out, err = cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("%s\n", out)
	// 	fmt.Println("GitHub CLI not available. Create a release manually and press <enter> to continue:")
	// 	fmt.Scanln()
	// 	return
	// }
	// fmt.Printf("creating release for sprint %s and tag %s...\n", app.sprint, app.tag)
	// cmd = exec.Command("gh", "release", "create", app.tag, "--title", fmt.Sprintf("Sprint %s", app.sprint))
	// out, err = cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("%s\n", out)
	// 	log.Fatal(err)
	// }

	cmd = exec.Command("git", "remote", "get-url", "origin")
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		log.Fatal(err)
	}

	fmt.Printf("Visit %s and create release for sprint %s and tag %s\n", out, app.sprint, app.tag)
	fmt.Println("press <enter> to continue:")
	fmt.Scanln()
}
