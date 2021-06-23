package l6470

import (
	"encoding/binary"
	"errors"
	"math"
	"time"

	"github.com/Yosh0124/go-mcp2210-l6470/config"
	"github.com/Yosh0124/go-mcp2210-l6470/devices/mcp2210"
	"github.com/sirupsen/logrus"
)

// パラメータをセットする
func SetParams(m *mcp2210.Mcp2210, val uint32, reg Register, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// コマンド 1byte + データバイト配列
	valLength := int(math.Ceil(float64(reg.Length) / 8))
	dataLength := 1 + valLength
	if dataLength > 4 {
		err := errors.New("invalid data length")
		logrus.Error(err.Error(), dataLength)
		return nil, err
	}
	// 送信用データ列
	txData := make([]byte, dataLength)
	// コマンドバイトの設定
	txData[0] = 0x00 | reg.Address
	// データをバイト配列へ変換
	tmpArray := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpArray, val)
	// データバイト配列の保存
	byteOffset := 4 - valLength
	for i := 0; i < valLength; i++ {
		txData[i+1] = tmpArray[byteOffset+i]
	}

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// パラメータを取得する
func GetParams(m *mcp2210.Mcp2210, reg Register, c config.Config) ([]byte, error) {
	txData := make([]byte, 4)
	txData[0] = 0x20 | reg.Address

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// 指定した向きとスピードで回転
func Run(m *mcp2210.Mcp2210, dir uint8, speed uint32, c config.Config) ([]byte, error) {
	// 方向の数値をチェック
	if dir > 0x01 {
		err := errors.New("invalid direction arg")
		logrus.Errorln(err.Error())
		return nil, err
	}

	// =================
	// コマンドを作成
	// コマンド 1byte + データバイト配列
	reg := SPEED
	valLength := int(math.Ceil(float64(reg.Length) / 8))
	dataLength := 1 + valLength
	if dataLength > 4 {
		err := errors.New("invalid data length")
		logrus.Error(err.Error(), dataLength)
		return nil, err
	}
	// 送信用データ列
	txData := make([]byte, dataLength)
	// コマンドバイトの設定
	txData[0] = 0x50 | dir
	// データをバイト配列へ変換
	tmpArray := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpArray, speed)
	// データバイト配列の保存
	byteOffset := 4 - valLength
	for i := 0; i < valLength; i++ {
		txData[i+1] = tmpArray[byteOffset+i]
	}

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// 方向とステップ数を指定して移動
func Move(m *mcp2210.Mcp2210, dir uint8, step uint32, c config.Config) ([]byte, error) {
	// 方向の数値をチェック
	if dir > 0x01 {
		err := errors.New("invalid direction arg")
		logrus.Errorln(err.Error())
		return nil, err
	}

	// =================
	// コマンドを作成
	// コマンド 1byte + データバイト配列
	reg := SPEED
	valLength := int(math.Ceil(float64(reg.Length) / 8))
	dataLength := 1 + valLength
	if dataLength > 4 {
		err := errors.New("invalid data length")
		logrus.Error(err.Error(), dataLength)
		return nil, err
	}
	// 送信用データ列
	txData := make([]byte, dataLength)
	// コマンドバイトの設定
	txData[0] = 0x50 | dir
	// データをバイト配列へ変換
	tmpArray := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpArray, step)
	// データバイト配列の保存
	byteOffset := 4 - valLength
	for i := 0; i < valLength; i++ {
		txData[i+1] = tmpArray[byteOffset+i]
	}

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// 方向と場所を指定して移動
func GoToDir(m *mcp2210.Mcp2210, dir uint8, pos uint32, c config.Config) ([]byte, error) {
	// 方向の数値をチェック
	if dir > 0x01 {
		err := errors.New("invalid direction arg")
		logrus.Errorln(err.Error())
		return nil, err
	}

	// =================
	// コマンドを作成
	// コマンド 1byte + データバイト配列
	reg := SPEED
	valLength := int(math.Ceil(float64(reg.Length) / 8))
	dataLength := 1 + valLength
	if dataLength > 4 {
		err := errors.New("invalid data length")
		logrus.Error(err.Error(), dataLength)
		return nil, err
	}
	// 送信用データ列
	txData := make([]byte, dataLength)
	// コマンドバイトの設定
	txData[0] = 0x70 | dir
	// データをバイト配列へ変換
	tmpArray := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpArray, pos)
	// データバイト配列の保存
	byteOffset := 4 - valLength
	for i := 0; i < valLength; i++ {
		txData[i+1] = tmpArray[byteOffset+i]
	}

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// ホーム位置へ戻る
func GoHome(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {

	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0x70

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// マークした場所へ移動
func GoMark(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {

	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0x78

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// ホーム位置に設定
func ResetPos(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xd8

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// デバイスを起動初期状態にリセットする
func ResetDevice(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xc0

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// 回転をゆっくり止める
func SoftStop(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xb0

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// 回転を直ちに止める
func HardStop(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xb8

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// ブリッジをゆっくりHiZにする
func SoftHiZ(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xa0

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// ブリッジを直ちにHiZにする
func HardHiZ(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 1)
	// コマンドバイトの設定
	txData[0] = 0xa8

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	return rxData, nil
}

// ステータスを取得する
func GetStatus(m *mcp2210.Mcp2210, c config.Config) ([]byte, error) {
	// =================
	// コマンドを作成
	// 送信用データ列
	txData := make([]byte, 4)
	// コマンドバイトの設定
	txData[0] = 0xd0

	rxData, err := SendCommand(m, txData, c)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, err
	}

	status := binary.BigEndian.Uint16(rxData[1:3])
	ParseStatus(status)

	return rxData, nil
}

// L6470は1byteずつ送らないといけない
func SendCommand(m *mcp2210.Mcp2210, txData []byte, c config.Config) (rxData []byte, err error) {

	length := len(txData)
	rxData = make([]byte, length)

	for i := 0; i < length; i++ {
		tx := []byte{txData[i]}
		rx, err := m.SendCommand(tx, c)
		if err != nil {
			logrus.Errorln(err.Error())
			return nil, err
		}
		rxData[i] = rx[0]
		time.Sleep(1 * time.Microsecond)
	}

	return rxData, nil
}
