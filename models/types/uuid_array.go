package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

// digunakan untuk mengscan ke struct
func (a *UUIDArray) Scan(value any) error{
	// bentuk data yg akan diproses {212232adasfd, dadfdaf, 234421} // diaggap string
	var str string

	switch v := value.(type){
		case []byte : 
			str = string(v)
		case string:
			str = v
		default:
			return errors.New("File to parse UUIDArray: unsupport data type")
	}

	// buat kurung kurawal depan dan belakang {}
	str = strings.TrimPrefix(str,"{") // buat {
	str = strings.TrimSuffix(str,"}") // buat } akhir
	parts := strings.Split(str, ",") // split/pisah data berdasarkan koma

	// make([]T, lengt, capacity)
	*a = make(UUIDArray,0,len(parts))
	for _ , s := range parts{
		s = strings.TrimSpace(strings.Trim(s,`"`)) // menghapus spasi dan kutip
		if s == ""{
			continue
		}
		u, err := uuid.Parse(s)
		if err != nil{
			return fmt.Errorf("Invalid UUID in Array: %v", err)
		} else{
			*a = append(*a, u)
		}
	}
	return nil
}

// berguna untuk mengscan dari struk ke postgres (agar terbaca sesuai format)
func (a UUIDArray) Value() (driver.Value, error){
	if len(a) == 0 {
		return "{}", nil
	}

	postgreFormat := make([]string,0,len(a))
	for _, value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s"`, value.String())) // sama seperti yang dibawah
		// postgreFormat = append(postgreFormat, `"` + value.String() +`"` )
	}
	return "{"+strings.Join(postgreFormat,",")+"}" , nil
}

// untuk memberitahukan gorm
func (UUIDArray) GormDataType() string{ //dipanggil otomatis saat buat skema atau data
	return "uuid[]"
}