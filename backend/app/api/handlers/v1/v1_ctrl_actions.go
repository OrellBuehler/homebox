package v1

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/thechosenlan/homebox/backend/internal/core/services"
	"github.com/thechosenlan/homebox/backend/internal/sys/validate"
	"github.com/thechosenlan/homebox/backend/pkgs/server"
)

type EnsureAssetIDResult struct {
	Completed int `json:"completed"`
}

// HandleGroupInvitationsCreate godoc
// @Summary  Get the current user
// @Tags     Group
// @Produce  json
// @Success  200     {object} EnsureAssetIDResult
// @Router   /v1/actions/ensure-asset-ids [Post]
// @Security Bearer
func (ctrl *V1Controller) HandleEnsureAssetID() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		totalCompleted, err := ctrl.svc.Items.EnsureAssetID(ctx, ctx.GID)
		if err != nil {
			log.Err(err).Msg("failed to ensure asset id")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, EnsureAssetIDResult{Completed: totalCompleted})
	}
}
