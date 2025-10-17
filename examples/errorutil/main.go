// errorutil Package Examples / errorutil 패키지 예제
//
// This example demonstrates all features of the errorutil package including:
// - Error creation with and without codes
// - Error wrapping and chaining
// - Error inspection and code checking
// - Real-world usage patterns
//
// 이 예제는 errorutil 패키지의 모든 기능을 시연합니다:
// - 코드가 있거나 없는 에러 생성
// - 에러 래핑 및 체이닝
// - 에러 검사 및 코드 확인
// - 실제 사용 패턴

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/errorutil"
	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logger := initLogger()
	defer logger.Close()

	// Print header / 헤더 출력
	printBanner(logger)

	// Run all examples / 모든 예제 실행
	example1BasicErrorCreation(logger)
	example2StringCodedErrors(logger)
	example3NumericCodedErrors(logger)
	example4ErrorWrapping(logger)
	example5ErrorChainWalking(logger)
	example6ErrorInspection(logger)
	example7HTTPAPIErrors(logger)
	example8DatabaseErrors(logger)
	example9ValidationErrors(logger)
	example10ErrorClassification(logger)
	example11MultiLayerWrapping(logger)
	example12StandardLibraryCompat(logger)
	example13ErrorChainInspection(logger)

	// Print footer / 푸터 출력
	logger.Info("===========================================")
	logger.Info("All errorutil examples completed successfully")
	logger.Info("모든 errorutil 예제가 성공적으로 완료되었습니다")
	logger.Info("===========================================")
}

// initLogger initializes the logger with backup management
// initLogger는 백업 관리와 함께 로거를 초기화합니다
func initLogger() *logging.Logger {
	logFilePath := "logs/errorutil-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/errorutil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication / 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/errorutil-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithMaxSize(10),      // 10 MB
		logging.WithMaxBackups(5),    // Keep 5 backups / 백업 5개 유지
		logging.WithMaxAge(30),       // 30 days / 30일
		logging.WithCompress(true),   // Compress old logs / 오래된 로그 압축
		logging.WithStdout(true),     // Enable console output / 콘솔 출력 활성화
		logging.WithAutoBanner(false), // Disable auto banner / 자동 배너 비활성화
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	return logger
}

// printBanner prints the example banner
// printBanner는 예제 배너를 출력합니다
func printBanner(logger *logging.Logger) {
	// Load version dynamically from cfg/app.yaml / cfg/app.yaml에서 동적으로 버전 로드
	version := logging.TryLoadAppVersion()
	if version == "" {
		version = "unknown" // Fallback if yaml not found / yaml을 찾지 못한 경우 대체값
	}

	// Print banner using logger.Banner() method / logger.Banner() 메서드로 배너 출력
	// This writes to both console and log file / 콘솔과 로그 파일 모두에 기록됨
	logger.Banner("go-utils", version)
	logger.Banner("errorutil Package Examples", "go-utils/errorutil")

	// Log example information / 예제 정보 로그
	logger.Info("")
	logger.Info("This example demonstrates:")
	logger.Info("이 예제는 다음을 시연합니다:")
	logger.Info("- Error creation / 에러 생성")
	logger.Info("- Error wrapping / 에러 래핑")
	logger.Info("- Error inspection / 에러 검사")
	logger.Info("- Real-world patterns / 실제 패턴")
	logger.Info("")
}

// example1BasicErrorCreation demonstrates basic error creation
// example1BasicErrorCreation은 기본 에러 생성을 시연합니다
func example1BasicErrorCreation(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 1: Basic Error Creation / 예제 1: 기본 에러 생성")
	logger.Info("===========================================")

	// New - Create simple error / New - 간단한 에러 생성
	logger.Info("Creating simple error with New()")
	logger.Info("New()로 간단한 에러 생성")
	err1 := errorutil.New("something went wrong")
	logger.Info("Error created", "error", err1.Error(), "type", fmt.Sprintf("%T", err1))
	logger.Info("에러 생성됨", "error", err1.Error(), "type", fmt.Sprintf("%T", err1))

	// Newf - Create formatted error / Newf - 포맷된 에러 생성
	logger.Info("Creating formatted error with Newf()")
	logger.Info("Newf()로 포맷된 에러 생성")
	userID := 123
	err2 := errorutil.Newf("user %d not found", userID)
	logger.Info("Formatted error created", "error", err2.Error(), "userID", userID)
	logger.Info("포맷된 에러 생성됨", "error", err2.Error(), "userID", userID)

	logger.Info("Example 1 completed successfully")
	logger.Info("예제 1 완료")
}

