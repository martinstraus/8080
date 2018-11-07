package vm8080_test

import (
	"hyperion/vm8080"
	"testing"
)

func TestINRA(t *testing.T) {
	checkINR(t, vm8080.INRA, regA)
}

func TestINRB(t *testing.T) {
	checkINR(t, vm8080.INRB, regB)
}

func TestINRC(t *testing.T) {
	checkINR(t, vm8080.INRC, regC)
}

func TestINRD(t *testing.T) {
	checkINR(t, vm8080.INRD, regD)
}

func TestINRE(t *testing.T) {
	checkINR(t, vm8080.INRE, regE)
}

func TestINRH(t *testing.T) {
	checkINR(t, vm8080.INRH, regH)
}

func TestINRL(t *testing.T) {
	checkINR(t, vm8080.INRL, regL)
}

func TestINRM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.INRM, vm8080.HLT, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.Mem[0x00006] != 0x01 {
			t.Errorf("We expected %X at address %X but it had %X", 0x01, 0x0006, v.Mem[0x0006])
		}
	})
}

func checkINR(t *testing.T, opcode vm8080.Word, f func(*vm8080.VirtualMachine) *vm8080.Register) {
	runAndCheckVM(t, &vm8080.Memory{opcode}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, f(v), 0x01)
	})
}

func TestDCRA(t *testing.T) {
	checkDCR(t, vm8080.DCRA, regA)
}

func TestDCRB(t *testing.T) {
	checkDCR(t, vm8080.DCRB, regB)
}

func TestDCRC(t *testing.T) {
	checkDCR(t, vm8080.DCRC, regC)
}

func TestDCRD(t *testing.T) {
	checkDCR(t, vm8080.DCRD, regD)
}

func TestDCRE(t *testing.T) {
	checkDCR(t, vm8080.DCRE, regE)
}

func TestDCRH(t *testing.T) {
	checkDCR(t, vm8080.DCRH, regH)
}

func TestDCRL(t *testing.T) {
	checkDCR(t, vm8080.DCRL, regL)
}

func checkDCR(t *testing.T, opcode vm8080.Word, f func(*vm8080.VirtualMachine) *vm8080.Register) {
	runAndCheckVM(t, &vm8080.Memory{opcode}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, f(v), 0xFF)
	})
}

func TestDRRM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.DCRM, vm8080.HLT, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.Mem[0x00006] != 0xFF {
			t.Errorf("We expected %X at address %X but it had %X", 0xFF, 0x0006, v.Mem[0x0006])
		}
	})
}
func regA(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.A
}

func regB(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.B
}

func regC(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.C
}

func regD(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.D
}

func regE(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.E
}

func regH(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.H
}

func regL(v *vm8080.VirtualMachine) *vm8080.Register {
	return &v.Regs.L
}

func TestDCR(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRA, vm8080.DCRA, vm8080.INRB, vm8080.DCRB, vm8080.INRC, vm8080.DCRC, vm8080.INRD, vm8080.DCRD}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
		checkRegHasValue(t, &v.Regs.B, 0x00)
		checkRegHasValue(t, &v.Regs.C, 0x00)
		checkRegHasValue(t, &v.Regs.D, 0x00)
	})
}

func TestMOVAA(t *testing.T) {
	checkMOVA(t, vm8080.MVIA, vm8080.MOVAA)
}
func TestMOVAB(t *testing.T) {
	checkMOVA(t, vm8080.MVIB, vm8080.MOVAB)
}

func TestMOVAC(t *testing.T) {
	checkMOVA(t, vm8080.MVIC, vm8080.MOVAC)
}

func TestMOVAD(t *testing.T) {
	checkMOVA(t, vm8080.MVID, vm8080.MOVAD)
}

func TestMOVAE(t *testing.T) {
	checkMOVA(t, vm8080.MVIE, vm8080.MOVAE)
}

func TestMOVAH(t *testing.T) {
	checkMOVA(t, vm8080.MVIH, vm8080.MOVAH)
}

func TestMOVAL(t *testing.T) {
	checkMOVA(t, vm8080.MVIL, vm8080.MOVAL)
}

func checkMOVA(t *testing.T, opcodeSetReg vm8080.Word, opcodeToTest vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x02)
	})
}

func TestMOVAM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.MOVAM, vm8080.HLT, 0xFF}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xFF)
	})
}

