/*
   參考資料：https://github.com/contribsys/faktory/blob/main/manager/manager.go
*/
package SryManager

import(
)

type MiddlewareFunc func(ctx context.Context, next func() error) error