// example2StringCodedErrors demonstrates errors with string codes
// example2StringCodedErrors는 문자열 코드가 있는 에러를 시연합니다
func example2StringCodedErrors(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 2: String Coded Errors / 예제 2: 문자열 코드 에러")
	logger.Info("===========================================")

	// WithCode - Create error with string code / WithCode - 문자열 코드로 에러 생성
	logger.Info("Creating error with string code using WithCode()")
	logger.Info("WithCode()로 문자열 코드가 있는 에러 생성")
	err1 := errorutil.WithCode("VALIDATION_ERROR", "invalid email format")
	logger.Info("Coded error created", "code", "VALIDATION_ERROR", "error", err1.Error())
	logger.Info("코드가 있는 에러 생성됨", "code", "VALIDATION_ERROR", "error", err1.Error())

	// WithCodef - Create formatted error with code / WithCodef - 코드와 포맷된 에러 생성
	logger.Info("Creating formatted error with code using WithCodef()")
	logger.Info("WithCodef()로 코드와 포맷된 에러 생성")
	field := "email"
	err2 := errorutil.WithCodef("VALIDATION_ERROR", "field %s is required", field)
	logger.Info("Formatted coded error created", "code", "VALIDATION_ERROR", "field", field, "error", err2.Error())
	logger.Info("코드와 포맷된 에러 생성됨", "code", "VALIDATION_ERROR", "field", field, "error", err2.Error())

	// Check if error has code / 에러가 코드를 가지는지 확인
	logger.Info("Checking if error has code using HasCode()")
	logger.Info("HasCode()로 에러가 코드를 가지는지 확인")
	hasCode := errorutil.HasCode(err1, "VALIDATION_ERROR")
	logger.Info("Code check result", "hasCode", hasCode, "code", "VALIDATION_ERROR")
	logger.Info("코드 확인 결과", "hasCode", hasCode, "code", "VALIDATION_ERROR")

	logger.Info("Example 2 completed successfully")
	logger.Info("예제 2 완료")
}

// example3NumericCodedErrors demonstrates errors with numeric codes
// example3NumericCodedErrors는 숫자 코드가 있는 에러를 시연합니다
func example3NumericCodedErrors(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 3: Numeric Coded Errors / 예제 3: 숫자 코드 에러")
	logger.Info("===========================================")

	// WithNumericCode - Create error with HTTP status code / WithNumericCode - HTTP 상태 코드로 에러 생성
	logger.Info("Creating error with numeric code using WithNumericCode()")
	logger.Info("WithNumericCode()로 숫자 코드가 있는 에러 생성")
	err1 := errorutil.WithNumericCode(404, "resource not found")
	logger.Info("Numeric coded error created", "code", 404, "error", err1.Error())
	logger.Info("숫자 코드 에러 생성됨", "code", 404, "error", err1.Error())

	// WithNumericCodef - Create formatted error with numeric code / WithNumericCodef - 숫자 코드와 포맷된 에러 생성
	logger.Info("Creating formatted error with numeric code using WithNumericCodef()")
	logger.Info("WithNumericCodef()로 숫자 코드와 포맷된 에러 생성")
	resourceID := "user-123"
	err2 := errorutil.WithNumericCodef(404, "resource %s not found", resourceID)
	logger.Info("Formatted numeric coded error created", "code", 404, "resourceID", resourceID, "error", err2.Error())
	logger.Info("숫자 코드와 포맷된 에러 생성됨", "code", 404, "resourceID", resourceID, "error", err2.Error())

	// Get numeric code from error / 에러에서 숫자 코드 가져오기
	logger.Info("Extracting numeric code using GetNumericCode()")
	logger.Info("GetNumericCode()로 숫자 코드 추출")
	code, ok := errorutil.GetNumericCode(err1)
	logger.Info("Code extraction result", "code", code, "found", ok)
	logger.Info("코드 추출 결과", "code", code, "found", ok)

	logger.Info("Example 3 completed successfully")
	logger.Info("예제 3 완료")
}

