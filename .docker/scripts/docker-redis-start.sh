#!/bin/bash
# Start Docker Redis for go-utils development
# go-utils 개발을 위한 Docker Redis 시작

set -e

# Colors for output / 출력용 색상
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting Docker Redis for go-utils...${NC}"
echo -e "${GREEN}go-utils를 위한 Docker Redis 시작 중...${NC}"
echo

# Check if Docker is installed / Docker 설치 확인
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed${NC}"
    echo -e "${RED}오류: Docker가 설치되어 있지 않습니다${NC}"
    echo
    echo "Please install Docker Desktop from:"
    echo "Docker Desktop을 다음에서 설치하세요:"
    echo "https://www.docker.com/products/docker-desktop"
    exit 1
fi

# Check if Docker is running / Docker 실행 확인
if ! docker info &> /dev/null; then
    echo -e "${RED}Error: Docker is not running${NC}"
    echo -e "${RED}오류: Docker가 실행 중이지 않습니다${NC}"
    echo
    echo "Please start Docker Desktop and try again"
    echo "Docker Desktop을 시작한 후 다시 시도하세요"
    exit 1
fi

# Get script directory / 스크립트 디렉토리 가져오기
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

# Change to project root / 프로젝트 루트로 이동
cd "${PROJECT_ROOT}"

# Check if container is already running / 컨테이너가 이미 실행 중인지 확인
if docker ps --filter "name=go-utils-redis" --format "{{.Names}}" | grep -q "go-utils-redis"; then
    echo -e "${YELLOW}Docker Redis is already running${NC}"
    echo -e "${YELLOW}Docker Redis가 이미 실행 중입니다${NC}"
    echo
    docker ps --filter "name=go-utils-redis"
    exit 0
fi

# Start Docker Compose / Docker Compose 시작
echo -e "${GREEN}Starting Docker Compose (Redis only)...${NC}"
echo -e "${GREEN}Docker Compose 시작 중 (Redis만)...${NC}"
docker compose up -d redis

echo
echo -e "${GREEN}Waiting for Redis to be ready...${NC}"
echo -e "${GREEN}Redis 준비 대기 중...${NC}"

# Wait for Redis to be ready / Redis 준비 대기
max_attempts=30
attempt=0
while [ $attempt -lt $max_attempts ]; do
    if docker exec go-utils-redis redis-cli ping &> /dev/null; then
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
        echo "To view logs / 로그 확인:"
        echo "  ./scripts/docker-redis-logs.sh"
        echo
        echo "To connect to Redis CLI / Redis CLI 연결:"
        echo "  ./scripts/docker-redis-cli.sh"
        echo
        exit 0
    fi
    attempt=$((attempt + 1))
    echo -n "."
    sleep 1
done

echo
echo -e "${RED}Error: Redis failed to become ready${NC}"
echo -e "${RED}오류: Redis가 준비되지 않았습니다${NC}"
echo
echo "Check logs with / 로그 확인:"
echo "  ./scripts/docker-redis-logs.sh"
exit 1
