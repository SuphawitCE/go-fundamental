package main

import "fmt"

// TODO: Demonstrate Inheritance by Composition in Go

type CPU struct {
	architecture string
}

func (cpu *CPU) setArchitecture(arch string) {
	cpu.architecture = arch
}

func (cpu *CPU) getArchitecture() string {
	return cpu.architecture
}

type RAM struct {
	size int
}

func (ram *RAM) setSize(size int) {
	ram.size = size
}

func (ram *RAM) getSize() int {
	return ram.size
}

type Motherboard struct {
	category string
}

func (m *Motherboard) setCategory(cat string) {
	m.category = cat
}

func (m *Motherboard) getCategory() string {
	return m.category
}

type Computer struct {
	cpu    CPU
	ram    RAM
	mboard Motherboard
}

func (comp *Computer) setSpecification(cpu CPU, ram RAM, mboard Motherboard) {
	comp.cpu.setArchitecture(cpu.getArchitecture())
	comp.ram.setSize(ram.getSize())
	comp.mboard.setCategory(mboard.getCategory())
}

func (comp *Computer) showSpecification() {
	fmt.Println("CPU:", comp.cpu.getArchitecture(), ", RAM: ", comp.ram.getSize(), "GB, Motherboard: ", comp.mboard.getCategory())
}

func main() {
	cpu := CPU{}
	cpu.setArchitecture("64-bit")

	ram := RAM{}
	ram.setSize(8)

	mboard := Motherboard{}
	mboard.setCategory("Micro ATX")

	computer1 := Computer{}
	computer1.setSpecification(cpu, ram, mboard)
	computer1.showSpecification()
}
