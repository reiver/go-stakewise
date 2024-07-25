package keeper

import (
	"github.com/reiver/go-erorr"
)

// RewardsRoot calls the "rewardsRoot" method on the StakeWise Keeper contract, and return the URI to the associate metadata.
func (receiver Client) RewardsRoot() (string, error) {
	const methodname string = "rewardsRoot"

	results, err := receiver.Call(methodname)
	if nil != err {
		return "", err
	}

	var value string
        {
		var result0 interface{}
		{
			const expected = 1
			var actual int = len(results)

			if expected != actual {
				return "", erorr.Errorf("stakewise: expected call to contract (%s) method %q to return %d result but actually got %d results", receiver.ContractAddress(), methodname, expected, actual)
			}

			result0 = results[0]
		}

		{
			var casted bool

			value, casted = result0.(string)
			if !casted {
				return "", erorr.Errorf("stakewise: expected call to contract (%s) method %q to return result of type %T but actually was %T", receiver.ContractAddress(), methodname, value, result0)
			}
		}
	}

	return value, nil
}
