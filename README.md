# go-bigquery-sample

## Run
You should set the following 3 shell environment variable.

- PROJECT=your-project
- DATASET=anyDataset
- TABLE=anyTable

And You can also specify the json file of the service account.

- GOOGLE_APPLICATION_CREDENTIALS

```sh
$ PROJECT=your-project DATASET=anyDataset TABLE=anyTable GOOGLE_APPLICATION_CREDENTIALS=key.json go run go-bigquery-sample.go
```
