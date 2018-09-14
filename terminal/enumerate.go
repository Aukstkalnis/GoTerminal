package terminal

import (
	"github.com/albenik/go-serial/enumerator"
)

func GetPortList() []string {
	list, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return nil
	}
	availablePorts := make([]string, len(list))
	for i, v := range list {
		availablePorts[i] = v.Name
		//fmt.Printf("%+v\n", v)
	}
	return availablePorts
}
