package config

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

type Config struct {
	BaudRate        uint32 `yaml:"baud_rate"`
	IdleCsVal       uint32 `yaml:"idle_cs_val"`
	ActiveCsVal     uint32 `yaml:"active_cs_val"`
	CsToDataDelay   uint32 `yaml:"cs_to_data_delay"`
	DataToDataDelay uint32 `yaml:"data_to_data_delay"`
	DataToCsDelay   uint32 `yaml:"data_to_cs_delay"`
	SpiMode         uint8  `yaml:"spi_mode"`
	CsMask          uint32 `yaml:"cs_mask"`
}
