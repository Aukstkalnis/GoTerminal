package terminal

import (
	"runtime"
	"sort"
	"strconv"

	"github.com/albenik/go-serial/enumerator"
)

// GetPortList returns available port list
func GetPortList() []string {
	list, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return nil
	}
	availablePorts := make([]string, len(list))
	for i, v := range list {
		availablePorts[i] = v.Name
	}
	switch runtime.GOOS {
	case "windows":
		// Sort windows serial ports
		sort.Slice(availablePorts, func(i, j int) bool {
			numA, _ := strconv.Atoi(availablePorts[i][3:])
			numB, _ := strconv.Atoi(availablePorts[j][3:])
			return numA < numB
		})
	case "linux":
		// Sort linux serial ports
		return []string{}
	}
	return availablePorts
}
