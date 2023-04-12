package SryClient

import(
   "fmt"
   "time"
   "strconv"
   cryptorand "crypto/rand"
)

type Failure struct {
   RetryCount     int      `json:"retry_count"`
   RetryRemaining int      `json:"remaining"`
   FailedAt       string   `json:"failed_at"`
   NextAt         string   `json:"next_at,omitempty"`
   ErrorMessage   string   `json:"message,omitempty"`
   ErrorType      string   `json:"errtype,omitempty"`
   Backtrace      []string `json:"backtrace,omitempty"`
}

type Job struct {
   Jid   string        `json:"jid"`
   Queue string        `json:"queue"`
   Type  string        `json:"jobtype"`
   Args  []interface{} `json:"args"`

   // optional
   CreatedAt  string                 `json:"created_at,omitempty"`
   EnqueuedAt string                 `json:"enqueued_at,omitempty"`
   At         string                 `json:"at,omitempty"`
   ReserveFor int                    `json:"reserve_for,omitempty"`
   Retry      *int                   `json:"retry"`
   Backtrace  int                    `json:"backtrace,omitempty"`
   Failure    *Failure               `json:"failure,omitempty"`
   Custom     map[string]interface{} `json:"custom,omitempty"`
}

func RandomJid() string {
   bytes := make([]byte, 12)
   _, err := cryptorand.Read(bytes)
   if err != nil {
      mathrand.Read(bytes)
   }
   return base64.RawURLEncoding.EncodeToString(bytes)
}

func NewJob(jobtype string, args ...interface{})(*Job, error) {
   RetryPolicyDefault := os.Getenv("RetryPolicy")
   if RetryPolicyDefault == "" {
      return nil, fmt.Errorf("Variable RetryPolicy not set")
   }
   retryPolicy, err := strconv.Atoi(RetryPolicyDefault)
   if err != nil {
      return nil, err
   }
   return &Job{
      Type:      jobtype,
      Queue:     "default",
      Args:      args,
      Jid:       RandomJid(),
      CreatedAt: time.Now().UTC().Format(time.RFC3339Nano),
      Retry:     *retryPolicy,
   }, nil
}
