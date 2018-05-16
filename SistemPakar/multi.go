package main

import (
	"fmt"
	"time"
)

type IPAddressCounter struct {
	ip    string
	time  time.Time
	count int
}

func (i *IPAddressCounter) IncrementCount() {
	i.count++
}

func makeCounter(ip string) IPAddressCounter {
	return IPAddressCounter{
		ip:   ip,
		time: time.Now(),
	}
}

func main() {
	mapOfIPCounters := make(map[string]IPAddressCounter)

	mapOfIPCounters["192.168.1.1"] = makeCounter("192.168.1.1")
	mapOfIPCounters["192.168.1.2"] = makeCounter("192.168.1.2")
	mapOfIPCounters["192.168.1.3"] = makeCounter("192.168.1.3")

	for key, value := range mapOfIPCounters {
		value.IncrementCount()
		mapOfIPCounters[key] = value

		fmt.Println("The counter for "+key+" is", mapOfIPCounters[key].count)
	}

}
