## 2023-04-13 for LTitan
* [doing] fix: using stand google protos library instead of gogo/protobuf
* [done] opt: using single instance mode for optimize `server` `controllers` some init functions
```
import "sync"
sync.Once{}.Do(something())
```
* [done] klog flag register into any components when process execute
