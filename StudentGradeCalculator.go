package main

import (
	"fmt"
)
func average(grades[]int) float64{
	var total int
	if len(grades)==0{
		return 0
	}
	for _, grade := range grades{
		total += grade
		
	}
	return  float64(total) / float64(len(grades))
}

func main() {
    fmt.Println("Hello you, welcome to Grade Calculator!")
	var name string
	var numberOfCourses uint
	subjectGradeMap := make(map[string]int)

	fmt.Println("Lets start  simple  ! Can you tell us Your name? ")
	fmt.Scan(&name)
	fmt.Println("Beautifull Name!, How many courses have you taken this semester?")
	fmt.Scan(&numberOfCourses)
	fmt.Printf("cool %v we will be right back with your average for %v courses! so you will be telling us the subject and the grade that you receive for that specific subject! ",name,numberOfCourses)
	
	var listOfGrades []int
	for i:=0; i<int(numberOfCourses); i++{
		var subject string
		var grade int

		fmt.Println("Write the course name and the grade you got (space-separated): ")
		fmt.Scan(&subject, &grade)

		if grade >= 0 && grade <= 100 {
			subjectGradeMap[subject] = grade
			
			listOfGrades=append(listOfGrades, grade)


		}else{
			fmt.Println("Invalid greade.please enter between 0 and 100")

		}
	

	

	}
	var avg=average(listOfGrades)

	fmt.Printf("hello again %v ! here are your results: \n",name)
	for subject,grade := range(subjectGradeMap){
		fmt.Printf("%v : %v \n",subject,grade)
	}
	
	fmt.Printf("you average turned out to be : %v",avg)



}