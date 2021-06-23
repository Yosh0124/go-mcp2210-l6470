package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Yosh0124/go-mcp2210-l6470/config"
	"github.com/Yosh0124/go-mcp2210-l6470/devices/l6470"
	"github.com/Yosh0124/go-mcp2210-l6470/devices/mcp2210"
	"github.com/sirupsen/logrus"
)

func main() {
	// ログをセット
	const filename string = "debug.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Panicln(err.Error())
	}
	logrus.SetOutput(f)
	// 設定値をロード
	conf := new(config.Config)
	if err := conf.Load("./config.yml"); err != nil {
		logrus.Panicln(err.Error())
	}

	// USB接続
	handler, err := mcp2210.New()
	if err != nil {
		logrus.Fatal(err)
	}
	// プログラム終了時に解放
	defer handler.Close()

	scanner := bufio.NewScanner(os.Stdin)

	// Start routine...
	logrus.Println("Established connection.")
	fmt.Println("Start!")
	for {
		fmt.Println("Input COMMAND...")
		fmt.Println("  01 : SetParam")
		fmt.Println("  02 : GetParams")
		fmt.Println("  03 : Run")
		fmt.Println("  04 : Move")
		fmt.Println("  05 : GoToDir")
		fmt.Println("  06 : GoHome")
		fmt.Println("  07 : GoMark")
		fmt.Println("  08 : ResetPos")
		fmt.Println("  09 : ResetDevice")
		fmt.Println("  10 : SoftStop")
		fmt.Println("  11 : HardStop")
		fmt.Println("  12 : SoftHiZ")
		fmt.Println("  13 : HardHiZ")
		fmt.Println("  14 : GetStatus")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "01":
			fmt.Println("Execute SetParams...")
			// パラメータを取得
			fmt.Print("param? : ")
			scanner.Scan()
			pStr := scanner.Text()
			p, e := hex.DecodeString(pStr)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// 値を取得
			fmt.Print("val? : ")
			scanner.Scan()
			vStr := scanner.Text()
			v, e := strconv.ParseUint(vStr, 10, 32)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// レジスタを取得
			reg, e := l6470.GetRegisterFromAddress(p[0])
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// データ送信
			rx, e := l6470.SetParams(handler, uint32(v), reg, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "02":
			fmt.Println("Execute GetParams...")
			// パラメータを取得
			fmt.Print("param? : ")
			scanner.Scan()
			pStr := scanner.Text()
			p, e := hex.DecodeString(pStr)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// レジスタを取得
			reg, e := l6470.GetRegisterFromAddress(p[0])
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// データ送信
			rx, e := l6470.GetParams(handler, reg, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "03":
			fmt.Println("Execute Run...")
			// 方向を取得
			fmt.Print("direction? 00 or 01 : ")
			scanner.Scan()
			dStr := scanner.Text()
			d, e := hex.DecodeString(dStr)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			} else if d[0] > 0x02 {
				e = errors.New("invalid direction")
				logrus.Errorln(e.Error())
				continue
			}
			// スピードを取得
			fmt.Print("speed? : ")
			scanner.Scan()
			sStr := scanner.Text()
			s, e := strconv.ParseUint(sStr, 10, 32)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// データ送信
			rx, e := l6470.Run(handler, d[0], uint32(s), *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "04":
			fmt.Println("Execute Move...")
			// 方向を取得
			fmt.Print("direction? 00 or 01 : ")
			scanner.Scan()
			dStr := scanner.Text()
			d, e := hex.DecodeString(dStr)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			} else if d[0] > 0x02 {
				e = errors.New("invalid direction")
				logrus.Errorln(e.Error())
				continue
			}
			// ステップを取得
			fmt.Print("step? : ")
			scanner.Scan()
			sStr := scanner.Text()
			s, e := strconv.ParseUint(sStr, 10, 32)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// データ送信
			rx, e := l6470.Move(handler, d[0], uint32(s), *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "05":
			fmt.Println("Execute GoToDir...")
			// 方向を取得
			fmt.Print("direction? 00 or 01 : ")
			scanner.Scan()
			dStr := scanner.Text()
			d, e := hex.DecodeString(dStr)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			} else if d[0] > 0x02 {
				e = errors.New("invalid direction")
				logrus.Errorln(e.Error())
				continue
			}
			// ポジションを取得
			fmt.Print("position? : ")
			scanner.Scan()
			pStr := scanner.Text()
			p, e := strconv.ParseUint(pStr, 10, 32)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			// データ送信
			rx, e := l6470.GoToDir(handler, d[0], uint32(p), *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "06":
			fmt.Println("Execute GoHome...")
			// データ送信
			rx, e := l6470.GoHome(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "07":
			fmt.Println("Execute GoMark...")
			// データ送信
			rx, e := l6470.GoMark(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "08":
			fmt.Println("Execute ResetPos...")
			// データ送信
			rx, e := l6470.ResetPos(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "09":
			fmt.Println("Execute ResetDevice...")
			// データ送信
			rx, e := l6470.ResetDevice(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "10":
			fmt.Println("Execute SoftStop...")
			// データ送信
			rx, e := l6470.SoftStop(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "11":
			fmt.Println("Execute HardStop...")
			// データ送信
			rx, e := l6470.HardStop(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "12":
			fmt.Println("Execute SoftHiZ...")
			// データ送信
			rx, e := l6470.SoftHiZ(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "13":
			fmt.Println("Execute HardHiZ...")
			// データ送信
			rx, e := l6470.HardHiZ(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		case "14":
			fmt.Println("Execute GetStatus...")
			// データ送信
			rx, e := l6470.GetStatus(handler, *conf)
			if e != nil {
				logrus.Errorln(e.Error())
				continue
			}
			fmt.Println("return val : ", rx)

		default:
			fmt.Println("Invalid input value:")
		}
	}
}
