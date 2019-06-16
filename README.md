# Rest-API-Monitoring---Introduction



#### Load test description
```
Duration = 5s
Rate = 5000
use MySQL
Query : select * from users where id = 2
Users Table row = 10000 row 

```

#### Load Test without implement circuit breaker
![Before CB](/Doc/beforeCB.jpg)
```
Success : 55.31% (13827/25000)
Tripped : 10971
No error code : 202 
```

#### Load Test after implement circuit breaker
![Added CB](/Doc/addedCB.jpg)
```
Success : 97.60% (24399/25000)
Tripped : 601
No error code : - 
```

Prometheus Monitoring with Grafana
![Monitoring](/Doc/monitoring.jpg)