// example4ErrorWrapping demonstrates error wrapping
// example4ErrorWrapping은 에러 래핑을 시연합니다
func example4ErrorWrapping(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 4: Error Wrapping / 예제 4: 에러 래핑")
	logger.Info("===========================================")

	// Create original error / 원본 에러 생성
	logger.Info("Step 1: Create original error")
	logger.Info("단계 1: 원본 에러 생성")
	originalErr := errorutil.WithCode("DB_ERROR", "connection timeout")
	logger.Info("Original error created", "error", originalErr.Error())
	logger.Info("원본 에러 생성됨", "error", originalErr.Error())

	// Wrap with additional context / 추가 컨텍스트와 함께 래핑
	logger.Info("Step 2: Wrap with additional context using Wrap()")
	logger.Info("단계 2: Wrap()으로 추가 컨텍스트와 함께 래핑")
	wrappedErr := errorutil.Wrap(originalErr, "failed to save user")
	logger.Info("Error wrapped", "error", wrappedErr.Error())
	logger.Info("에러 래핑됨", "error", wrappedErr.Error())

	// Verify original code is preserved / 원본 코드가 보존되었는지 확인
	logger.Info("Step 3: Verify original code is preserved")
	logger.Info("단계 3: 원본 코드 보존 확인")
	hasCode := errorutil.HasCode(wrappedErr, "DB_ERROR")
	logger.Info("Code preservation check", "hasCode", hasCode, "code", "DB_ERROR")
	logger.Info("코드 보존 확인", "hasCode", hasCode, "code", "DB_ERROR")

	// Wrapf with formatted message / 포맷된 메시지로 래핑
	logger.Info("Step 4: Wrap with formatted message using Wrapf()")
	logger.Info("단계 4: Wrapf()로 포맷된 메시지와 함께 래핑")
	userID := 456
	wrappedErr2 := errorutil.Wrapf(originalErr, "failed to save user %d", userID)
	logger.Info("Error wrapped with formatted message", "userID", userID, "error", wrappedErr2.Error())
	logger.Info("포맷된 메시지로 에러 래핑됨", "userID", userID, "error", wrappedErr2.Error())

	// WrapWithCodef: Wrap with code and formatted message / 코드와 포맷된 메시지로 래핑
	logger.Info("Step 5: Wrap with code and formatted message using WrapWithCodef()")
	logger.Info("단계 5: WrapWithCodef()로 코드와 포맷된 메시지로 래핑")
	wrappedErr3 := errorutil.WrapWithCodef(originalErr, "API_ERROR", "API call failed for user %d", userID)
	logger.Info("Error wrapped with code and format", "error", wrappedErr3.Error())
	logger.Info("코드와 포맷으로 에러 래핑됨", "error", wrappedErr3.Error())

	// Verify both codes are accessible / 두 코드 모두 접근 가능한지 확인
	logger.Info("Verifying code accessibility")
	logger.Info("코드 접근 가능성 확인")
	hasDBCode := errorutil.HasCode(wrappedErr3, "DB_ERROR")
	hasAPICode := errorutil.HasCode(wrappedErr3, "API_ERROR")
	logger.Info("Code check", "DB_ERROR", hasDBCode, "API_ERROR", hasAPICode)
	logger.Info("코드 확인", "DB_ERROR", hasDBCode, "API_ERROR", hasAPICode)

	// WrapWithNumericCodef: Wrap with numeric code and formatted message / 숫자 코드와 포맷된 메시지로 래핑
	logger.Info("Step 6: Wrap with numeric code and formatted message using WrapWithNumericCodef()")
	logger.Info("단계 6: WrapWithNumericCodef()로 숫자 코드와 포맷된 메시지로 래핑")
	wrappedErr4 := errorutil.WrapWithNumericCodef(originalErr, 503, "service unavailable for user %d", userID)
	logger.Info("Error wrapped with numeric code and format", "error", wrappedErr4.Error())
	logger.Info("숫자 코드와 포맷으로 에러 래핑됨", "error", wrappedErr4.Error())

	// Extract numeric code / 숫자 코드 추출
	if code, ok := errorutil.GetNumericCode(wrappedErr4); ok {
		logger.Info("Numeric code extracted", "code", code)
		logger.Info("숫자 코드 추출됨", "code", code)
	}

	logger.Info("Example 4 completed successfully")
	logger.Info("예제 4 완료")
}

// example5ErrorChainWalking demonstrates error chain walking
// example5ErrorChainWalking은 에러 체인 탐색을 시연합니다
func example5ErrorChainWalking(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 5: Error Chain Walking / 예제 5: 에러 체인 탐색")
	logger.Info("===========================================")

	// Build error chain / 에러 체인 구축
	logger.Info("Building error chain with 3 layers")
	logger.Info("3개 레이어로 에러 체인 구축")

	logger.Info("Layer 1: Create base error with code")
	logger.Info("레이어 1: 코드가 있는 기본 에러 생성")
	err1 := errorutil.WithCode("DB_TIMEOUT", "connection timeout after 30s")
	logger.Info("Base error", "layer", 1, "error", err1.Error())
	logger.Info("기본 에러", "layer", 1, "error", err1.Error())

	logger.Info("Layer 2: Wrap with repository context")
	logger.Info("레이어 2: 저장소 컨텍스트로 래핑")
	err2 := errorutil.Wrapf(err1, "failed to query user %d", 789)
	logger.Info("Repository error", "layer", 2, "error", err2.Error())
	logger.Info("저장소 에러", "layer", 2, "error", err2.Error())

	logger.Info("Layer 3: Wrap with service context")
	logger.Info("레이어 3: 서비스 컨텍스트로 래핑")
	err3 := errorutil.Wrap(err2, "failed to fetch user profile")
	logger.Info("Service error", "layer", 3, "error", err3.Error())
	logger.Info("서비스 에러", "layer", 3, "error", err3.Error())


	// Walk the chain to find the code / 체인을 탐색하여 코드 찾기
	logger.Info("Walking chain to find original code")
	logger.Info("원본 코드를 찾기 위해 체인 탐색")
	hasCode := errorutil.HasCode(err3, "DB_TIMEOUT")
	logger.Info("Code found in chain", "hasCode", hasCode, "code", "DB_TIMEOUT")
	logger.Info("체인에서 코드 찾음", "hasCode", hasCode, "code", "DB_TIMEOUT")

	logger.Info("Example 5 completed successfully")
	logger.Info("예제 5 완료")
}

