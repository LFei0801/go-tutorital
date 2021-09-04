package main

import "fmt"

// 切片与数组的区别不大。 可用相同的方式声明这两者
func sliceBasis()  {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months) //["January",....,"December"]
	fmt.Println(len(months)) //12
	fmt.Println(cap(months)) //12

   // 在months切片上创建四个新切片
   createNewSlice(months)
}

// 切片项促使
func createNewSlice(months []string){
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	// [January February March] 3 12
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	//  [April May June] 3 9
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	// [July August September] 3 6
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	// [October November December] 3 3
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

   // 测试切片的容量
	getCap(quarter1) //11
    getCap(quarter2) //8
    getCap(quarter3) //5
    getCap(quarter4) //2
 }

func getCap(quarters []string){
	newQuarter := quarters[1:2]
	fmt.Println(cap(newQuarter))
}

func appendSlice(){
	numbers := []int{}
	for i:=0;i<10;i++{
		numbers = append(numbers,i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

/***
 删除slice切片中的索引为index这一项
 @param slice 原始切片
 @param index 删除项的下标
 */
func deleteSlice(slice []int,index int){
	remove := 2
	if len(slice) > remove{
		fmt.Println("Before: ",slice,"remove:",slice[remove])
		slice = append(slice[:index],slice[index+1:]...)
		fmt.Println("After: ",slice)
	}
}

// 复制切片
func copySlice(){
	letters := []string{"A","B","C","D","E"}
	//Before [A B C D E]
	fmt.Println("Before",letters)

	slice1 := letters[0:2]
	slice2 := make([]string,3)
	copy(slice2,letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}