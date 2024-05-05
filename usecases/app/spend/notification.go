package usecases_spending

import (
	"fmt"
	"time"
)

func (m MySpending) HourlyNotification(category string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		m.SendAlertEmail(category)
		fmt.Println("Printed this hour")
	}
}

func (m MySpending) DailyNotification(hour, min, sec int, category string) {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
	if next.Before(now) {
		next = next.Add(24 * time.Hour)
	}
	fmt.Printf("Next daily print will be at %s\n", next)
	time.Sleep(next.Sub(now))
	m.SendAlertEmail(category)
	fmt.Println("Printed daily at the specified time")
}

func (m MySpending) WeeklyNotification(weekday time.Weekday, hour, min, sec int, category string) {
	now := time.Now()
	daysUntilNextWeekday := int((weekday - now.Weekday() + 7) % 7)
	next := time.Date(now.Year(), now.Month(), now.Day()+daysUntilNextWeekday, hour, min, sec, 0, now.Location())
	fmt.Printf("Next weekly print will be on %s at %s\n", weekday, next)
	time.Sleep(next.Sub(now))
	m.SendAlertEmail(category)
	fmt.Println("Printed weekly on the specified day and time")
}

func (m MySpending) MonthlyNotification(day, hour, min, sec int, category string) {
	now := time.Now()
	year, month, _ := now.Date()
	next := time.Date(year, month, day, hour, min, sec, 0, now.Location())
	if next.Before(now) {
		next = next.AddDate(0, 1, 0)
	}
	fmt.Printf("Next monthly print will be on the %dth day at %s\n", day, next)
	time.Sleep(next.Sub(now))
	m.SendAlertEmail(category)
	fmt.Println("Printed monthly on the specified day and time")
}