func TestMOVBA(t *testing.T) {
	checkMOVB(t, vm8080.MVIA, vm8080.MOVBA)
}
func TestMOVBB(t *testing.T) {
	checkMOVB(t, vm8080.MVIB, vm8080.MOVBB)
}

func TestMOVBC(t *testing.T) {
	checkMOVB(t, vm8080.MVIC, vm8080.MOVBC)
}

func TestMOVBD(t *testing.T) {
	checkMOVB(t, vm8080.MVID, vm8080.MOVBD)
}

func TestMOVBE(t *testing.T) {
	checkMOVB(t, vm8080.MVIE, vm8080.MOVBE)
}

func TestMOVBH(t *testing.T) {
	checkMOVB(t, vm8080.MVIH, vm8080.MOVBH)
}

func TestMOVBL(t *testing.T) {
	checkMOVB(t, vm8080.MVIL, vm8080.MOVBL)
}

func checkMOVB(t *testing.T, opcodeSetReg vm8080.Word, opcodeToTest vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0x02)
	})
}

func TestMOVBM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.MOVBM, vm8080.HLT, 0xFF}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0xFF)
	})
}

func TestMOVCA(t *testing.T) {
	checkMOVC(t, vm8080.MVIA, vm8080.MOVCA)
}
func TestMOVCB(t *testing.T) {
	checkMOVC(t, vm8080.MVIB, vm8080.MOVCB)
}

func TestMOVCC(t *testing.T) {
	checkMOVC(t, vm8080.MVIC, vm8080.MOVCC)
}

func TestMOVCD(t *testing.T) {
	checkMOVC(t, vm8080.MVID, vm8080.MOVCD)
}

func TestMOVCE(t *testing.T) {
	checkMOVC(t, vm8080.MVIE, vm8080.MOVCE)
}

func TestMOVCH(t *testing.T) {
	checkMOVC(t, vm8080.MVIH, vm8080.MOVCH)
}

func TestMOVCL(t *testing.T) {
	checkMOVC(t, vm8080.MVIL, vm8080.MOVCL)
}

func checkMOVC(t *testing.T, opcodeSetReg vm8080.Word, opcodeToTest vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0x02)
	})
}

func TestMOVCM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.MOVCM, vm8080.HLT, 0xFF}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0xFF)
	})
}

func TestMOVDA(t *testing.T) {
	checkMOVD(t, vm8080.MVIA, vm8080.MOVDA)
}
func TestMOVDB(t *testing.T) {
	checkMOVD(t, vm8080.MVIB, vm8080.MOVDB)
}

func TestMOVDC(t *testing.T) {
	checkMOVD(t, vm8080.MVIC, vm8080.MOVDC)
}

func TestMOVDD(t *testing.T) {
	checkMOVD(t, vm8080.MVID, vm8080.MOVDD)
}

func TestMOVDE(t *testing.T) {
	checkMOVD(t, vm8080.MVIE, vm8080.MOVDE)
}

func TestMOVDH(t *testing.T) {
	checkMOVD(t, vm8080.MVIH, vm8080.MOVDH)
}

func TestMOVDL(t *testing.T) {
	checkMOVD(t, vm8080.MVIL, vm8080.MOVDL)
}

func checkMOVD(t *testing.T, opcodeSetReg vm8080.Word, opcodeToTest vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{opcodeSetReg, 0x02, opcodeToTest}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x02)
	})
}

func TestMOVDM(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIH, 0x00, vm8080.MVIL, 0x06, vm8080.MOVDM, vm8080.HLT, 0xFF}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0xFF)
	})
}

func TestADDA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRA, vm8080.ADDA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x02)
	})
}

func TestADDB(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRB, vm8080.ADDB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDBSetsCarry(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0x01, vm8080.ADDB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestADDBSetsFlags(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x01, vm8080.MVIB, 0x01, vm8080.ADDB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkFlags(t, v, false, false, false)
	})
}

func TestADDC(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRC, vm8080.ADDC}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDCSetsCarry(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIC, 0x01, vm8080.ADDC}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestADDD(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRD, vm8080.ADDD}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestADDDSetsCarry(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVID, 0x01, vm8080.ADDD}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if !v.SR.Carry {
			t.Error("Expected Carry flag ON, but was OFF")
		}
	})
}

