package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	logger "github.com/sirupsen/logrus"
	"go_cache_example/src/model"
	"go_cache_example/src/repo"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CryptoService struct {
	mycache *cache.Cache
	repo    repo.CryptoRepo
}

func initCache() *cache.Cache {
	return cache.New(5*time.Minute, 10*time.Minute)
}

func CreateCryptoService(cryptoRepo repo.CryptoRepo) *CryptoService {
	return &CryptoService{initCache(), cryptoRepo}
}

//Get data by referencing their id (primary key). If the data has already been
//stored in cache then fetch data from cache.
func (cs *CryptoService) GetByID(c *gin.Context) {
	//Take start time of handler function call
	start := time.Now()
	//Bind POST Json
	var findid model.CryptoQuery
	err := c.BindJSON(&findid)
	if err != nil {
		log.Println(err)
		return
	}
	//The key for our cache
	id := fmt.Sprintf("%s:%d", "ID", findid.ID)
	//Get item from cache using above's key
	b, ok := cs.mycache.Get(id)
	if !ok {
		log.Println("Cannot get cache on", id)
	}
	//If item exist, then return it immediately
	if b != nil {
		log.Println("Get data for", id, "from cache")
		HandleSuccessWithData(c, b, id, "Success GetID", start)
		return
	}

	//If item in cache does not exist yet, then fetch data from DB
	data, err := cs.repo.GetByID(findid.ID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Cannot find data",
		})
		return
	}

	//Set cache with our key and the fetched data from DB, and set it to
	//be available for the next 2 minute
	cs.mycache.Set(id, data, 2*time.Minute)
	log.Println("Cache saved on ", id)
	HandleSuccessWithData(c, data, id, "Success GetID", start)
}

func (cs *CryptoService) GetAll(c *gin.Context) {
	start := time.Now()
	keyAll := "GetAll"
	b, ok := cs.mycache.Get(keyAll)
	if !ok {
		log.Println("Cannot get cache on", "GetAll")
	}

	if b != nil {
		log.Println("Get data for", keyAll, "from cache")
		HandleSuccessWithData(c, b, keyAll, "Success GetAll", start)
		return
	}

	datas, err := cs.repo.GetAll()
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
		return
	}

	cs.mycache.Set("GetAll", datas, 2*time.Minute)
	log.Println("Cache saved on GetAll")
	HandleSuccessWithData(c, datas, "GetAll", "Success GetAll", start)
}

//with data
func HandleSuccessWithData(c *gin.Context, data interface{}, logdata interface{}, logmsg string, start time.Time) {
	var returnData = model.ResponseBody{
		Success: true,
		Message: "Success",
		Data:    data,
	}
	logger.WithFields(logger.Fields{
		"detail":  logmsg,
		"data":    logdata,
		"elapsed": strconv.FormatInt(time.Since(start).Milliseconds(), 10) + "ms",
	}).Info(logmsg)
	c.JSON(http.StatusOK, returnData)

}
