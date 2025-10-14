#!/bin/bash
# Stop Docker MySQL for go-utils development
# go-utils 개발을 위한 Docker MySQL 중지

set -e

# Colors for output / 출력용 색상
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Stopping Docker MySQL for go-utils...${NC}"
echo -e "${YELLOW}go-utils를 위한 Docker MySQL 중지 중...${NC}"
echo

# Check if Docker is installed / Docker 설치 확인
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed${NC}"
    echo -e "${RED}오류: Docker가 설치되어 있지 않습니다${NC}"
    exit 1
fi

# Get script directory / 스크립트 디렉토리 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

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
