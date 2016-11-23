package cmds

import(
    "io"
    "os"
    "net/http"
    //"log"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "bytes"
    )
//common method
func GetURL(filename string)(url string,err error){
    f,err2 :=os.Open("/opt/ddcos.cnf")
    defer f.Close()
    if err2 !=nil{
        fmt.Println(err2)
        err=err2
        url=""
        return
    }
    buf :=make([]byte,1024)
    n,_ :=f.Read(buf)
    url =string(buf[:n])
    return 
}
//tasks
type Tasks struct{
    App_id string `json:"app_id"`
    Username string `json:"username"`
    User_id string `json:"user_id"`
}
type TasksResp struct{
    Msg string `json:"msg"`
    St  string `json:"st"`
    Tasknum int `json:"tasknum"` 
}
func(tasks Tasks) AppTasks(appid string)(num int,err error){
    result :=TasksResp{}
    url,_:=GetURL("/opt/ddcos.cnf")
    url2 :=url+"/v1.0/apps/appcount"
    b:=Tasks{Username:"cli",App_id:appid,User_id:"cli"}
    jb,err1:=json.Marshal(b)
    if err1!=nil{
        fmt.Println("err in RestartApp:",err1)
        num=0
        err =err1
        return
    }
    body :=bytes.NewBuffer([]byte(jb))
    res,err2:=http.Post(url2,"application/json;charset=utf-8",body)
    //fmt.Println("res:",res)
    if err2 != nil{
        fmt.Println(err2)
        num=0
        err =err2
        return
    }
    res2,err3 := ioutil.ReadAll(res.Body)
    if err3 != nil{
        fmt.Println(err3)
        num=0
        err =err3
        return
    }
    //fmt.Println("res2:",string(res2))
    defer res.Body.Close()
    err4 := json.Unmarshal(res2,&result)
    if err4 != nil{
        fmt.Println(err4)
        num=0
        err =err4
        return
    }
    fmt.Println("\n\tresult:",result)
    num=result.Tasknum
    return
}
//Status
type Status struct{
    App_id string `json:"app_id"`
}
type StatusResp struct{
    Msg string `json:"msg"`
    Data []map `json:"data"`
    St  string `json:"st"`
}
func(status Status) StatusApp(appid string)(status string,err error){
    result :=ScaleResp{}
    url,_:=GetURL("/opt/ddcos.cnf")
    url2 :=url+"/v1.0/apps/scale"
    b:=Scale{Username:"cli",App_id:appid,User_id:"cli",Scalenum:num}
    jb,err1:=json.Marshal(b)
    if err1!=nil{
        fmt.Println("err in RestartApp:",err1)
        status=""
        err =err1
        return
    }
    body :=bytes.NewBuffer([]byte(jb))
    res,err2:=http.Post(url2,"application/json;charset=utf-8",body)
    //fmt.Println("res:",res)
    if err2 != nil{
        fmt.Println(err2)
        status=""
        err =err2
        return
    }
    res2,err3 := ioutil.ReadAll(res.Body)
    if err3 != nil{
        fmt.Println(err3)
        status=""
        err =err3
        return
    }
    //fmt.Println("res2:",string(res2))
    defer res.Body.Close()
    err4 := json.Unmarshal(res2,&result)
    if err4 != nil{
        fmt.Println(err4)
        status=""
        err =err4
        return
    }
    fmt.Println("\n\tresult:",result)
    status=result.St
    return
}

