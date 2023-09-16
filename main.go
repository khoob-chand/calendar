package main

import (
"encoding/json"
"fmt"
"time"
"log"
"context"
"strconv"
"reflect"  
"github.com/gofiber/fiber/v2"
"github.com/gofiber/template/html/v2"
"github.com/jackc/pgx/v4/pgxpool"

)
// type Calender struct{
// 	Title string;
// 	Occurence int;
// 	Start_date string;
// 	End_date string;
// 	Mode string;
// 	Day string;
// 	Case string;
// 	Week string;
// }
type Calender struct{
		Title string;
		Occurence int;
		Start_date time.Time;
		End_date time.Time;
		Mode string;
		Day string;
		Case string;
		Week string;
	}
func main() {
	// test()
	getRecord := []Calender{}

	pool,_:= pgxpool.Connect(context.Background(),"postgres://postgres:root@localhost:5432/dbname")
	fmt.Println(reflect.TypeOf(pool))	

    app := fiber.New(fiber.Config{
        Views: html.New("./views", ".html"),
    })
	
	

	app.Get("/", func(c *fiber.Ctx) error {
		
		
		// rows,_:=pool.Query(context.Background(),`SELECT occurence,title,TO_CHAR(start_date, 'YYYY-MM-DD') ,TO_CHAR(end_date, 'YYYY-MM-DD'),mode ,day,selectcase,week FROM calender`)
		rows,_:=pool.Query(context.Background(),`SELECT occurence,title,start_date ,end_date,mode ,day,selectcase,week FROM calender`)

	
		for rows.Next() {
			var cal Calender
			// fmt.Println("kake")
			err:= rows.Scan(&cal.Occurence, &cal.Title,&cal.Start_date,&cal.End_date,&cal.Mode,&cal.Day,&cal.Case,&cal.Week)	
		if err != nil {		
			return err
		}
		getRecord = append(getRecord, cal)	
		}	

	

		jsonData, err :=daily(getRecord)
		if err != nil {
			fmt.Println("Error:", err)
			
		}













		return c.Render("calender", fiber.Map{
			"Calender":getRecord,
			"cal":string(jsonData),
		 })
    })
	app.Post("/", func(c *fiber.Ctx) error {
		cals := []Calender{}
		
		selectCase:=c.FormValue("selectcase[0]")
		fmt.Println(selectCase)
	
		if (selectCase=="on"){
			selectCase="1"
		}else{
			selectCase="2"
		}
		
		input_text,_:=strconv.Atoi(c.FormValue("daily[0]"))
		if(input_text==0){
		input_text,_=strconv.Atoi(c.FormValue("daily[1]"))

		}
		title_text:=c.FormValue("title")
		week:=c.FormValue("weekends")
		start_date:=c.FormValue("start_date")
		end_date:=c.FormValue("end_date")
		mode:=c.FormValue("mySelectName")
		day:=c.FormValue("mySelect[0]")
		if (day==""){
			day=c.FormValue("mySelect[1]")
		}
	
		
		res,error := pool.Exec(context.Background(), `INSERT INTO calender(occurence,title,start_date,end_date,mode,day,selectcase,week)VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,input_text,title_text,start_date,end_date,mode,day,selectCase,week)
		fmt.Println(res)
		fmt.Println(error)
		// rows,_:=pool.Query(context.Background(),`SELECT occurence,title,TO_CHAR(start_date, 'YYYY-MM-DD') ,TO_CHAR(end_date, 'YYYY-MM-DD'),mode ,day,selectcase,week FROM calender`)
		rows,_:=pool.Query(context.Background(),`SELECT occurence,title,start_date ,end_date,mode ,day,selectcase,week FROM calender`)

	
		for rows.Next() {
		
			var cal Calender
			
			// fmt.Println("kake")
			err:= rows.Scan(&cal.Occurence, &cal.Title,&cal.Start_date,&cal.End_date,&cal.Mode,&cal.Day,&cal.Case,&cal.Week)
			
		if err != nil {		
			return err
		}
		cals = append(cals, cal)	
	
		}	
		return c.Render("calender", fiber.Map{
			"Calender":cals,
		 })
	  })


	
	log.Fatal(app.Listen(":3006"),)
}
func daily(calendars []Calender)([]byte, error){
	events := []map[string]interface{}{}

	for i := 0; i < len(calendars); i++ {
		date:=calendars[i].Start_date
		startDate := date.Unix()
		unixTimestamp := int64(startDate)
		newdate := time.Unix(unixTimestamp, 0)
		endDate:=calendars[i].End_date
		endDate_2 := endDate.Unix()
		unixTimestampend := int64(endDate_2)
		newenddate := time.Unix(unixTimestampend, 0)

		for newdate.Before(newenddate)   || newdate.Equal(newenddate) {

            events = append(events, map[string]interface{}{
                "title": calendars[i].Title,
                "start": newdate.Format("2006-01-02"), 
                
            })
            newdate=newdate.AddDate(0,0,calendars[i].Occurence) // Add a 3-day gap for the next event
			// fmt.Println(events  ,"hello khoob chand jhariya")
        }
  
		

		
		
		

	}
	jsonData, err := json.Marshal(events)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		return jsonData,nil
}
// func getDaily() {
//     if mode == "daily" {
//         startDate := time.Unix(last, 0) // Replace with your start date
//         endDate := time.Unix(end, 0)   // Replace with your end date
//         events := []map[string]interface{}{}
//         currentDate := startDate
//         for currentDate.Before(endDate) {
//             eventEndDate := currentDate
//             eventEndDate = eventEndDate.AddDate(0, 0, 1) // Event ends 1 day after start
//             fmt.Println(eventEndDate)
//             events = append(events, map[string]interface{}{
//                 "title": "{{$title}}",
//                 "start": currentDate.Format("2006-01-02"), // Format as 'YYYY-MM-DD'
//                 "end":   eventEndDate.Format("2006-01-02"),
//             })
//             currentDate = currentDate.AddDate(0, 0, gap-1) // Add a 3-day gap for the next event
//         }
//         successCallback(events)
//     }
// }