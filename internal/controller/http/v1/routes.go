package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"tournament_scoring/internal/dto"
	"tournament_scoring/pkg/logger"
)

func (h *Handler) routes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {

		r.Post("/tournament", h.tournamentCreate)
		r.Delete("/tournament", h.tournamentDelete)
		r.Post("/division/{divisionName}/results", h.divisionResults)
		r.Post("/play-off/generate", h.playOffGenerateSchedule)
	})
}

func (h *Handler) tournamentCreate(w http.ResponseWriter, r *http.Request) {
	in := dto.CreateTournamentIn{}
	if err := in.Parse(r.Body); err != nil {
		h.l.Error("tournamentCreate: error on parsing request:", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := in.Validate(); err != nil {
		h.l.Error("tournamentCreate: error on validation request:", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	divisions, err := h.tournament.Create(r.Context(), in)
	if err != nil {
		h.l.Error("tournamentCreate: error on calling create tournament via use_case:", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, divisions, http.StatusOK)
}

func (h *Handler) tournamentDelete(w http.ResponseWriter, r *http.Request) {
	err := h.tournament.Delete(r.Context())
	if err != nil {
		h.l.Error("tournamentCreate: error on calling create tournament via use_case:", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, map[string]string{"status": "ok"}, http.StatusOK)
}

func (h *Handler) divisionResults(w http.ResponseWriter, r *http.Request) {
	in := dto.DivisionResultIn{}
	if err := in.Parse(r); err != nil {
		h.l.Error("divisionResults: error on parsing request:", logger.Err(err))
		h.Error(w, err, http.StatusBadRequest)
		return
	}

	response, err := h.division.Result(r.Context(), in)
	if err != nil {
		h.l.Error("divisionResults: error on processing division's result:", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, response, http.StatusOK)
}

func (h *Handler) playOffGenerateSchedule(w http.ResponseWriter, r *http.Request) {
	winner, err := h.playOff.GenerateSchedule(r.Context())
	if err != nil {
		h.l.Error("playOffGenerate: error on generating play-off schedule:", logger.Err(err))
		h.Error(w, err, http.StatusInternalServerError)
		return
	}

	h.Resp(w, map[string]interface{}{
		"id":     winner.ID,
		"name":   winner.Name,
		"rating": winner.Rating,
	}, http.StatusOK)
}
