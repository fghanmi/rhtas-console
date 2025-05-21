package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/securesign/rhtas-console/internal/services"
)

func RegisterRoutes(r *chi.Mux, as services.ArtifactService, rs services.RekorService, ts services.TrustService) {
	handler := NewHandler(as, rs, ts)

	r.Post("/v1/artifacts/sign", handler.PostV1ArtifactsSign)
	r.Post("/v1/artifacts/verify", handler.PostV1ArtifactsVerify)
	r.Get("/v1/artifacts/{artifact}/policies", handler.GetV1ArtifactsArtifactPolicies)
	r.Get("/v1/rekor/entries/{uuid}", handler.GetV1RekorEntriesUuid)
	r.Get("/v1/rekor/public-key", handler.GetV1RekorPublicKey)
	r.Get("/v1/trust/config", handler.GetV1TrustConfig)
}