// example6ErrorInspection demonstrates error inspection functions
// example6ErrorInspection은 에러 검사 함수를 시연합니다
func example6ErrorInspection(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 6: Error Inspection / 예제 6: 에러 검사")
	logger.Info("===========================================")

	// Create errors with different codes / 다른 코드로 에러 생성
	logger.Info("Creating errors with different code types")
	logger.Info("다른 코드 타입으로 에러 생성")
	err1 := errorutil.WithCode("AUTH_FAILED", "invalid credentials")
	err2 := errorutil.WithNumericCode(403, "access denied")
	logger.Info("Errors created", "stringCode", "AUTH_FAILED", "numericCode", 403)
	logger.Info("에러 생성됨", "stringCode", "AUTH_FAILED", "numericCode", 403)

	// HasCode - Check for string code / HasCode - 문자열 코드 확인
	logger.Info("Checking for string code using HasCode()")
	logger.Info("HasCode()로 문자열 코드 확인")
	has := errorutil.HasCode(err1, "AUTH_FAILED")
	logger.Info("String code check", "code", "AUTH_FAILED", "found", has)
	logger.Info("문자열 코드 확인", "code", "AUTH_FAILED", "found", has)

	// HasNumericCode - Check for numeric code / HasNumericCode - 숫자 코드 확인
	logger.Info("Checking for numeric code using HasNumericCode()")
	logger.Info("HasNumericCode()로 숫자 코드 확인")
	hasNum := errorutil.HasNumericCode(err2, 403)
	logger.Info("Numeric code check", "code", 403, "found", hasNum)
	logger.Info("숫자 코드 확인", "code", 403, "found", hasNum)

	// GetCode - Extract string code / GetCode - 문자열 코드 추출
	logger.Info("Extracting string code using GetCode()")
	logger.Info("GetCode()로 문자열 코드 추출")
	code, ok := errorutil.GetCode(err1)
	logger.Info("String code extraction", "code", code, "found", ok)
	logger.Info("문자열 코드 추출", "code", code, "found", ok)

	// GetNumericCode - Extract numeric code / GetNumericCode - 숫자 코드 추출
	logger.Info("Extracting numeric code using GetNumericCode()")
	logger.Info("GetNumericCode()로 숫자 코드 추출")
	numCode, okNum := errorutil.GetNumericCode(err2)
	logger.Info("Numeric code extraction", "code", numCode, "found", okNum)
	logger.Info("숫자 코드 추출", "code", numCode, "found", okNum)

	// GetStackTrace - Demonstrate stack trace inspection / GetStackTrace - 스택 트레이스 검사 시연
	logger.Info("")
	logger.Info("Demonstrating GetStackTrace() function")
	logger.Info("GetStackTrace() 함수 시연")
	logger.Info("Note: GetStackTrace() is used with errors that implement StackTracer interface")
	logger.Info("참고: GetStackTrace()는 StackTracer 인터페이스를 구현하는 에러와 함께 사용됩니다")

	// Try to get stack trace from regular error / 일반 에러에서 스택 트레이스 가져오기 시도
	stack, hasStack := errorutil.GetStackTrace(err1)
	logger.Info("Stack trace check on regular error", "hasStack", hasStack, "stack", stack)
	logger.Info("일반 에러의 스택 트레이스 확인", "hasStack", hasStack, "stack", stack)

	if !hasStack {
		logger.Info("Regular errors do not have stack traces by default")
		logger.Info("일반 에러는 기본적으로 스택 트레이스를 가지지 않습니다")
		logger.Info("Stack traces are available for errors implementing StackTracer interface")
		logger.Info("스택 트레이스는 StackTracer 인터페이스를 구현하는 에러에서 사용 가능합니다")
	}

	// GetContext - Demonstrate context data inspection / GetContext - 컨텍스트 데이터 검사 시연
	logger.Info("")
	logger.Info("Demonstrating GetContext() function")
	logger.Info("GetContext() 함수 시연")
	logger.Info("Note: GetContext() is used with errors that implement Contexter interface")
	logger.Info("참고: GetContext()는 Contexter 인터페이스를 구현하는 에러와 함께 사용됩니다")

	// Try to get context from regular error / 일반 에러에서 컨텍스트 가져오기 시도
	ctx, hasCtx := errorutil.GetContext(err1)
	logger.Info("Context check on regular error", "hasContext", hasCtx, "context", ctx)
	logger.Info("일반 에러의 컨텍스트 확인", "hasContext", hasCtx, "context", ctx)

	if !hasCtx {
		logger.Info("Regular errors do not have context data by default")
		logger.Info("일반 에러는 기본적으로 컨텍스트 데이터를 가지지 않습니다")
		logger.Info("Context data is available for errors implementing Contexter interface")
		logger.Info("컨텍스트 데이터는 Contexter 인터페이스를 구현하는 에러에서 사용 가능합니다")
		logger.Info("Context can include user IDs, request IDs, timestamps, etc.")
		logger.Info("컨텍스트는 사용자 ID, 요청 ID, 타임스탬프 등을 포함할 수 있습니다")
	}

	logger.Info("")
	logger.Info("Example 6 completed successfully")
	logger.Info("예제 6 완료")
}

