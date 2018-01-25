package main

import (
"fmt"
"io/ioutil"
"log"
"net/http"
	"net/url"
	"time"
)
type Arg struct {
	Msg string
	UID string
	Room string
}
var arg Arg
var timeout int
var flag=true
func findindex(s []byte,b byte)int  {
	var index int
	for i,char:=range s{
		if char==b {
			index=i
			break
		}
	}
	return index
}
func showroom()  {
	res, err := http.Get("http://localhost:12121/")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	s:=findindex(robots,'#')
	rooms:=robots[s+1:]
	fmt.Println(string(rooms))
}
func gethistroy()  {
	msg:=arg.UID+" come in\n"
	resp, err := http.PostForm("http://localhost:12121",
		url.Values{"Message": {msg}, "Uid": {"###历史###"},"Room":{arg.Room}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	/*fmt.Println("--------------------History--------------------")*/
}
func Reading()  {
	for true {
		var msgUid []byte
		var roomId []byte
		res, err := http.Get("http://localhost:12121/")
		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		l:=findindex(robots,'#')
		hint:=string(robots[:l])
		if hint!="齾" {
			i:=findindex(robots,'#')
			j:=findindex(robots,'&')
			k:=findindex(robots,'|')
			roomId=robots[i+1:j]
			msgUid=robots[j+1:k]
			message:=robots[k+1:]
			printing:=string(msgUid)+string(message)
			if string(roomId)== arg.Room{
					if string(msgUid)!="龖" {
						if string(msgUid)!=arg.UID {
							fmt.Println(printing)
							/*fmt.Println(msgUid)
							fmt.Println(i)
							fmt.Println(j)*/
						}
					}else {
						fmt.Println("This room is going to be closed#*#")
						timeout=-1
						break
					}
				}
		}/*else {
			m:=findindex(robots,'&')
			n:=findindex(robots,'|')
			s:=findindex(robots,'*')
			roomId:=string(robots[m+1:n])
			uId:=string(robots[n+1:s])
			if roomId==arg.Room {
				if uId!=arg.UID {
					if flag==true {
						h:=string(robots[l+1:m])
						fmt.Println(h)
						fmt.Println("--------------------History--------------------")
					}
				}
				flag=false
			}

		}*/

		time.Sleep(10*time.Millisecond)
	}

}
func httpPost() {
	fmt.Println("!!!!!!!!!!enjoy our chatting!!!!!!!!!!!!")
	for true {
		fmt.Scanln(&arg.Msg)
		if timeout!=-1 {
			t:=time.Now()
			fmt.Println(t.Format(time.RFC850))
			resp, err := http.PostForm("http://localhost:12121",
				url.Values{"Message": {arg.Msg}, "Uid": {arg.UID},"Room":{arg.Room}})
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
		}else {
			fmt.Println("Time up! See you in other room!")
			break
		}

	}

}
func main() {
	/*go Reading()*/
	fmt.Println("What is your name?")
	fmt.Scanln(&arg.UID)
	fmt.Println("which room do you want in?")
	showroom()
	fmt.Scanln(&arg.Room)
	fmt.Println("Welcome to Room: "+arg.Room+"! "+arg.UID)
	gethistroy()
	/*httpPost()*/
	go Reading()
	httpPost()
	/*time.Sleep(100*time.Minute)*/
}