func TestSUBA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRA, vm8080.INRA, vm8080.INRA, vm8080.SUBA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestSUBBHasCorrectResult(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x02, vm8080.MVIB, 0x01, vm8080.SUBB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestSUBBWithEqualBSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x01, vm8080.MVIB, 0x01, vm8080.SUBB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
		checkFlags(t, v, false, true, false)
	})
}

func TestSUBBWitBLessThanSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x02, vm8080.MVIB, 0x01, vm8080.SUBB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
		checkFlags(t, v, false, false, false)
	})
}

func TestSUBBWitBGreaterThanSetsCorrectFlags(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, vm8080.MVIB, 0x01, vm8080.SUBB}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
		checkFlags(t, v, true, false, true)
	})
}

func checkFlags(t *testing.T, v *vm8080.VirtualMachine, sign vm8080.Flag, zero vm8080.Flag, carry vm8080.Flag) {
	checkFlag(t, v.SR.Sign, sign, "Sign")
	checkFlag(t, v.SR.Zero, zero, "Zero")
	checkFlag(t, v.SR.Carry, carry, "Carry")
}

func checkFlag(t *testing.T, flag vm8080.Flag, expected vm8080.Flag, name string) {
	if flag != expected {
		t.Errorf("%s flag has unexpected value; expected '%t' and has '%t'.", name, expected, flag)
	}
}

func TestSUBC(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRA, vm8080.INRA, vm8080.INRC, vm8080.SUBC}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestSUBD(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.INRA, vm8080.INRA, vm8080.INRD, vm8080.SUBD}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x01)
	})
}

func TestMVIA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestMVIB(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIB, 0xff}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0xff)
	})
}

func TestMVIC(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIC, 0xff}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.C, 0xff)
	})
}

func TestMVID(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVID, 0xff}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0xff)
	})
}

func TestCMA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.CMA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestANAA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.ANAA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestANAB(t *testing.T) {
	checkANA(t, vm8080.MVIB, vm8080.ANAB)
}

func TestANAC(t *testing.T) {
	checkANA(t, vm8080.MVIC, vm8080.ANAC)
}

func TestANAD(t *testing.T) {
	checkANA(t, vm8080.MVID, vm8080.ANAD)
}

func TestANAE(t *testing.T) {
	checkANA(t, vm8080.MVIE, vm8080.ANAE)
}

func TestANAH(t *testing.T) {
	checkANA(t, vm8080.MVIH, vm8080.ANAH)
}

func TestANAL(t *testing.T) {
	checkANA(t, vm8080.MVIL, vm8080.ANAL)
}

func checkANA(t *testing.T, opcodeMVI vm8080.Word, opcodeANA vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xf0, opcodeMVI, 0xf0, opcodeANA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xf0)
	})
}

func TestXRAA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xf0, vm8080.XRAA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x00)
	})
}

func TestXRAB(t *testing.T) {
	checkXRA(t, vm8080.MVIB, vm8080.XRAB)
}

func TestXRAC(t *testing.T) {
	checkXRA(t, vm8080.MVIC, vm8080.XRAC)
}

func TestXRAD(t *testing.T) {
	checkXRA(t, vm8080.MVID, vm8080.XRAD)
}

func TestXRAE(t *testing.T) {
	checkXRA(t, vm8080.MVIE, vm8080.XRAE)
}

func TestXRAH(t *testing.T) {
	checkXRA(t, vm8080.MVIH, vm8080.XRAH)
}

func TestXRAL(t *testing.T) {
	checkXRA(t, vm8080.MVIL, vm8080.XRAL)
}

func checkXRA(t *testing.T, opcodeMVI vm8080.Word, opcodeXRA vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xf0, opcodeMVI, 0xff, opcodeXRA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0x0f)
	})

}

func TestORAA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.ORAA}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func TestORAB(t *testing.T) {
	checkOR(t, vm8080.MVIB, vm8080.ORAB)
}

func TestORAC(t *testing.T) {
	checkOR(t, vm8080.MVIC, vm8080.ORAC)
}

func TestORAD(t *testing.T) {
	checkOR(t, vm8080.MVID, vm8080.ORAD)
}

func TestORAE(t *testing.T) {
	checkOR(t, vm8080.MVIE, vm8080.ORAE)
}

func TestORAH(t *testing.T) {
	checkOR(t, vm8080.MVIH, vm8080.ORAH)
}

