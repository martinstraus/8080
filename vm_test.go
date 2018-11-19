package vm8080

import (
	"testing"
)

func TestINRA(t *testing.T) {
	checkINR(t, INRA, regA)
}

func TestINRB(t *testing.T) {
	checkINR(t, INRB, regB)
}

func TestINRC(t *testing.T) {
	checkINR(t, INRC, regC)
}

func TestINRD(t *testing.T) {
	checkINR(t, INRD, regD)
}

func TestINRE(t *testing.T) {
	checkINR(t, INRE, regE)
}

func TestINRH(t *testing.T) {
	checkINR(t, INRH, regH)
}

func TestINRL(t *testing.T) {
	checkINR(t, INRL, regL)
}

func TestINRM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, INRM, HLT, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.Mem[0x00006] != 0x01 {
			t.Errorf("We expected %X at address %X but it had %X", 0x01, 0x0006, v.Mem[0x0006])
		}
	})
}

func checkINR(t *testing.T, opcode Word, f func(*VirtualMachine) *Register) {
	runAndCheckVM(t, &Memory{opcode}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, f(v), 0x01)
	})
}

func TestDCRA(t *testing.T) {
	checkDCR(t, DCRA, regA)
}

func TestDCRB(t *testing.T) {
	checkDCR(t, DCRB, regB)
}

func TestDCRC(t *testing.T) {
	checkDCR(t, DCRC, regC)
}

func TestDCRD(t *testing.T) {
	checkDCR(t, DCRD, regD)
}

func TestDCRE(t *testing.T) {
	checkDCR(t, DCRE, regE)
}

func TestDCRH(t *testing.T) {
	checkDCR(t, DCRH, regH)
}

func TestDCRL(t *testing.T) {
	checkDCR(t, DCRL, regL)
}

func checkDCR(t *testing.T, opcode Word, f func(*VirtualMachine) *Register) {
	runAndCheckVM(t, &Memory{opcode}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, f(v), 0xFF)
	})
}

func TestDRRM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, DCRM, HLT, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.Mem[0x00006] != 0xFF {
			t.Errorf("We expected %X at address %X but it had %X", 0xFF, 0x0006, v.Mem[0x0006])
		}
	})
}
func regA(v *VirtualMachine) *Register {
	return &v.Regs.A
}

func regB(v *VirtualMachine) *Register {
	return &v.Regs.B
}

func regC(v *VirtualMachine) *Register {
	return &v.Regs.C
}

func regD(v *VirtualMachine) *Register {
	return &v.Regs.D
}

func regE(v *VirtualMachine) *Register {
	return &v.Regs.E
}

func regH(v *VirtualMachine) *Register {
	return &v.Regs.H
}

func regL(v *VirtualMachine) *Register {
	return &v.Regs.L
}

func TestDCR(t *testing.T) {
	runAndCheckVM(t, &Memory{INRA, DCRA, INRB, DCRB, INRC, DCRC, INRD, DCRD}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
		checkRegHasValue(t, &v.Regs.B, 0x00)
		checkRegHasValue(t, &v.Regs.C, 0x00)
		checkRegHasValue(t, &v.Regs.D, 0x00)
	})
}

func TestMOVAA(t *testing.T) {
	checkMOVA(t, MVIA, MOVAA)
}
func TestMOVAB(t *testing.T) {
	checkMOVA(t, MVIB, MOVAB)
}

func TestMOVAC(t *testing.T) {
	checkMOVA(t, MVIC, MOVAC)
}

func TestMOVAD(t *testing.T) {
	checkMOVA(t, MVID, MOVAD)
}

func TestMOVAE(t *testing.T) {
	checkMOVA(t, MVIE, MOVAE)
}

func TestMOVAH(t *testing.T) {
	checkMOVA(t, MVIH, MOVAH)
}

func TestMOVAL(t *testing.T) {
	checkMOVA(t, MVIL, MOVAL)
}

func checkMOVA(t *testing.T, opcodeSetReg Word, opcodeToTest Word) {
	runAndCheckVM(t, &Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x02)
	})
}

