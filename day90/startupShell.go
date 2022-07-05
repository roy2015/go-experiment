package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

//执行系统命令
func execCmdFunc(cmdWithPath string, cmdArgs []string) {
	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).
	binary, lookErr := exec.LookPath(cmdWithPath)
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` requires arguments in slice form (as
	// opposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	env := os.Environ()

	// Here's the actual `syscall.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	execErr := syscall.Exec(binary, cmdArgs, env)
	if execErr != nil {
		fmt.Println(execErr)
		panic(execErr)
	}
}

//startRedisServer
func mysqlServerStart() {
	mysqlServerStartCmd := flag.NewFlagSet("mysqlServerStart", flag.ExitOnError)
	mysqlServerStartCmd.Parse(os.Args[2:])
	cmdStr := "/usr/local/opt/mysql@5.7/bin/mysql.server"
	execCmdFunc(cmdStr, []string{cmdStr, "start"})
}

//startRedisServer
func mysqlServerStop() {
	mysqlServerStopCmd := flag.NewFlagSet("mysqlServerStop", flag.ExitOnError)
	mysqlServerStopCmd.Parse(os.Args[2:])
	cmdStr := "/usr/local/opt/mysql@5.7/bin/mysql.server"
	execCmdFunc(cmdStr, []string{cmdStr, "stop"})
}

//startRedisServer
func startRedisServer() {
	redisServerCmd := flag.NewFlagSet("redisServer", flag.ExitOnError)
	redisServerCmd.Parse(os.Args[2:])
	cmdStr := "/usr/local/opt/redis/bin/redis-server"
	execCmdFunc(cmdStr, []string{cmdStr, "/usr/local/etc/redis.conf"})
}

// kibana
func doKibana() {
	var cmdStr string
	kibanaCmd := flag.NewFlagSet("kibana", flag.ExitOnError)
	kibanaVersion := kibanaCmd.String("v", "7.6.2", "kibana version: 7.6.2 6.3.2")
	kibanaCmd.Parse(os.Args[2:])
	fmt.Println("kibana version is ", *kibanaVersion)
	switch *kibanaVersion {
	case "7.6.2":
		cmdStr = "/Users/apple/tools/kibana-7.6.2-darwin-x86_64/bin/kibana"
		execCmdFunc(cmdStr, []string{cmdStr})
		break
	case "6.3.2":
		cmdStr = "/Users/apple/tools/kibana-6.3.2-darwin-x86_64/bin/kibana"
		execCmdFunc(cmdStr, []string{cmdStr})
		break
	default:
		fmt.Println("version error, select 7.6.2 or 6.3.2")
		os.Exit(1)
	}
}

//redisCli
func doRedisCli() {
	var cmdStr string
	redisCliCmd := flag.NewFlagSet("redisCli", flag.ExitOnError)
	redisCliIp := redisCliCmd.String("h", "127.0.0.1", "redis ip")
	redisCliPort := redisCliCmd.Int("p", 6379, "redis ip")
	redisCliAuth := redisCliCmd.String("a", "foobar", "password")
	redisCliCluster := redisCliCmd.String("c", "", "cluster")
	cmdStr = "/usr/local/opt/redis/bin/redis-cli"
	redisCliCmd.Parse(os.Args[2:])
	fmt.Println("redis ip", *redisCliIp)
	fmt.Println("redis port", *redisCliPort)
	fmt.Println("redis auth", *redisCliAuth)
	fmt.Println("redis cluster", *redisCliCluster)
	fmt.Println("redis no-flag args", redisCliCmd.Args())
	cmdArgs := []string{cmdStr, "-h", *redisCliIp, "-p", strconv.Itoa(*redisCliPort), "-a", *redisCliAuth, "-c"}
	for _, s := range redisCliCmd.Args() {
		cmdArgs = append(cmdArgs, s)
	}
	//args参数要包含命名本身
	execCmdFunc(cmdStr, cmdArgs)
}

func main1() {

	// The subcommand is expected as the first argument
	// to the program.
	expectSubcommands := []string{"kibana", "redisCli", "redisServer", "mysqlServerStop", "mysqlServerStart"}
	if len(os.Args) < 2 {
		fmt.Println("expected  subcommands above :\n", expectSubcommands)
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	switch os.Args[1] {
	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	case "kibana":
		doKibana()
		break
	case "redisCli":
		doRedisCli()
	case "redisServer":
		startRedisServer()
	case "mysqlServerStop":
		mysqlServerStop()
	case "mysqlServerStart":
		mysqlServerStart()
	default:
		fmt.Println("expected  subcommands above :\n", expectSubcommands)
		os.Exit(1)
	}
}