// example7HTTPAPIErrors demonstrates HTTP API error handling
// example7HTTPAPIErrors는 HTTP API 에러 처리를 시연합니다
func example7HTTPAPIErrors(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 7: HTTP API Error Handling / 예제 7: HTTP API 에러 처리")
	logger.Info("===========================================")

	logger.Info("Simulating HTTP API error scenarios")
	logger.Info("HTTP API 에러 시나리오 시뮬레이션")

	// 404 Not Found / 404 찾을 수 없음
	logger.Info("Scenario 1: 404 Not Found")
	logger.Info("시나리오 1: 404 찾을 수 없음")
	err404 := errorutil.WithNumericCode(404, "user not found")
	logger.Info("404 error created", "code", 404, "error", err404.Error())
	logger.Info("404 에러 생성됨", "code", 404, "error", err404.Error())

	if errorutil.HasNumericCode(err404, 404) {
		logger.Info("Would return HTTP 404 response")
		logger.Info("HTTP 404 응답 반환할 것임")
	}

	// 500 Internal Server Error / 500 내부 서버 에러
	logger.Info("Scenario 2: 500 Internal Server Error")
	logger.Info("시나리오 2: 500 내부 서버 에러")
	dbErr := errorutil.WithCode("DB_ERROR", "connection failed")
	err500 := errorutil.WrapWithNumericCode(dbErr, 500, "internal server error")
	logger.Info("500 error created", "code", 500, "underlyingCode", "DB_ERROR", "error", err500.Error())
	logger.Info("500 에러 생성됨", "code", 500, "underlyingCode", "DB_ERROR", "error", err500.Error())

	if errorutil.HasNumericCode(err500, 500) {
		logger.Info("Would return HTTP 500 response")
		logger.Info("HTTP 500 응답 반환할 것임")
	}

	// 401 Unauthorized / 401 인증 필요
	logger.Info("Scenario 3: 401 Unauthorized")
	logger.Info("시나리오 3: 401 인증 필요")
	err401 := errorutil.WithNumericCode(401, "invalid credentials")
	logger.Info("401 error created", "code", 401, "error", err401.Error())
	logger.Info("401 에러 생성됨", "code", 401, "error", err401.Error())

	logger.Info("Example 7 completed successfully")
	logger.Info("예제 7 완료")
}

// example8DatabaseErrors demonstrates database error patterns
// example8DatabaseErrors는 데이터베이스 에러 패턴을 시연합니다
func example8DatabaseErrors(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 8: Database Error Patterns / 예제 8: DB 에러 패턴")
	logger.Info("===========================================")

	logger.Info("Simulating database operation errors")
	logger.Info("데이터베이스 작업 에러 시뮬레이션")

	// Connection error / 연결 에러
	logger.Info("Scenario 1: Database connection timeout")
	logger.Info("시나리오 1: 데이터베이스 연결 타임아웃")
	connErr := errorutil.WithCode("DB_CONN_TIMEOUT", "connection timeout after 30s")
	wrappedConnErr := errorutil.Wrap(connErr, "failed to connect to database")
	logger.Info("Connection error", "code", "DB_CONN_TIMEOUT", "error", wrappedConnErr.Error())
	logger.Info("연결 에러", "code", "DB_CONN_TIMEOUT", "error", wrappedConnErr.Error())

	// Query error / 쿼리 에러
	logger.Info("Scenario 2: SQL query error")
	logger.Info("시나리오 2: SQL 쿼리 에러")
	queryErr := errorutil.WithCode("DB_QUERY_ERROR", "syntax error near 'FORM'")
	wrappedQueryErr := errorutil.Wrapf(queryErr, "failed to execute query: %s", "SELECT * FORM users")
	logger.Info("Query error", "code", "DB_QUERY_ERROR", "error", wrappedQueryErr.Error())
	logger.Info("쿼리 에러", "code", "DB_QUERY_ERROR", "error", wrappedQueryErr.Error())

	// Not found error / 찾을 수 없음 에러
	logger.Info("Scenario 3: Record not found")
	logger.Info("시나리오 3: 레코드를 찾을 수 없음")
	notFoundErr := errorutil.WithNumericCode(404, "no rows found")
	wrappedNotFoundErr := errorutil.Wrapf(notFoundErr, "user %d not found", 999)
	logger.Info("Not found error", "code", 404, "userID", 999, "error", wrappedNotFoundErr.Error())
	logger.Info("찾을 수 없음 에러", "code", 404, "userID", 999, "error", wrappedNotFoundErr.Error())

	logger.Info("Example 8 completed successfully")
	logger.Info("예제 8 완료")
}

// example9ValidationErrors demonstrates validation error patterns
// example9ValidationErrors는 검증 에러 패턴을 시연합니다
func example9ValidationErrors(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 9: Validation Error Patterns / 예제 9: 검증 에러 패턴")
	logger.Info("===========================================")

	logger.Info("Simulating validation errors")
	logger.Info("검증 에러 시뮬레이션")

	// Required field / 필수 필드
	logger.Info("Scenario 1: Required field missing")
	logger.Info("시나리오 1: 필수 필드 누락")
	err1 := errorutil.WithCode("VALIDATION_ERROR", "email is required")
	logger.Info("Required field error", "code", "VALIDATION_ERROR", "field", "email", "error", err1.Error())
	logger.Info("필수 필드 에러", "code", "VALIDATION_ERROR", "field", "email", "error", err1.Error())

	// Format validation / 형식 검증
	logger.Info("Scenario 2: Invalid format")
	logger.Info("시나리오 2: 잘못된 형식")
	err2 := errorutil.WithCodef("VALIDATION_ERROR", "invalid email format: %s", "notanemail")
	logger.Info("Format validation error", "code", "VALIDATION_ERROR", "input", "notanemail", "error", err2.Error())
	logger.Info("형식 검증 에러", "code", "VALIDATION_ERROR", "input", "notanemail", "error", err2.Error())

	// Range validation / 범위 검증
	logger.Info("Scenario 3: Value out of range")
	logger.Info("시나리오 3: 범위를 벗어난 값")
	err3 := errorutil.WithCodef("VALIDATION_ERROR", "age must be between 18 and 120, got %d", 150)
	logger.Info("Range validation error", "code", "VALIDATION_ERROR", "value", 150, "error", err3.Error())
	logger.Info("범위 검증 에러", "code", "VALIDATION_ERROR", "value", 150, "error", err3.Error())

	logger.Info("Example 9 completed successfully")
	logger.Info("예제 9 완료")
}