func TestMOVAM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, MOVAM, HLT, 0xFF}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xFF)
	})
}

func TestMOVBA(t *testing.T) {
	checkMOVB(t, MVIA, MOVBA)
}
func TestMOVBB(t *testing.T) {
	checkMOVB(t, MVIB, MOVBB)
}

func TestMOVBC(t *testing.T) {
	checkMOVB(t, MVIC, MOVBC)
}

func TestMOVBD(t *testing.T) {
	checkMOVB(t, MVID, MOVBD)
}

func TestMOVBE(t *testing.T) {
	checkMOVB(t, MVIE, MOVBE)
}

func TestMOVBH(t *testing.T) {
	checkMOVB(t, MVIH, MOVBH)
}

func TestMOVBL(t *testing.T) {
	checkMOVB(t, MVIL, MOVBL)
}

func checkMOVB(t *testing.T, opcodeSetReg Word, opcodeToTest Word) {
	runAndCheckVM(t, &Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0x02)
	})
}

func TestMOVBM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, MOVBM, HLT, 0xFF}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0xFF)
	})
}

func TestMOVCA(t *testing.T) {
	checkMOVC(t, MVIA, MOVCA)
}
func TestMOVCB(t *testing.T) {
	checkMOVC(t, MVIB, MOVCB)
}

func TestMOVCC(t *testing.T) {
	checkMOVC(t, MVIC, MOVCC)
}

func TestMOVCD(t *testing.T) {
	checkMOVC(t, MVID, MOVCD)
}

func TestMOVCE(t *testing.T) {
	checkMOVC(t, MVIE, MOVCE)
}

func TestMOVCH(t *testing.T) {
	checkMOVC(t, MVIH, MOVCH)
}

func TestMOVCL(t *testing.T) {
	checkMOVC(t, MVIL, MOVCL)
}

func checkMOVC(t *testing.T, opcodeSetReg Word, opcodeToTest Word) {
	runAndCheckVM(t, &Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0x02)
	})
}

func TestMOVCM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, MOVCM, HLT, 0xFF}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0xFF)
	})
}

func TestMOVDA(t *testing.T) {
	checkMOVD(t, MVIA, MOVDA)
}
func TestMOVDB(t *testing.T) {
	checkMOVD(t, MVIB, MOVDB)
}

func TestMOVDC(t *testing.T) {
	checkMOVD(t, MVIC, MOVDC)
}

func TestMOVDD(t *testing.T) {
	checkMOVD(t, MVID, MOVDD)
}

func TestMOVDE(t *testing.T) {
	checkMOVD(t, MVIE, MOVDE)
}

func TestMOVDH(t *testing.T) {
	checkMOVD(t, MVIH, MOVDH)
}

func TestMOVDL(t *testing.T) {
	checkMOVD(t, MVIL, MOVDL)
}

func checkMOVD(t *testing.T, opcodeSetReg Word, opcodeToTest Word) {
	runAndCheckVM(t, &Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x02)
	})
}

func TestMOVDM(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIH, 0x00, MVIL, 0x06, MOVDM, HLT, 0xFF}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0xFF)
	})
}

func TestADDA(t *testing.T) {
	runAndCheckVM(t, &Memory{INRA, ADDA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x02)
	})
}

func TestADDB(t *testing.T) {
	runAndCheckVM(t, &Memory{INRB, ADDB}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDBSetsCarry(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0x01, ADDB}, func(v *VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestADDBSetsFlags(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x01, MVIB, 0x01, ADDB}, func(v *VirtualMachine, t *testing.T) {
		checkFlags(t, v, false, false, false)
	})
}

func TestADDC(t *testing.T) {
	runAndCheckVM(t, &Memory{INRC, ADDC}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDCSetsCarry(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIC, 0x01, ADDC}, func(v *VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestADDD(t *testing.T) {
	runAndCheckVM(t, &Memory{INRD, ADDD}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDDSetsCarry(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVID, 0x01, ADDD}, func(v *VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestSUBA(t *testing.T) {
	runAndCheckVM(t, &Memory{INRA, INRA, INRA, SUBA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestSUBBHasCorrectResult(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x02, MVIB, 0x01, SUBB}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestSUBBWithEqualBSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x01, MVIB, 0x01, SUBB}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
		checkFlags(t, v, false, true, false)
	})
}

func TestSUBBWitBLessThanSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x02, MVIB, 0x01, SUBB}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
		checkFlags(t, v, false, false, false)
	})
}

func TestSUBBWitBGreaterThanSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, MVIB, 0x01, SUBB}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
		checkFlags(t, v, true, false, true)
	})
}

