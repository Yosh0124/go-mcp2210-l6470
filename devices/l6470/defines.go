package l6470

type Register struct {
	Address byte
	Length  uint
}

var (
	// ABS_POS 絶対位置
	ABS_POS = Register{
		Address: 0x01,
		Length:  22,
	}
	// EL_POS 電気的位置
	EL_POS = Register{
		Address: 0x02,
		Length:  9,
	}
	// MARK マークした位置
	MARK = Register{
		Address: 0x03,
		Length:  22,
	}
	// SPEED 現在位置
	SPEED = Register{
		Address: 0x04,
		Length:  20,
	}
	// 加速
	ACC = Register{
		Address: 0x05,
		Length:  12,
	}
	// 減速
	DEC = Register{
		Address: 0x06,
		Length:  12,
	}
	// 最高速度
	MAX_SPEED = Register{
		Address: 0x07,
		Length:  10,
	}
	// 最低速度
	MIN_SPEED = Register{
		Address: 0x08,
		Length:  13,
	}
	// フルステップ時のスピード
	FS_SPD = Register{
		Address: 0x15,
		Length:  10,
	}
	KVAL_HOLD = Register{
		Address: 0x09,
		Length:  8,
	}
	KVAL_RUN = Register{
		Address: 0x0a,
		Length:  8,
	}
	KVAL_ACC = Register{
		Address: 0x0b,
		Length:  8,
	}
	KVAL_DEC = Register{
		Address: 0x0c,
		Length:  8,
	}
	INT_SPEED = Register{
		Address: 0x0d,
		Length:  14,
	}
	ST_SLP = Register{
		Address: 0x0e,
		Length:  8,
	}
	FN_SLP_ACC = Register{
		Address: 0x0f,
		Length:  8,
	}
	FN_SLP_DEC = Register{
		Address: 0x10,
		Length:  8,
	}
	K_THERM = Register{
		Address: 0x11,
		Length:  4,
	}
	ADC_OUT = Register{
		Address: 0x12,
		Length:  5,
	}
	OCD_TH = Register{
		Address: 0x13,
		Length:  4,
	}
	STALL_TH = Register{
		Address: 0x14,
		Length:  7,
	}
	STEP_MODE = Register{
		Address: 0x16,
		Length:  8,
	}
	ALARM_EN = Register{
		Address: 0x17,
		Length:  8,
	}
	CONFIG = Register{
		Address: 0x18,
		Length:  16,
	}
	STATUS = Register{
		Address: 0x19,
		Length:  16,
	}
)

const (
	MASK_STATUS_HIZ         uint16 = 0x0001 // ブリッジがハイ・インピーダンス状態であることを示す
	MASK_STATUS_BUSY        uint16 = 0x0002 // 定速、位置決め、または動作コマンドを実行中の際に[0]にされ、 コマンドが完了した後に解放[1]にされる。
	MASK_STATUS_SW_F        uint16 = 0x0004 // SW入力の状態を報告する。 （閉じられた場合は[1]、オープンの場合は[0]）
	MASK_STATUS_SW_EVN      uint16 = 0x0008 // スイッチ・オンの操作を示す。 （SW入力の立ち下がりエッジ）
	MASK_STATUS_DIR         uint16 = 0x0010 // 現在のモーターの回転方向を示す： [1] = 正方向 , [0] = 逆方向
	MASK_STATUS_MOT_STATUS  uint16 = 0x0060 // モーター電流の状態を示す
	MASK_STATUS_NOTPERF_CMD uint16 = 0x0080 // 発動が[1]で、SPIにより受信したコマンドが、それぞれ、実行できないか、または全く存在しないことを示します
	MASK_STATUS_WRONG_CMD   uint16 = 0x0100 // 発動が[1]で、SPIにより受信したコマンドが、それぞれ、実行できないか、または全く存在しないことを示します
	MASK_STATUS_UVLO        uint16 = 0x0200 // 電圧低下検出
	MASK_STATUS_TH_WRN      uint16 = 0x0400 // 温度警告
	MASK_STATUS_TH_SD       uint16 = 0x0800 // 温度の運転停止
	MASK_STATUS_OCD         uint16 = 0x1000 // 過電流
	MASK_STATUS_STEP_LOSS_A uint16 = 0x2000 // ブリッジAでの脱調検出
	MASK_STATUS_STEP_LOSS_B uint16 = 0x4000 // ブリッジBでの脱調検出
	MASK_STATUS_SCL_MOD     uint16 = 0x8000 // デバイスがステップ・クロック・モードで動作していることを示す
	STATUS_MOT_STOP         uint16 = 0x0000
	STATUS_MOT_ACC          uint16 = 0x0020
	STATUS_MOT_DEC          uint16 = 0x0040
	STATUS_MOT_RUN          uint16 = 0x0060
	STATUS_DIR_POS          uint16 = 0x0010
	STATUS_DIR_NEG          uint16 = 0x0000
)
