/*
   參考資料：https://github.com/contribsys/faktory/blob/main/manager/manager.go
*/
package SryManager

import(
   "fmt"
   "sync"
   "time"
   "context"
   "encoding/json"

   "github.com/asccclass/datahubjobworker/client"
   "github.com/asccclass/datahubjobworker/storage"
   // "github.com/contribsys/faktory/util"
   // "github.com/redis/go-redis/v9"
)

type Manager interface {
   Push(ctx context.Context, job *client.Job) error
   PauseQueue(ctx context.Context, qName string) error
   ResumeQueue(ctx context.Context, qName string) error
   RemoveQueue(ctx context.Context, qName string) error
   Fetch(ctx context.Context, wid string, queues ...string) (*client.Job, error)
   Acknowledge(ctx context.Context, jid string) (*client.Job, error)
   Fail(ctx context.Context, fail *FailPayload) error
   ExtendReservation(ctx context.Context, jid string, until time.Time) error
   WorkingCount() int
   ReapExpiredJobs(ctx context.Context, when time.Time) (int64, error)
   Purge(ctx context.Context, when time.Time) (int64, error)
   EnqueueScheduledJobs(ctx context.Context, when time.Time) (int64, error)
   RetryJobs(ctx context.Context, when time.Time) (int64, error)
   BusyCount(wid string) int
   AddMiddleware(fntype string, fn MiddlewareFunc)
   KV() storage.KV
   Redis() *redis.Client
   SetFetcher(f Fetcher)
}
