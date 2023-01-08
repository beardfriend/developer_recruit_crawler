# developer_recruit_crawler

요청하는 즉시 크롤러가 돌기 때문에 느릴 수 있습니다.(nginx 20분 캐싱)

URL : http://13.125.48.238:4000/?position=backend&career=1

position = frontend, backend  
career =  1 , 2 , .... 7

<img width="1659" alt="image" src="https://user-images.githubusercontent.com/97140962/208352203-fb74b4bb-d474-4989-8645-bb2e6071fc31.png">


![제목 없는 다이어그램](https://user-images.githubusercontent.com/97140962/208352682-fe95d77c-0107-4941-b748-f383afe2af23.jpg)


[Provider](https://github.com/beardfriend/developer_recruit_crawler/blob/main/provider/provider.go)

GetRecuriments 함수를 추상화 하였고, 추상화된 함수에 맞게 jumpit, saramin, wanted, programmers 크롤러 제작

[Engine](https://github.com/beardfriend/developer_recruit_crawler/blob/main/engine/recruitment.go#L40)
각 공급원들을 고루틴을 이용하여 비동기적으로 처리

[Model](https://github.com/beardfriend/developer_recruit_crawler/blob/main/model/recruitment.go)
여러 공급원에서 수집한 정보를 모델에 맞게끔 가공

[Template](https://github.com/beardfriend/developer_recruit_crawler/blob/main/templates/index.html)
 템플릿 엔진을 이용하여 렌더링


데이터베이스 도입과 크론탭에서 미리 크롤링을 하여, 요청 시 즉각적으로 데이터를 받아볼 수 있게끔 설계하려고 했으나,
사용도가 낮고 데이터베이스를 도입할 경우 비용이 추가적으로 발생하기 때문에 도입을 하지 않았음.
