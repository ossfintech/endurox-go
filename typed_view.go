package atmi

/*
** VIEW buffer support - dynamic access
**
** @file typed_view.go
**
** -----------------------------------------------------------------------------
** Enduro/X Middleware Platform for Distributed Transaction Processing
** Copyright (C) 2015, Mavimax, Ltd. All Rights Reserved.
** This software is released under one of the following licenses:
** GPL or Mavimax's license for commercial use.
** -----------------------------------------------------------------------------
** GPL license:
**
** This program is free software; you can redistribute it and/or modify it under
** the terms of the GNU General Public License as published by the Free Software
** Foundation; either version 2 of the License, or (at your option) any later
** version.
**
** This program is distributed in the hope that it will be useful, but WITHOUT ANY
** WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
** PARTICULAR PURPOSE. See the GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License along with
** this program; if not, write to the Free Software Foundation, Inc., 59 Temple
** Place, Suite 330, Boston, MA 02111-1307 USA
**
** -----------------------------------------------------------------------------
** A commercial use license is available from Mavimax, Ltd
** contact@mavimax.com
** -----------------------------------------------------------------------------
 */
import "C"
import (
	"unsafe"
)

///////////////////////////////////////////////////////////////////////////////////
// Buffer def, typedefs
///////////////////////////////////////////////////////////////////////////////////

//UBF Buffer
type TypedVIEW struct {
	Buf *ATMIBuf
}

//Return The ATMI buffer to caller
func (u *TypedVIEW) GetBuf() *ATMIBuf {
	return u.Buf
}

///////////////////////////////////////////////////////////////////////////////////
// VIEW API
///////////////////////////////////////////////////////////////////////////////////

//Allocate the UBF buffer
//@param size	Buffer size in bytes
//@return UBF Handler, ATMI Error
func (ac *ATMICtx) VIEWAlloc(view string, size int64) (TypedVIEW, ATMIError) {
	var err ATMIError
	var buf TypedVIEW
	buf.Buf, err = ac.TpAlloc("VIEW", view, size)
	return buf, err
}

//Get the UBF Handler
func (ac *ATMICtx) CastToVIEW(abuf *ATMIBuf) (*TypedVIEW, ATMIError) {
	var buf TypedVIEW

	//TODO: Check the buffer type!
	buf.Buf = abuf

	return &buf, nil
}

//Return int16 value from buffer
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return int16 val,	 UBF error
func (u *TypedVIEW) BVGetInt16(cname string, occ int) (int16, UBFError) {
	var c_val C.short

	//char *view, char *cname

	if ret := C.OCBvget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)),
		C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_SHORT); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return int16(c_val), nil
}

//Return int64 value from buffer
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return int64 val,	 UBF error
func (u *TypedVIEW) BVGetInt64(bfldid int, occ int) (int64, UBFError) {
	var c_val C.long
	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_LONG); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return int64(c_val), nil
}

//Return int (basicaly C long (int64) casted to) value from buffer
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return int64 val,	 UBF error
func (u *TypedVIEW) BVGetInt(bfldid int, occ int) (int, UBFError) {
	var c_val C.long
	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_INT); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return int(c_val), nil
}

//Return byte (c char) value from buffer
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return byte val, UBF error
func (u *TypedVIEW) BVGetByte(bfldid int, occ int) (byte, UBFError) {
	var c_val C.char
	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_CHAR); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return byte(c_val), nil
}

//Get float value from UBF buffer, see CBget(3)
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return float, UBF error
func (u *TypedVIEW) BVGetFloat32(bfldid int, occ int) (float32, UBFError) {
	var c_val C.float
	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_FLOAT); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return float32(c_val), nil
}

//Get double value
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return double, UBF error
func (u *TypedVIEW) BVGetFloat64(bfldid int, occ int) (float64, UBFError) {
	var c_val C.double
	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), nil, BFLD_DOUBLE); ret != SUCCEED {
		return 0, u.Buf.Ctx.NewUBFError()
	}
	return float64(c_val), nil
}

//Get string value
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return string val, UBF error
func (u *TypedVIEW) BVGetString(bfldid int, occ int) (string, UBFError) {
	var c_len C.BFLDLEN
	c_val := C.malloc(ATMI_MSG_MAX_SIZE)
	c_len = ATMI_MSG_MAX_SIZE

	if nil == c_val {
		return "", NewCustomUBFError(BEUNIX, "Cannot alloc memory")
	}

	defer C.free(c_val)

	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(c_val), &c_len, BFLD_STRING); ret != SUCCEED {
		return "", u.Buf.Ctx.NewUBFError()
	}

	return C.GoString((*C.char)(c_val)), nil
}

