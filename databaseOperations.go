package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type emp struct {
	id    int
	name  string
	phone string
}

func main() {

	fmt.Println("===Establishing Connection to database===")
	//Connecting to the database
	db, err := sql.Open("mysql", "root:maplelabs321#@tcp(127.0.0.1:3306)/employee")
	//Error check
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()1
	loop := true

	for loop {
		fmt.Println("Successfully Connected to the mysql Database...")
		fmt.Println()
		fmt.Println("============Welcome============")
		fmt.Println("Enter choice of Operation:")
		fmt.Println("1.Insertion")
		fmt.Println("2.Deletion")
		fmt.Println("3.Update")
		fmt.Println("4.Display")
		fmt.Println("5.Exit")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			//Inserting data to the database
			fmt.Println("==========Insert data==========")
			var id, phno int
			var name string
			fmt.Println("Enter Id: ")
			fmt.Scanln(&id)
			fmt.Println("Enter name: ")
			fmt.Scanln(&name)
			fmt.Println("Enter phno: ")
			fmt.Scanln(&phno)
			// insert, err :=
			stmt := fmt.Sprintf("INSERT INTO empdetails VALUES('%v','%s','%v')", id, name, phno)
			insert, err := db.Query(stmt)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
			fmt.Println("====Insert Successful====")
		} else if choice == 2 {
			//Deleting data from the database
			var id int
			fmt.Println("Enter Employee id to delete: ")
			fmt.Scanln(&id)
			fmt.Println("Deleting data...")
			stmt := fmt.Sprintf("DELETE FROM empdetails WHERE Emp_ID='%v'", id)
			delete, err := db.Query(stmt)
			if err != nil {
				panic(err.Error())
			}
			defer delete.Close()
			fmt.Println("=====Delete Successfull=====")
		} else if choice == 3 {
			//Updating data in database
			fmt.Println("======Update Data======")
			fmt.Println("Enter Employee Id: ")
			var id int
			fmt.Scanln(&id)
			var phoneNum int64
			fmt.Printf("Enter new Phone number: ")
			fmt.Scanln(&phoneNum)
			stmt := fmt.Sprintf("update empdetails set Phone_Num='%v' where Emp_ID='%v'", phoneNum, id)
			prep, err := db.Prepare(stmt)
			if err != nil {
				fmt.Print(err.Error())
			}
			a, _ := prep.Exec()
			fmt.Println("============Data Updated Successfully============")
			fmt.Println(a.RowsAffected())
			// defer sel.Close()
		} else if choice == 4 {
			//Displaying data
			fmt.Println("============Display Data============")
			sel, err := db.Query("SELECT * FROM empdetails")
			if err != nil {
				panic(err.Error())
			}
			e := emp{}
			for sel.Next() {
				val := sel.Scan(&e.id, &e.name, &e.phone)
				if val != nil {
					panic(val.Error())
				}
				fmt.Println("Emp ID: ", e.id, " Emp Name: ", e.name, " Phone: ", e.phone)
			}
		} else if choice == 5 {
			//Exiting the loop
			break
		} else {
			//Invalid choice
			fmt.Println("Entry not Recognized... Try Again")
		}
	}
}
