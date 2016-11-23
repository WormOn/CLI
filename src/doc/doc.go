package doc


type Judge struct{
    Name string
    Arglist string
}
func (judge *Judge) JudgeName(name string) string{
    if name == "ddcos"{
        description := "\tCommand line utility for the DDCOS Datacenter Operating System (DC/OS).\n\t The DDCOS is a distributed operating system built around Apache Mesos. \n\tThis utility provides tools for easy management of a DC/OS installation.\n\n\tAvailable DDCOS commands:\n\n\n\tlogin\t\tTo pass the auth of the DDCOS with the 'username' and 'password' that the DDCOS have assigned.\n\n\tlogout\t\tTo exit the DDCOS\n\n\ttasks\t\tGet the number of the corresponding app\n\n\tapplist\t\tGet the apps in the DDCOS\n\n\tconfig\t\tTo config the necessary information of the DDCOS,For example,url of dockbox,timeout of log in\n\n\trestart\t\trestart the specific app in the DDCOS\n\n\tscale\t\tScale out or scale in a specific app in the DDCOS\n\n\tstatus\t\tGet the status of the corresponding app\n\n\tGet detailed command description with 'ddcos <command>'.\n\n"
        return description
    }
    return "\n\n\t" + name + "Doesn't exist.\n\n"

}
func (judge *Judge) JudgeArgs(args...string) string{
    if len(args)==2{
        if args[0]=="ddcos" && args[1] == "logout"{
            dlogout := "\n\t-c\t\tTo exit the DDCOS.\n\n"
            return dlogout
        }else if args[0] == "ddcos" && args[1] == "applist"{
            dapplist := "\n\t-h\t\tThe marathon IP\n\t-p\t\tThe marathon port\n\n"
            return dapplist
        }else if args[0] == "ddcos" && args[1] =="config"{
            dconfig := "\n\tUsage:\tddcos config [Option] [IP]:[Port]\n\n\tOptions:\n\tdockbox_url\t\tConfig the dockbox_url\n\ttimeout\t\tConf the max time of log in\n\n\tParams:\n\tIP\t\tThe ip of the dockbox \n\tPort\t\tThe port of the dockbox\n\n"
            return dconfig
        }else if args[0] == "ddcos" && args[1] =="restart"{
            drestart :="\n\tUsage:\tddcos restart [appid] \n\tappid\tThe appid that have been deployed on DDCOS.\n"
            return drestart
        }else if args[0] == "ddcos" && args[1] =="scale"{
            dscale := "\n\tUsage:\tddcos scale [appid] [num]\n\tappid\tThe appid of the app to restart\n\tnum\tThe number to scale out or scale in\n\n"
            return dscale
        }else if args[0] == "ddcos" && args[1] == "login"{
            dlogin :="\n\tUsage:\tddcos login [username] \n\tusername\tThe username that have been authorized on DDCOS.\n\tpassword\tThe password that is assigned on DDCOS.\n"
            return dlogin
        }else if args[0] == "ddcos" && args[1] == "status"{
            dquery :="\n\tUsage:\tddcos status [appid] \n\tappid\tThe appid that have been deployed on DDCOS.\n"
            return dquery
        }else if args[0] == "ddcos" && args[1] == "tasks"{
            dtasks :="\n\tUsage:\tddcos tasks [appid] \n\tappid\tThe appid that have been deployed on DDCOS.\n"
            return dtasks
        }
        
    }else if len(args)==3{
        if args[0]=="ddcos" && args[1]=="config" && args[2]=="dockbox_url"{
            durl :="\n\tUsage:\n\n\tddcos config dockbox_url http://ip:port\n"
            return durl
        }else if args[0]=="ddcos" && args[1]=="config" && args[2]=="timeout"{
            dtimeout :="\n\tUsage:\n\n\tddcos config timeout [time(minute)]\n"
            return dtimeout
        }
    }
    return "\n\n\tThe command doesn't exist.\n\n"
     
}
