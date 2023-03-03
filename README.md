# splunk-cf-logdrain

![splunk-cf-logdrain excalidraw](https://user-images.githubusercontent.com/14123216/222737649-0f42e6f3-f7bc-4eb3-935e-1c98c59824c7.svg)

User deployable service which implements a pipeline consisting of a small Go app and a fluent-bit sidecar process. It presents a CF compatible logdrainer endpoint which accepts RFC5424 messages, forwards them to the fluent-bit sidecard process, which in turn forwards the log messages to splunk, done.

## Usage

* Deploy the logdrainer and create a user defined splunk service in your Cloud foundry Org
* Bind apps to the splunk service to activate forwarding 

## Splunk set up

* Configure a HTTP Event Collector (HEC) Data input (Settings -> Data inputs)
* Make sure to uncheck `Enable indexer acknowledgement`
* Copy the token for use in config below `SPLUNK_TOKEN`

## Logdrain set up

* Configure environment variables

| Environment     | Description                        | Recommended value                     |
|-----------------|------------------------------------|---------------------------------------|
 | FLUENT_BIT_PORT | The fluent-bit listen port metrics | 8080                                  |
 | LISTEN_PORT     | The logdrain HTTP listen port      | 2020                                  | 
 | SPLUNK_HOST     | The splunk host to forward to      | Example: `prd-p-xxxx.splunkcloud.com` |
 | SPLUNK_TOKEN    | The Splunk HEC token               |                                       |
 | TOKEN           | The logdrain TOKEN value           |                                       |

## License

License is MIT
