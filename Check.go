package main

func Check(MyArr []string) []string {
	Counter := 0
	if len(MyArr) > 1 {
		for i := 0; i < len(MyArr); i++ {
			if MyArr[i] == "" && i < len(MyArr)-1 && MyArr[i+1] != "" && Counter != 8 {
				MyArr = append(MyArr[:i], MyArr[i+1:]...)
			} else {
				Counter++
			}
		}
	}
	return MyArr
}