//Scale app
type Scale struct{
    App_id string `json:"app_id"`
    Username string `json:"username"`
    User_id string `json:"user_id"`
    Scalenum int    `json:"scalenum"`
}
type ScaleResp struct{
    Msg string `json:"msg"`
    St  string `json:"st"`
}
func(scale Scale) ScaleApp(appid string,num int)(status string,err error){
    result :=ScaleResp{}
    url,_:=GetURL("/opt/ddcos.cnf")
    url2 :=url+"/v1.0/apps/scale"
    b:=Scale{Username:"cli",App_id:appid,User_id:"cli",Scalenum:num}
    jb,err1:=json.Marshal(b)
    if err1!=nil{
        fmt.Println("err in RestartApp:",err1)
        status=""
        err =err1
        return
    }
    body :=bytes.NewBuffer([]byte(jb))
    res,err2:=http.Post(url2,"application/json;charset=utf-8",body)
    //fmt.Println("res:",res)
    if err2 != nil{
        fmt.Println(err2)
        status=""
        err =err2
        return
    }
    res2,err3 := ioutil.ReadAll(res.Body)
    if err3 != nil{
        fmt.Println(err3)
        status=""
        err =err3
        return
    }
    //fmt.Println("res2:",string(res2))
    defer res.Body.Close()
    err4 := json.Unmarshal(res2,&result)
    if err4 != nil{
        fmt.Println(err4)
        status=""
        err =err4
        return
    }
    fmt.Println("\n\tresult:",result)
    status=result.St
    return
}
    
//restart app
type Restart struct{
    Username string `json:"username"`
    App_id string   `json:"app_id"`
    User_id string `json:"user_id"`

}
type RestartResp struct{
    Msg string `json:"msg"`
    Status string `json:"st"`
}
func (restart Restart) RestartApp(appid string)(status string,err error){
    result :=RestartResp{}
    url,_:=GetURL("/opt/ddcos.cnf")
    url2 :=url+"/v1.0/apps/restart"
    b:=Restart{Username:"admin",App_id:appid,User_id:"admin"}
    jb,err1:=json.Marshal(b)
    if err1!=nil{
        fmt.Println("err in RestartApp:",err1)
        status=""
        err =err1
        return
    }
    body :=bytes.NewBuffer([]byte(jb))
    res,err2:=http.Post(url2,"application/json;charset=utf-8",body)
    //fmt.Println("res:",res)
    if err2 != nil{
        fmt.Println(err2)
        status=""
        err =err2
        return
    }
    res2,err3 := ioutil.ReadAll(res.Body)
    if err3 != nil{
        fmt.Println(err3)
        status=""
        err =err3
        return
    }
    //fmt.Println("res2:",string(res2))
    defer res.Body.Close()
    err4 := json.Unmarshal(res2,&result)
    if err4 != nil{
        fmt.Println(err4)
        status=""
        err =err4
        return
    }
    fmt.Println("\n\tresult:",result)
    status=result.Status
    return

}
//authority
type Log struct{
    UserName string `json:"username"`
    PassWord string `json:"password"`
}
type LogResp struct{
    St string `json:"st"`
    Judge bool `json:"judge"`
    Password string `json:"password"`
    Msg string `json:"msg"`
}
func(log Log)Auth(username ,password string)(rs bool,err error){
    clog :=Log{UserName:username,PassWord:password}
    b, err1:= json.Marshal(clog)
    if err1 != nil {
        fmt.Println("json err:", err1)
        err=err1
        return
    }
    body := bytes.NewBuffer([]byte(b))
    f,err2 :=os.Open("/opt/ddcos.cnf")
    defer f.Close()
    if err2 !=nil{
        fmt.Println(err2)
        err=err2
        return
    }
    buf :=make([]byte,1024)
    n,_ :=f.Read(buf)
    url :=string(buf[:n])
    //fmt.Println("url:",url)
    //url2 :=url+"/v1.0/apps/applist"
    url2 :=url+"/v1.0/goauth/auth"
    resp,err3 :=http.Post(url2,"application/json;charset=utf-8",body)
    if err3 != nil{
        fmt.Println("err:",err3)
        err=err3
        return
    }
    defer resp.Body.Close()
    rbody,err4 :=ioutil.ReadAll(resp.Body)
    if err4 != nil{
        fmt.Println("err:",err4)
        err=err4
        return
    }
    result :=LogResp{}
    err5 := json.Unmarshal(rbody,&result)
    if err5 != nil{
        fmt.Println(err5)
        rs=false
        err=err5
        return 
    }

    //fmt.Println("result:",result.Judge==true)
    if result.Judge !=true{
        fmt.Println("\n\tAuthority Failed.\n")
    }else{
        fmt.Println("\n\t\tAuthority Successful.\n")
    }
    if result.Judge==true{
        f,err6:=os.Create("/opt/.login")
        defer f.Close()
        if err6 != nil{
            fmt.Println(err6)
            err=err6
            rs=false
            return
        }
        _,err7:=io.WriteString(f,"true")
        if err7 !=nil{
            fmt.Println(err7)
            err=err7
            rs=false
            return
        }    
    }else{
        f,err8:=os.Create("/opt/.login")
        defer f.Close()
        if err8 != nil{
            fmt.Println(err8)
            err=err8
            rs=false
            return
        }
        _,err9:=io.WriteString(f,"false")
        if err9 !=nil{
            fmt.Println(err9)
            err=err9
            return 

        }
    }
    rs=true
    return
}
//url set and get

