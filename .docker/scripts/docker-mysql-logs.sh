#!/bin/bash
#
# Script Name: docker-mysql-logs.sh
# Description: Displays Docker MySQL container logs in real-time.
#              This script follows the MySQL container logs and displays them
#              in the terminal. Useful for debugging and monitoring.
#
#              Docker MySQL 컨테이너 로그를 실시간으로 표시합니다.
#              이 스크립트는 MySQL 컨테이너 로그를 따라가며 터미널에 표시합니다.
#              디버깅 및 모니터링에 유용합니다.
#
# Usage: ./docker-mysql-logs.sh [options]
#        사용법: ./docker-mysql-logs.sh [옵션]
#
# Options / 옵션:
#   --tail N    Show last N lines (default: all)
#               마지막 N줄만 표시 (기본값: 전체)
#   --no-follow Show logs without following (exit after display)
#               실시간 추적 없이 로그 표시 (표시 후 종료)
#
# Prerequisites / 사전 요구사항:
#   - MySQL container must be running
#     MySQL 컨테이너가 실행 중이어야 함
#
# Exit Codes / 종료 코드:
#   0 - Success / 성공
#   130 - Interrupted by user (Ctrl+C)
#         사용자 중단 (Ctrl+C)
#
# Examples / 예제:
#   # View all logs with real-time following / 실시간으로 전체 로그 보기
#   ./docker-mysql-logs.sh
#
#   # View last 50 lines / 마지막 50줄 보기
#   ./docker-mysql-logs.sh --tail 50
#
#   # View logs without following / 실시간 추적 없이 로그 보기
#   ./docker-mysql-logs.sh --no-follow
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
readonly GREEN='\033[0;32m'    # Success messages / 성공 메시지
readonly NC='\033[0m'          # No Color (reset) / 색상 없음 (리셋)

# ============================================================================
# Main Script / 메인 스크립트
# ============================================================================

echo -e "${GREEN}Viewing Docker MySQL logs...${NC}"
echo -e "${GREEN}Docker MySQL 로그 보기...${NC}"
echo
echo "Press Ctrl+C to exit / Ctrl+C를 눌러 종료"
echo

# Get script directory and project root
# 스크립트 디렉토리와 프로젝트 루트 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# View logs with follow by default / 기본적으로 실시간 추적하며 로그 보기
# Users can add --tail or --no-follow options to docker compose logs command
# 사용자는 docker compose logs 명령에 --tail 또는 --no-follow 옵션을 추가할 수 있습니다
docker compose logs -f mysql
