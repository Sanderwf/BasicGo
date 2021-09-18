package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strings"
)

type User struct {
	id               string
	phone            string
	distributor_name string
}

func ImportFromXLS() {
	// 首先读excel
	xlsx, err := excelize.OpenFile("import_user.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var strAll strings.Builder
	rows := xlsx.GetRows("导出结果")
	var users []User
	for i, row := range rows {
		// 去掉第一行，第一行是表头
		if i == 0 {
			continue
		}
		var user User
		for j, colCell := range row {
			// 去掉前后空格（自己封装的方法，去掉字符串前后的特殊字符）
			colCell = strings.Trim(colCell, " ")
			// 排除第一列为Null
			if j == 0 && colCell == "Null" {
				continue
			}
			// 第一列即是一级
			if j == 0 && colCell != "Null" {
				user.id = colCell
			}
			// 第二列即是二级
			if j == 1 {
				user.distributor_name = colCell
			}
			// 三级
			if j == 2 {
				user.phone = colCell
			}

		}
		fmt.Println(user)
		users = append(users, user)
		strAll.WriteString("INSERT INTO `user_import` (`phone`, `distributor_name`) VALUES ('" + user.phone + "', '" + user.distributor_name + "');\n")

	}
	fmt.Println(strAll.String())
	Write(strAll.String())
}

func Write(str string) {
	fileName := "test.sql"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(str + "\n")
}


func InsertUser_FLAG(){
	for i:= 1291;i<=1948;i++{
		
	}
}
