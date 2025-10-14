#!/bin/bash
# Start Docker MySQL for go-utils development
# go-utils 개발을 위한 Docker MySQL 시작

set -e

# Colors for output / 출력용 색상
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting Docker MySQL for go-utils...${NC}"
echo -e "${GREEN}go-utils를 위한 Docker MySQL 시작 중...${NC}"
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
if docker ps --filter "name=go-utils-mysql" --format "{{.Names}}" | grep -q "go-utils-mysql"; then
    echo -e "${YELLOW}Docker MySQL is already running${NC}"
    echo -e "${YELLOW}Docker MySQL이 이미 실행 중입니다${NC}"
    echo
    docker ps --filter "name=go-utils-mysql"
    exit 0
fi

# Start Docker Compose / Docker Compose 시작
echo -e "${GREEN}Starting Docker Compose...${NC}"
echo -e "${GREEN}Docker Compose 시작 중...${NC}"
docker-compose up -d

echo
echo -e "${GREEN}Waiting for MySQL to be ready...${NC}"
echo -e "${GREEN}MySQL 준비 대기 중...${NC}"

# Wait for MySQL to be ready / MySQL 준비 대기
max_attempts=30
attempt=0
while [ $attempt -lt $max_attempts ]; do
    if docker exec go-utils-mysql mysqladmin ping -h localhost -u root -prootpassword &> /dev/null; then
        echo
        echo -e "${GREEN}✅ Docker MySQL is ready!${NC}"
        echo -e "${GREEN}✅ Docker MySQL 준비 완료!${NC}"
        echo
        echo "Connection details / 연결 정보:"
        echo "  Host: localhost"
        echo "  Port: 3306"
        echo "  Database: testdb"
        echo "  User: root"
        echo "  Password: rootpassword"
        echo
        echo "To stop MySQL / MySQL 중지:"
        echo "  ./scripts/docker-mysql-stop.sh"
        echo
        exit 0
    fi
    attempt=$((attempt + 1))
    echo -n "."
    sleep 1
done

echo
echo -e "${RED}Error: MySQL failed to become ready${NC}"
echo -e "${RED}오류: MySQL이 준비되지 않았습니다${NC}"
echo
echo "Check logs with / 로그 확인:"
echo "  docker-compose logs mysql"
exit 1
