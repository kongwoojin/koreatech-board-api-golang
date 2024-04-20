## koreatech-board-api
비공식 한국기술교육대학교 공지사항, 게시판 클라이언트 REST API

## 사용된 기술 스택
* [Go](https://go.dev/)
* [EdgeDB](https://www.edgedb.com/)

## 사용된 라이브러리
* [echo](https://echo.labstack.com/)
* [edgedb-go](github.com/edgedb/edgedb-go)

## Swagger
* [Koreatech Board API](https://api.koreatech.kongjak.com/swagger/index.html)

## 설치
### Docker
1. Docker Hub에서 이미지를 다운로드합니다.
``` console
docker pull kongjak/koreatech-board-api:latest
```
2. Docker 컨테이너를 실행합니다.
``` console
docker run --env=EDGEDB_HOST=HOST --env=EDGEDB_PORT=PORT --env=EDGEDB_DBNAME=DBNAME --env=EDGEDB_USER=USER --env=EDGEDB_PASSWD=PASSWD -p 1323:1323 -d kongjak/koreatech-board-api:latest
```

### Binary
1. [GitHub Releases](https://github.com/kongwoojin/koreatech-board-api/releases)에서 적합한 최신 릴리즈를 다운로드합니다.
2. DB 정보를 .env 파일에 작성합니다.
```
EDGEDB_HOST=HOST
EDGEDB_PORT=PORT
EDGEDB_DBNAME=DBNAME
EDGEDB_USER=USER
EDGEDB_PASSWD=PASSWD
```
3. 다운로드 받은 바이너리 파일을 실행힙니다.
``` console
./koreatech-board-api
```

## 지원하는 게시판
* [한국기술교육대학교](https://koreatech.ac.kr/)
* [한국기술교육대학교 생활관](https://dorm.koreatech.ac.kr/)
* [한국기술교육대학교 컴퓨터공학부](https://www.koreatech.ac.kr/cse/)
* [한국기술교육대학교 기계공학부](https://www.koreatech.ac.kr/me/)
* [한국기술교육대학교 메카트로닉스공학부](https://www.koreatech.ac.kr/mecha/)
* [한국기술교육대학교 전기전자통신공학부](https://www.koreatech.ac.kr/ite/)
* 한국기술교육대학교 디자인·건축공학부
  * [디자인공학전공](https://www.koreatech.ac.kr/ide/)
  * [건축공학전공](https://www.koreatech.ac.kr/arch/)
* 한국기술교육대학교 에너지신소재화학공학부
  * [에너지신소재공학전공](https://www.koreatech.ac.kr/mse/)
  * ~~[화학생명공학전공](https://www.koreatech.ac.kr/ace/)~~[^1]
* [한국기술교육대학교 산업경영학부](https://www.koreatech.ac.kr/sim/)

## 관련 프로젝트
* [koreatech-board-android](https://github.com/kongwoojin/koreatech-board-android)
* [koreatech-board-api](https://github.com/kongwoojin/koreatech-board-api)
* [koreatech-board-crawler](https://github.com/kongwoojin/koreatech-board-crawler)

[^1]: 화학생명공학전공 게시판은 학부생만 접근 가능.
