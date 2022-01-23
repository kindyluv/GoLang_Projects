package main

import (
	"fmt"
	"golang-book/chapterOne/mapping"
)

func main() {

	cohortEight := make(map[int]string, 29)

	cohortEight[1] = "Olorunishola Emmanuel"
	cohortEight[2] = "Mojoyinola Kimberly"
	cohortEight[3] = "Gideon Udoh"
	cohortEight[4] = "Dairo Emmanuel"
	cohortEight[5] = "Yetife"
	cohortEight[6] = "Tife"
	cohortEight[7] = "Muqtar"
	cohortEight[8] = "Tosin"
	cohortEight[9] = "Ahmad"
	cohortEight[10] = "Toye David"
	cohortEight[11] = "Precious Lois"
	cohortEight[12] = "Israel Fasina"
	cohortEight[13] = "Abdulfatai"
	cohortEight[14] = "Saheed"
	cohortEight[15] = "Itunu"
	cohortEight[16] = "Goodnews Ozichukwu"
	cohortEight[17] = "John"
	cohortEight[18] = "Segun Emmanuel"
	cohortEight[19] = "Solomon Archibong"
	cohortEight[20] = "olaolu olowe"
	cohortEight[21] = "Teslim"
	cohortEight[22] = "Ademola"
	cohortEight[23] = "Jerry"
	cohortEight[24] = "Adesuwa"
	cohortEight[25] = "Femi"
	cohortEight[26] = "Dami Omolori"
	cohortEight[27] = "Ozioma Okoroafor"
	cohortEight[28] = "Harrison"
	cohortEight[29] = "Monsur Kareem"

	var k int
	fmt.Print("Enter native number to get native details: ")

	fmt.Scanf("%d", &k)

	maps := mapping.Maps{01, "precious"}
	id, details := maps.Phoenix(cohortEight, k)
	fmt.Printf("id is %d and details : %s", id, details)

}
