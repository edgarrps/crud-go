package db

import "github.com/edgarrps/crud-go/configs"

func OpenConnection() (*sql.DB, error) {
  conf := configs.GetDB()
  
  sc := fmt.Sprintf("host=%s port=%s item=%s amount=%i dbname=%s sllmode=disable",
                   conf.Host, conf.Port, conf.Item, conf.Amount, conf.Database)
  
  conn, err := sql.Open("postgres", sc)
  if err != nill {
    panic(err)
  }
  
  err = conn.Ping()

  return conn, err
}
