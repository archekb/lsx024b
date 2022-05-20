package device

import (
	"encoding/binary"
	"reflect"
	"strconv"
	"strings"

	"github.com/archekb/lsx024b/internal/log"
)

func Decode(data []byte) []uint16 {
	values := make([]uint16, len(data)/2)
	for i := 0; i < len(values); i++ {
		values[i] = binary.BigEndian.Uint16(data[i*2 : i*2+2])
	}

	return values
}

func Encode(data []uint16) []byte {
	values := make([]byte, len(data)*2)
	for i := 0; i < len(data); i++ {
		binary.BigEndian.PutUint16(values[i*2:i*2+2], data[i])
	}

	return values
}

func FillStruct(s interface{}, reader func(uint16, uint16) ([]byte, error), debug ...bool) error {
	refElem := reflect.ValueOf(s).Elem()

	t := refElem.Type()
	var paddr uint64
	var res []byte
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !refElem.FieldByName(field.Name).IsValid() {
			continue
		}

		// get field addr
		addr, err := strconv.ParseUint(field.Tag.Get("addr")[2:], 16, 16)
		if err != nil {
			log.Printf("%s: Parse addr (%s) err: %s", field.Name, field.Tag.Get("addr"), err)
			continue
		}

		// if addr not changed, do not read data again
		if addr != paddr {
			res, err = reader(uint16(addr), 0x01)
			if err != nil {
				log.Printf("%s: Read addr (%s) err: %s", field.Name, field.Tag.Get("addr"), err)
				continue
			}
			paddr = addr
		}

		var addrValue uint16

		if field.Tag.Get("mode") == "raw" {
			addrValue = uint16(res[0])
		} else {
			decoded := Decode(res)
			if len(decoded) < 1 {
				log.Printf("%s: Readed addr (%s) is empty", field.Name, field.Tag.Get("addr"))
				continue
			}
			addrValue = decoded[0]
		}

		var value float64

		// if we have a low bits of value, get it
		slAddr := field.Tag.Get("laddr")
		if slAddr != "" {
			lAddr, err := strconv.ParseUint(slAddr[2:], 16, 16)
			if err != nil {
				log.Println(field.Name, "Parse low bit addr err:", err)
				continue
			}

			res, err := reader(uint16(lAddr), 0x01)
			if err != nil {
				log.Println(field.Name, "Read low bit addr err:", err)
				continue
			}

			decoded := Decode(res)
			if len(decoded) < 1 {
				log.Println(field.Name, "Readed low bit addr is empty")
				continue
			}
			lAddrValue := decoded[0]

			value = float64(int(addrValue)<<16 | int(lAddrValue))
		} else {
			value = float64(addrValue)
		}

		if len(debug) > 0 && debug[0] {
			log.Println(value, Encode([]uint16{addrValue}))
		}

		// convert value to bits
		sBits := field.Tag.Get("bits")
		if sBits != "" {
			sBitsMSK, err := strconv.ParseInt(sBits, 2, 32)
			if err != nil {
				log.Printf("%s: Addr (%s) can't parse mask %s", field.Name, field.Tag.Get("addr"), sBits)
				continue
			}

			var shift int
			for i := len(sBits); i > 0; i-- {
				if sBits[i-1] == '0' {
					shift += 1
				} else {
					break
				}
			}

			if len(debug) > 0 && debug[0] {
				log.Println(field.Name, field.Tag.Get("addr"), sBits, value, shift, int64(value), sBitsMSK, (int64(value) & sBitsMSK), (int64(value)&sBitsMSK)>>shift)
			}
			value = float64((int64(value) & sBitsMSK) >> shift)
		}

		// convert value to enum
		sEnum := field.Tag.Get("enum")
		if sEnum != "" && field.Type.Kind() == reflect.String {
			ssEnum := strings.Split(sEnum, ",")
			for _, enum := range ssEnum {
				e := strings.Split(enum, ":")

				i, err := strconv.ParseInt(strings.TrimSpace(e[0]), 16, 16)
				if err != nil {
					log.Println(err)
					continue
				}

				if int64(value) == i && len(e) > 1 {
					refElem.FieldByName(field.Name).SetString(e[1])
					break
				}
			}
		}

		// write value to float64 field
		if field.Type.Kind() == reflect.Float64 {
			// if need divide
			sDivide := field.Tag.Get("divide")
			if len(sDivide) > 0 {
				delim, err := strconv.ParseFloat(sDivide, 64)
				if err != nil {
					return err
				}
				refElem.FieldByName(field.Name).SetFloat(value / delim)
			} else {
				refElem.FieldByName(field.Name).SetFloat(value)
			}
		}

		// write value to int field
		if field.Type.Kind() == reflect.Int {
			// if need divide
			sDivide := field.Tag.Get("divide")
			if len(sDivide) > 0 {
				delim, err := strconv.ParseFloat(sDivide, 64)
				if err != nil {
					return err
				}
				refElem.FieldByName(field.Name).SetInt(int64(value / delim))
			} else {
				refElem.FieldByName(field.Name).SetInt(int64(value))
			}
		}

		// write value to bool field
		if field.Type.Kind() == reflect.Bool {
			refElem.FieldByName(field.Name).SetBool(int(value) > 0)
		}
	}

	return nil
}
