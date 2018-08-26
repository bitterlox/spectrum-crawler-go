package util

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func DecodeHex(str string) uint64 {
	if len(str) < 2 {
		//log.Errorf("Invalid string: %v", str)
		return 0
	}
	if str == "0x0" || len(str) == 0 {
		return 0
	}

	if str[:2] == "0x" {
		str = str[2:]
	}

	// SOLUTION
	// https://play.golang.org/p/winrs9A7YHP
	//
	// if str[:2] == "0x" {
	// 	str = str[2:]
	// }

	i, err := strconv.ParseUint(str[2:], 16, 64)

	if err != nil {
		log.Errorf("Couldn't decode hex (%v): %v", str, err)
		return 0
	}

	// i, err := hexutil.DecodeBig(str)
	//
	// if err != nil {
	// 	if err == hexutil.ErrLeadingZero {
	//
	// 		// Backup conversion for 0-padded strings
	//
	// 		log.Debugf("util: decodeHex: using backup conversion method (%v)", str)
	//
	// 		x, error := strconv.ParseInt(str[2:], 16, 64)
	//
	// 		if error != nil {
	// 			log.Errorf("Couldn't decode hex (%v): %v", str, err)
	// 		}
	//
	// 		return big.NewInt(x)
	//
	// 	} else {
	// 		log.Errorf("Couldn't decode hex (%v): %v", str, err)
	// 		return big.NewInt(0)
	// 	}
	// }
	return i
}

func InputParamsToAddress(str string) string {
	return "0x" + strings.ToLower(str[24:])
}

func CaculateBlockReward(height uint64, uncleNo int) uint64 {
	baseReward := baseBlockReward(height)
	var uncleRewards uint64

	if uncleNo > 0 {
		uncleRewards = baseReward / uint64(32*uncleNo)
	}

	baseReward += uncleRewards
	return baseReward
}

func CaculateUncleReward(height uint64, uncleHeight uint64) uint64 {
	baseReward := baseBlockReward(height)
	uncleRewards := (((uncleHeight + 2) - height) * baseReward) / 2
	if uncleRewards < 0 {
		return 0
	}
	return uncleRewards
}

func baseBlockReward(height uint64) uint64 {
	if height > 2508545 {
		return 1000000000000000000
	} else if height > 2150181 {
		return 2000000000000000000
	} else if height > 1791818 {
		return 3000000000000000000
	} else if height > 1433454 {
		return 4000000000000000000
	} else if height > 1075090 {
		return 5000000000000000000
	} else if height > 716727 {
		return 6000000000000000000
	} else if height > 358363 {
		return 7000000000000000000
	} else if height > 0 {
		return 8000000000000000000
	} else {
		// genesis
		return 0
	}
}
