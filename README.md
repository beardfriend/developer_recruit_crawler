# Developer Recruitments Crawler

<div style="flex">
<img src="https://img.shields.io/badge/Go-gray?style=flat&logo=Go&logoColor=00ADD8"/>
<img src="https://img.shields.io/badge/gin-gray?style=flat"/>
</div>

<br/>
사람인, 프로그래머스, 점핏, 원티드의 채용공고를 수집하는 크롤러
<br/>
<img width="640" alt="image" src="https://github.com/beardfriend/developer_recruit_crawler/assets/97140962/08f2f62a-c0b0-429a-8b9c-a00fa72283d2">


## 동기

채용공고를 한 번에 볼 수 없을까?


## 서버 구조

![제목 없는 다이어그램](https://user-images.githubusercontent.com/97140962/208352682-fe95d77c-0107-4941-b748-f383afe2af23.jpg)

## 구현

- 고루틴을 사용하여 채용공고 비동기 수집
- 채용공고 크롤러 추상화
- CI / CD (github action)

<br/>

# 사용법

🙏

## 로컬

```bash
go run main.go

http://localhost:4000?position=backend&career=1
```

### 파라메터

position = frontend, backend, dataEngineer 
career= 1 ~ 7


<br/>



감사합니다
