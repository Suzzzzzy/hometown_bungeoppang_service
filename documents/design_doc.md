# 주 기능 범위 정리 및 설계
- 작은 덩어리를 진짜 모양으로 만들어 나가는 과정
- 데이터가 어디서 어떻게 만들어지고, 어디에 저장되고, 어떻게 처리되는지.. 데이터의 흐름에 따라 설계
- 작은 한 기능이 작동되는 것 까지 구현
- 작은 덩어리를 진짜 모양으로 만들어 나가는 과정
- 설계서를 보고 바로 코드로 만들어 기능이 작동할 수 있게 구현할만큼 상세한 설계 필요

# 우선순위가 높은 주 기능
1. 로그인
2. 회원 정보 저장(지역, 닉네임)
3. 중고거래 물건 피드 작성
4. 중고거래 피드 상세보기
5. 중고거래 피드 목록 조회
6. 로그아웃
7. 회원 탈퇴

# 로그인
- 카카오 로그인
- 앱을 다운받아 처음 실행했을 때 카카오 로그인 버튼이 있다

## API
1. [Get] https://kauth.kakao.com/oauth/authorize
- 인가 코드 받기
- **Parameter**
  - `client_id`: String/ 앱 키
  - `redirect_uri`: String/ 인가코드를 전달 받을 서비스 서버의 URI
  - `response_type`: String/ code로 고정
- **Response**: redirect_uri에 GET 요청으로 리다이렉트

2. [POST] https://kauth.kakao.com/oauth/token	
- **Header**
  - `Content_type`: 	Content-type: application/x-www-form-urlencoded;charset=utf-8
    요청 데이터 타입
- **Body**
  - `grant_type`: String/ authorization_code 고정
  - `client_id`: String/ 앱 REST API키
  - `redirect_uri`: String/ 인가 코드가 리다이렉트된 URI
  - `code`: String/ 인가 코드 받기 요청으로 얻은 인가코드
- **Response**
  - `token_type`: String/ 토큰 타입, bearer 고정
  - `access_token`: String/ 사용자 엑세스 토큰 값
  - `expires_in`: Integer/ 엑세스 토큰과 ID 토큰의 만료 시간
  - `refresh_token`: String/ 사용자 리프레시 토큰
  - `refresh_token_expires_in`: Integer/ 리프레시 토큰 만료 시간(초)

3. [Get] https://kapi.kakao.com/v2/user/me
토큰으로 kakao id, email 정보만 
- **Header**
  - `Authorization`: Bearer ${ACCESS_TOKEN}
  - `Content_type': Content-type: application/x-www-form-urlencoded;charset=utf-8
    요청 데이터 타입
- **Parameter**
  - `property_keys`: ["kakao_account.email"]
- **Response**
  - `id`: Long/ 회원번호
  - `email`: String/ 이메일

## DB
토큰에서 id, email 정보를 리턴 받아 DB에 저장한다
- `user_id`: Long/ 카카오 회원 번호
- `email`: Varchar/ 카카오 가입 이메일


# 회원 정보 저장
- 카카오 로그인 후 유저에게 추가 정보를 받아 저장
- 닉네임 설정, 거래 지역 설정

## 닉네임
- string 타입으로 저장
- 한글, 숫자로 10자 이내
- 닉네임 중복 체크 -> '다음'버튼 누를 때 사용 불가능한(중복된) 닉네임이라면 '이미 사용중인 닉네임 입니다' 알림

### API
- 닉네임 작성 칸 밑에 '다음' 버튼에 해당
- token 필요
- 유저 정보는 token 에서 파싱
- **Request**: `nickname` string
- 체크 항목: 한글과 숫자 포함 10자 이내, 중복 확인
- user_id로 데이터 조회하여 nickname 칼럼 업데이트

## 거래 지역 설정
- 사용자 입력값으로 설정
- 시-구-동 세개의 칸으로 입력받음
- **Request**: `area` string
- user 테이블 area 칼럼에 '시 구 동' 순서로 한 칸 띄어쓰기 하여 저장
- user_id로 데이터 조회하여 area 칼럼 업데이트


# 중고거래 피드 작성

<img width="250" alt="image" src="https://github.com/Suzzzzzy/transaction-app-service/assets/97580836/7015a75b-378d-4c10-a0f9-30fca59f5a4a">

- 중고거래 피드 작성 저장
- 사진은 한 장 저장

## [POST] app/feed
- **Header**
  - token: 유저정보 파싱을 위함
- **Body**
  - `title`: varchar/ 제목
  - `price`: float/ 가격
  - `description`: varchar/ 설명
  - `image`: file/ 사진 s3 url
- **Response**
  - title, price, description, image(s3 url), area(유저의 설정된 지역)

## 이미지 업로드
- AWS S3 서비스 이용
- s3 uploader 작성 필요