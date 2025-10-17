# Makefile for go-utils project
# go-utils 프로젝트용 Makefile
#
# This Makefile provides convenient targets for building, testing, and managing
# the go-utils project. It automates common development tasks.
#
# 이 Makefile은 go-utils 프로젝트를 빌드, 테스트 및 관리하기 위한
# 편리한 타겟을 제공합니다. 일반적인 개발 작업을 자동화합니다.
#
# Available targets / 사용 가능한 타겟:
#   make help       - Show this help message / 도움말 메시지 표시
#   make build      - Build all packages / 모든 패키지 빌드
#   make test       - Run all tests / 모든 테스트 실행
#   make test-v     - Run tests with verbose output / 상세 출력으로 테스트 실행
#   make coverage   - Run tests with coverage report / 커버리지 리포트와 함께 테스트 실행
#   make clean      - Remove build artifacts / 빌드 산출물 제거
#   make fmt        - Format code / 코드 포맷
#   make lint       - Run golangci-lint / golangci-lint 실행
#   make vet        - Run go vet / go vet 실행
#   make docker-mysql-start - Start MySQL container / MySQL 컨테이너 시작
#   make docker-mysql-stop  - Stop MySQL container / MySQL 컨테이너 중지
#   make docker-redis-start - Start Redis container / Redis 컨테이너 시작
#   make docker-redis-stop  - Stop Redis container / Redis 컨테이너 중지
#
# Examples / 예제:
#   make build        # Build all packages / 모든 패키지 빌드
#   make test         # Run all tests / 모든 테스트 실행
#   make coverage     # Run tests with coverage / 커버리지와 함께 테스트 실행
#   make clean        # Clean build artifacts / 빌드 산출물 정리
#
# Author: arkd0ng
# Created: 2025-10-17
#

# ============================================================================
# Configuration / 구성
# ============================================================================

# Go parameters / Go 매개변수
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOMOD=$(GOCMD) mod

# Project directories / 프로젝트 디렉토리
PACKAGES=$(shell go list ./... | grep -v /examples/)
DOCKER_SCRIPTS=.docker/scripts

# Test coverage output / 테스트 커버리지 출력
COVERAGE_OUT=coverage.out
COVERAGE_HTML=coverage.html

# ============================================================================
# Phony targets / Phony 타겟
# ============================================================================

.PHONY: all build test test-v coverage clean fmt lint vet help \
        docker-mysql-start docker-mysql-stop docker-mysql-logs \
        docker-redis-start docker-redis-stop docker-redis-logs docker-redis-cli \
        deps tidy

# ============================================================================
# Default target / 기본 타겟
# ============================================================================

# Default target: show help
# 기본 타겟: 도움말 표시
all: help

# ============================================================================
# Build targets / 빌드 타겟
# ============================================================================

# Build all packages
# 모든 패키지 빌드
#
# This target compiles all Go packages in the project to verify
# that the code builds successfully without creating executables.
#
# 이 타겟은 프로젝트의 모든 Go 패키지를 컴파일하여
# 실행 파일을 생성하지 않고 코드가 성공적으로 빌드되는지 확인합니다.
build:
	@echo "Building all packages..."
	@echo "모든 패키지 빌드 중..."
	@$(GOBUILD) ./...
	@echo "✅ Build successful!"
	@echo "✅ 빌드 성공!"

# ============================================================================
# Test targets / 테스트 타겟
# ============================================================================

# Run all tests
# 모든 테스트 실행
#
# Runs all unit tests across the project.
# 프로젝트 전체의 모든 단위 테스트를 실행합니다.
test:
	@echo "Running tests..."
	@echo "테스트 실행 중..."
	@$(GOTEST) ./...
	@echo "✅ Tests passed!"
	@echo "✅ 테스트 통과!"

# Run tests with verbose output
# 상세 출력으로 테스트 실행
#
# Runs all unit tests with detailed output showing each test case.
# 각 테스트 케이스를 보여주는 상세 출력으로 모든 단위 테스트를 실행합니다.
test-v:
	@echo "Running tests with verbose output..."
	@echo "상세 출력으로 테스트 실행 중..."
	@$(GOTEST) -v ./...

