
func main(){
    file:=“test.txt”
    initNcsFromFile(file)
}

func initNcsFromFile(file string) {
    utils.Logger.Infof("Load nc record from file:%s", file)
    var vmList []SimpleNcInfo
    if sigma_api.Exists(file) {
        file, err := os.Open(file)
        sigma_api.Check(err)
        defer file.Close()
        // create a new scanner and read the file line by line
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            var line = scanner.Text()
            var detail  SimpleNcInfo
            dec := json.NewDecoder(strings.NewReader(line))
            if err = dec.Decode(&detail); err != nil {
                utils.Logger.Infof("Decode failed. %s", line)
            } else {
                vmList=append(vmList,detail)
                utils.Logger.Infof(">>Cell:%s,ip:%s",detail.Cell,detail.SeverIp) // 读取内容
            }
        }
    }else{
        utils.Logger.Infof("%s not exist",file)
    }
    if vmList==nil || len(vmList)==0 {
        utils.Logger.Infof("nc record is empty....")
        return
    }else{
        UpdateNcMemory(vmList)
    }
}

func UpdateNcMemory(fullIps[] SimpleNcInfo){
   ipMap:=make(map[string]*SimpleNcInfo)
    ipMap2:=make(map[string]SimpleNcInfo)
    for _,item:= range fullIps{
        if _,ok := ipMap[item.SeverIp];ok{
            utils.Logger.Infof("Repeat nc from boss...%s",item.SeverIp)
        }else{
            ipMap[item.SeverIp] = &item // 这里&item 就是陷阱
            ipMap2[item.SeverIp]= item 
        }
    }
    for key,value:=range  ipMap1{
        utils.Logger.Infof("Map1 key:%s,value.cell:%s",key,value.Cell)
    }

       for key,value:=range  ipMap2{
        utils.Logger.Infof("Map2 key:%s,value.cell:%s",key,value.Cell)
    }
}
