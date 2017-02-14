# Send message to Kafka
This activity provides your flogo application the ability to send a message to an Apache Kafka broker.


## Installation

```bash
flogo add activity github.com/jvanderl/flogo-components/activity/kafka
```

## Schema
Inputs and Outputs:

```json
  "inputs":[
    {
      "name": "server",
      "type": "string"
    },
    {
      "name": "configid",
      "type": "string"
    },
    {
      "name": "topic",
      "type": "string"
    },
    {
      "name": "partition",
      "type": "int"
    },
    {
      "name": "message",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| server | The Kafka server [ipaddress]:[port] |         
| configid | The Kafka broker configuration name |
| topic | The Kafka topic name |
| partition | The Kafka partition |
| message  | The message content  |

## Configuration Examples
### Simple
Configure a task in flow to send 'hello from flogo' to kafka topic 'test', partition 0:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "kafka",
  "name": "Send Message",
  "attributes": [
    { "name": "server", "value": "192.168.178.41:2181" },
    { "name": "configid", "value": "flogo-test" },
    { "name": "topic", "value": "test" },
    { "name": "partition", "value": "0" },
    { "name": "message", "value": "hello from flogo" },
  ]
}
```