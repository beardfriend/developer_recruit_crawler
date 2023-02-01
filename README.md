# Developer Recruitments Crawler

<div style="flex">
<img src="https://img.shields.io/badge/GO-gray?style=flat&logo=Go&logoColor=00ADD8"/>
<img src="https://img.shields.io/badge/Gin-white?style=flat"/>
</div>

사람인, 프로그래머스, 점핏, 원티드의 채용공고를 수집하는 크롤러


## 동기

채용공고를 한 번에 볼 수 없을까?


## 서버 구조

![제목 없는 다이어그램](https://user-images.githubusercontent.com/97140962/208352682-fe95d77c-0107-4941-b748-f383afe2af23.jpg)

## 구현

- 고루틴을 사용하여 채용공고 비동기 수집
- 채용공고 크롤러 추상화
- CI / CD (github action)


# 사용법

🙏

## 프로덕션

주소 : http://13.125.48.238:4000/?position=backend&career=1

(요청 시 즉각 크롤러가 돌기에 느림)

<br/>

## 로컬

```bash
go run main.go

http://localhost:4000?position=backend&career=1
```

### 파라메터

position = frontend, backend
career= 1 ~ 7


# 기타

## 느림

소규모의 서버로 저비용 운영을 위해  
주기적인 크롤링 이후 데이터 저장 생략  

20분 간격으로 Nginx 캐싱


감사합니다
