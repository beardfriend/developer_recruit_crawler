# developer_recruit_crawler

URL : http://13.125.114.252:4000/?position=backend&career=1

position = frontend, backend
career = 0(신입) , 1 , 2

<img width="1659" alt="image" src="https://user-images.githubusercontent.com/97140962/208352203-fb74b4bb-d474-4989-8645-bb2e6071fc31.png">


크롤링 1시간에 한 번씩 몽고DB에 저장



![제목 없는 다이어그램](https://user-images.githubusercontent.com/97140962/208352682-fe95d77c-0107-4941-b748-f383afe2af23.jpg)


[Provider](https://github.com/beardfriend/developer_recruit_crawler/blob/main/provider/provider.go)

GetRecuriments interface에 맞게 jumpit, saramin, ... 크롤러 제작

[Engine](https://github.com/beardfriend/developer_recruit_crawler/blob/main/engine/recruitment.go#L40)
각 공급원들을 고루틴을 사용하여 처리


[Model](https://github.com/beardfriend/developer_recruit_crawler/blob/main/model/recruitment.go)
여러 공급원에서 수집한 정보를 모델에 맞게끔 가공

[Template](https://github.com/beardfriend/developer_recruit_crawler/blob/main/templates/index.html)
 템플릿 엔진을 이용하여 렌더링


