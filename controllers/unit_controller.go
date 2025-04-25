package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/models"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/services"
)

// UnitController handles unit-related HTTP requests
type UnitController struct {
	svc *services.UnitService
}

func NewUnitController(svc *services.UnitService) *UnitController {
	return &UnitController{svc: svc}
}

// Create godoc
// @Summary      Create a new organizational unit
// @Description  Adds a unit; if UnitID provided it's overridden, CreationDate auto-set
// @Tags         units
// @Accept       json
// @Produce      json
// @Param        unit  body      models.UnitBoundary  true  "Unit data"
// @Success      200   {object}  models.Unit
// @Failure      400   {object}  models.HTTPError
// @Router       /units [post]
func (c *UnitController) Create(w http.ResponseWriter, r *http.Request) {
	var b models.UnitBoundary
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	u, err := c.svc.CreateUnit(r.Context(), &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// GetOne godoc
// @Summary      Get a specific organizational unit
// @Tags         units
// @Produce      json
// @Param        unitId  path      string  true  "Unit ID"
// @Success      200     {object}  models.Unit
// @Failure      404     {object}  models.HTTPError
// @Router       /units/{unitId} [get]
func (c *UnitController) GetOne(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/units/")
	u, err := c.svc.GetUnitByID(r.Context(), id)
	if err == services.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// List godoc
// @Summary      List organizational units with pagination
// @Tags         units
// @Produce      json
// @Param        page  query     int  false  "Page number"
// @Param        size  query     int  false  "Page size"
// @Success      200   {array}   models.Unit
// @Router       /units [get]
func (c *UnitController) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.ParseInt(q.Get("page"), 10, 64)
	size, _ := strconv.ParseInt(q.Get("size"), 10, 64)
	units, _ := c.svc.ListUnits(r.Context(), page, size)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(units)
}

// Update godoc
// @Summary      Update an organizational unit
// @Tags         units
// @Accept       json
// @Param        unitId  path      string             true  "Unit ID"
// @Param        unit    body      models.UnitBoundary  true  "Updated unit data"
// @Success      204
// @Failure 400 {object} models.HTTPError
// @Failure      404   {object}  models.HTTPError
// @Router       /units/{unitId} [put]
func (c *UnitController) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/units/")
	var b models.UnitBoundary
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	err := c.svc.UpdateUnit(r.Context(), id, &b)
	if err == services.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteAll godoc
// @Summary      Delete all organizational units
// @Tags         units
// @Success      204
// @Router       /units [delete]
func (c *UnitController) DeleteAll(w http.ResponseWriter, r *http.Request) {
	c.svc.DeleteAllUnits(r.Context())
	w.WriteHeader(http.StatusNoContent)
}
