package user

import (
	"time"

	. "github.com/MihaiLupoiu/money-transferring-simulation/libs/constants"
	. "github.com/MihaiLupoiu/money-transferring-simulation/libs/util"
	"github.com/MihaiLupoiu/money-transferring-simulation/models"
	"github.com/gin-gonic/gin"
)

// TODO: Split DB queries and JSON return values.

func Post(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" && user.Mail != "" {
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": ErrorFieldsEmpty})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"mail\": \"jd@fake.com\"}" http://localhost:8080/api/v1/users
}

func GetUsers(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var users []models.User
	db.Find(&users)
	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func Update(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" && user.Mail != "" {
		if user.Id != 0 {
			var newUser models.User
			c.Bind(&newUser)

			result := models.User{
				Id:        user.Id,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
				Mail:      newUser.Mail,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": ErrorUserNotFound})
		}
	} else {
		c.JSON(422, gin.H{"error": ErrorFieldsEmpty})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\", \"mail\": \"tm@fake.com\" }" http://localhost:8080/api/v1/users/1
}

func Delete(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		db.Delete(&user)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func GetBalance(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		c.JSON(200, user.Balance)
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i http://localhost:8080/api/v1/balance/1
}

func AddDeposit(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		var deposit models.Deposit
		c.Bind(&deposit)

		if deposit.Amount > 0 {
			if deposit.Date == (time.Time{}) {
				deposit.Date = time.Now()
			}
			user.Balance += deposit.Amount
			db.Save(&deposit)
			db.Save(&user)

			c.JSON(200, gin.H{"success": user})
		} else {
			c.JSON(404, gin.H{"error": "Transfer amount not valid"})
		}
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"amount\": 100 }" http://localhost:8080/api/v1/deposit/1
}

func Transfer(c *gin.Context) {

	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var recvUser models.User
	db.First(&recvUser, id)

	if recvUser.Id != 0 {
		var transfer models.Transfer
		c.Bind(&transfer)
		if transfer.SenderId != 0 && transfer.Amount > 0 {
			var sender models.User
			db.First(&sender, transfer.SenderId)
			if sender.Id != 0 {
				if sender.Id == recvUser.Id {
					c.JSON(404, gin.H{"error": ErrorSenderUserSameAsReceiverUser})
					return
				}
				// TODO: Make atomic operation
				if sender.Balance >= transfer.Amount {
					sender.Balance -= transfer.Amount
					recvUser.Balance += transfer.Amount

					db.Save(&transfer)
					db.Save(&sender)
					db.Save(&recvUser)

					c.JSON(200, gin.H{"success": recvUser})

				} else {
					c.JSON(404, gin.H{"error": ErrorInsufficientFounds})
				}
			} else {
				c.JSON(404, gin.H{"error": ErrorSenderUserNotFound})
			}
		} else {
			c.JSON(422, gin.H{"error": ErrorFieldsEmpty})
		}
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"senderid\": 1, \"amount\": 10 }" http://localhost:8080/api/v1/transfer/1

}
