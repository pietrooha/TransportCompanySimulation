package main

import ("fmt" ; "container/list" ; "time" ; /*"reflect"*/)

type PartRoute struct {
	distance, speed int
}

type Train struct {
	maxSpeed, maxPassenger int
	sourceStation, goalStation string
}

/**
* This function create route from several part route.
*/
func createRoute(distances []int, speeds []int) list.List {
	var route list.List
	var routePart PartRoute

	for i := 0; i < len(distances); i++ {
		routePart.distance = distances[i]
		routePart.speed = speeds[i]
		route.PushBack(routePart)
	}
	return route
}

/**
* This function create map contents  all routes. All trains  will use this map  in order to  transport company simulation.
*/
func createMap() (list.List, list.List, list.List, list.List, list.List, list.List, list.List, list.List) {
	// Wr - Trz i Trz - Wr
	distancesWroclawTrzebnica := []int{-1, 11, 0, 2, 0, 11, -2}
	speedsWroclawTrzebnica := []int{-1, 110, 0, 60, 0, 110, -2}
	// Trz - Obor Sl i Obor Sl - Trz
	distancesTrzebnicaObornikiSl := []int{-1, 5, 0, 1, 0, 5, -2}
	speedsTrzebnicaObornikiSl := []int{-1, 100, 0, 20, 0, 100, -2}
	// Wr - Obor Sl i Obor Sl - Wr
	distancesWroclawObornikiSl := []int{-1, 12, 0, 1, 0, 12, -2}
	speedsWroclawObornikiSl := []int{-1, 120, 0, 45, 0, 120, -2}	
	// Trz - Wol i Wol - Trz
	distancesTrzebnicaWolow := []int{-1, 16, 0, 3, 0, 16, -2}
	speedsTrzebnicaWolow := []int{-1, 160, 0, 60, 0, 160, -2}
	// Trz - Zm i Zm - Trz
	distancesTrzebnicaZmigrod := []int{-1, 10, 0, 2, 0, 10, -2}
	speedsTrzebnicaZmigrod := []int{-1, 200, 0, 40, 0, 200, -2}
	// Trz - Mil i Mil - Trz
	distancesTrzebnicaMilicz := []int{-1, 15, 0, 2, 0, 15, -2}
	speedsTrzebnicaMilicz := []int{-1, 150, 0, 40, 0, 150, -2}
	// Trz - Zaw i Zaw - Trz
	distancesTrzebnicaZawonia := []int{-1, 5, 0, 2, 0, 5, -2}
	speedsTrzebnicaZawonia := []int{-1, 50, 0, 20, 0, 50, -2}
	// Trz - Oles i Oles - Trz
	distancesTrzebnicaOlesnica := []int{-1, 15, 0, 1, 0, 15, -2}
	speedsTrzebnicaOlesnica := []int{-1, 150, 0, 20, 0, 150, -2}

	var routeWroclawTrzebnica list.List = createRoute(distancesWroclawTrzebnica, speedsWroclawTrzebnica)
	var routeTrzebnicaObornikiSl list.List = createRoute(distancesTrzebnicaObornikiSl, speedsTrzebnicaObornikiSl)
	var routeWroclawObornikiSl list.List = createRoute(distancesWroclawObornikiSl, speedsWroclawObornikiSl)
	var routeTrzebnicaWolow list.List = createRoute(distancesTrzebnicaWolow, speedsTrzebnicaWolow)
	var routeTrzebnicaZmigrod list.List = createRoute(distancesTrzebnicaZmigrod, speedsTrzebnicaZmigrod)
	var routeTrzebnicaMilicz list.List = createRoute(distancesTrzebnicaMilicz, speedsTrzebnicaMilicz)
	var routeTrzebnicaZawonia list.List = createRoute(distancesTrzebnicaZawonia, speedsTrzebnicaZawonia)
	var routeTrzebnicaOlesnica list.List = createRoute(distancesTrzebnicaOlesnica, speedsTrzebnicaOlesnica)

	return routeWroclawTrzebnica, routeTrzebnicaObornikiSl, routeWroclawObornikiSl, routeTrzebnicaWolow, routeTrzebnicaZmigrod, routeTrzebnicaMilicz, routeTrzebnicaZawonia, routeTrzebnicaOlesnica
}

