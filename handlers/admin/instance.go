package admin

import (
	"cwc/admin"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.AdminDeleteInstance(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleAddInstance(user_email *string, name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	created_instance, err := client.AdminAddInstance(*user_email, *name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}

func HandleUpdateInstance(id *string, status *string) {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.AdminUpdateInstance(*id, *status)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully updated\n", *id)

}

func HandleGetInstances() {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instances, err := client.AdminGetAllInstances()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tpublic ip\tproject_id\n")
	for _, instance := range *instances {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.CreatedAt, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

	}

	return
}

func HandleGetInstance(id *string) {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instance, err := client.GetInstance(*id)

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tpublic ip\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.CreatedAt, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)


	return
}