//Get string value
//@param bfldid 	Field ID
//@param occ	Occurrance
//@return string val, UBF error
func (u *TypedVIEW) BVGetByteArr(bfldid int, occ int) ([]byte, UBFError) {
	var c_len C.BFLDLEN
	c_val := C.malloc(ATMI_MSG_MAX_SIZE)
	c_len = ATMI_MSG_MAX_SIZE

	if nil == c_val {
		return nil, NewCustomUBFError(BEUNIX, "Cannot alloc memory")
	}

	defer C.free(c_val)

	if ret := C.OCBget(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
		C.BFLDOCC(occ), (*C.char)(c_val), &c_len, BFLD_CARRAY); ret != SUCCEED {
		return nil, u.Buf.Ctx.NewUBFError()
	}

	g_val := make([]byte, c_len)

	for i := 0; i < int(c_len); i++ {
		g_val[i] = byte(*(*C.char)(unsafe.Pointer(uintptr(c_val) + uintptr(i))))
	}

	return g_val, nil
}

//Change field in buffer
//@param	bfldid	Field ID
//@param ival Input value
//@return UBF Error
func (u *TypedVIEW) BVChg(bfldid int, occ int, ival interface{}) UBFError {
	return u.BChgCombined(bfldid, occ, ival, false)
}

//Set the field value. Combined supports change (chg) or add mode
//@param	bfldid	Field ID
//@param occ	Field Occurrance
//@param ival Input value
//@param	 do_add Adding mode true = add, false = change
//@return UBF Error
func (u *TypedVIEW) BVChgCombined(bfldid int, occ int, ival interface{}, do_add bool) UBFError {

	switch ival.(type) {
	case int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64:
		/* Cast the value to integer... */
		var val int64
		switch ival.(type) {
		case int:
			val = int64(ival.(int))
		case int8:
			val = int64(ival.(int8))
		case int16:
			val = int64(ival.(int16))
		case int32:
			val = int64(ival.(int32))
		case int64:
			val = int64(ival.(int64))
		case uint:
			val = int64(ival.(uint))
		case uint8:
			val = int64(ival.(uint8))
		case uint16:
			val = int64(ival.(uint16))
		case uint32:
			val = int64(ival.(uint32))
		case uint64:
			val = int64(ival.(uint64))
		}
		c_val := C.long(val)

		if do_add {
			if ret := C.OCBadd(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				(*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_LONG); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		} else {
			if ret := C.OCBchg(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_LONG); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		}
	case float32:
		fval := ival.(float32)
		c_val := C.float(fval)
		if do_add {
			if ret := C.OCBadd(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				(*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_FLOAT); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		} else {
			if ret := C.OCBchg(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_FLOAT); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		}
	case float64:
		dval := ival.(float64)
		c_val := C.double(dval)
		if do_add {
			if ret := C.OCBadd(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				(*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_DOUBLE); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		} else {
			if ret := C.OCBchg(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				C.BFLDOCC(occ), (*C.char)(unsafe.Pointer(unsafe.Pointer(&c_val))), 0, BFLD_DOUBLE); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		}
	case string:
		str := ival.(string)
		c_val := C.CString(str)
		defer C.free(unsafe.Pointer(c_val))
		if do_add {
			if ret := C.OCBadd(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				c_val, 0, BFLD_STRING); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		} else {
			if ret := C.OCBchg(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				C.BFLDOCC(occ), c_val, 0, BFLD_STRING); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		}
	case []byte:
		arr := ival.([]byte)
		c_len := C.BFLDLEN(len(arr))
		c_arr := (*C.char)(unsafe.Pointer(&arr[0]))

		if do_add {
			if ret := C.OCBadd(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				c_arr, c_len, BFLD_CARRAY); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		} else {
			if ret := C.OCBchg(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid),
				C.BFLDOCC(occ), c_arr, c_len, BFLD_CARRAY); ret != SUCCEED {
				return u.Buf.Ctx.NewUBFError()
			}
		}
		/*
				- Currently not supported!
			case fmt.Stringer:
				str := ival.(fmt.Stringer).String()
				c_val := C.CString(str)
				defer C.free(unsafe.Pointer(c_val))
				if ret := C.CBchg((*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr), C.BFLDID(bfldid),
					C.BFLDOCC(occ), c_val, 0, BFLD_STRING); ret != SUCCEED {
					return NewUBFError()
				}
		*/
	default:
		/* TODO: Possibly we could take stuff from println to get string val... */
		return NewCustomUBFError(BEINVAL, "Cannot determine field type")
	}

	return nil
}

//Get the number of field occurrances in buffer
//@param bfldid	Field ID
//@return count (or -1 on error), UBF error
func (u *TypedVIEW) BVOccur(cname string) (int, UBFError) {
	c_ret := C.OBoccur(&u.Buf.Ctx.c_ctx,
		(*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)), C.BFLDID(bfldid))

	if FAIL == c_ret {
		return FAIL, u.Buf.Ctx.NewUBFError()
	}

	return int(c_ret), nil
}

//Get the total buffer size
//@return bufer size, UBF error
func (u *TypedVIEW) BVSizeof() (int64, UBFError) {
	c_ret := C.OBsizeof(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)))

	if FAIL == c_ret {
		return FAIL, u.Buf.Ctx.NewUBFError()
	}

	return int64(c_ret), nil
}