/**
* This function will calculate ride time for train distance and speed 
*/
func calculateRideTime(distance int, speed int) float64 {
	if distance < 1 || speed < 1 {
		return 0.0
	} else {
		return float64(distance) / float64(speed) * 60
	}
}

/** 
* Function which check max possible speed for train in specified track.
*/
func checkSpeed(trackMaxSpeed int, trainMaxSpeed int) int {
	var speed int
	if trainMaxSpeed <= trackMaxSpeed {
		speed = trainMaxSpeed
	} else {
		speed = trackMaxSpeed
	}
	return speed
}

func setNext(r *list.Element, direction int) *list.Element {
	var next *list.Element

	if direction == 1 {
		next = r.Next()
	} else if direction == 2 {
		next = r. Prev()
	}
	return next
}

func checkDirectionAndDistance(direction int, distance int) int {
	if direction == 2 && distance == -2 {
		distance = -1
	} else if direction == 2 && distance == -1 {
		distance = -2
	}
	return distance
}

func checkRideDirection(direction int, route list.List, train Train) (*list.Element, *list.Element, string, string) {
	var front, next *list.Element
	var startStation, endStation string

	if direction == 1 {
		front = route.Front()
		next = front.Next()
		startStation = train.sourceStation
		endStation = train.goalStation
	} else if direction == 2 {
		front = route.Back()
		next = front.Prev()
		startStation = train.goalStation
		endStation = train.sourceStation
	}
	return front, next, startStation, endStation
}

func changeRideDirection(direction int, route list.List, train Train) (int, *list.Element, *list.Element, string, string) {
	var front, next *list.Element
	var startStation, endStation string

	if direction == 1 {
		front = route.Back()
		next = front.Prev()
		startStation = train.goalStation
		endStation = train.sourceStation
		direction = 2
	} else if direction == 2 {
		front = route.Front()
		next = front.Next()
		startStation = train.sourceStation
		endStation = train.goalStation
		direction = 1
	}
	return direction, front, next, startStation, endStation
}

func saveCommunicationInCh(ch chan<- string, distance int, speed int, rideTime time.Duration, startStation string, endStation string, myDiff time.Time) time.Time {
	var diff time.Time = myDiff

	if distance == 0 {
		ch <- fmt.Sprint(myDiff.Format("15:04"), " [ZWROTNICA]", " Pociąg relacji ", startStation, " - ", endStation, " wjeżdża na zwrotnicę")
		twoMinuteSimulate := time.Minute * 5
		diff = myDiff.Add(twoMinuteSimulate)
		time.Sleep(5 * time.Second)
	} else if(distance == -1) {
		ch <- fmt.Sprint(myDiff.Format("15:04"), " [WYRUSZYŁ]", " Pociąg relacji ", startStation, " - ", endStation, " wyruszył ze stacji ", startStation)
	} else if(distance == -2) {
		ch <- fmt.Sprint(myDiff.Format("15:04"), " [DOJECHAŁ]", " Pociąg relacji ", startStation, " - ", endStation, " dojechał do stacji ", endStation)
		ch <- fmt.Sprint(myDiff.Format("15:04"), " [POSTÓJ]", " Pociąg relacji ", endStation, " - ", startStation, " stoi na stacji ", endStation,  " w oczekiwaniu na pasażerów")
		layoverTimeSimulate := time.Minute * 40
		diff = myDiff.Add(layoverTimeSimulate)
		time.Sleep(40 * time.Second)
		//fmt.Println("r is a type of: ", reflect.TypeOf(diff2))
	} else {
		ch <- fmt.Sprint(myDiff.Format("15:04"), " [W TRAKCIE JAZDY]",  " Pociąg relacji ", startStation, " - ", endStation, " pokonuje dystans ", distance, " km z prędkością ", speed, " km/h")
		rideTimeSimulate := time.Minute * rideTime
		diff = myDiff.Add(rideTimeSimulate)
		time.Sleep(rideTime * time.Second)
	}
	return diff
}

