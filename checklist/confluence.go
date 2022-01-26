package checklist

import (
	"fmt"
	"time"
)

func (app *appEnv) confluence() {
	fmt.Println("Go to the applications Confluence deployments page and press <enter> to continue:")
	fmt.Scanln()

	fmt.Println("Edit the page, duplicate the topmost row, and press <enter> to continue:")
	fmt.Scanln()

	fmt.Printf("Fill in date - today is %s - and press <enter> to continue:\n", time.Now().Format("02.01.2006"))
	fmt.Scanln()

	fmt.Printf("Fill in version - current version is %s - and press <enter> to continue:\n", app.tag)
	fmt.Scanln()

	fmt.Printf("Fill in sprint - current sprint is %s - and press <enter> to continue:\n", app.sprint)
}