func checkFlags(t *testing.T, v *VirtualMachine, sign Flag, zero Flag, carry Flag) {
	checkFlag(t, v.SR.Sign, sign, "Sign")
	checkFlag(t, v.SR.Zero, zero, "Zero")
	checkFlag(t, v.SR.Carry, carry, "Carry")
}

func checkFlag(t *testing.T, flag Flag, expected Flag, name string) {
	if flag != expected {
		t.Errorf("%s flag has unexpected value; expected '%t' and has '%t'.", name, expected, flag)
	}
}

func TestSUBC(t *testing.T) {
	runAndCheckVM(t, &Memory{INRA, INRA, INRC, SUBC}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestSUBD(t *testing.T) {
	runAndCheckVM(t, &Memory{INRA, INRA, INRD, SUBD}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestMVIA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestMVIB(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIB, 0xff}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0xff)
	})
}

func TestMVIC(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIC, 0xff}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0xff)
	})
}

func TestMVID(t *testing.T) {
	runAndCheckVM(t, &Memory{MVID, 0xff}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0xff)
	})
}

func TestCMA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, CMA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestANAA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, ANAA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestANAB(t *testing.T) {
	checkANA(t, MVIB, ANAB)
}

func TestANAC(t *testing.T) {
	checkANA(t, MVIC, ANAC)
}

func TestANAD(t *testing.T) {
	checkANA(t, MVID, ANAD)
}

func TestANAE(t *testing.T) {
	checkANA(t, MVIE, ANAE)
}

func TestANAH(t *testing.T) {
	checkANA(t, MVIH, ANAH)
}

func TestANAL(t *testing.T) {
	checkANA(t, MVIL, ANAL)
}

func checkANA(t *testing.T, opcodeMVI Word, opcodeANA Word) {
	runAndCheckVM(t, &Memory{MVIA, 0xf0, opcodeMVI, 0xf0, opcodeANA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xf0)
	})
}

func TestXRAA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xf0, XRAA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestXRAB(t *testing.T) {
	checkXRA(t, MVIB, XRAB)
}

func TestXRAC(t *testing.T) {
	checkXRA(t, MVIC, XRAC)
}

func TestXRAD(t *testing.T) {
	checkXRA(t, MVID, XRAD)
}

func TestXRAE(t *testing.T) {
	checkXRA(t, MVIE, XRAE)
}

func TestXRAH(t *testing.T) {
	checkXRA(t, MVIH, XRAH)
}

func TestXRAL(t *testing.T) {
	checkXRA(t, MVIL, XRAL)
}

func checkXRA(t *testing.T, opcodeMVI Word, opcodeXRA Word) {
	runAndCheckVM(t, &Memory{MVIA, 0xf0, opcodeMVI, 0xff, opcodeXRA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x0f)
	})

}

func TestORAA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, ORAA}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestORAB(t *testing.T) {
	checkOR(t, MVIB, ORAB)
}

func TestORAC(t *testing.T) {
	checkOR(t, MVIC, ORAC)
}

func TestORAD(t *testing.T) {
	checkOR(t, MVID, ORAD)
}

func TestORAE(t *testing.T) {
	checkOR(t, MVIE, ORAE)
}

func TestORAH(t *testing.T) {
	checkOR(t, MVIH, ORAH)
}

func TestORAL(t *testing.T) {
	checkOR(t, MVIL, ORAL)
}