func TestORAL(t *testing.T) {
	checkOR(t, vm8080.MVIL, vm8080.ORAL)
}

func checkOR(t *testing.T, opcodeMVI vm8080.Word, opcodeOR vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xf0, opcodeMVI, 0x0f, opcodeOR}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xff)
	})
}

func checkRegHasValue(t *testing.T, r *vm8080.Register, v vm8080.Word) {
	if !r.Has(v) {
		t.Errorf("Register %s has value %x but we were expecting %x", r.Name, r.Value, v)
	}
}

func TestCMPA(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.CMPA}, func(v *vm8080.VirtualMachine, t *testing.T) {
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
	testCMPXWithXLessThanA(t, vm8080.MVIB, vm8080.CMPB)
}

func TestCMPBWithBEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, vm8080.MVIB, vm8080.CMPB)
}

func TestCMPBWithBGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, vm8080.MVIB, vm8080.CMPB)
}

func TestCMPCWithCLessThanA(t *testing.T) {
	testCMPXWithXLessThanA(t, vm8080.MVIC, vm8080.CMPC)
}

func TestCMPCWithCEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, vm8080.MVIC, vm8080.CMPC)
}

func TestCMPCWithCGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, vm8080.MVIC, vm8080.CMPC)
}

func TestCMPDWithDLessThanA(t *testing.T) {
	testCMPXWithXLessThanA(t, vm8080.MVID, vm8080.CMPD)
}

func TestCMPDWithDEqualToA(t *testing.T) {
	testCMPXWithXEqualToA(t, vm8080.MVID, vm8080.CMPD)
}

func TestCMPDWithDGreaterThanA(t *testing.T) {
	testCMPXWithXGreaterThanA(t, vm8080.MVID, vm8080.CMPD)
}

func testCMPXWithXLessThanA(t *testing.T, instrMovX vm8080.Word, instrCmp vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, instrMovX, 0x01, instrCmp}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.SR.Zero {
			t.Error("Expected Z flag OFF but was ON")
		}
		if v.SR.Sign {
			t.Error("Expected S flag OFF but was ON")
		}
	})
}

func testCMPXWithXEqualToA(t *testing.T, instrMovX vm8080.Word, instrCmp vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, instrMovX, 0xff, instrCmp}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if !v.SR.Zero {
			t.Error("Expected Z flag ON but was OFF")
		}
		if v.SR.Sign {
			t.Error("Expected S flag OFF but was ON")
		}
	})
}

func testCMPXWithXGreaterThanA(t *testing.T, instrMovX vm8080.Word, instrCmp vm8080.Word) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, instrMovX, 0x01, instrCmp}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.SR.Zero {
			t.Error("Expected Z flag OFF but was ON")
		}
		if !v.SR.Sign {
			t.Error("Expected S flag ON but was OFF")
		}
	})
}

func TestNOPAndHLT(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.NOP, vm8080.NOP, vm8080.HLT, vm8080.NOP}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x02)
	})
}

func TestJMP(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.JMP, 0x05, 0x00, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x05)
	})
}

func TestJNZJumpsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0x0f, vm8080.CMPB, vm8080.JNZ, 0x11, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJNZDoesntJumpIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0xff, vm8080.CMPB, vm8080.JNZ, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJZJumpsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0xff, vm8080.CMPB, vm8080.JZ, 0x11, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJZDoesntJumpIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0x0f, vm8080.CMPB, vm8080.JZ, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJNCJumpsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x0f, vm8080.MVIB, 0x0f, vm8080.ADDB, vm8080.JNC, 0x11, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJNCDoesntJumpIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0x01, vm8080.ADDB, vm8080.JNC, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJCJumpsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0xff, vm8080.MVIB, 0x0f, vm8080.ADDB, vm8080.JC, 0x11, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x11)
	})
}

func TestJCDoesntJumpIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x0f, vm8080.MVIB, 0x0f, vm8080.ADDB, vm8080.JC, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJPJumpsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, vm8080.CMPA, vm8080.JP, 0x09, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x09)
	})
}

func TestJPDoesntJumpIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, vm8080.MVIB, 0xff, vm8080.CMPB, vm8080.JP, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestJMJumpsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, vm8080.MVIB, 0xff, vm8080.CMPB, vm8080.JM, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x10)
	})
}

func TestJMDoesntJumpIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.MVIA, 0x00, vm8080.MVIB, 0x00, vm8080.CMPB, vm8080.JM, 0x10, 0x00, vm8080.HLT, vm8080.NOP, vm8080.NOP, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x08)
	})
}

func TestLXIB(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXIB, 0x11, 0x22, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.B, 0x22)
		checkRegHasValue(t, &v.Regs.C, 0x11)
	})
}
func TestLXID(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXID, 0x11, 0x22, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x22)
		checkRegHasValue(t, &v.Regs.E, 0x11)
	})
}

func TestLXIH(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXIH, 0x11, 0x22, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.H, 0x22)
		checkRegHasValue(t, &v.Regs.L, 0x11)
	})
}
func TestLXISP(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0xAA, 0xFF, vm8080.HLT}, func(v *vm8080.VirtualMachine, t *testing.T) {
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
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0D, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.MVIA, 0xFF, vm8080.RET, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.SP != 0x000D {
			t.Errorf("SP should be %X but it's %X", 0x000D, v.SP)
		}
		if v.PC != 0x0006 {
			t.Errorf("CALL to 0x09 might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x06)
		}
	})
}

func TestCNZCallsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x0f, vm8080.MVIB, 0x01, vm8080.CMPB, vm8080.CNZ, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCNZDoesntCallsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x00, vm8080.CMPB, vm8080.CNZ, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCZCallsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x0f, vm8080.MVIB, 0x0f, vm8080.CMPB, vm8080.CZ, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCZDoesntCallsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x0f, vm8080.MVIB, 0x00, vm8080.CMPB, vm8080.CZ, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCNCCallsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x00, vm8080.ADDB, vm8080.CNC, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCNCDoesntCallIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0xff, vm8080.MVIB, 0x01, vm8080.ADDB, vm8080.CNC, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCCCallsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0xff, vm8080.MVIB, 0x01, vm8080.ADDB, vm8080.CC, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCCDoesntCallIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x00, vm8080.ADDB, vm8080.CC, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCPCallsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x01, vm8080.MVIB, 0x01, vm8080.ADDB, vm8080.CP, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCPDoesntCallIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x01, vm8080.SUBB, vm8080.CP, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C shouldn't have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestCMCallsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x01, vm8080.SUBB, vm8080.CM, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000C {
			t.Errorf("CALL to 0x0C might not have ocurred, since we ended up in %X but we expected to end up in %X", v.PC, 0x000C)
		}
	})
}

func TestCMDoesntCallIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x10, 0x00, vm8080.MVIA, 0x00, vm8080.MVIB, 0x01, vm8080.ADDB, vm8080.CM, 0x0C, 0x00, vm8080.HLT, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000B {
			t.Errorf("CALL to 0x0C shouldn't have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000B)
		}
	})
}

func TestRNZReturnsIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRB, vm8080.ADDB, vm8080.RNZ, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x0006 {
			t.Errorf("RNZ might not have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x0006)
		}
	})
}

func TestRNZDoesntReturnIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRA, vm8080.SUBA, vm8080.RNZ, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000A {
			t.Errorf("RNZ might have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRZReturnsIfZIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRA, vm8080.SUBA, vm8080.RZ, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x0006 {
			t.Errorf("RZ might not have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRZDoesntReturnIfZIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRB, vm8080.ADDB, vm8080.RZ, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		if v.PC != 0x000A {
			t.Errorf("RZ might have ocurred; we ended up in %X but we expected to end up in %X", v.PC, 0x000A)
		}
	})
}

func TestRNCReturnsIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0C, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.RNC, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRNCDoesntReturnIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0F, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.MVIA, 0xFF, vm8080.ADDA, vm8080.RNC, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x000B)
	})
}

func TestRCReturnsIfCIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0F, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.MVIA, 0xFF, vm8080.ADDA, vm8080.RC, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRCDoesntReturnIfCIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0C, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.RC, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0008)
	})
}

func TestRPReturnsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0D, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRA, vm8080.RP, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRPDoesntReturnIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRB, vm8080.SUBB, vm8080.RP, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x000A)
	})
}

func TestRMReturnsIfSIsOn(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0E, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.INRB, vm8080.SUBB, vm8080.RM, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0006)
	})
}

func TestRMDoesntReturnsIfSIsOff(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x0C, 0x00, vm8080.CALL, 0x07, 0x00, vm8080.HLT, vm8080.RM, vm8080.HLT, 0x00, 0x00, 0x00, 0x00}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, 0x0008)
	})
}

