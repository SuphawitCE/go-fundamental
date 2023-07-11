package main

import "fmt"

//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

// * Create functions to calculate averages for each dashboard component
func (bandWith *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(bandWith.amount); i++ {
		sum += int(bandWith.amount[i])
	}
	return sum / len(bandWith.amount)
}

func (cpu *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(cpu.temp); i++ {
		sum += int(cpu.temp[i])
	}
	return sum / len(cpu.temp)
}

func (memory *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(memory.amount); i++ {
		sum += int(memory.amount[i])
	}
	return sum / len(memory.amount)
}

//   - Using struct embedding, create a Dashboard structure that contains
//     each dashboard component
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}
	//* Print out a 5-second average from each component using promoted
	//  methods on the Dashboard
	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageBandwidth())
	fmt.Printf("Average temp usage: %v\n", dashboard.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dashboard.AverageMemoryUsage())
}
