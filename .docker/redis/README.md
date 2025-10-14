# Docker Redis Configuration / Docker Redis 설정

This directory contains Redis configuration files for go-utils development.

이 디렉토리에는 go-utils 개발을 위한 Redis 설정 파일이 포함되어 있습니다.

## Files / 파일

- **redis.conf** - Redis server configuration file / Redis 서버 설정 파일

## Usage / 사용법

### Start Redis / Redis 시작

```bash
./scripts/docker-redis-start.sh
```

### Stop Redis / Redis 중지

```bash
./scripts/docker-redis-stop.sh
```

### View Logs / 로그 확인

```bash
./scripts/docker-redis-logs.sh
```

### Connect to Redis CLI / Redis CLI 연결

```bash
./scripts/docker-redis-cli.sh
```

## Connection Details / 연결 정보

- **Host**: localhost
- **Port**: 6379
- **Password**: (none)
- **Databases**: 0-15

## Configuration / 설정

The `redis.conf` file is mounted at `/usr/local/etc/redis/redis.conf` in the Docker container.

`redis.conf` 파일은 Docker 컨테이너의 `/usr/local/etc/redis/redis.conf`에 마운트됩니다.

### Key Settings / 주요 설정

- **Persistence / 영속성**: RDB snapshots + AOF (Append Only File)
- **Snapshotting / 스냅샷**: Save after 900s if 1 key changed, 300s if 10 keys, 60s if 10000 keys
- **AOF**: Enabled with `everysec` fsync policy
- **Max Memory / 최대 메모리**: noeviction policy (no memory limit)
- **Databases / 데이터베이스**: 16 (0-15)

## Data Persistence / 데이터 영속성

Redis data is stored in a Docker volume named `go-utils-redis-data`.

Redis 데이터는 `go-utils-redis-data`라는 Docker 볼륨에 저장됩니다.

To completely remove the data:

데이터를 완전히 제거하려면:

```bash
./scripts/docker-redis-stop.sh
# Then choose 'y' when prompted to remove the volume
# 볼륨 제거 여부를 묻는 메시지에 'y' 선택
```

Or manually:

또는 수동으로:

```bash
docker volume rm go-utils-redis-data
```

## Testing / 테스트

To test the Redis connection:

Redis 연결을 테스트하려면:

```bash
./scripts/docker-redis-cli.sh
```

Then in Redis CLI:

Redis CLI에서:

```redis
PING
# Should return: PONG / PONG을 반환해야 함

SET test "Hello, Redis!"
GET test
# Should return: "Hello, Redis!"

DEL test
```

## Troubleshooting / 문제 해결

### Container won't start / 컨테이너가 시작되지 않음

Check if port 6379 is already in use:

포트 6379가 이미 사용 중인지 확인:

```bash
lsof -i :6379
```

### Can't connect to Redis / Redis에 연결할 수 없음

Check if the container is running:

컨테이너가 실행 중인지 확인:

```bash
docker ps | grep redis
```

Check the logs:

로그 확인:

```bash
./scripts/docker-redis-logs.sh
```

### Reset Redis data / Redis 데이터 초기화

Stop Redis and remove the volume:

Redis를 중지하고 볼륨 제거:

```bash
./scripts/docker-redis-stop.sh  # Choose 'y' to remove volume
./scripts/docker-redis-start.sh
```