func TestRST1JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST1, 0x08)
}

func TestRST2JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST2, 0x10)
}

func TestRST3JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST3, 0x18)
}

func TestRST4JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST4, 0x20)
}

func TestRST5JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST5, 0x28)
}

func TestRST6JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST6, 0x30)
}
func TestRST7JumpsAndPushesPCIntoStack(t *testing.T) {
	checkRSTJumpsToCorrectAddress(t, vm8080.RST7, 0x38)
}

func checkRSTJumpsToCorrectAddress(t *testing.T, opcode vm8080.Word, expectedFinishAddr uint16) {
	mem := vm8080.MakeMemory()
	mem.Load(map[uint16]vm8080.Memory{
		0x0000: setStackPointerToLastAddress(&mem),
		0x0003: vm8080.Memory{opcode, vm8080.HLT},
		0x0008: vm8080.Memory{vm8080.HLT},
		0x0010: vm8080.Memory{vm8080.HLT},
		0x0018: vm8080.Memory{vm8080.HLT},
		0x0020: vm8080.Memory{vm8080.HLT},
		0x0028: vm8080.Memory{vm8080.HLT},
		0x0030: vm8080.Memory{vm8080.HLT},
		0x0038: vm8080.Memory{vm8080.HLT},
	})
	runAndCheckVM(t, &mem, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkPCisAt(t, v, expectedFinishAddr)
	})
}

func setStackPointerToLastAddress(m *vm8080.Memory) vm8080.Memory {
	lastAddress := m.LastAddress()
	return vm8080.Memory{vm8080.LXISP, vm8080.Word(vm8080.LSB(lastAddress)), vm8080.Word(vm8080.MSB(lastAddress))}
}

func TestLastAddress(t *testing.T) {
	mem := vm8080.MakeMemory()
	lastAddress := mem.LastAddress()
	expected := uint16((1024 * 64) - 1)
	if lastAddress != expected {
		t.Errorf("Last address should be %X but it's %X", expected, lastAddress)
	}
}

func TestLSB(t *testing.T) {
	resultIsExpected(t, vm8080.LSB, 0xFFAE, 0xAE)
}

func TestMSB(t *testing.T) {
	resultIsExpected(t, vm8080.MSB, 0xAEFF, 0xAE)
}

func TestPUSHB(t *testing.T) {
	checkPUSHRegisters(t, vm8080.MVIB, vm8080.MVIC, vm8080.PUSHB)
}

func TestPUSHD(t *testing.T) {
	checkPUSHRegisters(t, vm8080.MVID, vm8080.MVIE, vm8080.PUSHD)
}

func TestPUSHH(t *testing.T) {
	checkPUSHRegisters(t, vm8080.MVIH, vm8080.MVIL, vm8080.PUSHH)
}

func checkPUSHRegisters(t *testing.T, opcodeSetRegister1 vm8080.Word,
	opcodeSetRegister2 vm8080.Word, opcodePush vm8080.Word) {
	checkPUSH(t,
		vm8080.Memory{opcodeSetRegister1, 0x01, opcodeSetRegister2, 0x02, opcodePush, vm8080.HLT},
		0x01,
		0x02)
}

func TestPUSHPSW(t *testing.T) {
	// We need a program that sets some distinctive values in SR.
	checkPUSH(t,
		vm8080.Memory{vm8080.MVIA, 0x00, vm8080.MVIB, 0x01, vm8080.SUBB, vm8080.PUSHPSW, vm8080.HLT},
		0x83,
		0xFF)
}

func checkPUSH(t *testing.T, program vm8080.Memory, expectedInStack1 vm8080.Word,
	expectedInStack2 vm8080.Word) {
	mem := vm8080.MakeMemory()
	mem.Load(map[uint16]vm8080.Memory{
		0x0000: setStackPointerToLastAddress(&mem),
		0x0003: program,
	})
	runAndCheckVM(t, &mem, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkMemoryContent(t, v, 0xFFFE, expectedInStack1)
		checkMemoryContent(t, v, 0xFFFD, expectedInStack2)
	})
}

func TestPOPB(t *testing.T) {
	checkPOPRegisters(t, vm8080.POPB, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.B
	}, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.C
	})
}