# Run tests with coverage report
# 커버리지 리포트와 함께 테스트 실행
#
# Runs all tests and generates a coverage report showing which
# lines of code are tested. Creates both text and HTML reports.
#
# 모든 테스트를 실행하고 어떤 코드 라인이 테스트되었는지
# 보여주는 커버리지 리포트를 생성합니다. 텍스트와 HTML 리포트를 모두 생성합니다.
coverage:
	@echo "Running tests with coverage..."
	@echo "커버리지와 함께 테스트 실행 중..."
	@$(GOTEST) -coverprofile=$(COVERAGE_OUT) ./...
	@$(GOCMD) tool cover -func=$(COVERAGE_OUT)
	@$(GOCMD) tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)
	@echo ""
	@echo "✅ Coverage report generated!"
	@echo "✅ 커버리지 리포트 생성 완료!"
	@echo "   Text report: $(COVERAGE_OUT)"
	@echo "   HTML report: $(COVERAGE_HTML)"

# ============================================================================
# Code quality targets / 코드 품질 타겟
# ============================================================================

# Format code
# 코드 포맷
#
# Formats all Go source files using gofmt standard.
# gofmt 표준을 사용하여 모든 Go 소스 파일을 포맷합니다.
fmt:
	@echo "Formatting code..."
	@echo "코드 포맷 중..."
	@$(GOFMT) ./...
	@echo "✅ Code formatted!"
	@echo "✅ 코드 포맷 완료!"

# Run go vet
# go vet 실행
#
# Examines Go source code and reports suspicious constructs.
# Go 소스 코드를 검사하고 의심스러운 구조를 보고합니다.
vet:
	@echo "Running go vet..."
	@echo "go vet 실행 중..."
	@$(GOVET) ./...
	@echo "✅ Vet check passed!"
	@echo "✅ Vet 검사 통과!"

# Run golangci-lint (if installed)
# golangci-lint 실행 (설치된 경우)
#
# Runs golangci-lint if available, otherwise shows installation instructions.
# golangci-lint가 사용 가능한 경우 실행하고, 그렇지 않으면 설치 방법을 표시합니다.
lint:
	@echo "Running linter..."
	@echo "린터 실행 중..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
		echo "✅ Lint check passed!"; \
		echo "✅ 린트 검사 통과!"; \
	else \
		echo "⚠️  golangci-lint not installed"; \
		echo "⚠️  golangci-lint가 설치되지 않음"; \
		echo "Install: https://golangci-lint.run/usage/install/"; \
	fi

# ============================================================================
# Dependency management / 의존성 관리
# ============================================================================

# Download dependencies
# 의존성 다운로드
#
# Downloads all Go module dependencies.
# 모든 Go 모듈 의존성을 다운로드합니다.
deps:
	@echo "Downloading dependencies..."
	@echo "의존성 다운로드 중..."
	@$(GOMOD) download
	@echo "✅ Dependencies downloaded!"
	@echo "✅ 의존성 다운로드 완료!"

# Tidy dependencies
# 의존성 정리
#
# Adds missing and removes unused module dependencies.
# 누락된 의존성을 추가하고 사용하지 않는 모듈 의존성을 제거합니다.
tidy:
	@echo "Tidying dependencies..."
	@echo "의존성 정리 중..."
	@$(GOMOD) tidy
	@echo "✅ Dependencies tidied!"
	@echo "✅ 의존성 정리 완료!"

# ============================================================================
# Cleanup targets / 정리 타겟
# ============================================================================

# Clean build artifacts and cache
# 빌드 산출물 및 캐시 정리
#
# Removes all build artifacts, test caches, and coverage reports.
# 모든 빌드 산출물, 테스트 캐시 및 커버리지 리포트를 제거합니다.
clean:
	@echo "Cleaning build artifacts..."
	@echo "빌드 산출물 정리 중..."
	@$(GOCLEAN)
	@rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)
	@rm -f validation_coverage.out
	@echo "✅ Clean complete!"
	@echo "✅ 정리 완료!"

# ============================================================================
# Docker targets / Docker 타겟
# ============================================================================

# Start MySQL container
# MySQL 컨테이너 시작
#
# Starts the MySQL Docker container for development and testing.
# 개발 및 테스트를 위한 MySQL Docker 컨테이너를 시작합니다.
docker-mysql-start:
	@$(DOCKER_SCRIPTS)/docker-mysql-start.sh