func checkOR(t *testing.T, opcodeMVI Word, opcodeOR Word) {
	runAndCheckVM(t, &Memory{MVIA, 0xf0, opcodeMVI, 0x0f, opcodeOR}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func checkRegHasValue(t *testing.T, r *Register, v Word) {
	if !r.Has(v) {
		t.Errorf("Register %s has value %x but we were expecting %x", r.Name, r.Value, v)
	}
}

func TestCMPA(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, CMPA}, func(v *VirtualMachine, t *testing.T) {
		if !v.SR.Zero {
			t.Error("Expected Z flag ON but was OFF")
		}
		if v.SR.Sign {
			t.Error("Expected S flag OFF but was ON")
		}
		if v.SR.Carry {
			t.Error("Expected C flag OFF but was ON")
		}
	})
}
func TestCMPBWithBLessThanA(t *testing.T) {
	testCMPXWithXLessThanA(t, MVIB, CMPB)
}

func TestCMPBWithBEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, MVIB, CMPB)
}

func TestCMPBWithBGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, MVIB, CMPB)
}

func TestCMPCWithCLessThanA(t *testing.T) {
	testCMPXWithXLessThanA(t, MVIC, CMPC)
}

func TestCMPCWithCEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, MVIC, CMPC)
}

func TestCMPCWithCGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, MVIC, CMPC)
}

func TestCMPDWithDLessThanA(t *testing.T) {
	testCMPXWithXLessThanA(t, MVID, CMPD)
}

func TestCMPDWithDEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, MVID, CMPD)
}

func TestCMPDWithDGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, MVID, CMPD)
}

func testCMPXWithXLessThanA(t *testing.T, instrMovX Word, instrCmp Word) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, instrMovX, 0x01, instrCmp}, func(v *VirtualMachine, t *testing.T) {
		if v.SR.Zero {
			t.Error("Expected Z flag OFF but was ON")
		}
		if v.SR.Sign {
			t.Error("Expected S flag OFF but was ON")
		}
	})
}

func testCMPXWithXEqualToA(t *testing.T, instrMovX Word, instrCmp Word) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, instrMovX, 0xff, instrCmp}, func(v *VirtualMachine, t *testing.T) {
		if !v.SR.Zero {
			t.Error("Expected Z flag ON but was OFF")
		}
		if v.SR.Sign {
			t.Error("Expected S flag OFF but was ON")
		}
	})
}

func testCMPXWithXGreaterThanA(t *testing.T, instrMovX Word, instrCmp Word) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, instrMovX, 0x01, instrCmp}, func(v *VirtualMachine, t *testing.T) {
		if v.SR.Zero {
			t.Error("Expected Z flag OFF but was ON")
		}
		if !v.SR.Sign {
			t.Error("Expected S flag ON but was OFF")
		}
	})
}

func TestNOPAndHLT(t *testing.T) {
	runAndCheckVM(t, &Memory{NOP, NOP, HLT, NOP}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x02)
	})
}

func TestJMP(t *testing.T) {
	runAndCheckVM(t, &Memory{JMP, 0x05, 0x00, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x05)
	})
}

func TestJNZJumpsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0x0f, CMPB, JNZ, 0x11, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJNZDoesntJumpIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0xff, CMPB, JNZ, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJZJumpsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0xff, CMPB, JZ, 0x11, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJZDoesntJumpIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0x0f, CMPB, JZ, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJNCJumpsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x0f, MVIB, 0x0f, ADDB, JNC, 0x11, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJNCDoesntJumpIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0x01, ADDB, JNC, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJCJumpsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0xff, MVIB, 0x0f, ADDB, JC, 0x11, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJCDoesntJumpIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x0f, MVIB, 0x0f, ADDB, JC, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJPJumpsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, CMPA, JP, 0x09, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x09)
	})
}

func TestJPDoesntJumpIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, MVIB, 0xff, CMPB, JP, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJMJumpsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, MVIB, 0xff, CMPB, JM, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x10)
	})
}

func TestJMDoesntJumpIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{MVIA, 0x00, MVIB, 0x00, CMPB, JM, 0x10, 0x00, HLT, NOP, NOP, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestLXIB(t *testing.T) {
	runAndCheckVM(t, &Memory{LXIB, 0x11, 0x22, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0x22)
		checkRegHasValue(t, &v.Regs.C, 0x11)
	})
}
func TestLXID(t *testing.T) {
	runAndCheckVM(t, &Memory{LXID, 0x11, 0x22, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x22)
		checkRegHasValue(t, &v.Regs.E, 0x11)
	})
}

