package user

import (
	"time"

	. "github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/constants"
	. "github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/util"
	"github.com/MihaiLupoiu/money-transferring-simulation/backend/models"
	"github.com/gin-gonic/gin"
)

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
