package main

import (
	"flag"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
	"github.com/r4ndsen/adventofcode/support"
	"math"
	"regexp"
	"strings"
	"sync"
)

var seedsRe = regexp.MustCompile(`seeds:( *(\d+))+`)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(support.GetInput())
		support.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(support.GetInput())
		support.CopyToClipboard(fmt.Sprintf("%v", ans))

		fmt.Println("Output:", ans)
	} else {
		ans := sample()
		fmt.Println("Output:", ans)
	}
}

func sample() int {
	input := support.InputType(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`)

	return part2(input)
}

func part1(input support.Input) int {
	return 0
}

func part2(input support.Input) int {

	mapdata := strings.Split(input.String(), "\n\n")

	re := regexp.MustCompile(`seeds: (.*)`)
	seedInput := strings.Split(re.FindStringSubmatch(mapdata[0])[1], " ")

	seeds := make(map[int]int)

	for idx, seed := range seedInput {
		if idx%2 == 0 {
			seeds[cast.ToInt(seed)] = cast.ToInt(seedInput[idx+1]) - 1
		}
	}

	boundaryGroups := make([][]*Boundary, 0)
	for _, m := range mapdata[1:] {
		boundaries := make([]*Boundary, 0)
		var name string
		for i, n := range strings.Split(m, "\n") {
			if i == 0 {
				name = n[:len(n)-5]
				continue
			}

			if n == "" {
				continue
			}

			parts := strings.Split(n, " ")

			boundaries = append(boundaries, NewBoundary(name, cast.ToInt(parts[0]), cast.ToInt(parts[1]), cast.ToInt(parts[2])))
		}
		boundaryGroups = append(boundaryGroups, boundaries)
	}

	result := math.MaxInt

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	var iteration int

	for lowerB, upperB := range seeds {
		wg.Add(1)
		go func(lowerB, upperB int, mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			for seed := lowerB; seed <= lowerB+upperB; seed++ {
				v := seed

			bgloop:
				for _, bG := range boundaryGroups {
					for _, b := range bG {
						iteration++
						if iteration%1000000000 == 0 {
							fmt.Println("iteration", iteration)
						}
						if b.Hit(v) {
							v = b.Convert(v)
							continue bgloop
						}
					}
				}

				mu.Lock()
				if v < result {
					result = v
				}
				mu.Unlock()
			}
		}(lowerB, upperB, mu, wg)
	}

	wg.Wait()

	fmt.Println("result:", result)

	return 0
}
