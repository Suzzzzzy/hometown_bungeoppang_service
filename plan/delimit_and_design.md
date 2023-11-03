# 주 기능 범위 정리 및 설계
- 작은 덩어리를 진짜 모양으로 만들어 나가는 과정
- 데이터가 어디서 어떻게 만들어지고, 어디에 저장되고, 어떻게 처리되는지.. 데이터의 흐름에 따라 설계
- 작은 한 기능이 작동되는 것 까지 구현
- 작은 덩어리를 진짜 모양으로 만들어 나가는 과정
- 설계서를 보고 바로 코드로 만들어 기능이 작동할 수 있게 구현할만큼 상세한 설계 필요

# 우선순위가 높은 주 기능
1. 로그인
2. 중고거래 물건 피드 작성
3. 중고거래 피드 상세보기
4. 중고거래 피드 목록 조회
5. 로그아웃
6. 회원 탈퇴

## 로그인
- 카카오 로그인
- 앱을 다운받아 처음 실행했을 때 카카오 로그인 버튼이 있다

### API
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

### DB
토큰에서 id, email 정보를 리턴 받아 DB에 저장한다
- `user_id`: Long/ 카카오 회원 번호
- `email`: Varchar/ 카카오 가입 이메일