func TestPOPD(t *testing.T) {
	checkPOPRegisters(t, vm8080.POPD, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.D
	}, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.E
	})
}

func TestPOPH(t *testing.T) {
	checkPOPRegisters(t, vm8080.POPH, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.H
	}, func(v *vm8080.VirtualMachine) *vm8080.Register {
		return &v.Regs.L
	})
}

func checkPOPRegisters(t *testing.T, opcodePOP vm8080.Word,
	fReg1 func(*vm8080.VirtualMachine) *vm8080.Register,
	fReg2 func(*vm8080.VirtualMachine) *vm8080.Register) {
	mem := vm8080.MakeMemory()
	mem.Load(map[uint16]vm8080.Memory{
		0x0000: vm8080.Memory{vm8080.LXISP, 0xFD, 0xFF},
		0x0003: vm8080.Memory{opcodePOP, vm8080.HLT},
		0xFFFD: vm8080.Memory{0x02, 0x01},
	})
	runAndCheckVM(t, &mem, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, fReg1(v), 0x01)
		checkRegHasValue(t, fReg2(v), 0x02)
	})
}

func TestPOPPSW(t *testing.T) {
	mem := vm8080.MakeMemory()
	mem.Load(map[uint16]vm8080.Memory{
		0x0000: vm8080.Memory{vm8080.LXISP, 0xFD, 0xFF},
		0x0003: vm8080.Memory{vm8080.POPPSW, vm8080.HLT},
		0xFFFD: vm8080.Memory{0xFF, 0xD7}, // 0xD7 = b11010111, setting all flags to 1
	})
	runAndCheckVM(t, &mem, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.A, 0xFF)
		checkFlags(t, v, true, true, true)
	})
}

func checkMemoryContent(t *testing.T, v *vm8080.VirtualMachine, addr uint16, value vm8080.Word) {
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

func checkPCisAt(t *testing.T, v *vm8080.VirtualMachine, expectedPC uint16) {
	if v.PC != expectedPC {
		t.Errorf("We expected the PC to be at %X but it's at %X", expectedPC, v.PC)
	}
}

func TestRegistersHL(t *testing.T) {
	regs := vm8080.Registers{H: vm8080.Register{0x00, "H"}, L: vm8080.Register{0x05, "L"}}
	if regs.HL() != 0x0005 {
		t.Errorf("HL expected to be %X but was %X", 0x005, regs.HL())
	}
}

type DummyDevice struct {
	Data vm8080.Word
}

func (d *DummyDevice) Read() vm8080.Word {
	return d.Data
}

func (d *DummyDevice) Write(data vm8080.Word) {
	d.Data = data
}

func TestIN(t *testing.T) {
	v := vm8080.Make8080()
	v.Load(&vm8080.Memory{vm8080.IN, 0x01, vm8080.HLT})
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
	v := vm8080.Make8080()
	v.Load(&vm8080.Memory{vm8080.MVIA, 0xFF, vm8080.OUT, 0x01, vm8080.HLT})
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
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXID, 0x11, 0x22, vm8080.LXIH, 0x33, 0x44, vm8080.XCHG}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.D, 0x44)
		checkRegHasValue(t, &v.Regs.E, 0x33)
		checkRegHasValue(t, &v.Regs.H, 0x22)
		checkRegHasValue(t, &v.Regs.L, 0x11)
	})
}

func TestXTHL(t *testing.T) {
	runAndCheckVM(t, &vm8080.Memory{vm8080.LXISP, 0x08, 0x00, vm8080.LXIH, 0x11, 0x22, vm8080.XTHL, vm8080.HLT, 0x33, 0x44}, func(v *vm8080.VirtualMachine, t *testing.T) {
		checkRegHasValue(t, &v.Regs.H, 0x44)
		checkRegHasValue(t, &v.Regs.L, 0x33)
		checkMemoryContent(t, v, 0x08, 0x11)
		checkMemoryContent(t, v, 0x09, 0x22)
	})
}

func runAndCheckVM(t *testing.T, memory *vm8080.Memory, test func(v *vm8080.VirtualMachine, t *testing.T)) {
	v := vm8080.Make8080()
	v.Load(memory)
	//	fmt.Print("Running\n")
	//	defer fmt.Print("Ended\n")
	err := v.Run()
	if err != nil {
		t.Error(err)
	}
	test(&v, t)
}
