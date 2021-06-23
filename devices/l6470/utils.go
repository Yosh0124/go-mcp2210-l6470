package l6470

import (
	"errors"

	"github.com/sirupsen/logrus"
)

/*
GetRegisterFromAddress コマンドアドレスからレジスタを取得します。
*/
func GetRegisterFromAddress(address uint8) (Register, error) {
	var ret Register
	switch address {
	case 0x01:
		ret = ABS_POS
	case 0x02:
		ret = EL_POS
	case 0x03:
		ret = MARK
	case 0x04:
		ret = SPEED
	case 0x05:
		ret = ACC
	case 0x06:
		ret = DEC
	case 0x07:
		ret = MAX_SPEED
	case 0x08:
		ret = MIN_SPEED
	case 0x09:
		ret = KVAL_HOLD
	case 0x0a:
		ret = KVAL_RUN
	case 0x0b:
		ret = KVAL_ACC
	case 0x0c:
		ret = KVAL_DEC
	case 0x0d:
		ret = INT_SPEED
	case 0x0e:
		ret = ST_SLP
	case 0x0f:
		ret = FN_SLP_ACC
	case 0x10:
		ret = FN_SLP_DEC
	case 0x11:
		ret = K_THERM
	case 0x12:
		ret = ADC_OUT
	case 0x13:
		ret = OCD_TH
	case 0x14:
		ret = STALL_TH
	case 0x15:
		ret = FS_SPD
	case 0x16:
		ret = STEP_MODE
	case 0x17:
		ret = ALARM_EN
	case 0x18:
		ret = CONFIG
	case 0x19:
		ret = STATUS
	default:
		err := errors.New("invalid register")
		return Register{}, err
	}
	logrus.Println("found register :", ret)
	return ret, nil
}

/*
ParseStatus L6470から受信したSTATUSをログに出力します。
*/
func ParseStatus(status uint16) {
	if status&MASK_STATUS_HIZ > 0 {
		logrus.Println("Status HiZ.")
	}
	if status&MASK_STATUS_BUSY > 0 {
		logrus.Println("Status Busy.")
	}
	if status&MASK_STATUS_SW_F > 0 {
		logrus.Println("Switch closed.")
	}
	if status&MASK_STATUS_SW_EVN > 0 {
		logrus.Println("SW Edge is detected.")
	}
	if status&MASK_STATUS_DIR > 0 {
		logrus.Println("Rotation Clockwise.")
	} else {
		logrus.Println("Rotation Conter-clockwise.")
	}
	if status&MASK_STATUS_NOTPERF_CMD > 0 {
		logrus.Println("Invalid Command.")
	}
	if status&MASK_STATUS_WRONG_CMD > 0 {
		logrus.Println("Wrong Command.")
	}
	if status&MASK_STATUS_UVLO > 0 {
		logrus.Println("Status UVLO.")
	}
	if status&MASK_STATUS_TH_WRN > 0 {
		logrus.Println("Thermal Warning.")
	}
	if status&MASK_STATUS_TH_SD > 0 {
		logrus.Println("Thermal Stop.")
	}
	if status&MASK_STATUS_OCD > 0 {
		logrus.Println("Status OCD.")
	}
	if status&MASK_STATUS_STEP_LOSS_A > 0 {
		logrus.Println("Status Step Loss A.")
	}
	if status&MASK_STATUS_STEP_LOSS_B > 0 {
		logrus.Println("Status Step Loss B.")
	}
	if status&MASK_STATUS_SCL_MOD > 0 {
		logrus.Println("Status SCL Mode.")
	}

	motorStatus := status & MASK_STATUS_MOT_STATUS
	switch motorStatus {
	case STATUS_MOT_STOP:
		logrus.Println("Motor is stopped.")
	case STATUS_MOT_ACC:
		logrus.Println("Motor is accelerating.")
	case STATUS_MOT_DEC:
		logrus.Println("Motor is decelerating.")
	case STATUS_MOT_RUN:
		logrus.Println("Motor is running.")
	}
}
