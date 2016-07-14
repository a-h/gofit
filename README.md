# gofit

Retrieves health data from Google Fit.

# What does it look like?

![1](images/app1.png)

# How do I use it?

Visit the Google Developer console at https://console.cloud.google.com/apis/credentials and setup access to the 
Google Fitness API by creating credentials.

For Fitness, you will need an "OAuth Client id" because we're requesting access to data in a customer's account.

![1](images/screen1.png)

![2](images/screen2.png)

![3](images/screen3.png)

Download the credentials to client_id.json

![4](images/screen4.png)

You can now run the Website by typing: `go run main.go`

# Exploring the API

Visit the https://developers.google.com/oauthplayground/ and generate a key to try out the API in the browser.

Alternatively, once you've created the key, you can use CURL:

```
# list data sources
curl --header "Authorization: Bearer ya29.sdjasdhhjasdghjasgdhjdr5dsadaksdghajs" --verbose -X GET \
  --header "Content-Type: application/json;encoding=utf-8" \
  "https://www.googleapis.com/fitness/v1/users/me/dataSources"

# Get data source
curl --header "Authorization: Bearer ya29.sdjasdhhjasdghjasgdhjdr5dsadaksdghajs" --verbose -X GET \
  --header "Content-Type: application/json;encoding=utf-8" \
  "https://www.googleapis.com/fitness/v1/users/me/dataSources/derived%3Acom.google.step_count.delta%3Acom.google.android.gms%3Aestimated_steps"

curl --header "Authorization: Bearer ya29.sdjasdhhjasdghjasgdhjdr5dsadaksdghajs" --verbose -X GET \
  --header "Content-Type: application/json;encoding=utf-8" \
  "https://www.googleapis.com/fitness/v1/users/userId/dataSources/derived%3Acom.google.step_count.delta%3Acom.google.android.gms%3Aestimated_steps/datasets/"

# Aggregate:
curl --header "Authorization: Bearer ya29.sdjasdhhjasdghjasgdhjdr5dsadaksdghajs" \
  --verbose -X POST \
  --header "Content-Type: application/json;encoding=utf-8" \
  "https://www.googleapis.com/fitness/v1/users/me/dataset:aggregate" \
  -d '{ "startTimeMillis": 1464735600000, "endTimeMillis": 1467919580307, "aggregateBy": [ { "dataTypeName": "com.google.step_count.delta", "dataSourceId": "derived:com.google.step_count.delta:com.google.android.gms:estimated_steps" }  ], "bucketByActivityType": { "minDurationMillis": 1000, "activityDataSourceId": 7  } }'

# List Sessions
curl --header "Authorization: Bearer ya29.sdjasdhhjasdghjasgdhjdr5dsadaksdghajs" --verbose -X GET \
  --header "Content-Type: application/json;encoding=utf-8" \
  "https://www.googleapis.com/fitness/v1/users/me/sessions?startTime=2016-07-01T00:01:00.00Z&endTime=2016-07-07T23:59:00.00Z"
```

Note: You won't have a "session" unless you've manually clicked "add" in the Google fit application.
