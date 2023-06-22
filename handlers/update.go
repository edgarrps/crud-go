package handlers

func Update(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(chi.URLParam(r, "id"))
  if err != nil {
  log.Printf("Erro ao fazer parse do id: %v", err)
    http.Error(w, http.StatusInternalServerError), StatusInternalServerError
    return
  }
}
