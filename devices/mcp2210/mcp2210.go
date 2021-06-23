package mcp2210

/*
#cgo windows LDFLAGS: -L . -lmcp2210_dll_um_x64
#include <stdio.h>
#include <stdlib.h>
#include "mcp2210_dll_um.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/Yosh0124/go-mcp2210-l6470/config"
	"github.com/sirupsen/logrus"
)

const (
	DllPath                         string   = "mcp2210_dll_um_x86.dll"
	DEFAULT_VID                     C.ushort = 0x4d8
	DEFAULT_PID                     C.ushort = 0xde
	E_SUCCESS                       C.int    = 0
	E_ERR_INVALID_PARAMETER         C.int    = -2
	E_ERR_BUFFER_TOO_SMALL          C.int    = -3
	E_ERR_NULL                      C.int    = -10
	E_ERR_INVALID_HANDLE_VALUE      C.int    = -30
	E_ERR_NO_SUCH_INDEX             C.int    = -101
	E_ERR_CONNECTION_ALREADY_OPENED C.int    = -106
	E_ERR_CLOSE_FAILED              C.int    = -107
	E_ERR_NO_SUCH_SERIALNR          C.int    = -108
	E_ERR_HID_RW_FILEIO             C.int    = -111
	E_ERR_SPI_EXTERN_MASTER         C.int    = -204
	E_ERR_SPI_TIMEOUT               C.int    = -205
	E_ERR_SPI_XFER_ONGOING          C.int    = -207
	E_ERR_BLOCKED_ACCESS            C.int    = -300
	E_ERR_NVRAM_LOCKED              C.int    = -350
	E_ERR_WRONG_PASSWD              C.int    = -351
	E_ERR_ACCESS_DENIED             C.int    = -352
	E_ERR_NVRAM_PROTECTED           C.int    = -353
	E_ERR_PASSWD_CHANGE             C.int    = -354
	E_ERR_STRING_TOO_LARGE          C.int    = -401
	CSMASK_NOCHANGE                 C.uint   = 0x10000000
)

type Mcp2210 struct {
	Handler unsafe.Pointer
	first   bool
}

// New : MCP2210にUSB接続
func New() (*Mcp2210, error) {
	var devCount = C.Mcp2210_GetConnectedDevCount(DEFAULT_VID, DEFAULT_PID)
	logrus.Println(devCount, " devices found")

	if devCount <= 0 {
		err := errors.New("device not found")
		return nil, err
	}

	var path *C.ushort
	var pathSize C.ulong = 0

	deviceHandle := C.Mcp2210_OpenByIndex(DEFAULT_VID, DEFAULT_PID, 0, path, &pathSize)
	res := C.Mcp2210_GetLastError()
	if res != E_SUCCESS {
		err := errors.New("failed to open connection")
		logrus.Error(err.Error())
		return nil, err
	}

	m := new(Mcp2210)
	m.Handler = deviceHandle
	m.first = true

	return m, nil
}

// New : MCP2210と接続解除
func (m *Mcp2210) Close() {
	C.Mcp2210_Close(m.Handler)
}

func (m *Mcp2210) SendCommand(txData []byte, c config.Config) ([]byte, error) {
	txSize := C.uint(len(txData))
	rxData := make([]byte, txSize)

	logrus.Println("Send data : ", txData)

	// 初回アクセスだけコンフィグが必要
	if m.first {
		logrus.Println("Execute Mcp2210_xferSpiDataEx...")
		res := C.Mcp2210_xferSpiDataEx(
			m.Handler,
			(*C.uchar)(&txData[0]),
			(*C.uchar)(&rxData[0]),
			(*C.uint)(&c.BaudRate),
			(*C.uint)(&txSize),
			(C.uint)(c.CsMask),
			(*C.uint)(&c.IdleCsVal),
			(*C.uint)(&c.ActiveCsVal),
			(*C.uint)(&c.CsToDataDelay),
			(*C.uint)(&c.DataToCsDelay),
			(*C.uint)(&c.DataToDataDelay),
			(*C.uchar)(&c.SpiMode),
		)
		if res != E_SUCCESS {
			err := errors.New("transfer error")
			logrus.Error(err.Error())
			return nil, err
		}
		m.first = false
	} else {
		logrus.Println("Execute Mcp2210_xferSpiData...")
		res := C.Mcp2210_xferSpiData(
			m.Handler,
			(*C.uchar)(&txData[0]),
			(*C.uchar)(&rxData[0]),
			(*C.uint)(&c.BaudRate),
			(*C.uint)(&txSize),
			CSMASK_NOCHANGE,
		)
		if res != E_SUCCESS {
			err := errors.New("transfer error")
			logrus.Error(err.Error())
			return nil, err
		}
	}

	logrus.Println("received data :", rxData)

	return rxData, nil
}
