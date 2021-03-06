package main

import (
	"atmi"
	"fmt"
	"os"
	"ubftab"
)

const (
	SUCCEED = 0
	FAIL    = -1
)

//TESTSVC service
func TESTSVC(ac *atmi.ATMICtx, svc *atmi.TPSVCINFO) {

	ret := SUCCEED

	//Get UBF Handler
	ub, _ := ac.CastToUBF(&svc.Data)

	//Print the buffer to stdout
	fmt.Println("Incoming request:")
	ub.BPrint()

	//Return to the caller
	defer func() {
		if SUCCEED == ret {
			ac.TpReturn(atmi.TPSUCCESS, 0, ub, 0)
		} else {
			ac.TpReturn(atmi.TPFAIL, 0, ub, 0)
		}
	}()

	//Resize buffer, to have some more space
	if err := ub.TpRealloc(1024); err != nil {
		fmt.Printf("Got error: %d:[%s]\n", err.Code(), err.Message())
		ret = FAIL
		return
	}

	s_val, _ := ub.BGetInt16(ubftab.T_SHORT_FLD, 0)

	if s_val != 100 {
		fmt.Printf("Got error, T_SHORT_FLD not 100!")
		ret = FAIL
		return
	}

	//Set some field
	if err := ub.BChg(ubftab.T_STRING_FLD, 0,
		"Hello World from Enduro/X service"); err != nil {
		fmt.Printf("Got error: %d:[%s]\n", err.Code(), err.Message())
		ret = FAIL
		return
	}

	b := []byte{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if err := ub.BChg(ubftab.T_CARRAY_FLD, 0, b); err != nil {
		fmt.Printf("Got error: %d:[%s]\n", err.Code(), err.Message())
		ret = FAIL
		return
	}
}

//Server init
func Init(ac *atmi.ATMICtx) int {

	//Advertize TESTSVC
	if err := ac.TpAdvertise("TESTSVC", "TESTSVC", TESTSVC); err != nil {
		fmt.Println(err)
		return atmi.FAIL
	}

	return atmi.SUCCEED
}

//Server shutdown
func Uninit(ac *atmi.ATMICtx) {
	fmt.Println("Server shutting down...")
}

//Executable main entry point
func main() {
	//Have some context
	ac, err := atmi.NewATMICtx()

	if nil != err {
		fmt.Errorf("Failed to allocate cotnext!", err)
		os.Exit(atmi.FAIL)
	} else {
		//Run as server
		ac.TpRun(Init, Uninit)
		ac.FreeATMICtx()
	}

}
