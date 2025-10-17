#!/bin/bash
#
# Script Name: docker-redis-start.sh
# Description: Starts Docker Redis container for go-utils development and testing.
#              This script checks Docker availability, starts the Redis container
#              using Docker Compose, and waits for Redis to be ready for connections.
#
#              go-utils 개발 및 테스트를 위한 Docker Redis 컨테이너를 시작합니다.
#              이 스크립트는 Docker 사용 가능 여부를 확인하고, Docker Compose를 사용하여
#              Redis 컨테이너를 시작한 후, Redis가 연결 준비될 때까지 대기합니다.
#
# Usage: ./docker-redis-start.sh
#        사용법: ./docker-redis-start.sh
#
# Prerequisites / 사전 요구사항:
#   - Docker Desktop installed and running
#     Docker Desktop 설치 및 실행 중
#   - docker-compose.yml in project root
#     프로젝트 루트에 docker-compose.yml 존재
#
# Exit Codes / 종료 코드:
#   0 - Success (Redis started and ready)
#       성공 (Redis 시작 및 준비 완료)
#   1 - Docker not installed or not running
#       Docker 미설치 또는 미실행
#   1 - Redis failed to become ready within timeout
#       타임아웃 내 Redis 준비 실패
#
# Examples / 예제:
#   # Start Redis container / Redis 컨테이너 시작
#   ./docker-redis-start.sh
#
#   # Check if already running / 이미 실행 중인지 확인
#   docker ps --filter "name=go-utils-redis"
#
# Author: arkd0ng
# Created: 2024
# Modified: 2025-10-17
#

# Exit on error / 에러 시 종료
set -e

# Exit on undefined variable / 미정의 변수 사용 시 종료
set -u

# Pipe failure causes exit / 파이프 실패 시 종료
set -o pipefail

# ============================================================================
# Configuration / 구성
# ============================================================================

# ANSI color codes for terminal output
# 터미널 출력을 위한 ANSI 색상 코드
readonly RED='\033[0;31m'      # Error messages / 에러 메시지
readonly GREEN='\033[0;32m'    # Success messages / 성공 메시지
readonly YELLOW='\033[1;33m'   # Warning messages / 경고 메시지
readonly NC='\033[0m'          # No Color (reset) / 색상 없음 (리셋)

# Redis container configuration
# Redis 컨테이너 구성
readonly CONTAINER_NAME="go-utils-redis"
readonly MAX_WAIT_ATTEMPTS=30  # Maximum seconds to wait for Redis / Redis 대기 최대 초

# ============================================================================
# Functions / 함수
# ============================================================================

# Function: check_docker_installed
# Description: Checks if Docker command is available in the system.
#              시스템에서 Docker 명령을 사용할 수 있는지 확인합니다.
#
# Returns / 반환값:
#   0 - Docker is installed / Docker 설치됨
#   1 - Docker is not installed / Docker 미설치
check_docker_installed() {
    if ! command -v docker &> /dev/null; then
        echo -e "${RED}Error: Docker is not installed${NC}" >&2
        echo -e "${RED}오류: Docker가 설치되어 있지 않습니다${NC}" >&2
        echo
        echo "Please install Docker Desktop from:"
        echo "Docker Desktop을 다음에서 설치하세요:"
        echo "https://www.docker.com/products/docker-desktop"
        return 1
    fi
    return 0
}

# Function: check_docker_running
# Description: Checks if Docker daemon is running and accessible.
#              Docker 데몬이 실행 중이고 접근 가능한지 확인합니다.
#
# Returns / 반환값:
#   0 - Docker is running / Docker 실행 중
#   1 - Docker is not running / Docker 미실행
check_docker_running() {
    if ! docker info &> /dev/null; then
        echo -e "${RED}Error: Docker is not running${NC}" >&2
        echo -e "${RED}오류: Docker가 실행 중이지 않습니다${NC}" >&2
        echo
        echo "Please start Docker Desktop and try again"
        echo "Docker Desktop을 시작한 후 다시 시도하세요"
        return 1
    fi
    return 0
}

