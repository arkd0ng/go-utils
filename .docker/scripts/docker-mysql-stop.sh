#!/bin/bash
#
# Script Name: docker-mysql-stop.sh
# Description: Stops and removes the Docker MySQL container for go-utils.
#              This script gracefully stops the MySQL container using Docker Compose
#              and removes associated volumes and networks.
#
#              go-utils를 위한 Docker MySQL 컨테이너를 중지하고 제거합니다.
#              이 스크립트는 Docker Compose를 사용하여 MySQL 컨테이너를 안전하게 중지하고
#              관련 볼륨 및 네트워크를 제거합니다.
#
# Usage: ./docker-mysql-stop.sh
#        사용법: ./docker-mysql-stop.sh
#
# Prerequisites / 사전 요구사항:
#   - Docker Desktop installed and running
#     Docker Desktop 설치 및 실행 중
#   - MySQL container previously started
#     MySQL 컨테이너가 이전에 시작됨
#
# Exit Codes / 종료 코드:
#   0 - Success (MySQL stopped)
#       성공 (MySQL 중지됨)
#   1 - Docker not installed
#       Docker 미설치
#
# Examples / 예제:
#   # Stop MySQL container / MySQL 컨테이너 중지
#   ./docker-mysql-stop.sh
#
#   # Verify container is stopped / 컨테이너 중지 확인
#   docker ps -a --filter "name=go-utils-mysql"
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

# ============================================================================
# Main Script / 메인 스크립트
# ============================================================================

echo -e "${YELLOW}Stopping Docker MySQL for go-utils...${NC}"
echo -e "${YELLOW}go-utils를 위한 Docker MySQL 중지 중...${NC}"
echo

# Check prerequisites / 사전 요구사항 확인
check_docker_installed || exit 1

# Get script directory and project root
# 스크립트 디렉토리와 프로젝트 루트 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# Stop Docker Compose / Docker Compose 중지
echo -e "${YELLOW}Stopping Docker Compose...${NC}"
echo -e "${YELLOW}Docker Compose 중지 중...${NC}"
docker compose down

echo
echo -e "${GREEN}✅ Docker MySQL stopped successfully${NC}"
echo -e "${GREEN}✅ Docker MySQL이 성공적으로 중지되었습니다${NC}"
echo
echo "To start MySQL again / MySQL 다시 시작:"
echo "  ./scripts/docker-mysql-start.sh"
echo
