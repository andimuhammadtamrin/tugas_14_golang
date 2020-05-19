package main

import "fmt"
import "os/exec"

func main(){
  var hasil,_ = exec.Command("ls").Output()
  fmt.Println(" ls\n\n",string(hasil))
  var hasil2,_ = exec.Command("pwd").Output()
  fmt.Println(" pwd\n\n",string(hasil2))
  var hasil3,_ = exec.Command("git","config","user.name").Output()
  fmt.Println("git config username\n\n",string(hasil3))
}
