package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func BoolPtrToSQLNullBool(v *bool) sql.NullBool {
	if v == nil {
		return sql.NullBool{
			Valid: false,
		}
	}
	return sql.NullBool{
		Valid: true,
		Bool:  *v,
	}
}

func SQLNullBoolToBoolPtr(ts sql.NullBool) *bool {
	if ts.Valid {
		return &ts.Bool
	}
	return nil
}

func SQLNullBoolToBool(ts sql.NullBool) bool {
	if ts.Valid {
		return ts.Bool
	}
	return false
}

func StringToSQLNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		Valid:  true,
		String: v,
	}
}

func StringPtrToSQLNullString(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		Valid:  true,
		String: *v,
	}
}

func SQLNullStringToStringPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func IntPtrToSQLNullInt32(v *int) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{
			Valid: false,
		}
	}
	return sql.NullInt32{
		Valid: true,
		Int32: int32(*v),
	}
}

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

func SQLNullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func SQLNullInt32ToInt(ni sql.NullInt32) int {
	if ni.Valid {
		return int(ni.Int32)
	}
	return 0
}

func IntToSQLNullInt32(i int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(i), Valid: i != 0}
}

func TimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func SQLNullUUIDToUUID(nu uuid.NullUUID) uuid.UUID {
	if nu.Valid {
		return nu.UUID
	}
	return uuid.Nil
}
