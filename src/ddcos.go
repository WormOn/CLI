package main


import(
    "os"
    "os/exec"
    "fmt"
    "cmds"
    "doc"
    "time"
    "strconv"
)
//counter for Timer
var count int
func main(){
    if len(os.Args) ==1{
        CliTip(len(os.Args))
    }else if len(os.Args) ==2{
        if os.Args[1]=="applist"{
            SelectMethod(os.Args)
            return
        }else if os.Args[1]=="logout"{
            SelectMethod(os.Args)
            return
        }else{
            CliTip(len(os.Args))
            return
        }
    }else{
        //CliTip(len(os.Args))
        SelectMethod(os.Args)
    }

}
//Timer
func Timer(){
    for{
        time.Sleep(1*time.Second)
        count+=1
        if count == 15{
            break
        }
    }
    os.Remove("/opt/.login")
    fmt.Println("Timeout,please login again!")
}
func LoginVerifyTips()(status int){
    f_status,_:= os.Stat("/opt/.login")
    //fmt.Println("===========================================auth========================================")
    if f_status ==nil{
        fmt.Println("Please login:\r\n\t\tddcos login [username] [password]\t\tThe username and password is authorized by DDCOS")
        status=1
        return
    }else{
        f,err:=os.Open("/opt/.login")
        defer f.Close()
        if err!=nil{
            fmt.Println(err)
            status=1
            return
        }
        buf :=make([]byte,1024)
        n,_:=f.Read(buf)
        sstatus :=string(buf[:n])
        //fmt.Println("sstatus:",sstatus)
        //fmt.Println(sstatus=="true")
        if sstatus!="true"{
            fmt.Println("Please login:\r\n\t\tddcos login [username]\t\tThe username and password is authorized by DDCOS")
            status=1
            return
        }
    }
    status=0
    return
}
//Timer for login

//Auth for Login
func Authority(username,password string)(rs bool){
    //fmt.Println("args in authority:",args)
    var log cmds.Log
    rs,_ = log.Auth(username,password)
    //fmt.Println("err in Authority:",err)
    return
}
//Select correspod method by the commd 
func SelectMethod(args []string){
        if args[0]=="ddcos" && args[1]=="config"{
            status :=LoginVerifyTips()
            count=0
            if status==1{
                return
            }
            var conf cmds.UrlConf
            err := conf.SetUrlConf(args[3])
            if err != nil{
                fmt.Println("Set dockboxurl failed.")
            }
            fmt.Println("\n\tSet dockboxurl success at '/opt/ddcos.cnf'.\n")
            fmt.Println("The method is applist")
        }else if args[0]=="ddcos" && args[1]=="applist"{
            status :=LoginVerifyTips()
            count=0
            if status==1{
                return
            }
            apps :=cmds.Apps{} 
            err := apps.AppListQuery()
            if err!=nil{
                fmt.Println("err",err)
            }
        }else if args[0] == "ddcos" && args[1] == "restart"{
            status :=LoginVerifyTips()
            count=0
            if status==1{
                return
            }
            var restart cmds.Restart
            res,_:=restart.RestartApp(args[2])
            //fmt.Println("res:",res)
            //fmt.Println("err:",err)
            if res!= "ok"{
                fmt.Println("\n\tRestart failed:",args[2])
            }else{
                fmt.Println("\n\tRestart successful:",args[2])
            }
        }else if args[0]=="ddcos" && args[1]=="login"{
            var password string
            fmt.Println("Please input your password:")
            exec.Command("/bin/stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
            exec.Command("/bin/stty", "-F", "/dev/tty", "-echo").Run()
            defer exec.Command("/bin/stty", "-F", "/dev/tty", "echo").Run()
            fmt.Scan(&password)
            //fmt.Println("password:",password)
            astatus:= Authority(args[2],password)
            if astatus==false{
                //fmt.Println("err",err)
                fmt.Println("Authority failed!")
                return 
            }
           // Timer()
            //fmt.Println("Authority successful!")
        }else if args[0]=="ddcos" && args[1]=="logout"{
            status :=LoginVerifyTips()
            if status==1{
                return
            }
            os.Remove("/opt/.login")
            fmt.Println("\n\tLogout Success!\n")
        }else if args[0] == "ddcos" && args[1] == "scale"{
            status :=LoginVerifyTips()
            count=0
            if status==1{
                return
            }
            var scale cmds.Scale
            num,_:=strconv.Atoi(args[3])
            res,_:=scale.ScaleApp(args[2],num)
            //fmt.Println("res:",res)
            //fmt.Println("err:",err)
            if res!= "ok"{
                fmt.Println("\n\tScale failed:",args[2])
            }else{
                fmt.Println("\n\tScale successful:",args[2])
            }
        }else if args[0] == "ddcos" && args[1] == "tasks"{
            status :=LoginVerifyTips()
            count=0
            if status==1{
                return
            }
            var tasks cmds.Tasks
            num,_:=tasks.AppTasks(args[2])
            //fmt.Println("res:",res)
            //fmt.Println("err:",err)
            if num==0{
                fmt.Println("\n\tquery tasks failed:",args[2])
            }else{
                fmt.Println("\n\tquery tasks successful:",num)
            }
        }
        
    
}
//Cli Tips
func CliTip(arg_num int){
    var judge doc.Judge
    //var apps cmds.Apps
    //var conf cmds.UrlConf
    //arg_num := len(os.Args)
    if arg_num == 1{
        parent := os.Args[0]
        judge.Name=string(parent)
        result:=judge.JudgeName(judge.Name)
        fmt.Println(result)
    }else if arg_num == 2{
        judge.Name=os.Args[0]
        judge.Arglist=os.Args[1]
        cresult := judge.JudgeArgs(judge.Name,judge.Arglist)
        fmt.Println(cresult)
    }else if arg_num == 3{
        //judge.Name=os.Args[0]
        //judge.Arglist=os.Args[1]
        cresult := judge.JudgeArgs(os.Args[0],os.Args[1],os.Args[2])
        fmt.Println(cresult)
    }
}
