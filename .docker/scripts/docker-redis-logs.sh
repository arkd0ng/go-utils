#!/bin/bash
#
# Script Name: docker-redis-logs.sh
# Description: Displays Docker Redis container logs in real-time.
#              This script follows the Redis container logs and displays them
#              in the terminal. Useful for debugging and monitoring.
#
#              Docker Redis 컨테이너 로그를 실시간으로 표시합니다.
#              이 스크립트는 Redis 컨테이너 로그를 따라가며 터미널에 표시합니다.
#              디버깅 및 모니터링에 유용합니다.
#
# Usage: ./docker-redis-logs.sh [--tail N]
#        사용법: ./docker-redis-logs.sh [--tail N]
#
# Options / 옵션:
#   --tail N    Show last N lines (default: all)
#               마지막 N줄만 표시 (기본값: 전체)
#
# Prerequisites / 사전 요구사항:
#   - Redis container must be running
#     Redis 컨테이너가 실행 중이어야 함
#
# Exit Codes / 종료 코드:
#   0 - Success / 성공
#   1 - Docker not installed or container not running
#       Docker 미설치 또는 컨테이너 미실행
#   130 - Interrupted by user (Ctrl+C)
#         사용자 중단 (Ctrl+C)
#
# Examples / 예제:
#   # View all logs with real-time following / 실시간으로 전체 로그 보기
#   ./docker-redis-logs.sh
#
#   # View last 100 lines / 마지막 100줄 보기
#   docker logs --tail 100 -f go-utils-redis
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

echo -e "${GREEN}Viewing Docker Redis logs...${NC}"
echo -e "${GREEN}Docker Redis 로그 확인 중...${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop / Ctrl+C를 눌러 중지${NC}"
echo

# Follow logs with real-time updates / 실시간 업데이트로 로그 팔로우
docker logs -f "${CONTAINER_NAME}"
