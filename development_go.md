### files

#### go.mod

#### go.sum




### go command

코드를 검색하여 사용하지 않는 package는 지운다.
```
go mod tidy
```

외부 모듈을 gopath가 아닌 내부 디렉토리에 vendor에 복사하여 이용한다.
(모듈이 참조한 외부 모듈을 모두 %projectdir%/vendor 디렉토리에 복사한다.)
```
go mod vendor
```

모듈의 캐시를 지운다.
```
go clean -modcache
```