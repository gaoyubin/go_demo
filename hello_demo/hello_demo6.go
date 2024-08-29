package main

import "fmt"

/*
func main() {
	//fmt.Println(strings.Contains("浏览差异化品类频道", "差异化品类频道"))  // true
	//fmt.Println(strings.Contains("主端运营+浏览直播间", "差异化品类频道")) // false
	//
	//a := 3
	//if a == 3 {
	//	return
	//}
	//
	//defer func() {
	//	fmt.Println("hit")
	//}()
	//
	//fmt.Println("end")
	strTmp := "用户12345"
	r := []rune(strTmp)
	result := string(r[0:3]) + strings.Repeat("*", len(r)-3)

	fmt.Println(result)
}
*/

/*
func main() {
	group := new(errgroup.Group)

	nums := []int{-1, 0, 1}
	for _, num := range nums {
		num := num
		group.Go(func() error {
			res, err := output(num)
			fmt.Println(res)
			return err
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}
}

func output(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("math: square root error!")
	}
	return num, nil
}
*/
/*
func main() {
	var saasid int64
	saasid = 3400141315520407
	LotteryDailyFirstDrawKey := "lottery_daily_first_draw:%d:%s:%d"
	var activityID int64
	activityID = 7372073287158989094
	date := "2024-05-27"
	str := fmt.Sprintf(LotteryDailyFirstDrawKey, activityID, date, saasid)
	fmt.Println(str)
}
*/

/*
	func main() {
		runtime.GOMAXPROCS(2)
		//testMap := make(map[int64]int64)
		var group errgroup.Group
		var sum int64
		//var lock sync.Mutex
		for i := int64(0); i < 100000; i++ {
			//num := i
			group.Go(func() error {
				//lock.Lock()
				//testMap[num] = num
				//lock.Unlock()
				sum++
				return nil
			})
		}
		if err := group.Wait(); err != nil {
			fmt.Printf("group err: err=%v", err)
		}
		fmt.Printf("succ sum=%v", sum)
	}
*/
type People struct {
	Age  int
	Name string
}

func main() {
	//	defer fmt.Println("天才第一步")
	//	defer fmt.Println("雀氏纸尿裤")
	//	defer fmt.Println("战神第一步")
	//	defer fmt.Println("盖亚纸尿裤")

	// url := "https://p1-tt-gip.byteimg.com/origin/top-static-files-outer/d4e3ba359ad768953241d22fcb56df47/eade0d12-15f6-4608-ab3a-6178294afcdb"
	// index := strings.Index(url, "top-static-files-outerdfdf")
	// if index != -1 {
	// 	uri := url[index:]
	// 	fmt.Println(index, uri)
	// }
	// fmt.Println("end")
	//
	// nowTime := time.Now()
	// yes := nowTime.AddDate(0, 0, -1)
	// fmt.Println(yes)

	// str := "ab.c.d"
	// tmp := strings.Split(str, ".")
	// fmt.Println(tmp)

	// p1 := &People{
	// 	Age:  11,
	// 	Name: "lowbin",
	// }
	//
	// str1 := tools.ToJson(p1)
	// fmt.Println(str1)
	//
	// p2 := People{}
	// err := sonic.UnmarshalString(str1, &p2)
	// // err := json.Unmarshal([]byte(str1), &p2)
	// fmt.Println(p2, err)

	str1 := "hello"
	str2 := str1
	// str1 = str1 + "world"
	fmt.Println(str1[0])
	fmt.Println(str1, str2)

	s1 := []int{1, 2, 3}
	s2 := s1
	// s1 = append(s1, 4)
	s1[0] = 0
	fmt.Println(s1, s2)
}
