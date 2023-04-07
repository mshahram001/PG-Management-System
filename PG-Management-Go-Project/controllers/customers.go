package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"PG-Management-Go-Project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Get_All_PG() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		results, err := db.Query("SELECT * FROM PropertyDetails")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_month int
			var price_day int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'  Advance_Deposit_Month='%d' Advance_Deposit_Day='%d'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Book_PG() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var add_booking models.Booking_pg
		err = c.BindJSON(&add_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf(`INSERT INTO BookingDetails (Customer_ID,Customer_Name,Cus_Contact_No,Property_ID,From_Date,To_date) VALUES(%d,"%s","%s",%d,"%s","%s")`, add_booking.Customer_ID, add_booking.Customer_Name, add_booking.Cus_Contact_No, add_booking.Property_ID, add_booking.From_Date, add_booking.To_Date)

		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.IndentedJSON(200, "Yes, PG Book Successfully!")

	}
}

func Update_booking() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var edit_booking models.Booking_pg
		err = c.BindJSON(&edit_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf("UPDATE BookingDetails SET Customer_ID=%d,Customer_Name='%s',Cus_Contact_No='%s',From_Date='%s',To_Date='%s' WHERE Booking_ID=%d", edit_booking.Customer_ID, edit_booking.Customer_Name, edit_booking.Cus_Contact_No, edit_booking.From_Date, edit_booking.To_Date, edit_booking.Booking_ID)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, "Yes, Booking Update Successfully!")
	}
}

func Delete_booking() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var delete_booking models.Booking_pg
		err = c.BindJSON(&delete_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf("DELETE FROM BookingDetails WHERE Booking_ID=%d", delete_booking.Booking_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(200, "Yes, Booking Delete Successfully!")
		defer results.Close()
	}
}

func Get_PG_By_location() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checklocation models.User
		err = c.BindJSON(&checklocation)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM propertydetails WHERE Landmark='%s' AND  City='%s'", *checklocation.Landmark, *checklocation.City)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_day int
			var price_month int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'  Advance_Deposit_Month='%d' Advance_Deposit_Day='%d' ", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG_by_location")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG By Location")
	}
}

func Get_PG_By_Price_Month() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_byprice_month models.User
		err = c.BindJSON(&checkpg_byprice_month)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM propertydetails WHERE Price_Month < '%d'", *checkpg_byprice_month.Price_Month)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_month int
			var price_day int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'   Advance_Deposit_Month='%d' Advance_Deposit_Day='%d'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG_by_Price_Month")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG By Price")
	}
}

func Get_PG_By_Price_Day() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_byprice_day models.User
		err = c.BindJSON(&checkpg_byprice_day)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM propertydetails WHERE Price_Day < '%d'", *checkpg_byprice_day.Price_Day)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_month int
			var price_day int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'   Advance_Deposit_Month='%d' Advance_Deposit_Day='%d'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG_by_Price_Day")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG By Price")
	}
}

func Get_PG_By_Type() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_bytype models.User
		err = c.BindJSON(&checkpg_bytype)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM propertydetails WHERE Property_Type='%s'", *checkpg_bytype.Property_Type)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_month int
			var price_day int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'   Advance_Deposit_Month='%d' Advance_Deposit_Day='%d'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG_by_Type")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG By Type")
	}
}

func Get_PG_By_Ammeneties() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_byammeneties models.User
		err = c.BindJSON(&checkpg_byammeneties)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM propertydetails WHERE Ammeneties='%s'", *checkpg_byammeneties.Ammeneties)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_month int
			var price_day int
			var advancedeposit_month int
			var advancedeposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &price_day, &advancedeposit_month, &advancedeposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month='%d' Price_Day='%d'   Advance_Deposit_Month='%d' Advance_Deposit_Day='%d'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, price_day, advancedeposit_month, advancedeposit_day)

			c.IndentedJSON(200, "PG_by_Ammeneties")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your PG By Ammeneties")
	}
}

func See_booking_by_customerid() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var checkpg_bycustid models.Booking_pg
		err = c.BindJSON(&checkpg_bycustid)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM bookingdetails WHERE Customer_ID='%d'", checkpg_bycustid.Customer_ID)
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
			output = fmt.Sprintf(" Booking_ID='%d' Customer_ID='%d',  CUSTOMER_Name='%s'  Cus_Contact_No='%s'  Property_ID='%d'     From_Date='%s' To_Date='%s'", booking_id, customer_id, customer_name, cus_contact_no, property_id, from_date, to_date)

			c.IndentedJSON(200, "See_bookings_by_customerID")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
		fmt.Println("You successfully fetch your Booking By your Customer_ID")
	}
}
