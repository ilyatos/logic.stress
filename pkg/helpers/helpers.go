package helpers

import (
	"fmt"
	"github.com/ilyatos/logic.stress/pkg/client"
)

func PrintLabState(u *client.User, ls *client.LabStatus) {
	fmt.Printf("User: %s; \n\tState: %s; \n\tStatus %d\n", u.Subdomain, ls.State, ls.Status)
}

func PrintStructureWithFields(v interface{}) {
	fmt.Printf("%+v\n", v)
}
