# developer_recruit_crawler

URL : http://43.201.147.22:4000/?position=backend&career=1

position = frontend, backend
career = 0(신입) , 1 , 2

<img width="1659" alt="image" src="https://user-images.githubusercontent.com/97140962/208352203-fb74b4bb-d474-4989-8645-bb2e6071fc31.png">




요청 시 크롤링을 해서 데이터를 가공하기 때문에
속도가 느리다는 단점이 있다.

대안으로 캐싱을 선택했다.

DB에 저장하여 불러오면 모든 유저가 동일한 속도를 낼 수 있지만,
패스..


![제목 없는 다이어그램](https://user-images.githubusercontent.com/97140962/208352682-fe95d77c-0107-4941-b748-f383afe2af23.jpg)


[Provider](https://github.com/beardfriend/developer_recruit_crawler/blob/main/provider/provider.go)

GetRecuriments interface에 맞게 jumpit, saramin, ... 크롤러 제작

[Engine](https://github.com/beardfriend/developer_recruit_crawler/blob/main/engine/recruitment.go#L40)
각 공급원들을 고루틴을 사용하여 처리


[Model](https://github.com/beardfriend/developer_recruit_crawler/blob/main/model/recruitment.go)
여러 공급원에서 수집한 정보를 모델에 맞게끔 가공

[Template](https://github.com/beardfriend/developer_recruit_crawler/blob/main/templates/index.html)
 템플릿 엔진을 이용하여 렌더링


