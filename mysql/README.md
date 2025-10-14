# Docker MySQL for go-utils Development

This directory contains the Docker MySQL configuration for go-utils development and testing.

이 디렉토리는 go-utils 개발 및 테스트를 위한 Docker MySQL 설정을 포함합니다.

## Directory Structure / 디렉토리 구조

```
mysql/
├── README.md          # This file / 이 파일
├── init/              # SQL initialization scripts / SQL 초기화 스크립트
│   └── 01-create-tables.sql   # Create tables and sample data / 테이블 및 샘플 데이터 생성
└── conf/              # MySQL configuration files / MySQL 설정 파일
    └── my.cnf         # Custom MySQL configuration / 사용자 정의 MySQL 설정
```

## Quick Start / 빠른 시작

### Prerequisites / 전제 조건

1. **Install Docker Desktop** / Docker Desktop 설치
   - Download from: https://www.docker.com/products/docker-desktop
   - 다운로드: https://www.docker.com/products/docker-desktop

2. **Start Docker Desktop** / Docker Desktop 시작
   - Make sure Docker is running before proceeding
   - 진행하기 전에 Docker가 실행 중인지 확인하세요

### Start MySQL / MySQL 시작

```bash
# From project root / 프로젝트 루트에서
./scripts/docker-mysql-start.sh

# Or use docker-compose directly / 또는 docker-compose를 직접 사용
docker-compose up -d
```

### Stop MySQL / MySQL 중지

```bash
# From project root / 프로젝트 루트에서
./scripts/docker-mysql-stop.sh

# Or use docker-compose directly / 또는 docker-compose를 직접 사용
docker-compose down
```

### View Logs / 로그 보기

```bash
# From project root / 프로젝트 루트에서
./scripts/docker-mysql-logs.sh

# Or use docker-compose directly / 또는 docker-compose를 직접 사용
docker-compose logs -f mysql
```

## Connection Details / 연결 정보

```yaml
Host: localhost
Port: 3306
Database: testdb
User: root
Password: rootpassword
```

## Database Schema / 데이터베이스 스키마

The MySQL container is initialized with the following schema:

MySQL 컨테이너는 다음 스키마로 초기화됩니다:

### Users Table / Users 테이블

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    age INT NOT NULL,
    city VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_email (email),
    INDEX idx_city (city),
    INDEX idx_age (age)
);
```

The table is pre-populated with 10 sample users for testing.

테이블은 테스트를 위해 10명의 샘플 사용자로 미리 채워집니다.

## Custom Configuration / 사용자 정의 설정

You can customize MySQL configuration by editing `mysql/conf/my.cnf`.

`mysql/conf/my.cnf`를 편집하여 MySQL 설정을 사용자 정의할 수 있습니다.

After making changes, restart MySQL:

변경 후 MySQL을 다시 시작하세요:

```bash
./scripts/docker-mysql-stop.sh
./scripts/docker-mysql-start.sh
```

## Adding Initialization Scripts / 초기화 스크립트 추가

To add more initialization scripts, create `.sql` files in `mysql/init/` directory.
Scripts are executed in alphabetical order.

더 많은 초기화 스크립트를 추가하려면 `mysql/init/` 디렉토리에 `.sql` 파일을 만드세요.
스크립트는 알파벳 순서로 실행됩니다.

Example:

```bash
mysql/init/
├── 01-create-tables.sql    # Executed first / 먼저 실행됨
├── 02-seed-data.sql        # Executed second / 두 번째로 실행됨
└── 03-create-indexes.sql   # Executed third / 세 번째로 실행됨
```

## Persistent Data / 영구 데이터

MySQL data is persisted in a Docker volume named `go-utils-mysql-data`.

MySQL 데이터는 `go-utils-mysql-data`라는 Docker 볼륨에 영구적으로 저장됩니다.

To remove all data and start fresh:

모든 데이터를 제거하고 새로 시작하려면:

```bash
# Stop MySQL / MySQL 중지
./scripts/docker-mysql-stop.sh

# Remove volume / 볼륨 제거
docker volume rm go-utils-mysql-data

# Start MySQL again / MySQL 다시 시작
./scripts/docker-mysql-start.sh
```

## Troubleshooting / 문제 해결

### Port 3306 already in use / 포트 3306이 이미 사용 중

If you have a local MySQL installation running on port 3306:

로컬 MySQL 설치가 포트 3306에서 실행 중인 경우:

```bash
# Stop local MySQL / 로컬 MySQL 중지
brew services stop mysql

# Or change port in docker-compose.yml / 또는 docker-compose.yml에서 포트 변경
ports:
  - "3307:3306"  # Use port 3307 instead / 3307 포트 사용
```

### Container fails to start / 컨테이너 시작 실패

Check Docker logs:

Docker 로그 확인:

```bash
./scripts/docker-mysql-logs.sh
```

### Cannot connect to MySQL / MySQL에 연결할 수 없음

Make sure MySQL is ready:

MySQL이 준비되었는지 확인:

```bash
# Check container status / 컨테이너 상태 확인
docker ps --filter "name=go-utils-mysql"

# Check health status / 헬스 상태 확인
docker inspect go-utils-mysql | grep -A 5 Health
```

## Running Examples / 예제 실행

After starting MySQL, you can run the examples:

MySQL을 시작한 후 예제를 실행할 수 있습니다:

```bash
cd examples/mysql
go run .
```

The examples will automatically detect and use the Docker MySQL container.

예제는 자동으로 Docker MySQL 컨테이너를 감지하고 사용합니다.
