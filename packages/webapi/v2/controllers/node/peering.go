package node

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/webapi/v2/apierrors"
	"github.com/iotaledger/wasp/packages/webapi/v2/models"
)

func (c *Controller) getRegisteredPeers(e echo.Context) error {
	peers := c.peeringService.GetRegisteredPeers()
	peerModels := make([]models.PeeringNodeStatusResponse, len(peers))

	for k, v := range peers {
		peerModels[k] = models.PeeringNodeStatusResponse{
			IsAlive:   v.IsAlive,
			NetID:     v.NetID,
			NumUsers:  v.NumUsers,
			PublicKey: v.PublicKey.String(),
			IsTrusted: v.IsTrusted,
		}
	}

	return e.JSON(http.StatusOK, peerModels)
}

func (c *Controller) getTrustedPeers(e echo.Context) error {
	peers, err := c.peeringService.GetTrustedPeers()
	if err != nil {
		return apierrors.InternalServerError(err)
	}

	peerModels := make([]models.PeeringNodeIdentityResponse, len(peers))
	for k, v := range peers {
		peerModels[k] = models.PeeringNodeIdentityResponse{
			NetID:     v.NetID,
			PublicKey: v.PublicKey.String(),
			IsTrusted: v.IsTrusted,
		}
	}

	return e.JSON(http.StatusOK, peerModels)
}

func (c *Controller) getIdentity(e echo.Context) error {
	self := c.peeringService.GetIdentity()

	peerModel := models.PeeringNodeIdentityResponse{
		NetID:     self.NetID,
		PublicKey: self.PublicKey.String(),
		IsTrusted: self.IsTrusted,
	}

	return e.JSON(http.StatusOK, peerModel)
}

func (c *Controller) trustPeer(e echo.Context) error {
	var trustedPeer models.PeeringTrustRequest

	if err := e.Bind(&trustedPeer); err != nil {
		return apierrors.InvalidPropertyError("body", err)
	}

	publicKey, err := cryptolib.NewPublicKeyFromString(trustedPeer.PublicKey)
	if err != nil {
		return apierrors.InvalidPropertyError("publicKey", err)
	}

	_, err = c.peeringService.TrustPeer(publicKey, trustedPeer.NetID)
	if err != nil {
		return apierrors.InternalServerError(err)
	}

	return e.NoContent(http.StatusOK)
}

func (c *Controller) distrustPeer(e echo.Context) error {
	var trustedPeer models.PeeringTrustRequest

	if err := e.Bind(&trustedPeer); err != nil {
		return apierrors.InvalidPropertyError("body", err)
	}

	publicKey, err := cryptolib.NewPublicKeyFromString(trustedPeer.PublicKey)
	if err != nil {
		return apierrors.InvalidPropertyError("publicKey", err)
	}

	_, err = c.peeringService.DistrustPeer(publicKey)
	if err != nil {
		return apierrors.InternalServerError(err)
	}

	return e.NoContent(http.StatusOK)
}
