package main

import (
    "github.com/gin-gonic/gin"
    "github.com/miko2009/gogogo_follow_me/lib/config"
    "github.com/miko2009/gogogo_follow_me/lib/mysql"
    "github.com/miko2009/gogogo_follow_me/middleware"
    v1Router "github.com/miko2009/gogogo_follow_me/routers/v1"
)

func main() {
    err := Init()
    if err != nil {
        panic(err)
    }
    router := gin.New()

    router.Use(middleware.ErrHandler())
    v1Router.RegisterMusicRouter(router)
    router.Run(config.GetPort())
}

func Init() error {
    var err error

    //初始化配置
    err = config.InitConfig()
    if err != nil {
        return err
    }

    //初始化数据库
    err = mysql.InitMysqlPool()
    if err != nil {
        return err
    }

    return nil
}