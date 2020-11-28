
## 단위테스트
- 첫 번째 법칙: 실제 코드를 작성하기 전에 실패하는 단위 테스트를 먼저 작성한다.
- 두 번째 법칙: 컴파일에 실패하지 않으면서 실패하기에 충분한 수준으로만 단위테스트를 작성한다.
- 세 번째 법칙: 현재 실패하는 테스트를 통과할 수 있는 수준으로만 실제 코드를 작성한다.

> Go에서 마이크로서비스를 테스트하는 가장 효과적인 방법 중 하나는 HTTP 인터페이스를 통해 모든 테스트를 실행하려고 하는 함정에 빠지지 않는 것이다. 이를 위해 물리적인 웹 서버를 만들지 않아도 핸들러를 테스트할 수 있는 패턴을 개발할 필요가 있다.

## net/http/httptest 패키지
`NewRequest`와 `NewResponse`라는 두 개의 메소드가 있다.
의존성이 있는 코드를 실행하지 않고도 코드의 동작을 테스트할 수 있도록 하기 위해 종종 모의객체(Mock, `테스트를 위해 만든 동일한 인터페이스를 가진 가짜 객체`)나 스파이(`실제 인스턴스를 테스트용 목으로 사용`)로 의존성을 대체한다.
`NewRequest`와 `NewResponse` 메소드의 역할이 정확이 이것이다. 이들은 의존성이 있는 객체인 `http.Request`와 `http.ResponseWriter`의 모의 객체 버전을 생성한다.

```
func NewRequest(method, target string, body io.Reader) *http.Reqeust
```

## Errorf
```
func (c *T) Errorf(format string, args ...interface{})
```
Errorf 함수는 형식 문자열과 가변 매개 변수 목록을 매개 변수로 받는다. 내부적으로 이것은 Logf 메소드를 호출한 후에 Fail을 호출한다.

## go test -v -race ./...
`-v` 플래그는 상세 출력 형식, 테스트가 성공하더라도 모든 텍스트를 보여준다.
`-race` 플래그는 동시성 문제가 있는 버그를 발견하는 Go의 레이스 탐지기를 활성화.
마지막 매개 변수 `./...` 를 사용하면 현재 폴더와 하위 폴더의 모든 테스트를 실행할 수 있으므로 테스트할 패키지 또는 파일 목록을 수동으로 작성하지 않아도 된다.

## Testify
https://github.com/stretchr/testify.git
모의 객체 프레임워크

## 기타 패키지
### labix.org/v2/mgo 
```
❯ go get labix.org/v2/mgo
go: missing Bazaar command. See https://golang.org/s/gogetcmd
package labix.org/v2/mgo: exec: "bzr": executable file not found in $PATH
```

`Bazaar`: git, svn 과 같은 버전 제어 시스템
```
brew install bazaar
```
설치 후 다시 패키지 다운

### github.com/DATA-DOG/godog/cmd/godog
```
❯ go get github.com/cucumber/godog/cmd/godog
package github.com/cucumber/messages-go/v10: cannot find package "github.com/cucumber/messages-go/v10" in any of:
        /usr/local/go/src/github.com/cucumber/messages-go/v10 (from $GOROOT)
        /Users/falcon/go/src/github.com/cucumber/messages-go/v10 (from $GOPATH)
package github.com/cucumber/gherkin-go/v11: cannot find package "github.com/cucumber/gherkin-go/v11" in any of:
        /usr/local/go/src/github.com/cucumber/gherkin-go/v11 (from $GOROOT)
        /Users/falcon/go/src/github.com/cucumber/gherkin-go/v11 (from $GOPATH)
```

go mod init 실행 후 다시 명령어 실행