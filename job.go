package faktory

type Failure struct {
	RetryCount   int      `json:"retry_count"`
	FailedAt     string   `json:"failed_at"`
	NextAt       string   `json:"next_at",omitempty`
	ErrorMessage string   `json:"message",omitempty`
	ErrorType    string   `json:"errtype",omitempty`
	Backtrace    []string `json:"backtrace",omitempty`
}

type Job struct {
	// required
	Jid   string        `json:"jid"`
	Queue string        `json:"queue"`
	Type  string        `json:"jobtype"`
	Args  []interface{} `json:"args"`

	// optional
	Gid        string                 `json:"gid,omitempty"`
	CreatedAt  string                 `json:"created_at,omitempty"`
	EnqueuedAt string                 `json:"enqueued_at,omitempty"`
	RunAt      string                 `json:"run_at,omitempty"`
	ReserveFor int                    `json:"reserve_for,omitempty"`
	Retry      int                    `json:"retry,omitempty"`
	Backtrace  int                    `json:"backtrace,omitempty"`
	Failure    *Failure               `json:"failure,omitempty"`
	Custom     map[string]interface{} `json:"custom,omitempty"`
}
