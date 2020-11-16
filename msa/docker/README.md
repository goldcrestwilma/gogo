## Docker 컨테이너 실행
```
$ docker run --rm hello-world
```

`--rm` 플래그는 Docker 엔진에게 종료 시 컨테이너를 제거하고 볼륨과 같은 사용이 끝난 리소스를 삭제하도록 지시한다.

## 컨테이너 시작한 후 내부에 쉘 만들기
```
$ docker run -it --rm alpine:latest sh
```

Alpine은 Linux의 경량화된 버전이며 Go Application을 실행하는 데 적합.
`-it` 플래그는 대화형 터미널(interactive terminal)을 의미
`sh`는 시작할 때 컨테이너에서 실행하고자 하는 명령의 이름

## 명령어
### 컨테이너 목록
```
$ docker ps -a
```

`-a` 플래그 명령을 추가하면 이전에 시작시켰던 컨테이너 조회 가능

### 컨테이너 재시작
```
$ docker start -it [컨테이너 ID] sh
```

### 모든 컨테이너 제거

```
$ docker rm -v $(docker ps -a -q)
```

`-a` 플래그는 중지된 것을 포함한 모든 컨테이너를 나열
`-q` 플래그는 전체 세부 사항이 아닌 컨테이너 ID의 목록만 리턴

## 볼륨 마운트

```
$ docker run -it -v $(pwd):/host alpine:latest /bin/sh
```

`host` 폴더로 이동해 보면 docker run 명령을 실행한 폴더와 동일한 폴더로 접근했음을 알 수 있음.

## Docker Port
```
$ docker run -it --rm -v $(pwd):/src -p 8080:8080 -w /src golang:alpine /bin/sh
```

`-w` 플래그는 작업 디렉토리를 설정하는 것으로 컨테이너에서 실행하는 모든 명령이 이 폴더 내에서 실행된다는 것을 의미

## Docker용 애플리케이션 코드 빌드하기
```
$ CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cog -ldflags '-s' -o server
```

vi Dockerfile
```
FROM scratch
MAINTAINER aga6023496@gmail.com

EXPOSE 8080

COPY ./server ./server

ENTRYPOINT ["./server"]
```

## Dockerfile로 이미지 빌드하기
```
$ docker build -t testserver .
$ docker run --rm -p 8080:8080 testserver
```