#!/bin/bash
# Stop and cleanup Docker Redis for go-utils development
# go-utils 개발을 위한 Docker Redis 중지 및 정리

set -e

# Colors for output / 출력용 색상
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Stopping Docker Redis for go-utils...${NC}"
echo -e "${GREEN}go-utils를 위한 Docker Redis 중지 중...${NC}"
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

# Check if container is running / 컨테이너 실행 확인
if ! docker ps --filter "name=go-utils-redis" --format "{{.Names}}" | grep -q "go-utils-redis"; then
    echo -e "${YELLOW}Docker Redis is not running${NC}"
    echo -e "${YELLOW}Docker Redis가 실행 중이지 않습니다${NC}"
    echo
    # Still try to remove stopped container / 중지된 컨테이너 제거 시도
    if docker ps -a --filter "name=go-utils-redis" --format "{{.Names}}" | grep -q "go-utils-redis"; then
        echo -e "${YELLOW}Removing stopped container...${NC}"
        echo -e "${YELLOW}중지된 컨테이너 제거 중...${NC}"
        docker compose down redis
    fi
    exit 0
fi

# Stop Docker Compose / Docker Compose 중지
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
    docker volume rm go-utils-redis-data 2>/dev/null || true
    echo -e "${GREEN}✅ Redis data volume removed${NC}"
    echo -e "${GREEN}✅ Redis 데이터 볼륨 제거됨${NC}"
else
    echo -e "${YELLOW}Redis data volume preserved${NC}"
    echo -e "${YELLOW}Redis 데이터 볼륨 보존됨${NC}"
fi
