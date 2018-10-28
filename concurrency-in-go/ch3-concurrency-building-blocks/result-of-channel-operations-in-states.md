# Result of channel operations given a channel’s state

| Operation | Channel state | Result | 
|---|---|---|
| Read | nil | Block | 
| | Open and Not Empty | Value |
| | Open and Empty | Block |
| | Closed | \<default value>, false |
| | Write Only | Compilation Error |
| Write | nil | Block | 
| | Open and Full | Block |
| | Open and Not Full | Write Value |
| | Closed | __panic__ |
| | Receive Only | Compilation Error |
| Close | nil | __panic__ | 
| | Open and Not Empty | Closes Channel; reads succeed until channel is drained, then reads produce default value |
| | Open and Empty | Closes Channel; reads produces default value |
| | Closed | __panic__ |
| | Receive Only | Compilation Error |
