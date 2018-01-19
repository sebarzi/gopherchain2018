package main

import (
    "fmt"
    "strconv"
    "time"
)

type Job struct {
    JobId   string
    JobCont string
}

func main() {
    stopWork := make(chan bool)
    finish := make(chan bool)
        defer close(stopWork)
        defer close(finish)
    var jobs []*Job
    var tmps []*Job
    for i := 0; i < 5; i++ {
        str := strconv.Itoa(i)
        job := &Job{JobId: str, JobCont: str}
        jobs = append(jobs, job)
        tmps = append(tmps, job)
    }
    go doWork(finish, stopWork, jobs)

    stopMain := false
    time_out := time.After(1 * time.Minute)
    quitTimeOut := false

    for {
        if stopMain {
            break
        }
        select {
        case s := <-finish:
            if s == true {
                fmt.Printf("finish..%v\n", s)
                stopMain = true
                quitTimeOut = false
                break
            }

        case <-time_out:
            fmt.Println("You cost too much time quit now.")
            stopMain = true
            quitTimeOut = true
            stopWork <- true
            break
        }
    }

    if quitTimeOut {
        fmt.Printf("Quit timeout...\n")
    } else {
        fmt.Printf("Quit common...\n")
    }
    size := len(tmps)
    for i := 0; i < size; i++ {
        if jobs[i] == nil {
            fmt.Printf("Already Check jobId:%d\n", i)
        }
    }

    time.Sleep(time.Duration(1) * time.Minute) // 10秒 check 一次
    fmt.Printf("Quit main...\n")
}

func doWork(result chan bool, stopSignal chan bool, jobs []*Job) {
    size := len(jobs)
    stop := false
    rount := 1
    for {
        if stop {
            break
        }
        select {
        case stop = <-stopSignal:
            fmt.Printf("Notity timeout stop...")
        default:
            checkNum := 0

            fmt.Printf("Check....rount:%d\n", rount)
            for i := 0; i < size; i++ {
                if jobs[i] != nil {
                    if i%2 == 0 {
                        jobs[i] = nil
                        fmt.Printf("check i:%d\n", i)
                        checkNum++
                    } else {
                        fmt.Printf("pass check i:%d\n", i)
                    }
                } else {
                    checkNum++
                }

                if checkNum == size {
                    stop = true
                }
            }
            rount++
            time.Sleep(time.Duration(2) * time.Second) // 10秒 check 一次
        }
    }
    result <- true
}