func TestLXIH(t *testing.T) {
	runAndCheckVM(t, &Memory{LXIH, 0x11, 0x22, HLT}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.H, 0x22)
		checkRegHasValue(t, &v.Regs.L, 0x11)
	})
}
func TestLXISP(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0xAA, 0xFF, HLT}, func(v *VirtualMachine, t *testing.T) {
		if v.SP != 0xFFAA {
			t.Errorf("Expected SP to be %X but was %X", 0xFFAA, v.SP)
		}
	})
}

func TestCALL(t *testing.T) {
	/* To test CALL, we create a program that CALLs; in the subroutine, we set
	   A <- 0xFF, and return. Then we test A - B; if Z, we jump to the last
		 instruction; otherwise halt.*/
	/* This looks complicated. Could we add an observer to the VirtualMachine,
	notify when then CALL is performed, and then verify the PC? It'd add
	more complexity to the VM but it seems like a more direct approach */
	runAndCheckVM(t, &Memory{LXISP, 0x0D, 0x00, CALL, 0x07, 0x00, HLT, MVIA, 0xFF, RET, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.SP != 0x000D {
			t.Errorf("SP should be %X but it's %X", 0x000D, v.SP)
		}
		if v.PC != 0x0006 {
			t.Errorf("CALL to 0x09 might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x06)
		}
	})
}

func TestCNZCallsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x0f, MVIB, 0x01, CMPB, CNZ, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCNZDoesntCallsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x00, CMPB, CNZ, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCZCallsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x0f, MVIB, 0x0f, CMPB, CZ, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCZDoesntCallsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x0f, MVIB, 0x00, CMPB, CZ, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCNCCallsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x00, ADDB, CNC, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCNCDoesntCallIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0xff, MVIB, 0x01, ADDB, CNC, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCCCallsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0xff, MVIB, 0x01, ADDB, CC, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCCDoesntCallIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x00, ADDB, CC, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCPCallsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x01, MVIB, 0x01, ADDB, CP, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCPDoesntCallIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x01, SUBB, CP, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C shouldn't have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCMCallsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x01, SUBB, CM, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCMDoesntCallIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x10, 0x00, MVIA, 0x00, MVIB, 0x01, ADDB, CM, 0x0C, 0x00, HLT, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C shouldn't have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestRNZReturnsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRB, ADDB, RNZ, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x0006 {
			t.Errorf("RNZ might not have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x0006)
		}
	})
}

func TestRNZDoesntReturnIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRA, SUBA, RNZ, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000A {
			t.Errorf("RNZ might have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRZReturnsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRA, SUBA, RZ, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x0006 {
			t.Errorf("RZ might not have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRZDoesntReturnIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRB, ADDB, RZ, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		if v.PC != 0x000A {
			t.Errorf("RZ might have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRNCReturnsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0C, 0x00, CALL, 0x07, 0x00, HLT, RNC, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRNCDoesntReturnIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0F, 0x00, CALL, 0x07, 0x00, HLT, MVIA, 0xFF, ADDA, RNC, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x000B)
	})
}

func TestRCReturnsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0F, 0x00, CALL, 0x07, 0x00, HLT, MVIA, 0xFF, ADDA, RC, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRCDoesntReturnIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0C, 0x00, CALL, 0x07, 0x00, HLT, RC, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0008)
	})
}

func TestRPReturnsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0D, 0x00, CALL, 0x07, 0x00, HLT, INRA, RP, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRPDoesntReturnIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRB, SUBB, RP, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x000A)
	})
}

func TestRMReturnsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0E, 0x00, CALL, 0x07, 0x00, HLT, INRB, SUBB, RM, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRMDoesntReturnsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x0C, 0x00, CALL, 0x07, 0x00, HLT, RM, HLT, 0x00, 0x00, 0x00, 0x00}, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0008)
	})
}

func TestRST1JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST1, 0x08)
}

func TestRST2JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST2, 0x10)
}

