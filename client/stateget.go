package client

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/webapi/v1/routes"
)

// StateGet fetches the raw value associated with the given key in the chain state
func (c *WaspClient) StateGet(chainID *isc.ChainID, key string) ([]byte, error) {
	var res []byte
	if err := c.do(http.MethodGet, routes.StateGet(chainID.String(), hexutil.Encode([]byte(key))), nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}
