package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleAddProject(user_email *string, name *string, host *string, token *string, git_username *string, namespace *string) {

	c, _ := admin.NewClient()
	created_project, err := c.AdminAddProject(*user_email, *name, *host, *token, *git_username, *namespace)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_project)
	} else {
		utils.PrintRow(*created_project)
	}

}

func HandleDeleteProject(id *string, name *string, url *string) {
	admin, _ := admin.NewClient()
	if *id != "" {
		err := admin.AdminDeleteProjectById(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else if *name != "" {
		err := admin.AdminDeleteProjectByName(*name)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else {
		err := admin.AdminDeleteProjectByUrl(*url)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("project successfully deleted\n")
}

func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	c, _ := admin.NewClient()
	var err error
	if *project_id == "" && *project_name == "" && *project_url == "" {
		projects, err := c.AdminGetAllProjects()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		if admin.GetDefaultFormat() == "json" {
			utils.PrintJson(projects)
		} else {
			utils.PrintMultiRow(admin.Project{}, *projects)
		}

	} else {
		project := &admin.Project{}
		if *project_id != "" {
			project, err = c.AdminGetProjectById(*project_id)
		} else if *project_name != "" {
			project, err = c.AdminGetProjectByName(*project_name)

		} else {
			project, err = c.AdminGetProjectByUrl(*project_url)
		}

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		if admin.GetDefaultFormat() == "json" {
			utils.PrintJson(project)
		} else {
			utils.PrintRow(*project)
		}

	}

	return
}
