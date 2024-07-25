package keeper

import (
	_ "embed"
	"io"
	"strings"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/reiver/go-erorr"
)

//go:embed abi.json
var abijson string

var abi ethabi.ABI

func init () {
	var reader io.Reader = strings.NewReader(abijson)

	var err error

	abi, err = ethabi.JSON(reader)
	if nil != err {
		var e error = erorr.Errorf("stakewise: problem parsing 'keeper' contract JSON ABI: %w", err)
		panic(e)
	}
}
