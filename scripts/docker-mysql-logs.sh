#!/bin/bash
# View Docker MySQL logs
# Docker MySQL 로그 보기

set -e

# Colors for output / 출력용 색상
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "${GREEN}Viewing Docker MySQL logs...${NC}"
echo -e "${GREEN}Docker MySQL 로그 보기...${NC}"
echo
echo "Press Ctrl+C to exit / Ctrl+C를 눌러 종료"
echo

# Get script directory / 스크립트 디렉토리 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# View logs with follow / 로그 보기 (실시간)
docker-compose logs -f mysql