func TestRST3JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST3, 0x18)
}

func TestRST4JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST4, 0x20)
}

func TestRST5JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST5, 0x28)
}

func TestRST6JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST6, 0x30)
}
func TestRST7JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, RST7, 0x38)
}

func checkRSTJumpsToCorrectAddress(t *testing.T, opcode Word, expectedFinishAddr uint16) {
	mem := MakeMemory()
	mem.Load(map[uint16]Memory{
		0x0000: setStackPointerToLastAddress(&mem),
		0x0003: Memory{opcode, HLT},
		0x0008: Memory{HLT},
		0x0010: Memory{HLT},
		0x0018: Memory{HLT},
		0x0020: Memory{HLT},
		0x0028: Memory{HLT},
		0x0030: Memory{HLT},
		0x0038: Memory{HLT},
	})
	runAndCheckVM(t, &mem, func(v *VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, expectedFinishAddr)
	})
}

func setStackPointerToLastAddress(m *Memory) Memory {
	lastAddress := m.LastAddress()
	return Memory{LXISP, Word(LSB(lastAddress)), Word(MSB(lastAddress))}
}

func TestLastAddress(t *testing.T) {
	mem := MakeMemory()
	lastAddress := mem.LastAddress()
	expected := uint16((1024 * 64) - 1)
	if lastAddress != expected {
		t.Errorf("Last address should be %X but it's %X", expected, lastAddress)
	}
}

func TestLSB(t *testing.T) {
	resultIsExpected(t, LSB, 0xFFAE, 0xAE)
}

func TestMSB(t *testing.T) {
	resultIsExpected(t, MSB, 0xAEFF, 0xAE)
}

func TestPUSHB(t *testing.T) {
	checkPUSHRegisters(t, MVIB, MVIC, PUSHB)
}

func TestPUSHD(t *testing.T) {
	checkPUSHRegisters(t, MVID, MVIE, PUSHD)
}

func TestPUSHH(t *testing.T) {
	checkPUSHRegisters(t, MVIH, MVIL, PUSHH)
}

func checkPUSHRegisters(t *testing.T, opcodeSetRegister1 Word,
	opcodeSetRegister2 Word, opcodePush Word) {
	checkPUSH(t,
		Memory{opcodeSetRegister1, 0x01, opcodeSetRegister2, 0x02, opcodePush, HLT},
		0x01,
		0x02)
}

func TestPUSHPSW(t *testing.T) {
	// We need a program that sets some distinctive values in SR.
	checkPUSH(t,
		Memory{MVIA, 0x00, MVIB, 0x01, SUBB, PUSHPSW, HLT},
		0x83,
		0xFF)
}

func checkPUSH(t *testing.T, program Memory, expectedInStack1 Word,
	expectedInStack2 Word) {
	mem := MakeMemory()
	mem.Load(map[uint16]Memory{
		0x0000: setStackPointerToLastAddress(&mem),
		0x0003: program,
	})
	runAndCheckVM(t, &mem, func(v *VirtualMachine, t *testing.T) {
		checkMemoryContent(t, v, 0xFFFE, expectedInStack1)
		checkMemoryContent(t, v, 0xFFFD, expectedInStack2)
	})
}

func TestPOPB(t *testing.T) {
	checkPOPRegisters(t, POPB, func(v *VirtualMachine) *Register {
		return &v.Regs.B
	}, func(v *VirtualMachine) *Register {
		return &v.Regs.C
	})
}

func TestPOPD(t *testing.T) {
	checkPOPRegisters(t, POPD, func(v *VirtualMachine) *Register {
		return &v.Regs.D
	}, func(v *VirtualMachine) *Register {
		return &v.Regs.E
	})
}

func TestPOPH(t *testing.T) {
	checkPOPRegisters(t, POPH, func(v *VirtualMachine) *Register {
		return &v.Regs.H
	}, func(v *VirtualMachine) *Register {
		return &v.Regs.L
	})
}

