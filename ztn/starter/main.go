package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

type serversType map[string]int

func printServersStatus(servers serversType) {
	fmt.Printf("There are %d servers\n", len(servers))
	fmt.Println(servers)
	stats := make(map[int]int)

	for _, status := range servers {

		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Unknown server")
		}

	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omricon", "w359", "baseline"}

	serversStatus := make(serversType)

	for _, server := range servers {
		serversStatus[server] = Online
	}
	fmt.Println("--------------------------------")
	fmt.Println(servers)
	fmt.Println(serversStatus)
	fmt.Println("--------------------------------")
	printServersStatus(serversStatus)
	fmt.Println("--------------------------------")
	serversStatus["darkstar"] = Retired
	serversStatus["aiur"] = Offline
	printServersStatus(serversStatus)

	for k, _ := range serversStatus {
		serversStatus[k] = Maintenance
	}
	printServersStatus(serversStatus)

}
