#!/bin/bash
#
# Script Name: docker-redis-stop.sh
# Description: Stops and optionally removes the Docker Redis container for go-utils.
#              This script gracefully stops the Redis container using Docker Compose.
#              Optionally removes the Redis data volume if requested by user.
#
#              go-utils를 위한 Docker Redis 컨테이너를 중지하고 선택적으로 제거합니다.
#              이 스크립트는 Docker Compose를 사용하여 Redis 컨테이너를 안전하게 중지합니다.
#              사용자 요청 시 Redis 데이터 볼륨을 선택적으로 제거합니다.
#
# Usage: ./docker-redis-stop.sh
#        사용법: ./docker-redis-stop.sh
#
# Prerequisites / 사전 요구사항:
#   - Docker Desktop installed and running
#     Docker Desktop 설치 및 실행 중
#   - Redis container previously started
#     Redis 컨테이너가 이전에 시작됨
#
# Exit Codes / 종료 코드:
#   0 - Success (Redis stopped)
#       성공 (Redis 중지됨)
#   1 - Docker not installed
#       Docker 미설치
#
# Examples / 예제:
#   # Stop Redis container / Redis 컨테이너 중지
#   ./docker-redis-stop.sh
#
#   # Verify container is stopped / 컨테이너 중지 확인
#   docker ps -a --filter "name=go-utils-redis"
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
readonly VOLUME_NAME="go-utils-redis-data"

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

# Function: check_container_exists
# Description: Checks if Redis container exists (running or stopped).
#              Redis 컨테이너가 존재하는지 확인합니다 (실행 중 또는 중지됨).
#
# Parameters / 매개변수:
#   $1 - Container name / 컨테이너 이름
#
# Returns / 반환값:
#   0 - Container exists / 컨테이너 존재
#   1 - Container does not exist / 컨테이너 존재하지 않음
check_container_exists() {
    local container_name="$1"
    docker ps -a --filter "name=${container_name}" --format "{{.Names}}" | grep -q "${container_name}"
}

# ============================================================================
# Main Script / 메인 스크립트
# ============================================================================

echo -e "${GREEN}Stopping Docker Redis for go-utils...${NC}"
echo -e "${GREEN}go-utils를 위한 Docker Redis 중지 중...${NC}"
echo

# Check prerequisites / 사전 요구사항 확인
check_docker_installed || exit 1

# Get script directory and project root
# 스크립트 디렉토리와 프로젝트 루트 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# Check if container exists / 컨테이너 존재 확인
if ! check_container_exists "${CONTAINER_NAME}"; then
    echo -e "${YELLOW}Docker Redis container not found${NC}"
    echo -e "${YELLOW}Docker Redis 컨테이너를 찾을 수 없습니다${NC}"
    exit 0
fi

# Stop Docker Compose (Redis only) / Docker Compose 중지 (Redis만)
echo -e "${GREEN}Stopping Docker Compose (Redis only)...${NC}"
echo -e "${GREEN}Docker Compose 중지 중 (Redis만)...${NC}"
docker compose down redis

echo
echo -e "${GREEN}✅ Docker Redis stopped successfully!${NC}"
echo -e "${GREEN}✅ Docker Redis 성공적으로 중지되었습니다!${NC}"
echo

# Optional: Remove volume / 선택사항: 볼륨 제거
read -p "Do you want to remove the Redis data volume? (y/N) / Redis 데이터 볼륨을 제거하시겠습니까? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}Removing Redis data volume...${NC}"
    echo -e "${GREEN}Redis 데이터 볼륨 제거 중...${NC}"
    docker volume rm "${VOLUME_NAME}" 2>/dev/null || true
    echo -e "${GREEN}✅ Redis data volume removed${NC}"
    echo -e "${GREEN}✅ Redis 데이터 볼륨 제거됨${NC}"
else
    echo -e "${YELLOW}Redis data volume preserved${NC}"
    echo -e "${YELLOW}Redis 데이터 볼륨 보존됨${NC}"
fi

echo
echo "To start Redis again / Redis 다시 시작:"
echo "  ./scripts/docker-redis-start.sh"
echo
