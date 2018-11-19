package vm

import (
	"fmt"
)

type VirtualMachine struct {
	Regs    Registers
	Mem     Memory
	PC      uint16 // Program counter
	SR      StatusRegister
	SP      uint16 // Stack pointer
	Devices map[uint8]Device
	INTE    bool // interrupts enabled / disabled
}

func (vm *VirtualMachine) Load(mem *Memory) {
	vm.Mem = *mem
}

func (vm *VirtualMachine) Run() (err error) {
	err = nil
	halt := false
	vm.PC = 0
	for !halt && vm.PC <= vm.Mem.LastAddress() {
		opcode := vm.Mem[vm.PC]
		if isINR(opcode) {
			vm.inr(opcode)
		} else if isDCR(opcode) {
			vm.dcr(opcode)
		} else if isRST(opcode) {
			vm.jumpTo(uint16(opcode & 0x38))
		} else if isPUSH(opcode) {
			vm.push(opcode)
		} else if isPOP(opcode) {
			vm.pop(opcode)
		} else if isMOV(opcode) {
			vm.mov(opcode)
		} else if isAccumulatorInstr(opcode) {
			vm.onAccumulator(opcode)
		} else if isIO(opcode) {
			vm.io(opcode)
		} else if isLXI(opcode) {
			vm.lxi(opcode)
		} else {
			switch opcode {
			case NOP:
			case HLT:
				halt = true
			case MVIA:
				vm.PC = vm.PC + 1
				vm.Regs.A.load(vm.Mem[vm.PC])
			case MVIB:
				vm.PC = vm.PC + 1
				vm.Regs.B.load(vm.Mem[vm.PC])
			case MVIC:
				vm.PC = vm.PC + 1
				vm.Regs.C.load(vm.Mem[vm.PC])
			case MVID:
				vm.PC = vm.PC + 1
				vm.Regs.D.load(vm.Mem[vm.PC])
			case MVIE:
				vm.PC = vm.PC + 1
				vm.Regs.E.load(vm.Mem[vm.PC])
			case MVIH:
				vm.PC = vm.PC + 1
				vm.Regs.H.load(vm.Mem[vm.PC])
			case MVIL:
				vm.PC = vm.PC + 1
				vm.Regs.L.load(vm.Mem[vm.PC])
			case CMA:
				vm.Regs.A.NOT()
			case JMP:
				vm.PC = vm.getAddressFromNextAddress()
			case JNZ:
				if !vm.SR.Zero {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case JZ:
				if vm.SR.Zero {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case JNC:
				if !vm.SR.Carry {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case JC:
				if vm.SR.Carry {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case JP:
				if !vm.SR.Sign {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case JM:
				if vm.SR.Sign {
					vm.PC = vm.getAddressFromNextAddress()
				} else {
					vm.PC = vm.PC + 2
				}
			case CALL:
				vm.callNextSubroutine()
			case CNZ:
				if !vm.SR.Zero {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case CZ:
				if vm.SR.Zero {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case CNC:
				if !vm.SR.Carry {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case CC:
				if vm.SR.Carry {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case CP:
				if !vm.SR.Sign {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case CM:
				if vm.SR.Sign {
					vm.callNextSubroutine()
				} else {
					vm.PC = vm.PC + 3
				}
			case RET:
				vm.PC = vm.popPCFromStack()
			case RNZ:
				if !vm.SR.Zero {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case RZ:
				if vm.SR.Zero {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case RNC:
				if !vm.SR.Carry {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case RC:
				if vm.SR.Carry {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case RP:
				if !vm.SR.Sign {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case RM:
				if vm.SR.Sign {
					vm.PC = vm.popPCFromStack()
				} else {
					vm.PC = vm.PC + 1
				}
			case XCHG:
				vm.exchangeDEandHL()
			case XTHL:
				vm.exchangeHLandSP()
			case EI:
				vm.enableInterrupts()
			case DI:
				vm.disableInterrupts()
			default:
				err = fmt.Errorf("Invalid OPCODE %X in address %X", opcode, vm.PC)
				halt = true
			}
		}
		// Simple way of incrementing PC in all opcodes but conditionals
		if opcode != JMP && opcode != JNZ && opcode != JZ && opcode != JNC &&
			opcode != JC && opcode != JP && opcode != JM && opcode != CALL &&
			opcode != CNZ && opcode != CZ && opcode != RET &&
			opcode != HLT && opcode != CNC && opcode != CC && opcode != CPO &&
			opcode != CPE && opcode != CP && opcode != CM && opcode != RNZ &&
			opcode != RZ && opcode != RNC && opcode != RC && opcode != RP &&
			opcode != RM && !isRST(opcode) && !isIO(opcode) && !isLXI(opcode) {
			vm.PC = vm.PC + 1
		}
	}
	return
}

func isINR(opcode Word) bool {
	return (uint8(opcode) & 0xC7) == 0x04
}

func isDCR(opcode Word) bool {
	return (uint8(opcode) & 0xC7) == 0x05
}

func isMOV(opcode Word) bool {
	// We need to test for != HLT because the bits 6 and 7 are the same for MOV and HLT
	return opcode != HLT && (uint8(opcode)&0xC0) == 0x40
}

func isRST(opcode Word) bool {
	return (uint8(opcode) & 0xC7) == 0xC7
}

func isPUSH(opcode Word) bool {
	return (uint8(opcode) & 0xCF) == 0xC5
}

func isPOP(opcode Word) bool {
	return (uint8(opcode) & 0xCF) == 0xC1
}

func isAccumulatorInstr(opcode Word) bool {
	return uint8(opcode)&0xC0 == 0x80
}

func isIO(opcode Word) bool {
	return uint8(opcode)&0xF7 == 0xD3
}

func isLXI(opcode Word) bool {
	return uint8(opcode)&0xCF == 0x01
}

func (vm *VirtualMachine) dump() {
	fmt.Printf("Mem: %d bytes PC=%X SP=%X S=%t;C=%t;Z=%t regs=%s HL=%X\n", len(vm.Mem), vm.PC, vm.SP, vm.SR.Sign, vm.SR.Carry, vm.SR.Zero, vm.Regs, vm.Regs.HL())
}

func (vm *VirtualMachine) inr(opcode Word) {
	// bits 3, 4 and 5 represent the register to increment
	switch opcode & 0x38 {
	case 0:
		vm.Regs.B.inc()
	case 8:
		vm.Regs.C.inc()
	case 16:
		vm.Regs.D.inc()
	case 24:
		vm.Regs.E.inc()
	case 32:
		vm.Regs.H.inc()
	case 40:
		vm.Regs.L.inc()
	case 48:
		vm.Mem.inc(vm.Regs.HL())
	case 56:
		vm.Regs.A.inc()
	}
}

func (vm *VirtualMachine) dcr(opcode Word) {
	// bits 3, 4 and 5 represent the register to increment
	switch opcode & 0x38 {
	case 0:
		vm.Regs.B.dec()
	case 8:
		vm.Regs.C.dec()
	case 16:
		vm.Regs.D.dec()
	case 24:
		vm.Regs.E.dec()
	case 32:
		vm.Regs.H.dec()
	case 40:
		vm.Regs.L.dec()
	case 48:
		vm.Mem.dec(vm.Regs.HL())
	case 56:
		vm.Regs.A.dec()
	}
}

func (vm *VirtualMachine) onAccumulator(opcode Word) {
	reg := vm.reg(uint8(opcode) & 0x07) // bits 0, 1, and 2 codify the register
	switch uint8(opcode) & 0x38 {       // Bits 3, 4, and 5 codify the operation on the accumulator
	case add:
		vm.addSetFlags(reg)
	case adc:
		// not implemented
	case sub:
		vm.subtractAndSetFlags(reg)
	case sbb:
		// not implemented
	case ana:
		vm.Regs.A.AND(reg)
	case xra:
		vm.Regs.A.XOR(reg)
	case ora:
		vm.Regs.A.OR(reg)
	case cmp:
		vm.cmp(&vm.Regs.A, reg)
	}
}

const (
	add = uint8(0)
	adc = uint8(1 << 3)
	sub = uint8(2 << 3)
	sbb = uint8(3 << 3)
	ana = uint8(4 << 3)
	xra = uint8(5 << 3)
	ora = uint8(6 << 3)
	cmp = uint8(7 << 3)
)

func (vm *VirtualMachine) cmp(a *Register, b *Register) {
	vm.SR.Zero = a.Value == b.Value
	vm.SR.Sign = b.Value > a.Value
	// TODO Count bits and set vm.SR.Parity
}

func (vm *VirtualMachine) getAddressFromNextAddress() uint16 {
	return vm.readDoubleWordFrom(vm.PC + 1)
}

func (vm *VirtualMachine) readDoubleWordFrom(addr uint16) uint16 {
	return uint16(vm.Mem[addr+1])<<8 + uint16(vm.Mem[addr])
}

func (vm *VirtualMachine) addSetFlags(reg *Register) {
	vm.SR.Carry, vm.SR.Zero, vm.SR.Sign = vm.Regs.A.add(reg)
}

func (vm *VirtualMachine) subtractAndSetFlags(reg *Register) {
	vm.SR.Carry, vm.SR.Zero, vm.SR.Sign = vm.Regs.A.subtract(reg)
}

func (vm *VirtualMachine) push(opcode Word) {
	// From 8080 assembly programming manual (page 22), the bits 4 and 5 from
	// the PUSH opcode codify which registers are pushed.
	switch opcode & 0x30 {
	case 0:
		vm.pushWordIntoStack(vm.Regs.B.Value)
		vm.pushWordIntoStack(vm.Regs.C.Value)
	case 16:
		vm.pushWordIntoStack(vm.Regs.D.Value)
		vm.pushWordIntoStack(vm.Regs.E.Value)
	case 32:
		vm.pushWordIntoStack(vm.Regs.H.Value)
		vm.pushWordIntoStack(vm.Regs.L.Value)
	case 48:
		vm.pushWordIntoStack(vm.SR.Value())
		vm.pushWordIntoStack(vm.Regs.A.Value)
	}
}

func (vm *VirtualMachine) pushIntoStack(value uint16) {
	vm.pushWordIntoStack(Word(value >> 8))
	vm.pushWordIntoStack(Word(0x00FF & value))
}

func (vm *VirtualMachine) pushWordIntoStack(value Word) {
	vm.SP = vm.SP - 1
	vm.Mem[vm.SP] = value
}

func (vm *VirtualMachine) popPCFromStack() uint16 {
	val := vm.readDoubleWordFrom(vm.SP)
	vm.SP = vm.SP + 2
	return val
}

func (vm *VirtualMachine) pop(opcode Word) {
	// From 8080 assembly programming manual (page 23), the bits 4 and 5 from
	// the POP opcode codify which registers are poped.
	switch opcode & 0x30 {
	case 0:
		vm.popWordIntoRegister(&vm.Regs.C)
		vm.popWordIntoRegister(&vm.Regs.B)
	case 16:
		vm.popWordIntoRegister(&vm.Regs.E)
		vm.popWordIntoRegister(&vm.Regs.D)
	case 32:
		vm.popWordIntoRegister(&vm.Regs.L)
		vm.popWordIntoRegister(&vm.Regs.H)
	case 48:
		vm.popWordIntoRegister(&vm.Regs.A)
		vm.SR.setFromValue(vm.popWord())
	}
}

func (vm *VirtualMachine) popWordIntoRegister(reg *Register) {
	reg.Value = vm.popWord()
}

func (vm *VirtualMachine) popWord() Word {
	value := vm.Mem[vm.SP]
	vm.SP = vm.SP + 1
	return value
}

func (vm *VirtualMachine) mov(opcode Word) {
	dst := uint8(opcode) & 0x38 // register indicated by bits  bits 3, 4, and 5
	src := uint8(opcode) & 0x07 // register indicated by bits  bits 0, 1, and 2
	if dst == 0x30 && src == 0x06 {
		panic("MOV cannot apply to M in both dst and src")
	}
	var srcVal Word
	if src != 0x06 {
		srcVal = vm.reg(src).Value
	} else {
		srcVal = vm.Mem[vm.Regs.HL()]
	}
	if dst != 0x30 {
		vm.reg(dst).load(srcVal)
	} else {
		vm.Mem[vm.Regs.HL()] = srcVal
	}

}

func (vm *VirtualMachine) callNextSubroutine() {
	vm.call(vm.getAddressFromNextAddress())
}

func (vm *VirtualMachine) call(addr uint16) {
	vm.pushNextPCIntoStack()
	vm.jumpTo(addr)
}

func (vm *VirtualMachine) pushNextPCIntoStack() {
	vm.pushIntoStack(vm.PC + 3)
}

func (vm *VirtualMachine) jumpToNextAddress() {
	vm.jumpTo(vm.getAddressFromNextAddress())
}

func (vm *VirtualMachine) jumpTo(addr uint16) {
	vm.PC = addr
}

func (vm *VirtualMachine) io(opcode Word) {
	// bit 3 holds 1 for IN and 0 for OUT
	devNum := uint8(vm.Mem[vm.PC+1])
	switch opcode & 0x08 {
	case 0:
		vm.Devices[devNum].Write(vm.Regs.A.Value)
	case 8:
		vm.Regs.A.load(vm.Devices[devNum].Read())
	}
	vm.PC = vm.PC + 2
}

func (vm *VirtualMachine) lxi(opcode Word) {
	// bits 4 and 5 codify the register pair loaded
	addr := vm.getAddressFromNextAddress()
	switch opcode & 0x30 {
	case 0:
		loadRegPair(&vm.Regs.B, &vm.Regs.C, addr)
	case 0x10:
		loadRegPair(&vm.Regs.D, &vm.Regs.E, addr)
	case 0x20:
		loadRegPair(&vm.Regs.H, &vm.Regs.L, addr)
	case 0x30:
		vm.SP = addr
	}
	vm.PC = vm.PC + 3
}

func loadRegPair(high *Register, low *Register, value uint16) {
	low.load(Word(LSB(value)))
	high.load(Word(MSB(value)))
}

func (vm *VirtualMachine) exchangeDEandHL() {
	exchange(&vm.Regs.D, &vm.Regs.H)
	exchange(&vm.Regs.E, &vm.Regs.L)
}

func exchange(r1 *Register, r2 *Register) {
	temp := r1.Value
	r1.load(r2.Value)
	r2.load(temp)
}

func (vm *VirtualMachine) exchangeHLandSP() {
	exchangeRegWithAddr(&vm.Regs.L, &vm.Mem, vm.SP)
	exchangeRegWithAddr(&vm.Regs.H, &vm.Mem, vm.SP+1)
}

func exchangeRegWithAddr(reg *Register, mem *Memory, addr uint16) {
	t := reg.Value
	reg.load((*mem)[addr])
	(*mem)[addr] = t
}

func (vm *VirtualMachine) enableInterrupts() {
	vm.INTE = true
}

func (vm *VirtualMachine) disableInterrupts() {
	vm.INTE = false
}

type Memory []Word

func (m *Memory) inc(addr uint16) {
	(*m)[addr] = (*m)[addr] + 1
}

func (m *Memory) dec(addr uint16) {
	(*m)[addr] = (*m)[addr] - 1
}

func (m *Memory) getNext2Words(addr uint16) uint16 {
	return m.get2Words(addr + 1)
}

func (m *Memory) get2Words(addr uint16) uint16 {
	return uint16((*m)[addr]) + uint16((*m)[addr+1])>>8
}

func (m *Memory) Load(sections map[uint16]Memory) {
	// Each section entry in the map is a section of the memory; it's key is
	// the start address, and the array is the content.
	// The result is a full 64Kb memory with all the sections copied
	for address, content := range sections {
		copy((*m)[address:int(address)+len(content)], content)
	}
}

func (m *Memory) LastAddress() uint16 {
	return uint16(cap(*m) - 1)
}

type Device interface {
	Read() Word
	Write(data Word)
}

func LSB(addr uint16) uint8 {
	return uint8(addr)
}

func MSB(addr uint16) uint8 {
	return uint8(addr >> 8)
}

const (
	NOP     = 0x00 // Do nothing
	HLT     = 0x76 // NOP; PC <- PC-1
	INRA    = 0x3C // A <- A + 1
	INRB    = 0x04 // B <- B + 1
	INRC    = 0x0C // C <- C + 1
	INRD    = 0x14 // D <- D + 1
	INRE    = 0x1C // E <- E + 1
	INRH    = 0x24 // H <- H + 1
	INRL    = 0x2C // L <- L + 1
	INRM    = 0x34 // (HL) <- (HL) + 1
	DCRA    = 0x3D // A <- A - 1
	DCRB    = 0x05 // B <- B - 1
	DCRC    = 0x0D // C <- C - 1
	DCRD    = 0x15 // D <- D - 1
	DCRE    = 0x1D // E <- E - 1
	DCRH    = 0x25 // H <- H - 1
	DCRL    = 0x2D // L <- L - 1
	DCRM    = 0x35 // (HL) <- (HL) - 1
	MOVAA   = 0x7F // A <- A
	MOVAB   = 0x78 // A <- B
	MOVAC   = 0x79 // A <- C
	MOVAD   = 0x7A // A <- D
	MOVAE   = 0x7B // A <- E
	MOVAH   = 0x7C // A <- H
	MOVAL   = 0x7D // A <- L
	MOVAM   = 0x7E // A <- (HL)
	MOVBA   = 0x47 // B <- A
	MOVBB   = 0x40 // B <- B
	MOVBC   = 0x41 // B <- C
	MOVBD   = 0x42 // B <- D
	MOVBE   = 0x43 // B <- E
	MOVBH   = 0x44 // B <- H
	MOVBL   = 0x45 // B <- L
	MOVBM   = 0x46 // B <- (HL)
	MOVCA   = 0x4F // C <- A
	MOVCB   = 0x48 // C <- B
	MOVCC   = 0x49 // C <- C
	MOVCD   = 0x4A // C <- D
	MOVCE   = 0x4B // C <- E
	MOVCH   = 0x4C // C <- H
	MOVCL   = 0x4D // C <- L
	MOVCM   = 0x4E // C <- (HL)
	MOVDA   = 0x57 // D <- A
	MOVDB   = 0x50 // D <- B
	MOVDC   = 0x51 // D <- C
	MOVDD   = 0x52 // D <- D
	MOVDE   = 0x53 // D <- E
	MOVDH   = 0x54 // D <- H
	MOVDL   = 0x55 // D <- L
	MOVDM   = 0x56 // D <- (HL)
	MOVEA   = 0x5F // E <- A
	MOVEB   = 0x58 // E <- B
	MOVEC   = 0x59 // E <- C
	MOVED   = 0x5A // E <- D
	MOVEE   = 0x5B // E <- E
	MOVEH   = 0x5C // E <- H
	MOVEL   = 0x5D // E <- L
	MOVEM   = 0x5E // E <- (HL)
	MOVHA   = 0x67 // H <- A
	MOVHB   = 0x60 // H <- B
	MOVHC   = 0x61 // H <- C
	MOVHD   = 0x62 // H <- D
	MOVHE   = 0x63 // H <- E
	MOVHH   = 0x64 // H <- H
	MOVHL   = 0x65 // H <- L
	MOVHM   = 0x66 // H <- (HL)
	MOVLA   = 0x6F // L <- A
	MOVLB   = 0x68 // L <- B
	MOVLC   = 0x69 // L <- C
	MOVLD   = 0x6A // L <- D
	MOVLE   = 0x6B // L <- E
	MOVLH   = 0x6C // L <- H
	MOVLL   = 0x6D // L <- L
	MOVLM   = 0x6E // L <- (HL)
	MOVMA   = 0x77 // (HL) <- A
	MOVMB   = 0x70 // (HL) <- B
	MOVMC   = 0x71 // (HL) <- C
	MOVMD   = 0x72 // (HL) <- D
	MOVME   = 0x73 // (HL) <- E
	MOVMH   = 0x74 // (HL) <- H
	MOVML   = 0x75 // (HL) <- L
	ADDA    = 0x87 // A <- A + A
	ADDB    = 0x80 // A <- A + B
	ADDC    = 0x81 // A <- A + C
	ADDD    = 0x82 // A <- A + D
	ADDE    = 0x83 // A <- A + E
	ADDH    = 0x84 // A <- A + H
	ADDL    = 0x85 // A <- A + L
	ADDM    = 0x86 // A <- A + (HL)
	SUBA    = 0x97 // A <- A - A
	SUBB    = 0x90 // A <- A - B
	SUBC    = 0x91 // A <- A - C
	SUBD    = 0x92 // A <- A - D
	SUBE    = 0x93 // A <- A - E
	SUBH    = 0x94 // A <- A - H
	SUBL    = 0x95 // A <- A - L
	SUBM    = 0x96 // A <- A - M
	SUI     = 0xD6 // A <- A - byte (not implemented)
	MVIA    = 0x3E // A <- byte
	MVIB    = 0x06 // B <- byte
	MVIC    = 0x0e // C <- byte
	MVID    = 0x16 // D <- byte
	MVIE    = 0x1E // E <- byte
	MVIH    = 0x26 // H <- byte
	MVIL    = 0x2E // L <- byte
	CMA     = 0x27 // A <- not A
	ANAA    = 0xA7 // A <- A and A
	ANAB    = 0xA0 // A <- A and B
	ANAC    = 0xA1 // A <- A and C
	ANAD    = 0xA2 // A <- A and D
	ANAE    = 0xA3 // A <- A and E
	ANAH    = 0xA4 // A <- A and H
	ANAL    = 0xA5 // A <- A and L
	ANAM    = 0xA6 // A <- A and (HL) (not implemented)
	ANAI    = 0xE6 // A <- A and byte (not implemented)
	XRAA    = 0xAF // A <- A xor A
	XRAB    = 0xA8 // A <- A xor B
	XRAC    = 0xA9 // A <- A xor C
	XRAD    = 0xAA // A <- A xor D
	XRAE    = 0xAB // A <- A xor E
	XRAH    = 0xAC // A <- A xor H
	XRAL    = 0xAD // A <- A xor L
	XRAM    = 0xAE // A <- A xor (HL)
	XRAI    = 0xEE // A <- A xor byte (not implemented)
	ORAA    = 0xB7 // A <- A or A
	ORAB    = 0xB0 // A <- A or B
	ORAC    = 0xB1 // A <- A or C
	ORAD    = 0xB2 // A <- A or D
	ORAE    = 0xB3 // A <- A or E
	ORAH    = 0xB4 // A <- A or H
	ORAL    = 0xB5 // A <- A or L
	ORAM    = 0xB6 // A <- A or (HL)
	ORI     = 0xF6 // A <- A or byte (not implemented)
	CMPA    = 0xBf // A - A (SR is affected)
	CMPB    = 0xB8 // A - B (SR is affected)
	CMPC    = 0xB9 // A - C (SR is affected)
	CMPD    = 0xBA // A - D (SR is affected)
	CMPE    = 0xBB // A - E (SR is affected)
	CMPH    = 0xBC // A - H (SR is affected)
	CMPL    = 0xBD // A - L (SR is affected)
	CMPM    = 0xBE // A - (HL) (SR is affected)
	CMPI    = 0xFE // A - byte
	JMP     = 0xC3 // PC <- address
	JNZ     = 0xC2 // If Z flag is off, PC <- address
	JZ      = 0xCA // If Z flag is on, PC <- address
	JNC     = 0xD2 // If C flag is off, PC <- address
	JC      = 0xDA // If C flag is on, PC <- address
	JPO     = 0xE2 // If P flag is odd, PC <- address (not implemented)
	JPE     = 0xEA // If P flag is even, PC <- address (not implemented)
	JP      = 0xF2 // If S flag is off, PC <- address
	JM      = 0xFA // If S flag is on, PC <- address
	LXIB    = 0x01 // BC <- word
	LXID    = 0x11 // DE <- word
	LXIH    = 0x21 // HL <- word
	LXISP   = 0x31 // SP <- address
	XCHG    = 0xEB // HL <-> DE
	XTHL    = 0xE3 // H <-> (SP+1); L <-> (SP)
	CALL    = 0xCD // push PC to stack; PC <- address
	CNZ     = 0xC4 // If NZ, CALL address
	CZ      = 0xCC //	If Z, CALL address
	CNC     = 0xD4 // If C flag is off, CALL address
	CC      = 0xDC // If C flag is off, CALL address
	CPO     = 0xE4 // If PO, CALL address (not implemented)
	CPE     = 0xEC // If PE, CALL address (not implemented)
	CP      = 0xF4 // If S is off (indicates positive result), CALL address
	CM      = 0xFC // If S is on (indicates negative result), CALL address
	RET     = 0xC9 // Pop PC from stack
	RNZ     = 0xC0 // If Z is off, pop PC from stack
	RZ      = 0xC8 // If Z is on, pop PC from stack
	RNC     = 0xD0 // If C is off, pop PC from stack
	RC      = 0xD8 // If C is on, pop PC from stack
	RPO     = 0xE0 // if Parity is Odd, pop PC from stack (not implemented)
	RPE     = 0xE8 // If Parity is Even, pop PC from stack (not implemented)
	RP      = 0xF0 // If P (S is off), pop PC from stack
	RM      = 0xF8 // If M (S is on), pop PC from stack
	RST0    = 0xC7 // CALL 0
	RST1    = 0xCF // CALL 8
	RST2    = 0xD7 // CALL 10H
	RST3    = 0xDF // CALL 18H
	RST4    = 0xE7 // CALL 20H
	RST5    = 0xEF // CALL 28H
	RST6    = 0xF7 // CALL 30H
	RST7    = 0xFF // CALL 38H
	PUSHB   = 0xC5 // (SP-2) <- C; (SP-1) <- B; SP <- SP - 2
	PUSHD   = 0xD5 // (SP-2) <- E; (SP-1) <- D; SP <- SP - 2
	PUSHH   = 0xE5 // (SP-2) <- L; (SP-1) <- H; SP <- SP - 2
	PUSHPSW = 0xF5 // (SP-2) <- Flags; (SP-1) <- A; SP <- SP - 2
	POPB    = 0xC1 // B <- (SP+1); C <- (SP); SP <- SP + 2
	POPD    = 0xD1 // D <- (SP+1); E <- (SP); SP <- SP + 2
	POPH    = 0xE1 // H <- (SP+1); L <- (SP); SP <- SP + 2
	POPPSW  = 0xF1 // A <- (SP+1); Flags <- (SP); SP <- SP + 2
	IN      = 0xDB // A <- [byte]
	OUT     = 0xD3 // [byte] <- A
	DI      = 0xF3 // IFF <- 0
	EI      = 0xFB // IFF <- 1
)

type Registers struct {
	A Register
	B Register
	C Register
	D Register
	E Register
	H Register
	L Register
}

func (r *Registers) HL() uint16 {
	return uint16(r.H.Value)<<7 + uint16(r.L.Value)
}

func (v *VirtualMachine) reg(index uint8) *Register {
	// Many opcodes codify an index to an operand that is a register; this
	// function returns a reference the Register corresponding to such an index.
	var reg *Register = nil
	switch index {
	case 0:
		reg = &v.Regs.B
	case 1:
		fallthrough
	case 1 << 3:
		reg = &v.Regs.C
	case 2:
		fallthrough
	case 2 << 3:
		reg = &v.Regs.D
	case 3:
		fallthrough
	case 3 << 3:
		reg = &v.Regs.E
	case 4:
		fallthrough
	case 4 << 3:
		reg = &v.Regs.H
	case 5:
		fallthrough
	case 5 << 3:
		reg = &v.Regs.L
	case 7:
		fallthrough
	case 7 << 3:
		reg = &v.Regs.A
	default:
		panic(fmt.Sprintf("Invalid register: %d", index))
	}
	return reg
}

type Register struct {
	Value Word
	Name  string
}

func (r *Register) inc() {
	r.Value = r.Value + 1
}

func (r *Register) dec() {
	r.Value = r.Value - 1
}

func (r *Register) copyFrom(other *Register) {
	r.Value = other.Value
}

func (r *Register) add(other *Register) (carry Flag, zero Flag, sign Flag) {
	res := r.Value + other.Value
	zero = res == 0
	sign = res < 0
	carry = res < r.Value && res < other.Value
	r.Value = res
	return
}

func (r *Register) subtract(other *Register) (carry Flag, zero Flag, sign Flag) {
	// As per 8080 docs, subtract is adding (complement + 1)
	o := (0xFF ^ other.Value) + 0x01
	res := r.Value + o
	carry = res >= r.Value || res >= o
	zero = res == 0
	sign = carry
	r.Value = res
	return
}

func (r *Register) load(value Word) {
	r.Value = value
}

func (r *Register) Has(value Word) bool {
	return r.Value == value
}

func (r *Register) NOT() {
	r.Value = 0xff ^ r.Value
}

func (r *Register) AND(other *Register) {
	r.Value = r.Value & other.Value
}

func (r *Register) XOR(other *Register) {
	r.Value = r.Value ^ other.Value
}

func (r *Register) OR(other *Register) {
	r.Value = r.Value | other.Value
}

func (r Register) String() string {
	return fmt.Sprintf("%s=%X", r.Name, r.Value)
}

type Word uint8

type StatusRegister struct {
	Sign           Flag // set if the result is negative
	Zero           Flag // set if the result is zero
	Parity         Flag // set if the number if 1 bits in the result is even
	Carry          Flag // set if the last addition resulted in a carry or if the last subtraction required a borrow
	AuxiliaryCarry Flag // used for binary-coded decimal arithmetic (?!)
}

func (sr StatusRegister) Value() Word {
	var value uint8 = 0x02 // as per 8080 assembly manual (page 22) bit 1 is always 1
	if sr.Carry {
		value = value | 1
	}
	if sr.Parity {
		value = value | 1<<2
	}
	if sr.AuxiliaryCarry {
		value = value | 1<<4
	}
	if sr.Zero {
		value = value | 1<<6
	}
	if sr.Sign {
		value = value | 1<<7
	}
	return Word(value)
}

func (sr *StatusRegister) setFromValue(value Word) {
	v := uint8(value)
	sr.Carry = v&1 == 1
	sr.Parity = v&1<<2 == 1<<2
	sr.AuxiliaryCarry = v&1<<4 == 1<<4
	sr.Zero = v&1<<6 == 1<<6
	sr.Sign = v&1<<7 == 1<<7
}

type Flag bool

func Make8080() VirtualMachine {
	return VirtualMachine{
		Regs:    makeRegisters(),
		Mem:     MakeMemory(),
		Devices: makeDevices()}
}

func makeRegisters() Registers {
	return Registers{
		Register{Name: "A"},
		Register{Name: "B"},
		Register{Name: "C"},
		Register{Name: "D"},
		Register{Name: "E"},
		Register{Name: "H"},
		Register{Name: "L"},
	}
}

func MakeMemory() Memory {
	return make(Memory, 1024*64, 1024*64)
}

const _64K = 1024 * 64

func makeDevices() map[uint8]Device {
	return make(map[uint8]Device)
}
