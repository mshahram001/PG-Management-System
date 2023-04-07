package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"PG-Management-Go-Project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Add_property() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "addproperty.html", nil)
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		var add_property models.User
		err = c.BindJSON(&add_property)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO PropertyDetails (Property_Name, Contact_No, Property_Type, Property_Address, City, Pin_Code, Landmark, Ammeneties, Price_Month, Price_Day, Advance_Deposit_Month, Advance_Deposit_Day)
				          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		stmt, err := db.Prepare(query)
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(add_property.Property_Name, add_property.Contact_No, add_property.Property_Type, add_property.Property_Address, add_property.City, add_property.Pin_Code, add_property.Landmark, add_property.Ammeneties, add_property.Price_Month, add_property.Price_Day, add_property.Advance_Deposit_Month, add_property.Advance_Deposit_Day)
		if err != nil {
			panic(err.Error())
		}

		c.IndentedJSON(http.StatusOK, "Yes, PG added successfully")
	}
}

func Update_property() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var edit_property models.User
		err = c.BindJSON(&edit_property)
		if err != nil {
			return
		}
		query := fmt.Sprintf("UPDATE PropertyDetails SET Property_Name='%s',Contact_No='%s',Property_Type='%s',Property_Address='%s',City='%s',Pin_Code='%s',LandMark='%s',Ammeneties='%s',Price_Month='%d',Price_Day='%d',Advance_Deposit_Month='%d', Advance_Deposit_Day='%d' WHERE Property_ID ='%d' ", *edit_property.Property_Name, *edit_property.Contact_No, *edit_property.Property_Type, *edit_property.Property_Address, *edit_property.City, *edit_property.Pin_Code, *edit_property.Landmark, *edit_property.Ammeneties, *edit_property.Price_Month, *edit_property.Price_Day, *edit_property.Advance_Deposit_Month, *edit_property.Advance_Deposit_Day, *edit_property.Property_ID)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, "Yes, PG Update Successfully!")
	}
}

func Delete_property() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var delete_property models.User
		err = c.BindJSON(&delete_property)
		if err != nil {
			return
		}
		query := fmt.Sprintf("DELETE FROM PropertyDetails WHERE Property_ID=%d", *delete_property.Property_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(200, "Yes, PG Delete Successfully!")
		defer results.Close()
	}
}

func See_bookings() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var check_bookings models.Booking_pg
		err = c.BindJSON(&check_bookings)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM BookingDetails WHERE Property_ID=%d", check_bookings.Property_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Booking_id int
			var Customer_id int
			var Customer_name string
			var Cus_Contact_no string
			var Property_id int
			// var Booking_time string
			var From_date string
			var To_date string
			err = results.Scan(&Booking_id, &Customer_id, &Customer_name, &Cus_Contact_no, &Property_id, &From_date, &To_date)
			if err != nil {
				panic(err.Error())
			}
			c.IndentedJSON(200, "See Your Booking")
			output = fmt.Sprintf("Booking_Id=%d, Customer_ID=%d, Customer_Name='%s', Cus_Contact_No='%s', Property_ID=%d,  From_Date='%s', To_date='%s' ", Booking_id, Customer_id, Customer_name, Cus_Contact_no, Property_id, From_date, To_date)
			c.JSON(http.StatusOK, gin.H{"": output})
		}

	}
}

func See_booking_by_propertyid() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_bypropertyid models.Booking_pg
		err = c.BindJSON(&checkpg_bypropertyid)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM bookingdetails WHERE Property_ID='%d'", checkpg_bypropertyid.Property_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var customer_id int
			var customer_name string
			var cus_contact_no string
			var property_id int
			var booking_id int
			var from_date string
			var to_date string
			err = results.Scan(&booking_id, &customer_id, &customer_name, &cus_contact_no, &property_id, &from_date, &to_date)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Booking_ID='%d' Customer_ID='%d',  CUSTOMER_Name='%s'  Cus_Contact_No='%s'  Property_ID='%d' From_Date='%s' To_Date='%s'", booking_id, customer_id, customer_name, cus_contact_no, property_id, from_date, to_date)

			c.IndentedJSON(200, "See_bookings_by_Property_ID")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG Bookings By Property_ID")
	}
}
