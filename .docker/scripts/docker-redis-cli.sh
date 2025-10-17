#!/bin/bash
#
# Script Name: docker-redis-cli.sh
# Description: Opens an interactive Redis CLI session in the Docker container.
#              This script connects to the running Redis container and provides
#              an interactive command-line interface for Redis operations.
#
#              Docker 컨테이너에서 대화형 Redis CLI 세션을 엽니다.
#              이 스크립트는 실행 중인 Redis 컨테이너에 연결하고
#              Redis 작업을 위한 대화형 명령줄 인터페이스를 제공합니다.
#
# Usage: ./docker-redis-cli.sh [redis-cli arguments]
#        사용법: ./docker-redis-cli.sh [redis-cli 인수]
#
# Redis CLI Commands / Redis CLI 명령:
#   PING              Test connection / 연결 테스트
#   SET key value     Set a key / 키 설정
#   GET key           Get a key / 키 가져오기
#   KEYS pattern      List keys / 키 목록
#   DEL key           Delete a key / 키 삭제
#   FLUSHDB           Clear current database / 현재 데이터베이스 지우기
#   INFO              Server information / 서버 정보
#
# Prerequisites / 사전 요구사항:
#   - Redis container must be running
#     Redis 컨테이너가 실행 중이어야 함
#
# Exit Codes / 종료 코드:
#   0 - Success / 성공
#   1 - Docker not installed or container not running
#       Docker 미설치 또는 컨테이너 미실행
#
# Examples / 예제:
#   # Open interactive CLI / 대화형 CLI 열기
#   ./docker-redis-cli.sh
#
#   # Execute single command / 단일 명령 실행
#   docker exec go-utils-redis redis-cli PING
#
#   # Set and get a value / 값 설정 및 가져오기
#   In CLI:
#   > SET mykey "Hello"
#   > GET mykey
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

# Container name / 컨테이너 이름
readonly CONTAINER_NAME="go-utils-redis"

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
        return 1
    fi
    return 0
}

# Function: check_container_running
# Description: Checks if Redis container is running.
#              Redis 컨테이너가 실행 중인지 확인합니다.
#
# Parameters / 매개변수:
#   $1 - Container name / 컨테이너 이름
#
# Returns / 반환값:
#   0 - Container is running / 컨테이너 실행 중
#   1 - Container is not running / 컨테이너 미실행
check_container_running() {
    local container_name="$1"
    if ! docker ps --filter "name=${container_name}" --format "{{.Names}}" | grep -q "${container_name}"; then
        echo -e "${RED}Error: Docker Redis is not running${NC}" >&2
        echo -e "${RED}오류: Docker Redis가 실행 중이지 않습니다${NC}" >&2
        echo
        echo "Start Redis with / Redis 시작:"
        echo "  ./scripts/docker-redis-start.sh"
        return 1
    fi
    return 0
}

# ============================================================================
# Main Script / 메인 스크립트
# ============================================================================

# Check prerequisites / 사전 요구사항 확인
check_docker_installed || exit 1
check_container_running "${CONTAINER_NAME}" || exit 1

echo -e "${GREEN}Connecting to Docker Redis CLI...${NC}"
echo -e "${GREEN}Docker Redis CLI 연결 중...${NC}"
echo -e "${YELLOW}Type 'exit' or press Ctrl+D to quit / 'exit' 입력 또는 Ctrl+D를 눌러 종료${NC}"
echo

# Connect to Redis CLI interactively / 대화형으로 Redis CLI 연결
docker exec -it "${CONTAINER_NAME}" redis-cli
