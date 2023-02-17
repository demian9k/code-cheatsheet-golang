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



### go workspace
여러개의 gomodule 프로젝트를 상위개념인 workspace로 묶어서 module간 import한다.

이전에는 go.mod 파일에서 directive를 replace 하는 방식으로 사용했다.
go mod edit -replace github.com/gowkr=../my_mod

```bash

//해당 경로를 workspace로 만든다. go.work 파일이 만들어진다.
go work init  

//go.work 파일의 .는 현재 디렉토리의 go.mod 파일을 가리키는 것이다.
use (
.
)


디렉토리를 만들고 go module을 만든다.

mkdir ./mygomodule1  
mkdir ./mygomodule2 
mkdir ./mygomodule2/mygomodule2sub1  //nested module

#그리고 
$ go work use ./mygomodule1
$ go work use ./mygomodule2
$ go work use ./mygomodule2/mygomodule2sub1 //nested module을 추가하는 방법

// 입력하면 go.work 파일에 아래처럼 입력되어 workspace 아래 속한 모듈로 인식된다.

use (
.
./mygomodule1
./mygomodule2
./mygomodule2/mygomodule2sub1
)
  
```
