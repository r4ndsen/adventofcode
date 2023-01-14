package main

import (
	"github.com/r4ndsen/adventofcode/support"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}.*)`)

	valveCache := make(ValveCache)

	for _, line := range support.GetInputFor(16) {
		if len(line) == 0 {
			continue
		}

		res := re.FindAllStringSubmatch(string(line), -1)

		v := valveCache.get(res[0][1])
		v.flowRate = support.ToInt(res[0][2])

		for _, l := range strings.Split(res[0][3], ", ") {
			refValve := valveCache.get(l)
			v.Add(refValve)
		}
	}
}
