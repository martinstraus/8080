package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/martinstraus/8080/vm"
)

func main() {
	if len(os.Args) < 2 {
		panic("You must enter the name of the memory file to load.")
	}

	filename := os.Args[1]
	memory := loadMemoryFromFile(filename)
	machine := vm.Make8080()
	machine.Load(memory)

	fmt.Println("Running 8080 emulator.")
	defer fmt.Println("8080 emulator finished.")
	machine.Run()
	machine.Dump()
}

func loadMemoryFromFile(filename string) *vm.Memory {
	dat, err := ioutil.ReadFile(filename)
	panicIfError(err)
	if len(dat) > vm.MEMORY_SIZE {
		panic("The input file must less than 64Kb in length.")
	}
	memory := (*memoryFromBytes(&dat))
	randomize(&memory, len(dat), vm.MEMORY_SIZE)
	return &memory
}

func memoryFromBytes(bytes *[]byte) *vm.Memory {
	memory := vm.MakeMemory()
	for i, v := range *bytes {
		memory[i] = vm.Word(v)
	}
	return &memory
}

func randomize(m *vm.Memory, start int, end int) {
	for i := start; i < end; i++ {
		(*m)[i] = vm.Word(rand.Intn(256))
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
