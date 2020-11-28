### go module
 - go.mod 
파일로 패키지 관리

 - go.sum
설치된 모듈의 해시 값을 저장해두고, 매 go 커맨드가 실행되기 전에 설치 되어있는 모듈의 해시 값과 go.sum에 저장된 해시 값을 비교하여 설치된 모듈의 유효성을 검증합니다.

> `go.mod`, `go.sum` 파일 모두 깃저장소에 포함시켜 커밋한다.
go.sum 파일을 체크인하면 다른 사람들이 당신과 같은 모듈을 사용하고 있는지 확인하는 데 도움이됩니다

## go module command
- go mod init
모듈을 생성한다.
- go get 패키지@버전
버전을 지정하여 모듈을 추가, 예) github.com/gorilla/mux@latest
- go mod tiny -v
불필요한 종속성을 제거, 소스를 확인하여 import 되지 않은 패키지는 제거하고, import 되어있지만 go.mod 파일에 존재하지 않다면 추가한다.
`-v` 옵션은 더 자세한 내용을 볼 수 있다.
- go mod verify
로컬에 설치된 모듈의 해시 값과 `go.sum`을 비교하여 모듈의 유효성을 검증합니다.
- go mod vendor
vendor 디렉토리를 생성