/*
 * @Author: feng.pan@trip.com
 * @Date: 2022-03-27 01:40:00
 * @LastEditors: feng.pan
 * @LastEditTime: 2022-08-08 00:16:52
 * @FilePath: /pain-utils/perr/error.go
 * @Description: todo...
 */
package perr

import "fmt"

func PanicErrDouble[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintErr(err error) {
	if err != nil {
		fmt.Printf("err: %s", err)
	}
}
