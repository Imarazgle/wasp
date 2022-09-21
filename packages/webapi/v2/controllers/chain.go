package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/webapi/v1/httperrors"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/models"
	"github.com/iotaledger/wasp/packages/webapi/v2/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/v2/routes"
)

type ChainController struct {
	log *loggerpkg.Logger

	chainService    interfaces.Chain
	nodeService     interfaces.Node
	registryService interfaces.Registry
}

func NewChainController(log *loggerpkg.Logger, chainService interfaces.Chain, nodeService interfaces.Node, registryService interfaces.Registry) interfaces.APIController {
	return &ChainController{
		log:             log,
		chainService:    chainService,
		nodeService:     nodeService,
		registryService: registryService,
	}
}

func (c *ChainController) Name() string {
	return "chains"
}

func (c *ChainController) activateChain(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	if err := c.chainService.ActivateChain(chainID); err != nil {
		return err
	}

	return e.NoContent(http.StatusOK)
}

func (c *ChainController) deactivateChain(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	if err := c.chainService.DeactivateChain(chainID); err != nil {
		return err
	}

	return e.NoContent(http.StatusOK)
}

func (c *ChainController) getCommitteeInfo(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	chainRecord, err := c.registryService.GetChainRecordByChainID(chainID)
	if err != nil {
		return err
	}

	if chainRecord == nil {
		return httperrors.NotFound("")
	}

	chain := c.chainService.GetChainByID(chainID)

	if chain == nil {
		return httperrors.NotFound("")
	}

	committeeInfo := chain.GetCommitteeInfo()
	chainNodeInfo, err := c.nodeService.GetNodeInfo(chain)
	if err != nil {
		return err
	}

	chainInfo := models.CommitteeInfoResponse{
		ChainID:        chainID.String(),
		Active:         chainRecord.Active,
		StateAddress:   committeeInfo.Address.String(),
		CommitteeNodes: chainNodeInfo.CommitteeNodes,
		AccessNodes:    chainNodeInfo.AccessNodes,
		CandidateNodes: chainNodeInfo.CandidateNodes,
	}

	return e.JSON(http.StatusOK, chainInfo)
}

func (c *ChainController) getContracts(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	contracts, err := c.chainService.GetContracts(chainID)
	if err != nil {
		return err
	}

	contractList := make([]models.ContractInfoResponse, 0, len(contracts))

	for hName, contract := range contracts {
		contractInfo := models.ContractInfoResponse{
			Description: contract.Description,
			HName:       hName,
			Name:        contract.Name,
			ProgramHash: contract.ProgramHash,
		}

		contractList = append(contractList, contractInfo)
	}

	return e.JSON(http.StatusOK, contractList)
}

func (c *ChainController) getChainInfo(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	chainInfo, err := c.chainService.GetChainInfoByChainID(chainID)
	if err != nil {
		return err
	}

	chainInfoResponse := models.MapChainInfoResponse(chainInfo)

	return e.JSON(http.StatusOK, chainInfoResponse)
}

func (c *ChainController) getChainList(e echo.Context) error {
	chainIDs, err := c.chainService.GetAllChainIDs()
	if err != nil {
		return err
	}

	chainList := models.ChainListResponse{}

	for _, chainID := range chainIDs {
		chainInfo, err := c.chainService.GetChainInfoByChainID(chainID)
		if err != nil {
			return err
		}

		chainInfoResponse := models.MapChainInfoResponse(chainInfo)

		chainList = append(chainList, chainInfoResponse)
	}

	return e.JSON(http.StatusOK, chainList)
}

func (c *ChainController) RegisterExampleData(mock interfaces.Mocker) {
	mock.AddModel(&models.ChainInfoResponse{})
	mock.AddModel(&models.ContractListResponse{})
	mock.AddModel(&models.CommitteeInfoResponse{})
	mock.AddModel(&models.ChainListResponse{})
	mock.AddModel(&models.ContractInfoResponse{})
}

func (c *ChainController) RegisterPublic(publicAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	publicAPI.POST(routes.ActivateChain(":chainID")+"/test", c.activateChain).
		AddParamPath("", "chainID", "ChainID (string)").
		SetOperationId("activateChainTest").
		SetSummary("Activate a chain")
}

func (c *ChainController) RegisterAdmin(adminAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	adminAPI.POST(routes.ActivateChain(":chainID"), c.activateChain).
		AddParamPath("", "chainID", "ChainID (string)").
		AddResponse(http.StatusOK, "Chain was successfully activated", nil, nil).
		AddResponse(http.StatusNotModified, "Chain was not activated", nil, nil).
		SetOperationId("activateChain").
		SetSummary("Activate a chain")

	adminAPI.POST(routes.DeactivateChain(":chainID"), c.deactivateChain).
		AddParamPath("", "chainID", "ChainID (string)").
		AddResponse(http.StatusOK, "Chain was successfully deactivated", nil, nil).
		AddResponse(http.StatusNotModified, "Chain was not deactivated", nil, nil).
		SetOperationId("deactivateChain").
		SetSummary("Deactivate a chain")

	adminAPI.GET(routes.GetChainCommitteeInfo(":chainID"), c.getCommitteeInfo).
		AddParamPath("", "chainID", "ChainID (string)").
		AddResponse(http.StatusOK, "A list of all nodes tied to the chain.", mocker.GetMockedStruct(models.CommitteeInfoResponse{}), nil).
		SetOperationId("getChainCommitteeInfo").
		SetSummary("Get basic chain info.")

	adminAPI.GET(routes.GetChainContracts(":chainID"), c.getContracts).
		AddParamPath("", "chainID", "ChainID (string)").
		AddResponse(http.StatusOK, "A list of all available contracts.", mocker.GetMockedStruct(models.ContractListResponse{}), nil).
		SetOperationId("getChainContracts").
		SetSummary("Get all available chain contracts.")

	adminAPI.GET(routes.GetChainList(), c.getChainList).
		AddResponse(http.StatusOK, "A list of all available chains.", mocker.GetMockedStruct(models.ChainListResponse{}), nil).
		SetOperationId("getChainList").
		SetSummary("Get a list of all chains.")

	adminAPI.GET(routes.GetChainInfo(":chainID"), c.getChainInfo).
		AddParamPath("", "chainID", "ChainID (string)").
		AddResponse(http.StatusOK, "Information about a specific chain.", mocker.GetMockedStruct(models.ChainInfoResponse{}), nil).
		SetOperationId("getChainInfo").
		SetSummary("Get information about a specific chain.")
}
