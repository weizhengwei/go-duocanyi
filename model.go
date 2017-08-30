package main

import (
)

type RetItem struct {
	RET_CODE string
	SERVICE_CODE string
	RET_MSG string
}

type SysHead struct {
	RET []RetItem
	RET_STATUS string
}

type Result struct {
	SYS_HEAD SysHead
	BODY interface{}
}

