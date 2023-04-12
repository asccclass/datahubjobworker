/*
   參考資料：https://github.com/contribsys/faktory/blob/main/manager/manager.go
*/
package SryManager

import(
   // "fmt"
)

type FailPayload struct {
   Jid          string   `json:"jid"`
   ErrorMessage string   `json:"message"`
   ErrorType    string   `json:"errtype"`
   Backtrace    []string `json:"backtrace"`
}
