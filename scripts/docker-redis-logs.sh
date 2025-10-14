#!/bin/bash
# View Docker Redis logs for go-utils development
# go-utils 개발을 위한 Docker Redis 로그 확인

set -e

# Colors for output / 출력용 색상
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if Docker is installed / Docker 설치 확인
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed${NC}"
    echo -e "${RED}오류: Docker가 설치되어 있지 않습니다${NC}"
    exit 1
fi

# Check if container is running / 컨테이너 실행 확인
if ! docker ps --filter "name=go-utils-redis" --format "{{.Names}}" | grep -q "go-utils-redis"; then
    echo -e "${RED}Error: Docker Redis is not running${NC}"
    echo -e "${RED}오류: Docker Redis가 실행 중이지 않습니다${NC}"
    echo
    echo "Start Redis with / Redis 시작:"
    echo "  ./scripts/docker-redis-start.sh"
    exit 1
fi

echo -e "${GREEN}Viewing Docker Redis logs...${NC}"
echo -e "${GREEN}Docker Redis 로그 확인 중...${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop / Ctrl+C를 눌러 중지${NC}"
echo

# Follow logs / 로그 팔로우
docker logs -f go-utils-redis
