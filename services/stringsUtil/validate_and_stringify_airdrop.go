package stringsUtil

import (
	"errors"
	"math/big"
	"strconv"
	"valorize-app/services/ethereum"
)

func ValidateAndStringifyMap(addressToClaimAmountMap [][]string) (string, error) {
	finalStringifiedMap :="["
	for i := 0; i < len(addressToClaimAmountMap); i++ {
		address := addressToClaimAmountMap[i][0]
		claimString := addressToClaimAmountMap[i][1]
		_, ok := new(big.Int).SetString(claimString, 10)
		if !ok || !ethereum.CheckAddress(address) {
			return "", errors.New("invalid value in input for address: " + address + ", claimAmount: " + claimString + "in index: " + strconv.Itoa(i))
		}
		//create a string array of the address and claim amount
		finalStringifiedMap += "[\"" + address + "\",\"" + claimString + "\"]"
		//only add a comma if there are more elements in the map
		if i < len(addressToClaimAmountMap)-1 {
			finalStringifiedMap += ","
		}
	}
	finalStringifiedMap += "]"

	return finalStringifiedMap, nil
}