// example10ErrorClassification demonstrates error classification system
// example10ErrorClassification은 에러 분류 시스템을 시연합니다
func example10ErrorClassification(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 10: Error Classification / 예제 10: 에러 분류")
	logger.Info("===========================================")

	logger.Info("Building error classification system")
	logger.Info("에러 분류 시스템 구축")

	// Define error categories / 에러 카테고리 정의
	errors := []error{
		errorutil.WithCode("VALIDATION_ERROR", "invalid input"),
		errorutil.WithCode("AUTH_ERROR", "unauthorized"),
		errorutil.WithCode("DB_ERROR", "database failure"),
		errorutil.WithNumericCode(404, "not found"),
		errorutil.WithNumericCode(500, "internal error"),
	}

	logger.Info("Classifying errors", "count", len(errors))
	logger.Info("에러 분류 중", "count", len(errors))


	for i, err := range errors {
		logger.Info("Processing error", "index", i+1, "error", err.Error())
		logger.Info("에러 처리 중", "index", i+1, "error", err.Error())

		// Check string codes / 문자열 코드 확인
		if code, ok := errorutil.GetCode(err); ok {
			logger.Info("String code found", "code", code)
			logger.Info("문자열 코드 찾음", "code", code)

			switch code {
			case "VALIDATION_ERROR":
				logger.Info("Classification: Client error (validation)")
				logger.Info("분류: 클라이언트 에러 (검증)")
			case "AUTH_ERROR":
				logger.Info("Classification: Client error (auth)")
				logger.Info("분류: 클라이언트 에러 (인증)")
			case "DB_ERROR":
				logger.Info("Classification: Server error (database)")
				logger.Info("분류: 서버 에러 (데이터베이스)")
			}
			continue
		}

		// Check numeric codes / 숫자 코드 확인
		if numCode, ok := errorutil.GetNumericCode(err); ok {
			logger.Info("Numeric code found", "code", numCode)
			logger.Info("숫자 코드 찾음", "code", numCode)

			if numCode >= 400 && numCode < 500 {
				logger.Info("Classification: Client error (HTTP)")
				logger.Info("분류: 클라이언트 에러 (HTTP)")
			} else if numCode >= 500 {
				logger.Info("Classification: Server error (HTTP)")
				logger.Info("분류: 서버 에러 (HTTP)")
			}
		}
	}

	logger.Info("Example 10 completed successfully")
	logger.Info("예제 10 완료")
}

// example11MultiLayerWrapping demonstrates multi-layer error wrapping
// example11MultiLayerWrapping은 다중 레이어 에러 래핑을 시연합니다
func example11MultiLayerWrapping(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 11: Multi-Layer Wrapping / 예제 11: 다중 레이어 래핑")
	logger.Info("===========================================")

	logger.Info("Demonstrating error propagation through application layers")
	logger.Info("애플리케이션 레이어를 통한 에러 전파 시연")

	// Layer 1: Database / 레이어 1: 데이터베이스
	logger.Info("Layer 1: Database error occurs")
	logger.Info("레이어 1: 데이터베이스 에러 발생")
	dbErr := errorutil.WithCode("DB_TIMEOUT", "query timeout after 30s")
	logger.Info("Database error", "layer", "database", "code", "DB_TIMEOUT", "error", dbErr.Error())
	logger.Info("데이터베이스 에러", "layer", "database", "code", "DB_TIMEOUT", "error", dbErr.Error())

	// Layer 2: Repository / 레이어 2: 저장소
	logger.Info("Layer 2: Repository wraps database error")
	logger.Info("레이어 2: 저장소가 데이터베이스 에러 래핑")
	repoErr := errorutil.WrapWithCode(dbErr, "REPO_ERROR", "failed to fetch user from database")
	logger.Info("Repository error", "layer", "repository", "code", "REPO_ERROR", "error", repoErr.Error())
	logger.Info("저장소 에러", "layer", "repository", "code", "REPO_ERROR", "error", repoErr.Error())

	// Layer 3: Service / 레이어 3: 서비스
	logger.Info("Layer 3: Service wraps repository error")
	logger.Info("레이어 3: 서비스가 저장소 에러 래핑")
	serviceErr := errorutil.Wrap(repoErr, "user service error")
	logger.Info("Service error", "layer", "service", "error", serviceErr.Error())
	logger.Info("서비스 에러", "layer", "service", "error", serviceErr.Error())

	// Layer 4: HTTP Handler / 레이어 4: HTTP 핸들러
	logger.Info("Layer 4: HTTP handler wraps service error with status code")
	logger.Info("레이어 4: HTTP 핸들러가 상태 코드와 함께 서비스 에러 래핑")
	httpErr := errorutil.WrapWithNumericCode(serviceErr, 503, "service unavailable")
	logger.Info("HTTP error", "layer", "http", "code", 503, "error", httpErr.Error())
	logger.Info("HTTP 에러", "layer", "http", "code", 503, "error", httpErr.Error())


	// Verify all codes are accessible / 모든 코드에 접근 가능한지 확인
	logger.Info("Verifying code accessibility through chain")
	logger.Info("체인을 통한 코드 접근성 확인")

	hasDBTimeout := errorutil.HasCode(httpErr, "DB_TIMEOUT")
	hasRepoError := errorutil.HasCode(httpErr, "REPO_ERROR")
	has503 := errorutil.HasNumericCode(httpErr, 503)

	logger.Info("Code accessibility", "DB_TIMEOUT", hasDBTimeout, "REPO_ERROR", hasRepoError, "503", has503)
	logger.Info("코드 접근성", "DB_TIMEOUT", hasDBTimeout, "REPO_ERROR", hasRepoError, "503", has503)


	logger.Info("Example 11 completed successfully")
	logger.Info("예제 11 완료")
}

