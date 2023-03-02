# splunk-cf-logdrain

Logdrainer to forward Cloud foundry logs to Splunk

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
