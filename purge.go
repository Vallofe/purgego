package purge

import (
    "os/exec"
    "fmt"
)

func dualMode(gid string, authtoken string, app string, sub string, sub2 string){
    data := fmt.Sprintf("node /root/dual.js gid=%s %s %s token=%s app=%s",gid,sub,sub2,authtoken,app)
    data2 := strings.Fields(data)
    command, _ := exec.Command(data2[0],data2[1:]...).Output()
    fmt.Println(string(command))