//Delete field (all occurrances) from buffer
//@param bfldid field ID
//@return UBF error
func (u *TypedVIEW) BVSetOccur(cname string, occ int) UBFError {
	if ret := C.OBdelall(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)),
		C.BFLDID(bfldid)); SUCCEED != ret {
		return u.Buf.Ctx.NewUBFError()
	}
	return nil
}

//Allocate the new UBF buffer
//NOTE: realloc or other ATMI ops you can do with TypedVIEW.Buf
//@param size - buffer size
//@return Typed UBF, ATMI error
func (ac *ATMICtx) NewVIEW(view stirng, size int64) (*TypedVIEW, ATMIError) {

	var buf TypedVIEW

	if ptr, err := ac.TpAlloc("UBF", "", size); nil != err {
		return nil, err
	} else {
		buf.Buf = ptr
		buf.Buf.Ctx = ac
		return &buf, nil
	}
}

//Converts string JSON buffer passed in 'buffer' to UBF buffer. This function will
//automatically allocate the free space in UBF to fit the JSON. The size will be
//determinated by string length. See tpjsontoubf(3) C call for more information.
//@param buffer	String buffer containing JSON message. The format must be one level
//JSON containing UBF_FIELD:Value. The value can be array, then it is loaded into
//occurrences.
//@return UBFError ('BEINVAL' if failed to convert, 'BMALLOC' if buffer resize failed)
func (u *TypedVIEW) TpJSONToVIEW(buffer string) UBFError {
	size := int64(len(buffer))
	sizeof, _ := u.BSizeof()
	unused, _ := u.BUnused()
	alloc := size - unused

	c_buffer := C.CString(buffer)

	defer C.free(unsafe.Pointer(c_buffer))

	u.Buf.Ctx.ndrxLog(LOG_INFO, "Data size: %d, UBF sizeof: %d, "+
		"unused: %d, about to alloc (if >0) %d",
		size, sizeof, unused, alloc)

	if alloc > 0 {
		if err := u.TpRealloc(sizeof + alloc); nil != err {
			return NewCustomUBFError(BMALLOC, err.Message())
		}
	}

	if ret := C.Otpjsontoubf(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)),
		c_buffer); ret != 0 {
		return NewCustomUBFError(BEINVAL, "Failed to convert JSON 2 UBF "+
			"(tpjsontoubf() failed see UBF/ATMI logs")
	}

	return nil
}

//Convert given UBF buffer to JSON block, see tpubftojson(3) C call
//Output string is automatically allocated
//@return JSON string (if converted ok), ATMIError in case of failure. More detailed
//infos in case of error is found in 'ubf' and 'ndrx' facility logs.
func (u *TypedVIEW) TpVIEWToJSON() (string, ATMIError) {

	used, _ := u.BUsed()

	ret_size := used * 10

	u.Buf.Ctx.ndrxLog(LOG_INFO, "TpUBFToJSON: used %d allocating %d", used, ret_size)

	c_buffer := C.malloc(C.size_t(ret_size))

	if nil == c_buffer {
		return "", NewCustomUBFError(BEUNIX, "Cannot alloc memory")
	}

	defer C.free(c_buffer)

	if ret := C.Otpubftojson(&u.Buf.Ctx.c_ctx, (*C.UBFH)(unsafe.Pointer(u.Buf.C_ptr)),
		(*C.char)(unsafe.Pointer(c_buffer)), C.int(ret_size)); ret != 0 {
		return "", NewCustomUBFError(BEINVAL, "Failed to convert UBF2JSON "+
			"(tpubftojson() failed see UBF/ATMI logs")
	}

	return C.GoString((*C.char)(c_buffer)), nil

}

///////////////////////////////////////////////////////////////////////////////////
// Wrappers for memory management
///////////////////////////////////////////////////////////////////////////////////

func (v *TypedVIEW) TpRealloc(size int64) ATMIError {
	return v.Buf.TpRealloc(size)
}