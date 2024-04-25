package database

import (
	"database/sql"
	"time"
)

func TimePtrToSQLNullTime(ts *time.Time) sql.NullTime {
	if ts == nil {
		return sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}
	return sql.NullTime{
		Time:  *ts,
		Valid: true,
	}
}

func SQLNullTimeToTimePtr(ts sql.NullTime) *time.Time {
	if ts.Valid {
		return &ts.Time
	}
	return nil
}

func SQLNullTimeToTime(ts sql.NullTime) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func TimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}
