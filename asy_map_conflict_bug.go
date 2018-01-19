package main

import (
    "encoding/json"
    "fmt"
    seelog "github.com/cihub/seelog"
    "sync"
    "time"
)

//https://github.com/cihub/seelog/wiki/Logger-types
/* file seelog-main.xml
<seelog>  
    <outputs formatid="main">  
        <buffered size="10" flushperiod="1000">  
            <rollingfile type="date" filename="gologs/main.log" datepattern="2006.01.02" maxrolls="30"/>  
        </buffered>  
    </outputs>  
    <formats>  
        <format id="main" format="%Msg%n"/>  
    </formats>  
</seelog>  
*/

func main() {
    var logger, _ = seelog.LoggerFromConfigAsFile("conf/seelog-main.xml")
    seelog.ReplaceLogger(logger)
    defer seelog.Flush()
    seelog.Info("需要输入的日志")

    for i := 0; i < 40; i++ {
        //i := 0
        go workerSimlutor(i)
    }
    fmt.Printf("40 worker simlutor running....")
    time.Sleep(time.Duration(3) * 60 * time.Second)
    fmt.Printf("40 worker simlutor finish....")
}

func workerSimlutor(i int) {
    bmap := make(map[string]string)
    name := fmt.Sprintf("text_%d", i)

    ei := EventInfo{
        Name:           name,
        BusinessKey:    name,
        BusinessParams: bmap,
        lockBizParams:  &sync.RWMutex{},
    }

    for id := 0; id < 100; id++ {
        //strOut := ei.String() // 这里提前主动string，那么异步log 里面不会有map的read操作。
        seelog.Infof("i:%d,ei:%v", i, ei)// 这里是对象传入log，log里面 异步，异步marshal 对象，触发对对象里面map的read操作
    }

    for id := 0; id < 100; id++ {
        ei.SetVariable("num", fmt.Sprintf("%d", i)) // 这里更新map，触发write操作，
    }
}

type EventInfo struct {
    // 注册的 handler 的名称
    Name           string            `json:"name"`
    BusinessKey    string            `json:"businessKey"`
    BusinessParams map[string]string `json:"businessParams"`
    lockBizParams  *sync.RWMutex     `json:"-"`
}

func (ei *EventInfo) String() string {
    //ei.lockBizParams.Lock()
    //defer ei.lockBizParams.Unlock()  // 这里加锁能解决问题，不过复杂场景下，会导致死锁
    if b, err := json.Marshal(ei); err != nil {
        return fmt.Sprintf("%s", ei)
    } else {
        return fmt.Sprintf("%s", b)
    }
}

func (ei *EventInfo) SetVariable(key, val string) {
    ei.BusinessParams[key] = val
}

func (ei *EventInfo) GetVariable(key string) string {
    val, ok := ei.BusinessParams[key]
    if !ok {
        return ""
    }
    return val
}