# Stop MySQL container
# MySQL 컨테이너 중지
#
# Stops the MySQL Docker container.
# MySQL Docker 컨테이너를 중지합니다.
docker-mysql-stop:
	@$(DOCKER_SCRIPTS)/docker-mysql-stop.sh

# View MySQL logs
# MySQL 로그 보기
#
# Displays MySQL container logs in real-time.
# MySQL 컨테이너 로그를 실시간으로 표시합니다.
docker-mysql-logs:
	@$(DOCKER_SCRIPTS)/docker-mysql-logs.sh

# Start Redis container
# Redis 컨테이너 시작
#
# Starts the Redis Docker container for development and testing.
# 개발 및 테스트를 위한 Redis Docker 컨테이너를 시작합니다.
docker-redis-start:
	@$(DOCKER_SCRIPTS)/docker-redis-start.sh

# Stop Redis container
# Redis 컨테이너 중지
#
# Stops the Redis Docker container.
# Redis Docker 컨테이너를 중지합니다.
docker-redis-stop:
	@$(DOCKER_SCRIPTS)/docker-redis-stop.sh

# View Redis logs
# Redis 로그 보기
#
# Displays Redis container logs in real-time.
# Redis 컨테이너 로그를 실시간으로 표시합니다.
docker-redis-logs:
	@$(DOCKER_SCRIPTS)/docker-redis-logs.sh

# Connect to Redis CLI
# Redis CLI 연결
#
# Opens an interactive Redis CLI session.
# 대화형 Redis CLI 세션을 엽니다.
docker-redis-cli:
	@$(DOCKER_SCRIPTS)/docker-redis-cli.sh

# ============================================================================
# Help target / 도움말 타겟
# ============================================================================

# Show help message
# 도움말 메시지 표시
#
# Displays all available make targets with descriptions.
# 설명과 함께 사용 가능한 모든 make 타겟을 표시합니다.
help:
	@echo "================================================================================"
	@echo "go-utils Makefile"
	@echo "================================================================================"
	@echo ""
	@echo "📦 Build targets / 빌드 타겟:"
	@echo "  make build              Build all packages / 모든 패키지 빌드"
	@echo ""
	@echo "🧪 Test targets / 테스트 타겟:"
	@echo "  make test               Run all tests / 모든 테스트 실행"
	@echo "  make test-v             Run tests with verbose output / 상세 출력으로 테스트 실행"
	@echo "  make coverage           Run tests with coverage / 커버리지와 함께 테스트 실행"
	@echo ""
	@echo "✨ Code quality / 코드 품질:"
	@echo "  make fmt                Format code / 코드 포맷"
	@echo "  make vet                Run go vet / go vet 실행"
	@echo "  make lint               Run golangci-lint / golangci-lint 실행"
	@echo ""
	@echo "📚 Dependencies / 의존성:"
	@echo "  make deps               Download dependencies / 의존성 다운로드"
	@echo "  make tidy               Tidy dependencies / 의존성 정리"
	@echo ""
	@echo "🐳 Docker MySQL:"
	@echo "  make docker-mysql-start Start MySQL container / MySQL 컨테이너 시작"
	@echo "  make docker-mysql-stop  Stop MySQL container / MySQL 컨테이너 중지"
	@echo "  make docker-mysql-logs  View MySQL logs / MySQL 로그 보기"
	@echo ""
	@echo "🐳 Docker Redis:"
	@echo "  make docker-redis-start Start Redis container / Redis 컨테이너 시작"
	@echo "  make docker-redis-stop  Stop Redis container / Redis 컨테이너 중지"
	@echo "  make docker-redis-logs  View Redis logs / Redis 로그 보기"
	@echo "  make docker-redis-cli   Connect to Redis CLI / Redis CLI 연결"
	@echo ""
	@echo "🧹 Cleanup / 정리:"
	@echo "  make clean              Remove build artifacts / 빌드 산출물 제거"
	@echo ""
	@echo "❓ Help / 도움말:"
	@echo "  make help               Show this help message / 도움말 표시"
	@echo ""
	@echo "================================================================================"
	@echo "Examples / 예제:"
	@echo "  make build && make test           # Build and test / 빌드 및 테스트"
	@echo "  make docker-mysql-start && make test  # Start MySQL and test / MySQL 시작 및 테스트"
	@echo "================================================================================"
