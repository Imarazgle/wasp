package corecontracts

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/webapi/v2/models"

	"github.com/iotaledger/wasp/packages/webapi/v2/params"

	"github.com/iotaledger/wasp/packages/webapi/v2/apierrors"
	"github.com/labstack/echo/v4"
)

func (c *Controller) getAccounts(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	accounts, err := c.accounts.GetAccounts(chainID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	accountsResponse := &models.AccountListResponse{
		Accounts: make([]string, len(accounts)),
	}

	for k, v := range accounts {
		accountsResponse.Accounts[k] = v.String()
	}

	return e.JSON(http.StatusOK, accountsResponse)
}

func (c *Controller) getTotalAssets(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	assets, err := c.accounts.GetTotalAssets(chainID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	assetsResponse := &models.AssetsResponse{
		BaseTokens: assets.BaseTokens,
		Tokens:     models.MapNativeTokens(assets.Tokens),
	}

	return e.JSON(http.StatusOK, assetsResponse)
}

func (c *Controller) getAccountBalance(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	agentID, err := params.DecodeAgentID(e)
	if err != nil {
		return err
	}

	assets, err := c.accounts.GetAccountBalance(chainID, agentID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	assetsResponse := &models.AssetsResponse{
		BaseTokens: assets.BaseTokens,
		Tokens:     models.MapNativeTokens(assets.Tokens),
	}

	return e.JSON(http.StatusOK, assetsResponse)
}

func (c *Controller) getAccountNFTs(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	agentID, err := params.DecodeAgentID(e)
	if err != nil {
		return err
	}

	nfts, err := c.accounts.GetAccountNFTs(chainID, agentID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	nftsResponse := &models.AccountNFTsResponse{
		NFTIDs: make([]string, len(nfts)),
	}

	for k, v := range nfts {
		nftsResponse.NFTIDs[k] = v.ToHex()
	}

	return e.JSON(http.StatusOK, nftsResponse)
}

func (c *Controller) getAccountNonce(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	agentID, err := params.DecodeAgentID(e)
	if err != nil {
		return err
	}

	nonce, err := c.accounts.GetAccountNonce(chainID, agentID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	nonceResponse := &models.AccountNonceResponse{
		Nonce: nonce,
	}

	return e.JSON(http.StatusOK, nonceResponse)
}

func (c *Controller) getNFTData(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	nftID, err := params.DecodeNFTID(e)
	if err != nil {
		return err
	}

	nftData, err := c.accounts.GetNFTData(chainID, *nftID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	nftDataResponse := models.MapNFTDataResponse(nftData)

	return e.JSON(http.StatusOK, nftDataResponse)
}

func (c *Controller) getNativeTokenIDRegistry(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	registries, err := c.accounts.GetNativeTokenIDRegistry(chainID)

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	nativeTokenIDRegistryResponse := &models.NativeTokenIDRegistryResponse{
		NativeTokenRegistryIDs: make([]string, len(registries)),
	}

	for k, v := range registries {
		nativeTokenIDRegistryResponse.NativeTokenRegistryIDs[k] = v.String()
	}

	return e.JSON(http.StatusOK, nativeTokenIDRegistryResponse)
}

func (c *Controller) getFoundryOutput(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	serialNumber, err := params.DecodeUInt(e, "serialNumber")
	if err != nil {
		return err
	}

	foundryOutput, err := c.accounts.GetFoundryOutput(chainID, uint32(serialNumber))

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	foundryOutputID, err := foundryOutput.ID()

	if err != nil {
		return apierrors.InvalidPropertyError("FoundryOutput.ID", err)
	}

	foundryOutputResponse := &models.FoundryOutputResponse{
		FoundryID: foundryOutputID.ToHex(),
		Token: models.AssetsResponse{
			BaseTokens: foundryOutput.Amount,
			Tokens:     models.MapNativeTokens(foundryOutput.NativeTokens),
		},
	}

	return e.JSON(http.StatusOK, foundryOutputResponse)
}
