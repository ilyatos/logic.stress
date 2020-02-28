package helpers

import (
	"fmt"
	"github.com/ilyatos/logic.stress/pkg/client"
)

func PrintLabState(u *client.User, launch int, ls *client.LabStatus) {
	fmt.Printf("User: %s and Launch: %d \n\tState: %s \n\tStatus %d\n", u.Subdomain, launch, ls.State, ls.Status)
}

func PrintStructureWithFields(v interface{}) {
	fmt.Printf("%+v\n", v)
}