func ride(route list.List, train Train, ch chan<- string, startTime float64, direction int) {
	var front, next *list.Element
	var startStation, endStation string

	then := time.Now()
	now := time.Now()
	diff := now.Sub(then)
	front, next, startStation, endStation = checkRideDirection(direction, route, train)

	for diff.Seconds() < startTime {
		diff = now.Sub(then)
		now = time.Now()
	}

	var timeDiff time.Time = now
	test := time.Minute * time.Duration(startTime)
	timeDiff = timeDiff.Add(test)
	var tDiff time.Time
	for {
		for r := front; r != nil; r = next {
			trackMaxSpeed := r.Value.(PartRoute).speed
			trainMaxSpeed := train.maxSpeed

			speed := checkSpeed(trackMaxSpeed, trainMaxSpeed)
			distance := r.Value.(PartRoute).distance

			var rideTime time.Duration = time.Duration(calculateRideTime(distance, speed))
			distance = checkDirectionAndDistance(direction, distance)
			tDiff = saveCommunicationInCh(ch, distance, speed, rideTime, startStation, endStation, timeDiff)
			timeDiff = tDiff

			next = setNext(r, direction)
		}
		direction, front, next, startStation, endStation = changeRideDirection(direction, route, train)
	}
}

func simulateTrainNetwork() {
	ch := make(chan string)
	route1, route2 , route3, route4, route5, route6, route7, route8 := createMap()
	_, _, _, _, _, _, _, _ = route1, route2, route3, route4, route5, route6, route7, route8
	var trainWrTrz Train = Train{110, 150, "WROCŁAW", "TRZEBNICA"}
	var trainTrzOborSl Train = Train{100, 200, "TRZEBNICA", "OBORNIKI ŚLĄSKIE"}
	var trainWrOborSl Train = Train{120, 200, "WROCŁAW", "OBORNIKI ŚLĄSKIE"}
	var trainTrzWol Train = Train{165, 140, "TRZEBNICA", "WOŁÓW"}
	var trainTrzZm Train = Train{210, 100, "TRZEBNICA", "ŻMIGRÓD"}
	var trainTrzMil Train = Train{150, 100, "TRZEBNICA", "MILICZ"}
	var trainTrzZaw Train = Train{85, 70, "TRZEBNICA", "ZAWONIA"}
	var trainTrzOles Train = Train{170, 150, "TRZEBNICA", "OLEŚNICA"}
	go ride(route1, trainWrTrz, ch, 5, 1)
	//go ride(route1, trainWrTrz, ch, 5, 2)
	go ride(route2, trainTrzOborSl, ch, 10, 1)
	//go ride(route2, trainTrzOborSl, ch, 10, 2)
	go ride(route3, trainWrOborSl, ch, 15, 1)
	//go ride(route3, trainWrOborSl, ch, 15, 2)
	go ride(route4, trainTrzWol, ch, 20, 1)
	//go ride(route4, trainTrzWol, ch, 20, 2)
	go ride(route5, trainTrzZm, ch, 25, 1)
	//go ride(route5, trainTrzZm, ch, 25, 2)
	go ride(route6, trainTrzMil, ch, 30, 1)
	//go ride(route6, trainTrzMil, ch, 30, 2)
	go ride(route7, trainTrzZaw, ch, 35, 1)
	//go ride(route7, trainTrzZaw, ch, 35, 2)
	go ride(route8, trainTrzOles, ch, 40, 1)
	//go ride(route8, trainTrzOles, ch, 40, 2)	
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	simulateTrainNetwork()
}