type UrlConf struct{
    IP string
    Port string
}
func (conf UrlConf) SetUrlConf(url string)(ret error){
        //url :="http://"+ip+":"+port
        f,err:=os.Create("/opt/ddcos.cnf")
        defer f.Close()
        if err != nil{
            fmt.Println(err)
           return err
        }
        _,err2:=io.WriteString(f,url)
        if err2 !=nil{
            fmt.Println(err2)
            return err2
        }
        return 
}
func (conf UrlConf) ReadUrlConf()(url string,err error){
    f,err := os.Open("/etc/ddcos.cnf")
    defer f.Close()
    if err !=nil{
        fmt.Println(err)
        return "Open file Failed.",err
    }
    buf := make([]byte,1024)
    n,_ := f.Read(buf)
    fmt.Printf("buf:%d,%s",n,buf)
    fmt.Println("buf:",string(buf[:n]))
    return string(buf[:n]),nil
}
//app list 
type Apps struct{
    Ip string
    Port string
}
type AppsResp struct{
    St string `json:"st"`
    Data []map[string]string `json:"data"`
    Msg string `json:"msg"`
}
func(apps Apps) AppListQuery()(err error){
    //fmt.Println("IP:",IP,"Port:",Port)
    //url := "http://" + IP + ":" + Port + "/v1.0/apps/applist"
    f,err1 :=os.Open("/opt/ddcos.cnf")
    defer f.Close()
    if err !=nil{
        fmt.Println(err)
        err=err1
        return
    }
    buf :=make([]byte,1024)
    n,_ :=f.Read(buf)
    url :=string(buf[:n])
    //fmt.Println("url:",url)
    url2 :=url+"/v1.0/apps/applist"
    //fmt.Println("=========================================Begin_Call=====================================================")
    resp,err2 :=http.Post(url2,"application/json;charset=utf-8",nil)
    //fmt.Println("=========================================End_Call=====================================================")
    if err2!= nil{
        fmt.Println("err2:",err2)
        err=err2
        return
    }
    //fmt.Println("resp:",resp)
    defer resp.Body.Close()
    body,err3 :=ioutil.ReadAll(resp.Body)
    if err3 != nil{
        fmt.Println("err3:",err3)
        err=err3
        return
    }
    //fmt.Println("body:",body)
    result :=AppsResp{}
    err4 := json.Unmarshal(body,&result)
    if err4 != nil{
        fmt.Println("err4===========:",err4)
        err=err4
        return
    }

    //fmt.Println("result:",result.Data)
    for i,item :=range result.Data{
        fmt.Println("================================APP:",i,"=================================================")
        fmt.Printf("sysname:%15s||\tapp_id:%15s\t||appname:%15s\n",item["sysname"],item["app_id"],item["appname"])
        fmt.Println("=========================================================================================")
        /*
        fmt.Println("||","\t\t\t\t\t","||")
        fmt.Println("||","\tapp_id:","\t\t",item["app_id"],"\t","||")
        fmt.Println("||","\t\t\t\t\t","||")
        fmt.Println("||","\tappname:","\t\t",item["appname"],"\t","||")
        fmt.Println("||=======================================||")
        */
        fmt.Println("\n")
    }
    return

}
