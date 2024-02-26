package lib

import "fmt"

type City struct {
	Name  int
	Trips []*Trip
}

type Trip struct {
	From  *City
	To    *City
	Price int
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// n = number of cities
	// src = trip start
	// dst = trip destiny
	// k = allowed stops to take
	hashCity := buildCities(n)
	fillTrips(&flights, hashCity)
	var curr *City = nil
	// var next *City = nil
	curr, ok := (*hashCity)[src]
	level, nextPath := 0, 0
	path := map[int][]int{}
	alreadyAnEntry := false
	if ok {
		path[src] = []int{}
		for {
			for _, x := range curr.Trips {
				path[level] = append(path[level], x.To.Name)
			}
			level += 1
			nextPath = 0
			if level >= n {
				break
			}
			if len(curr.Trips) > 0 {
				for {
					_, alreadyAnEntry = path[curr.Trips[nextPath].To.Name]
					if alreadyAnEntry && (nextPath+1) < len(curr.Trips) {
						nextPath += 1
					} else {
						break
					}
				}
				if !alreadyAnEntry {
					curr = curr.Trips[nextPath].To
				} else {
					curr, ok = (*hashCity)[level]
					if !ok {
						break
					}
				}
			} else {
				curr, ok = (*hashCity)[level]
				if !ok {
					break
				}
			}
			path[level] = []int{}
		}
	}

	fmt.Println(path)

	return 0
}

func fillTrips(flights *[][]int, hashCity *map[int]*City) {
	for _, x := range *flights {
		src, ok1 := (*hashCity)[x[0]]
		dst, ok2 := (*hashCity)[x[1]]
		if ok1 && ok2 {
			src.Trips = append(src.Trips, &Trip{
				From:  src,
				To:    dst,
				Price: x[2],
			})
		}
	}
	fmt.Println(len(*flights))
}

func buildCities(n int) *map[int]*City {
	hashCity := map[int]*City{}
	for i := 0; i < n; i += 1 {
		c := City{
			Name:  i,
			Trips: []*Trip{},
		}
		hashCity[i] = &c
	}
	return &hashCity
}
