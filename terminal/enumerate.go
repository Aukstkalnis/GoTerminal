package terminal

import (
	"sort"
	"strconv"
	"strings"

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
	// sort com ports
	if strings.Contains(availablePorts[0], "COM") {
		// Sort windows serial ports
		sort.Slice(availablePorts, func(i, j int) bool {
			numA, _ := strconv.Atoi(availablePorts[i][3:])
			numB, _ := strconv.Atoi(availablePorts[j][3:])
			return numA < numB
		})
	}
	return availablePorts
}
