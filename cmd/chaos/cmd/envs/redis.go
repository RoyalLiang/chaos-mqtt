package envs

import (
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rAddress string
	redisPwd string
	redisDB  int
)

var RedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "配置redis",
	Run: func(cmd *cobra.Command, args []string) {
		if rAddress == "" && redisPwd == "" {
			_ = cmd.Help()
		} else {
			configs.Chaos.Redis = &configs.RedisConfig{
				Address:  rAddress,
				Password: redisPwd,
				DB:       redisDB,
			}

			if err := configs.WriteFMSConfig("redis", configs.Chaos.Redis); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println("Redis 配置成功...")
		}
	},
}

func init() {
	RedisCmd.Flags().StringVarP(&rAddress, "address", "a", "", "Redis Host (127.0.0.1:6379)")
	RedisCmd.Flags().StringVar(&redisPwd, "password", "", "Redis Password")
	RedisCmd.Flags().IntVarP(&redisDB, "db", "d", 0, "Redis DB")
}
