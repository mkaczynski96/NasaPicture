# GogoSpace
TWljaGFsIEdvZ29BcHBzIE5BU0E=


<b>HOW TO RUN:</b>

```clone project```
``` cd to project```
```docker build . -t gogo-app```  
```docker run -p 8080:8080 gogo-app```

<b>AVAILABLE ENDPOINTS</b>

>GET /pictures?start_date={start_date}&end_date={end_date}

start_date and end_date must contain date in YYYY-MM-DD format. Start_date must be earlier than end_date.

<b>EXAMPLE REQUEST:</b>

curl -X GET 'localhost:8080/pictures?start_date=2020-12-13&end_date=2020-12-14'
