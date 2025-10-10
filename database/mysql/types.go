package mysql

import (
	"context"
	"database/sql"
)

// CredentialRefreshFunc is a user-provided function that returns a new DSN
// CredentialRefreshFunc는 새 DSN을 반환하는 사용자 제공 함수입니다
//
// Users can implement this function to fetch credentials from:
// - HashiCorp Vault
// - AWS Secrets Manager
// - Environment variables
// - Configuration files
// - Any other credential store
//
// 사용자는 다음에서 자격 증명을 가져오기 위해 이 함수를 구현할 수 있습니다:
// - HashiCorp Vault
// - AWS Secrets Manager
// - 환경 변수
// - 설정 파일
// - 기타 자격 증명 저장소
//
// Example / 예제:
//
//	func getDSN() (string, error) {
//	    user := os.Getenv("DB_USER")
//	    pass := os.Getenv("DB_PASS")
//	    return fmt.Sprintf("%s:%s@tcp(localhost:3306)/mydb", user, pass), nil
//	}
type CredentialRefreshFunc func() (dsn string, err error)

// Tx represents a database transaction
// Tx는 데이터베이스 트랜잭션을 나타냅니다
type Tx struct {
	tx       *sql.Tx
	client   *Client
	finished bool
}

// Query executes a query within the transaction
// Query는 트랜잭션 내에서 쿼리를 실행합니다
func (t *Tx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}
	return t.tx.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that returns a single row within the transaction
// QueryRow는 트랜잭션 내에서 단일 행을 반환하는 쿼리를 실행합니다
func (t *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if t.finished {
		return nil
	}
	return t.tx.QueryRowContext(ctx, query, args...)
}

// Exec executes a query without returning rows within the transaction
// Exec는 트랜잭션 내에서 행을 반환하지 않는 쿼리를 실행합니다
func (t *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}
	return t.tx.ExecContext(ctx, query, args...)
}

// Commit commits the transaction
// Commit은 트랜잭션을 커밋합니다
func (t *Tx) Commit() error {
	if t.finished {
		return ErrTransactionFailed
	}
	t.finished = true
	return t.tx.Commit()
}

// Rollback rolls back the transaction
// Rollback은 트랜잭션을 롤백합니다
func (t *Tx) Rollback() error {
	if t.finished {
		return nil // Already finished, no error
	}
	t.finished = true
	return t.tx.Rollback()
}
