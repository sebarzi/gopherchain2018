package main

import "sort"
import (
    "fmt"
    "math/rand"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
    var interns_pod []*Candidate
    var interns []*Candidate
    podListMap := make(map[string]PodCandidates)

    r := rand.New(rand.NewSource(99))
    podList := []string{"", "", "", ""}
    for i := 0; i < 10; i++ {
        one := &Candidate{Pod: podList[r.Intn(4)], Weight: r.Float64()}
        interns = append(interns, one)
    }
    for _, item := range interns {
        fmt.Printf("item:%+v\n", item)
    }

    for _, candidate := range interns {
        podId := candidate.Pod
        if val, ok := podListMap[podId]; ok {
            list := val.CandidateList
            list = append(list, candidate)
            val.CandidateList = list
            podListMap[podId] = val
        } else {
            var list []*Candidate
            list = append(list, candidate)
            pod := PodCandidates{End: false, Index: 0, Pod: candidate.Pod, CandidateList: list}
            podListMap[podId] = pod
        }
    }

    var clist []PodCandidates
    for key, val := range podListMap {
        clist = append(clist, val)
        fmt.Printf("key:%s,", key)
        for _, c := range val.CandidateList {
            fmt.Printf("%+v", c)
        }
        fmt.Printf("\n")
    }

    size := len(interns)
    for i := 0; i < size; {
        sort.Sort(ByTopScore(clist))
        fmt.Printf("After Sort\n")
        for _, c := range clist {
            for _, it := range c.CandidateList {
                fmt.Printf("%+v", it)
            }
            fmt.Printf("\n")
        }

        newCandidatas := fetchTop(clist)
        for _, newc := range newCandidatas {
            interns_pod = append(interns_pod, newc)
        }
        i = i + len(newCandidatas)
        removeFirst(clist)
    }
    for _, newItem := range interns_pod {
        fmt.Printf("new_item:%+v\n", newItem)
    }
    fmt.Printf("interns_pod:%+v\n", interns_pod)
}

type ByTopScore []PodCandidates

func (w ByTopScore) Len() int      { return len(w) }
func (w ByTopScore) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w ByTopScore) Less(i, j int) bool {
    index_i := w[i].Index
    index_j := w[j].Index
    val_i := w[i].CandidateList[index_i].Weight
    val_j := w[j].CandidateList[index_j].Weight
    return val_i > val_j
}

type Candidate struct {
    Pod    string
    Weight float64
}

type PodCandidates struct {
    End           bool
    Index         int
    Pod           string
    CandidateList []*Candidate
}

func fetchTop(clist []PodCandidates) []*Candidate {
    var topItems []*Candidate
    for _, item := range clist {
        if item.End {
            continue
        } else {
            index := item.Index
            first := item.CandidateList[index]
            topItems = append(topItems, first)
        }
    }
    fmt.Printf("OutPut Index new Index \n")
    for _, item := range clist {
        fmt.Printf("%d ", item.Index)
    }
    fmt.Printf("\nEnd Output Index\n")
    return topItems
}

func removeFirst(clist []PodCandidates) {
    size := len(clist)
    for i := 0; i < size; i++ {
        nextIndex := clist[i].Index + 1
        if nextIndex < len(clist[i].CandidateList) {
            clist[i].Index = nextIndex
            clist[i].End = false
        } else {
            clist[i].End = true
        }
    }
}