func checkPOPRegisters(t *testing.T, opcodePOP Word,
	fReg1 func(*VirtualMachine) *Register,
	fReg2 func(*VirtualMachine) *Register) {
	mem := MakeMemory()
	mem.Load(map[uint16]Memory{
		0x0000: Memory{LXISP, 0xFD, 0xFF},
		0x0003: Memory{opcodePOP, HLT},
		0xFFFD: Memory{0x02, 0x01},
	})
	runAndCheckVM(t, &mem, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, fReg1(v), 0x01)
		checkRegHasValue(t, fReg2(v), 0x02)
	})
}

func TestPOPPSW(t *testing.T) {
	mem := MakeMemory()
	mem.Load(map[uint16]Memory{
		0x0000: Memory{LXISP, 0xFD, 0xFF},
		0x0003: Memory{POPPSW, HLT},
		0xFFFD: Memory{0xFF, 0xD7}, // 0xD7 = b11010111, setting all flags to 1
	})
	runAndCheckVM(t, &mem, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xFF)
		checkFlags(t, v, true, true, true)
	})
}

func checkMemoryContent(t *testing.T, v *VirtualMachine, addr uint16, value Word) {
	if v.Mem[addr] != value {
		t.Errorf("Expected value %X in address %X but found %X", value, addr, v.Mem[addr])
	}
}

func resultIsExpected(t *testing.T, f func(uint16) uint8, param uint16, expected uint8) {
	result := f(param)
	if result != expected {
		t.Errorf("Expected %X, got %X", expected, result)
	}
}

func checkPCisAt(t *testing.T, v *VirtualMachine, expectedPC uint16) {
	if v.PC != expectedPC {
		t.Errorf("We expected the PC to be at %X but it's at %X", expectedPC, v.PC)
	}
}

func TestRegistersHL(t *testing.T) {
	regs := Registers{H: Register{0x00, "H"}, L: Register{0x05, "L"}}
	if regs.HL() != 0x0005 {
		t.Errorf("HL expected to be %X but was %X", 0x005, regs.HL())
	}
}

type DummyDevice struct {
	Data Word
}

func (d *DummyDevice) Read() Word {
	return d.Data
}

func (d *DummyDevice) Write(data Word) {
	d.Data = data
}

func TestIN(t *testing.T) {
	v := Make8080()
	v.Load(&Memory{IN, 0x01, HLT})
	dev := new(DummyDevice)
	dev.Data = 0xFF
	v.Devices[0x01] = dev
	err := v.Run()
	if err != nil {
		t.Error(err)
	}
	checkRegHasValue(t, &v.Regs.A, 0xFF)
}

func TestOUT(t *testing.T) {
	v := Make8080()
	v.Load(&Memory{MVIA, 0xFF, OUT, 0x01, HLT})
	dev := new(DummyDevice)
	v.Devices[0x01] = dev
	err := v.Run()
	if err != nil {
		t.Error(err)
	}
	if dev.Data != 0xFF {
		t.Errorf("Expected %X in device %X but it has %X", 0xFF, 0x01, dev.Data)
	}
}

func TestXCHG(t *testing.T) {
	runAndCheckVM(t, &Memory{LXID, 0x11, 0x22, LXIH, 0x33, 0x44, XCHG}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x44)
		checkRegHasValue(t, &v.Regs.E, 0x33)
		checkRegHasValue(t, &v.Regs.H, 0x22)
		checkRegHasValue(t, &v.Regs.L, 0x11)
	})
}

func TestXTHL(t *testing.T) {
	runAndCheckVM(t, &Memory{LXISP, 0x08, 0x00, LXIH, 0x11, 0x22, XTHL, HLT, 0x33, 0x44}, func(v *VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.H, 0x44)
		checkRegHasValue(t, &v.Regs.L, 0x33)
		checkMemoryContent(t, v, 0x08, 0x11)
		checkMemoryContent(t, v, 0x09, 0x22)
	})
}

func runAndCheckVM(t *testing.T, memory *Memory, test func(v *VirtualMachine, t *testing.T)) {
	v := Make8080()
	v.Load(memory)
	//	fmt.Print("Running\n")
	//	defer fmt.Print("Ended\n")
	err := v.Run()
	if err != nil {
		t.Error(err)
	}
	test(&v, t)
}