// example12StandardLibraryCompat demonstrates standard library compatibility
// example12StandardLibraryCompat은 표준 라이브러리 호환성을 시연합니다
func example12StandardLibraryCompat(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 12: Standard Library Compatibility / 예제 12: 표준 라이브러리 호환성")
	logger.Info("===========================================")

	logger.Info("Demonstrating compatibility with errors.Is and errors.As")
	logger.Info("errors.Is 및 errors.As와의 호환성 시연")

	// Create sentinel error / 센티널 에러 생성
	logger.Info("Creating sentinel error")
	logger.Info("센티널 에러 생성")
	var ErrNotFound = errorutil.WithCode("NOT_FOUND", "resource not found")

	// Create wrapped error / 래핑된 에러 생성
	logger.Info("Creating wrapped error")
	logger.Info("래핑된 에러 생성")
	err := errorutil.Wrap(ErrNotFound, "failed to fetch user")
	logger.Info("Wrapped error created", "error", err.Error())
	logger.Info("래핑된 에러 생성됨", "error", err.Error())

	// Test errors.Is / errors.Is 테스트
	logger.Info("Testing errors.Is()")
	logger.Info("errors.Is() 테스트")
	isNotFound := errors.Is(err, ErrNotFound)
	logger.Info("errors.Is result", "isNotFound", isNotFound)
	logger.Info("errors.Is 결과", "isNotFound", isNotFound)

	// Test errors.As / errors.As 테스트
	logger.Info("Testing errors.As() with Coder interface")
	logger.Info("Coder 인터페이스로 errors.As() 테스트")
	var coder interface{ Code() string }
	asCoder := errors.As(err, &coder)
	if asCoder {
		code := coder.Code()
		logger.Info("errors.As succeeded", "code", code)
		logger.Info("errors.As 성공", "code", code)
	}

	// Test with NumericCoder / NumericCoder로 테스트
	logger.Info("Testing with numeric code")
	logger.Info("숫자 코드로 테스트")
	numErr := errorutil.WithNumericCode(404, "not found")
	wrappedNumErr := errorutil.Wrap(numErr, "wrapped")

	var numCoder interface{ Code() int }
	asNumCoder := errors.As(wrappedNumErr, &numCoder)
	if asNumCoder {
		code := numCoder.Code()
		logger.Info("errors.As with NumericCoder succeeded", "code", code)
		logger.Info("NumericCoder로 errors.As 성공", "code", code)
	}

	logger.Info("Example 12 completed successfully")
	logger.Info("예제 12 완료")
}

