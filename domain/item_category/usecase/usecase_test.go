package usecase_test

import (
	"context"
	"time"
)

var ctx = context.Background()
var timeoutContext = time.Duration(5) * time.Second

var createdAtString string = "2022-08-13T04:06:16.312789Z"
var updatedAtString string = "2022-09-14T05:07:17.463431Z"
var deletedAtString string = "2022-10-15T04:08:18.538807Z"