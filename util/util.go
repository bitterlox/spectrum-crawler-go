package util

import (
	"math/big"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/ubiq/go-ubiq/common/hexutil"
)

func DecodeHex(str string) *big.Int {
	if len(str) < 2 {
		//log.Errorf("Invalid string: %v", str)
		return big.NewInt(0)
	}
	if str == "0x0" || len(str) == 0 {
		return big.NewInt(0)
	}

	if str[:2] != "0x" {
		str = "0x" + str
	}

	// SOLUTION
	// https://play.golang.org/p/winrs9A7YHP
	//
	// if str[:2] == "0x" {
	// 	str = str[2:]
	// }

	i, err := hexutil.DecodeBig(str)

	if err != nil {
		if err == hexutil.ErrLeadingZero {

			// Backup conversion for 0-padded strings

			log.Debugf("util: decodeHex: using backup conversion method (%v)", str)

			x, error := strconv.ParseInt(str[2:], 16, 64)

			if error != nil {
				log.Errorf("Couldn't decode hex (%v): %v", str, err)
			}

			return big.NewInt(x)

		} else {
			log.Errorf("Couldn't decode hex (%v): %v", str, err)
			return big.NewInt(0)
		}
	}
	return i
}

func InputParamsToAddress(str string) string {
	return "0x" + strings.ToLower(str[24:])
}

func CaculateBlockReward(height int64, uncleNo int) *big.Int {
	baseReward := baseBlockReward(height)
	uncleRewards := big.NewInt(0)

	if uncleNo > 0 {
		uncleRewards.Div(baseReward, big.NewInt(int64(32*uncleNo)))
	}

	baseReward.Add(baseReward, uncleRewards)
	return baseReward
}

func CaculateUncleReward(height int64, uncleHeight int64) *big.Int {
	baseReward := baseBlockReward(height)
	uncleRewards := big.NewInt((((uncleHeight + 2) - height) * baseReward.Int64()) / 2)
	r := uncleRewards.Cmp(big.NewInt(0))
	if r == -1 {
		return big.NewInt(0)
	}
	return uncleRewards
}

func baseBlockReward(height int64) *big.Int {
	if height > 2508545 {
		return big.NewInt(1000000000000000000)
	} else if height > 2150181 {
		return big.NewInt(2000000000000000000)
	} else if height > 1791818 {
		return big.NewInt(3000000000000000000)
	} else if height > 1433454 {
		return big.NewInt(4000000000000000000)
	} else if height > 1075090 {
		return big.NewInt(5000000000000000000)
	} else if height > 716727 {
		return big.NewInt(6000000000000000000)
	} else if height > 358363 {
		return big.NewInt(7000000000000000000)
	} else if height > 0 {
		return big.NewInt(8000000000000000000)
	} else {
		// genesis
		return big.NewInt(0)
	}
}
