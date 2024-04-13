package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
  "marcos": {
    AuthToken: "123ABC",
    Username: "Marcos",
  },
  "bel": {
    AuthToken: "321BCA",
    Username: "Bel",
  },
  "maria": {
    AuthToken: "456ABC",
    Username: "Maria",
  },
}

var mockCoinDetails = map[string]CoinDetails{
  "Marcos": {
    Coins: 0,
    Username: "Marcos",
  }, 
  "Bel": {
    Coins: 4000,
    Username: "Bel",
  },
  "Maria": {
    Coins: 4219313123,
    Username: "Maria",
  },
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
  time.Sleep(time.Second * 1)

  var clientData = LoginDetails{}
  clientData, ok := mockLoginDetails[username]

  if !ok {
    return nil
  }
  return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails{
  time.Sleep(time.Second * 1)

  var clientData = CoinDetails{}

  clientData, ok := mockCoinDetails[username]
  if !ok {
    return nil
  }

  return &clientData
}

func (d *mockDB) SetupDatabase() error {
  return nil
}
