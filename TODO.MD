## 2023-04-13 for LTitan
* fix: using stand google protos library instead of gogo/protobuf
* opt: using single instance mode for optimize `server` `controllers` some init functions
```
import "sync"
sync.Once{}.Do(something())
```