# Function: check_container_running
# Description: Checks if Redis container is already running.
#              Redis 컨테이너가 이미 실행 중인지 확인합니다.
#
# Parameters / 매개변수:
#   $1 - Container name / 컨테이너 이름
#
# Returns / 반환값:
#   0 - Container is running / 컨테이너 실행 중
#   1 - Container is not running / 컨테이너 미실행
check_container_running() {
    local container_name="$1"
    docker ps --filter "name=${container_name}" --format "{{.Names}}" | grep -q "${container_name}"
}

# Function: wait_for_redis
# Description: Waits for Redis to be ready to accept connections.
#              Attempts to ping Redis server until it responds or timeout.
#
#              Redis가 연결을 수락할 준비가 될 때까지 대기합니다.
#              응답할 때까지 또는 타임아웃될 때까지 Redis 서버에 ping을 시도합니다.
#
# Parameters / 매개변수:
#   $1 - Container name / 컨테이너 이름
#   $2 - Maximum wait attempts in seconds / 최대 대기 시도 (초)
#
# Returns / 반환값:
#   0 - Redis is ready / Redis 준비 완료
#   1 - Timeout reached / 타임아웃 도달
wait_for_redis() {
    local container_name="$1"
    local max_attempts="$2"
    local attempt=0
    
    echo -e "${GREEN}Waiting for Redis to be ready...${NC}"
    echo -e "${GREEN}Redis 준비 대기 중...${NC}"
    
    while [ $attempt -lt $max_attempts ]; do
        if docker exec "${container_name}" redis-cli ping &> /dev/null; then
            return 0
        fi
        attempt=$((attempt + 1))
        echo -n "."
        sleep 1
    done
    
    return 1
}

# ============================================================================
# Main Script / 메인 스크립트
# ============================================================================

echo -e "${GREEN}Starting Docker Redis for go-utils...${NC}"
echo -e "${GREEN}go-utils를 위한 Docker Redis 시작 중...${NC}"
echo

# Check prerequisites / 사전 요구사항 확인
check_docker_installed || exit 1
check_docker_running || exit 1

# Get script directory and project root
# 스크립트 디렉토리와 프로젝트 루트 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# Check if container is already running
# 컨테이너가 이미 실행 중인지 확인
if check_container_running "${CONTAINER_NAME}"; then
    echo -e "${YELLOW}Docker Redis is already running${NC}"
    echo -e "${YELLOW}Docker Redis가 이미 실행 중입니다${NC}"
    echo
    docker ps --filter "name=${CONTAINER_NAME}"
    echo
    echo "Connection details / 연결 정보:"
    echo "  Host: localhost"
    echo "  Port: 6379"
    echo "  Password: (none)"
    echo
    echo "To connect to Redis CLI / Redis CLI 연결:"
    echo "  ./scripts/docker-redis-cli.sh"
    exit 0
fi

# Start Docker Compose (Redis only) / Docker Compose 시작 (Redis만)
echo -e "${GREEN}Starting Docker Compose (Redis only)...${NC}"
echo -e "${GREEN}Docker Compose 시작 중 (Redis만)...${NC}"
docker compose up -d redis

echo

# Wait for Redis to be ready / Redis 준비 대기
if wait_for_redis "${CONTAINER_NAME}" "${MAX_WAIT_ATTEMPTS}"; then
    echo
    echo -e "${GREEN}✅ Docker Redis is ready!${NC}"
    echo -e "${GREEN}✅ Docker Redis 준비 완료!${NC}"
    echo
    echo "Connection details / 연결 정보:"
    echo "  Host: localhost"
    echo "  Port: 6379"
    echo "  Password: (none)"
    echo
    echo "To stop Redis / Redis 중지:"
    echo "  ./scripts/docker-redis-stop.sh"
    echo
    echo "To view logs / 로그 보기:"
    echo "  ./scripts/docker-redis-logs.sh"
    echo
    echo "To connect to Redis CLI / Redis CLI 연결:"
    echo "  ./scripts/docker-redis-cli.sh"
    echo
    exit 0
else
    echo
    echo -e "${RED}Error: Redis failed to become ready within ${MAX_WAIT_ATTEMPTS} seconds${NC}"
    echo -e "${RED}오류: ${MAX_WAIT_ATTEMPTS}초 내에 Redis가 준비되지 않았습니다${NC}"
    echo
    echo "Check logs with / 로그 확인:"
    echo "  ./scripts/docker-redis-logs.sh"
    exit 1
fi
