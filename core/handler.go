package core

import (
	"relayer/indexer"
)

// HandleInterchainRequest handles the interchain request
func (r *Relayer) HandleInterchainRequest(chainID string, request InterchainRequest) error {
	r.Logger.Infof("got the interchain request on %s: %+v", chainID, request)

	// TODO
	indexer.OnInterchainRequestReceived()

	callback := func(icRequestID string, response ResponseI) {
		r.Logger.Infof(
			"got the response of the interchain request on %s: %+v",
			r.HubChain.GetChainID(),
			response,
		)

		// TODO
		indexer.OnInterchainRequestHandled()

		err := r.AppChains[chainID].SendResponse(request.ID, response)
		if err != nil {
			r.Logger.Errorf(
				"failed to send the response to %s: %s",
				chainID,
				err,
			)

			return
		}

		r.Logger.Infof(
			"response sent to %s successfully",
			chainID,
		)
	}

	err := r.HubChain.SendInterchainRequest(request, callback)
	if err != nil {
		r.Logger.Errorf(
			"failed to handle the interchain request %+v on %s: %s",
			request,
			r.HubChain.GetChainID(),
			err,
		)

		return err
	}

	return nil
}
