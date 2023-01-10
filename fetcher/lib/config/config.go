package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	FetcherApiToken  string
	FetcherApiListen string
	FetcherApiDebug  string

	BTDownloadDir string

	TransmissionRpcHost   string
	TransmissionRpcHttps  bool
	TransmissionRpcPort   uint16
	TransmissionRpcUri    string
	TransmissionRpcUser   string
	TransmissionRpcPasswd string

	RcloneConfigPath string
	RcloneDriver     string
	RcloneRootDir    string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	FetcherApiToken = os.Getenv("FETCHER_API_TOKEN")
	FetcherApiListen = os.Getenv("FETCHER_API_LISTEN")
	fetcherApiDebugBool, err := strconv.ParseBool(os.Getenv("FETCHER_API_DEBUG"))
	if err != nil {
		log.Fatalln(err)
	}
	if fetcherApiDebugBool {
		FetcherApiDebug = "debug"
	} else {
		FetcherApiDebug = "release"
	}

	BTDownloadDir = os.Getenv("BT_DOWNLOAD_DIR")

	TransmissionRpcHost = os.Getenv("TRANSMISSION_RPC_HOST")
	TransmissionRpcHttps, err = strconv.ParseBool(os.Getenv("TRANSMISSION_RPC_HTTPS"))
	if err != nil {
		log.Fatalln(err)
	}
	transmissionRpcPortInt, err := strconv.Atoi(os.Getenv("TRANSMISSION_RPC_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	TransmissionRpcPort = uint16(transmissionRpcPortInt)
	TransmissionRpcUri = os.Getenv("TRANSMISSION_RPC_URI")
	TransmissionRpcUser = os.Getenv("TRANSMISSION_RPC_USER")
	TransmissionRpcPasswd = os.Getenv("TRANSMISSION_RPC_PASSWD")

	RcloneConfigPath = os.Getenv("RCLONE_CONFIG_PATH")
	RcloneDriver = os.Getenv("RCLONE_DRIVER")
	RcloneRootDir = os.Getenv("RCLONE_ROOT_DIR")
}
