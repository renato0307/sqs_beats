# Elastic beats - SQS Output
The objective of this project is to build a new output to SQS to be used with the Elastic beats.

## How to Build
Before starting you must have Go Lang installed in your system, download it here https://golang.org/doc/install

To avoid issues, use the same version as Elastic are using for Beats Development.
You can check the version here: https://www.elastic.co/guide/en/beats/devguide/current/beats-contributing.html

Go to the beat folder you want to build
```bash
cd beats/filebeat
go build
```

If you have issues with dependencies, execute the go get command, example:
```bash
go get github.com/elastic/beats/filebeat/cmd
```

By default, go build will build using the head of https://github.com/elastic/beats.
If you want to build on a specific commit/release:

1. Clone the beats repo into your gopath
```bash
mkdir -p ${GOPATH}/src/github.com/elastic
git clone https://github.com/elastic/beats ${GOPATH}/src/github.com/elastic/beats
```
2. Checkout the commit you want to build on
3. Download the dependencies for the beats repo using:
```bash
cd ${GOPATH}/src/github.com/elastic/beats
go mod vendor
```


## How to Configure and install
1. Download the beat from the elastic website, https://www.elastic.co/beats
2. Extract the content
3. Replace the binary with the builded one
4. Change the configuration file (filebeat.yml, metricbeat.yml...)

## Sample Configuration
```yaml
output.sqs:
  access_key_id: YOUR_ACCESS_KEY
  access_secret_key: YOUR_SECRET_KEY
  region: eu-west-1
  queue_url: FULL_URL_OF_YOUR_QUEUE
```

The user must have the following permissions to be able to write send messages to SQS:
* SQS:SendMessage
* SQS:GetQueueUrl
