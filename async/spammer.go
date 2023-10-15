package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	in := make(chan interface{})
	wg := &sync.WaitGroup{}
	for _, command := range cmds {
		wg.Add(1)
		nextOut := make(chan interface{})
		go worker(wg, command, in, nextOut)
		in = nextOut
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup, command cmd, in, out chan interface{}) {
	defer func() {
		close(out)
		wg.Done()
	}()
	
	command(in, out)
}

func SelectUsers(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	isAlias := &sync.Map{}
	for email := range in {
		strEmail, ok := email.(string)
		if !ok {
			fmt.Println("Error happened in SelectUsers in type assertion")
			return
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			outputUser := GetUser(strEmail)
			if outputUser.Email != strEmail {
				if _, ok = isAlias.Load(outputUser.ID); ok {
					return
				}
				isAlias.Store(outputUser.ID, strEmail)
				return
			}

			out <- outputUser
			isAlias.Store(outputUser.ID, false)
		}()
	}
	wg.Wait()

	isAlias.Range(func(key, value interface{}) bool {
		keyUint, ok := key.(uint64)
		if !ok {
			fmt.Println("Error happened in SelectUsers in isAlias.Range in type assertion")
			return false
		}

		valueString, ok := value.(string)
		if !ok {
			return true
		}

		out <- User{ID: keyUint, Email: valueString}
		return true
	})
}

func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	batchUsers := make([]User, 0)
	for userObject := range in {
		user, ok := userObject.(User)
		if !ok {
			fmt.Println("Error happened in SelectMessages in type assertion")
			return
		}

		batchUsers = append(batchUsers, user)
		userObject, ok = <-in
		if ok {
			user, ok = userObject.(User)
			batchUsers = append(batchUsers, user)
		}

		wg.Add(1)
		go func(batchUsers []User) {
			defer wg.Done()
			msgs, err := GetMessages(batchUsers...)
			if err != nil {
				fmt.Println("Error happened in GetMessages ", err)
				return
			}
			for _, msg := range msgs {
				out <- msg
			}
		}(batchUsers)

		batchUsers = make([]User, 0)
	}
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	semaphore := make(chan struct{}, HasSpamMaxAsyncRequests)
	wg := sync.WaitGroup{}
	for msgID := range in {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(msgID interface{}) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			msg, ok := msgID.(MsgID)
			if !ok {
				fmt.Println("Error happened in CheckSpam in type assertion")
				return
			}

			isSpam, err := HasSpam(msg)
			if err != nil {
				fmt.Println("Error happened in HasSpam ", err)
				return
			}

			out <- MsgData{
				ID:      msg,
				HasSpam: isSpam,
			}
		}(msgID)
	}
	wg.Wait()
	close(semaphore)
}

func CombineResults(in, out chan interface{}) {
	var msgDataList []MsgData
	for msgData := range in {
		msg, ok := msgData.(MsgData)
		if !ok {
			fmt.Println("Error happened in CombineResults in type assertion")
			return
		}

		msgDataList = append(msgDataList, msg)
	}

	sort.Slice(msgDataList, func(i, j int) bool {
		if msgDataList[i].HasSpam != msgDataList[j].HasSpam {
			return msgDataList[i].HasSpam
		}

		return msgDataList[i].ID < msgDataList[j].ID
	})

	for _, msgData := range msgDataList {
		out <- fmt.Sprintf("%v %d", msgData.HasSpam, msgData.ID)
	}
}
