package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"github.com/odmrs/api/api"
	"github.com/odmrs/api/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
  var params = api.CoinBalanceParams{}
  var decoder *schema.Decoder = schema.NewDecoder()
  var err error

  err = decoder.Decode(&params, r.URL.Query())

  if err != nil{
        log.Error("Erro ao decodificar parâmetros da URL: ", err)

    api.InternalErrorHandler(w)
    return
  }
   log.Info("Parâmetros decodificados: ", params)

  var database tools.DatabaseInterface
  database, err = tools.NewDatabase()

  if err != nil{
    api.InternalErrorHandler(w)
    return
  }
  
tokenDetails := database.GetUserCoins(params.Username)
  
  if tokenDetails == nil{
    log.Error("Token details not found for the specified username")
    api.InternalErrorHandler(w)
    return
  }
 log.Info("Detalhes do token encontrados: ", tokenDetails)

  var response = api.CoinBalanceResponse{
    Balance: tokenDetails.Coins,
    Code: http.StatusOK,
  }

  w.Header().Set("Content-Type", "application/json")

  err = json.NewEncoder(w).Encode(response)

  if err != nil{
    log.Error(err)
    api.InternalErrorHandler(w)
    return
  }
}