// example13ErrorChainInspection demonstrates Root, UnwrapAll, and Contains functions
// example13ErrorChainInspection은 Root, UnwrapAll, Contains 함수를 시연합니다
func example13ErrorChainInspection(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 13: Error Chain Inspection / 예제 13: 에러 체인 검사")
	logger.Info("===========================================")

	logger.Info("Demonstrating Root(), UnwrapAll(), and Contains() functions")
	logger.Info("Root(), UnwrapAll(), Contains() 함수 시연")

	// Create a multi-layer error chain / 다층 에러 체인 생성
	logger.Info("")
	logger.Info("Creating multi-layer error chain")
	logger.Info("다층 에러 체인 생성")

	baseErr := errors.New("database connection failed")
	logger.Info("Base error created", "error", baseErr.Error())
	logger.Info("기본 에러 생성됨", "error", baseErr.Error())

	err1 := errorutil.Wrap(baseErr, "failed to connect to primary database")
	logger.Info("Layer 1 wrapped", "error", err1.Error())
	logger.Info("레이어 1 래핑됨", "error", err1.Error())

	err2 := errorutil.WrapWithCode(err1, "DB_ERROR", "database operation failed")
	logger.Info("Layer 2 wrapped with code", "error", err2.Error())
	logger.Info("코드와 함께 레이어 2 래핑됨", "error", err2.Error())

	err3 := errorutil.Wrap(err2, "failed to fetch user data")
	logger.Info("Layer 3 wrapped", "error", err3.Error())
	logger.Info("레이어 3 래핑됨", "error", err3.Error())

	// Test Root() function / Root() 함수 테스트
	logger.Info("")
	logger.Info("Testing Root() function")
	logger.Info("Root() 함수 테스트")

	root := errorutil.Root(err3)
	logger.Info("Root error found", "root", root.Error())
	logger.Info("루트 에러 발견", "root", root.Error())
	logger.Info("Root matches base error", "matches", root.Error() == baseErr.Error())
	logger.Info("루트가 기본 에러와 일치", "matches", root.Error() == baseErr.Error())

	// Test UnwrapAll() function / UnwrapAll() 함수 테스트
	logger.Info("")
	logger.Info("Testing UnwrapAll() function")
	logger.Info("UnwrapAll() 함수 테스트")

	chain := errorutil.UnwrapAll(err3)
	logger.Info("Total errors in chain", "count", len(chain))
	logger.Info("체인의 총 에러 개수", "count", len(chain))

	for i, e := range chain {
		logger.Info("Error chain level", "level", i, "error", e.Error())
		logger.Info("에러 체인 레벨", "level", i, "error", e.Error())
	}

	// Test Contains() function / Contains() 함수 테스트
	logger.Info("")
	logger.Info("Testing Contains() function")
	logger.Info("Contains() 함수 테스트")

	// Create sentinel errors / 센티널 에러 생성
	var ErrNotFound = errors.New("not found")
	var ErrTimeout = errors.New("timeout")

	// Create error chain with sentinel error / 센티널 에러로 에러 체인 생성
	notFoundErr := errorutil.Wrap(ErrNotFound, "user not found")
	wrappedNotFound := errorutil.Wrap(notFoundErr, "failed to get user profile")

	containsNotFound := errorutil.Contains(wrappedNotFound, ErrNotFound)
	logger.Info("Contains ErrNotFound", "result", containsNotFound)
	logger.Info("ErrNotFound 포함", "result", containsNotFound)

	containsTimeout := errorutil.Contains(wrappedNotFound, ErrTimeout)
	logger.Info("Contains ErrTimeout", "result", containsTimeout)
	logger.Info("ErrTimeout 포함", "result", containsTimeout)

	// Real-world use case: Error chain analysis / 실제 사용 사례: 에러 체인 분석
	logger.Info("")
	logger.Info("Real-world use case: Error chain analysis")
	logger.Info("실제 사용 사례: 에러 체인 분석")

	// Create a complex error scenario / 복잡한 에러 시나리오 생성
	dbErr := errorutil.WithNumericCode(500, "internal database error")
	serviceErr := errorutil.WrapWithCode(dbErr, "SVC_ERROR", "service unavailable")
	apiErr := errorutil.WrapWithNumericCode(serviceErr, 503, "API temporarily unavailable")

	logger.Info("Complex error created", "error", apiErr.Error())
	logger.Info("복잡한 에러 생성됨", "error", apiErr.Error())

	// Analyze the error chain / 에러 체인 분석
	logger.Info("Analyzing error chain:")
	logger.Info("에러 체인 분석:")

	allErrors := errorutil.UnwrapAll(apiErr)
	logger.Info("Chain depth", "depth", len(allErrors))
	logger.Info("체인 깊이", "depth", len(allErrors))

	rootCause := errorutil.Root(apiErr)
	logger.Info("Root cause", "error", rootCause.Error())
	logger.Info("근본 원인", "error", rootCause.Error())

	// Check for specific error codes in the chain / 체인에서 특정 에러 코드 확인
	if errorutil.HasCode(apiErr, "SVC_ERROR") {
		logger.Info("Found service error in chain")
		logger.Info("체인에서 서비스 에러 발견")
	}

	if errorutil.HasNumericCode(apiErr, 500) {
		logger.Info("Found HTTP 500 error in chain")
		logger.Info("체인에서 HTTP 500 에러 발견")
	}

	// Use case: Detailed error logging / 사용 사례: 상세 에러 로깅
	logger.Info("")
	logger.Info("Use case: Detailed error logging for debugging")
	logger.Info("사용 사례: 디버깅을 위한 상세 에러 로깅")

	testErr := simulateComplexOperation()
	if testErr != nil {
		logger.Info("Operation failed, analyzing error chain:")
		logger.Info("작업 실패, 에러 체인 분석:")

		// Log all errors in the chain / 체인의 모든 에러 로깅
		errorChain := errorutil.UnwrapAll(testErr)
		for i, e := range errorChain {
			logger.Info("Chain analysis",
				"depth", i,
				"error", e.Error(),
				"type", fmt.Sprintf("%T", e))
			logger.Info("체인 분석",
				"깊이", i,
				"에러", e.Error(),
				"타입", fmt.Sprintf("%T", e))
		}

		// Get root cause for reporting / 보고를 위한 근본 원인 가져오기
		root := errorutil.Root(testErr)
		logger.Info("Root cause for error report", "root", root.Error())
		logger.Info("에러 보고서를 위한 근본 원인", "root", root.Error())
	}

	logger.Info("")
	logger.Info("Example 13 completed successfully")
	logger.Info("예제 13 완료")
}

// simulateComplexOperation simulates a complex operation that can fail at multiple levels
// simulateComplexOperation은 여러 레벨에서 실패할 수 있는 복잡한 작업을 시뮬레이션합니다
func simulateComplexOperation() error {
	// Simulate a low-level error / 저수준 에러 시뮬레이션
	lowLevelErr := errors.New("network timeout")

	// Wrap at middleware layer / 미들웨어 레이어에서 래핑
	middlewareErr := errorutil.WrapWithCode(lowLevelErr, "MIDDLEWARE_ERROR", "request processing failed")

	// Wrap at service layer / 서비스 레이어에서 래핑
	serviceErr := errorutil.WrapWithNumericCode(middlewareErr, 504, "gateway timeout")

	// Wrap at API layer / API 레이어에서 래핑
	apiErr := errorutil.Wrap(serviceErr, "failed to complete user request")

	return apiErr
}
