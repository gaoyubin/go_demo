package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

// useridList 长度<=50, 接口耗时：50ms
func BatchGetUserInfo(useridList []uint64) ([]UserInfo, error) {
	// userInfo := &UserInfo{
	// 	UserName: "lucas",
	// 	UserID:123,
	// }
	userInfos := make([]UserInfo, len(useridList))
	for i := 0; i < len(useridList); i++ {
		userInfos[i] = UserInfo{
			// UserName: "lucas",
			UserID: useridList[i],
		}
	}
	return userInfos, nil
}

type UserInfo struct {
	UserName string
	UserID   uint64
}

func BatchGetUserInfoV2(useridList []uint64) ([]UserInfo, error) {
	n := len(useridList)
	group := errgroup.Group{}
	userInfoRes := make([]UserInfo, n, n)
	// mutex sync.Mutex
	for i := 0; i <= n/50; i++ {
		startI := i
		group.Go(func() error {
			end := (startI + 1) * 50
			if (startI+1)*50 > n {
				end = n
			}
			userInfos, err := BatchGetUserInfo(useridList[startI*50 : end])
			// fmt.Println(end, userInfos, err)
			if err != nil {
				return err
			}
			for j := 0; j < len(userInfos); j++ {
				userInfoRes[startI*50+j] = userInfos[j]
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return userInfoRes, nil
}

func main() {
	useridList := make([]uint64, 0, 2000)
	for i := 0; i < 51; i++ {
		useridList = append(useridList, uint64(i))
	}
	res, err := BatchGetUserInfoV2(useridList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
