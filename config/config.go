package config

import (
	"encoding/json"
	"github.com/giantliao/beatles-protocol/licenses"
	"github.com/giantliao/beatles-protocol/miners"
	"github.com/kprc/libeth/account"
	"github.com/kprc/nbsnetwork/tools"
	"log"
	"os"
	"path"
	"strconv"
	"sync"
)

const (
	BTLC_HomeDir      = ".btlclient"
	BTLC_CFG_FileName = "btlclient.json"
)

type ClientPrice struct {
	Sig          *licenses.NoncePriceSig `json:"sig"`
	Gas          float64                 `json:"gas"`
	EstimatedFee float64                 `json:"estimated_fee"`
}

type BtlClientConf struct {
	BeatlesMasterAddr account.BeatleAddress `json:"beatles_master_addr"`
	BeatlesEthAddr    string                `json:"beatles_eth_addr"`
	BeatlesTrxAddr    string                `json:"beatles_trx_addr"`
	EthAccPoint       string                `json:"eth_acc_point"`
	TrxAccPoint       string                `json:"trx_acc_point"`

	CmdListenPort       string `json:"cmdlistenport"`
	HttpServerPort      int    `json:"http_server_port"`
	StreamServerPort    int    `json:"stream_server_port"`
	StreamServerPacPort int    `json:"stream_server_pac_port"`
	WalletSavePath      string `json:"wallet_save_path"`
	LicenseSavePath     string `json:"license_save_path"`
	TransactionSavePath string `json:"transaction_save_path"`
	ScriptSavePath      string `json:"script_save_path"`

	ApiPath        string            `json:"api_path"`
	NoncePricePath string            `json:"nonce_price"`
	PurchasePath   string            `json:"purchase_path"`
	ListMinerPath  string            `json:"list_miner_path"`
	EthBalance     float64           `json:"-"`
	TrxBalance     float64           `json:"-"`
	MemLicense     *licenses.License `json:"-"`
	MemPrice       *ClientPrice      `json:"-"`
	CurrentMiner   int               `json:"current_miner"`
	VPNMode        int               `json:"vpn_mode"` //1 global, 0 pac

	GithubAddress []*miners.GithubDownLoadPoint `json:"github_address"`
	Miners        []*miners.Miner               `json:"miners"`
}

var (
	btlccfgInst     *BtlClientConf
	btlccfgInstLock sync.Mutex
)

func (bc *BtlClientConf) InitCfg() *BtlClientConf {
	bc.HttpServerPort = 50102
	bc.StreamServerPacPort = 50211
	bc.StreamServerPort = 50212
	bc.CmdListenPort = "127.0.0.1:50502"
	bc.WalletSavePath = "wallet.json"
	bc.LicenseSavePath = "license.db"
	bc.TransactionSavePath = "tx.db"
	bc.ScriptSavePath = "gfw"

	bc.ApiPath = "api"
	bc.NoncePricePath = "price"
	bc.PurchasePath = "purchase"
	bc.ListMinerPath = "list"

	gd := &miners.GithubDownLoadPoint{}
	gd.Path = "beatles.bootstrap"
	gd.Repository = "beatleslist"
	gd.Owner = "youpipe001"
	gd.ReadToken = "at34KdLw4XkwffHVECnD1UcU1PjW1Y"

	bc.GithubAddress = append(bc.GithubAddress, gd)

	return bc
}

func (bc *BtlClientConf) Load() *BtlClientConf {
	if !tools.FileExists(GetBtlcCFGFile()) {
		return nil
	}

	jbytes, err := tools.OpenAndReadAll(GetBtlcCFGFile())
	if err != nil {
		log.Println("load file failed", err)
		return nil
	}

	err = json.Unmarshal(jbytes, bc)
	if err != nil {
		log.Println("load configuration unmarshal failed", err)
		return nil
	}

	return bc

}

func newBtlcCfg() *BtlClientConf {

	bc := &BtlClientConf{}

	bc.InitCfg()

	return bc
}

func GetCBtlc() *BtlClientConf {
	if btlccfgInst == nil {
		btlccfgInstLock.Lock()
		defer btlccfgInstLock.Unlock()
		if btlccfgInst == nil {
			btlccfgInst = newBtlcCfg()
		}
	}

	return btlccfgInst
}

func PreLoad() *BtlClientConf {
	bc := &BtlClientConf{}

	return bc.Load()
}

func LoadFromCfgFile(file string) *BtlClientConf {
	bc := &BtlClientConf{}

	bc.InitCfg()

	bcontent, err := tools.OpenAndReadAll(file)
	if err != nil {
		log.Fatal("Load Config file failed")
		return nil
	}

	err = json.Unmarshal(bcontent, bc)
	if err != nil {
		log.Fatal("Load Config From json failed")
		return nil
	}

	btlccfgInstLock.Lock()
	defer btlccfgInstLock.Unlock()
	btlccfgInst = bc

	return bc

}

func LoadFromCmd(initfromcmd func(cmdbc *BtlClientConf) *BtlClientConf) *BtlClientConf {
	btlccfgInstLock.Lock()
	defer btlccfgInstLock.Unlock()

	lbc := newBtlcCfg().Load()

	if lbc != nil {
		btlccfgInst = lbc
	} else {
		lbc = newBtlcCfg()
	}

	btlccfgInst = initfromcmd(lbc)

	return btlccfgInst
}

func GetBtlcHomeDir() string {
	curHome, err := tools.Home()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(curHome, BTLC_HomeDir)
}

func GetBtlcCFGFile() string {
	return path.Join(GetBtlcHomeDir(), BTLC_CFG_FileName)
}

func (bc *BtlClientConf) Save() {
	jbytes, err := json.MarshalIndent(*bc, " ", "\t")

	if err != nil {
		log.Println("Save BASD Configuration json marshal failed", err)
	}

	if !tools.FileExists(GetBtlcHomeDir()) {
		os.MkdirAll(GetBtlcHomeDir(), 0755)
	}

	err = tools.Save2File(jbytes, GetBtlcCFGFile())
	if err != nil {
		log.Println("Save BASD Configuration to file failed", err)
	}

}

func (bc *BtlClientConf) GetNoncePriceWebPath(ip string, port int) string {
	url := "http://" + ip
	url += ":" + strconv.Itoa(port)

	pricePath := path.Join(bc.ApiPath, bc.NoncePricePath)

	url += "/" + pricePath

	return url

}

func (bc *BtlClientConf) GetWalletSavePath() string {
	return path.Join(GetBtlcHomeDir(), bc.WalletSavePath)
}

func (bc *BtlClientConf) GetPurchasePath(ip string, port int) string {
	url := "http://" + ip
	url += ":" + strconv.Itoa(port)

	pricePath := path.Join(bc.ApiPath, bc.PurchasePath)

	url += "/" + pricePath

	return url
}

func (bc *BtlClientConf) GetListMinerPath(ip string, port int) string {
	url := "http://" + ip
	url += ":" + strconv.Itoa(port)

	pricePath := path.Join(bc.ApiPath, bc.ListMinerPath)

	url += "/" + pricePath

	return url
}

func (bc *BtlClientConf) GetLicenseDBPath() string {
	return path.Join(GetBtlcHomeDir(), bc.LicenseSavePath)
}

func (bc *BtlClientConf) GetTransactionDBPath() string {
	return path.Join(GetBtlcHomeDir(), bc.TransactionSavePath)
}

func (bc *BtlClientConf) GetScriptPath() string {
	sp := path.Join(GetBtlcHomeDir(), bc.ScriptSavePath)
	if !tools.FileExists(sp) {
		os.MkdirAll(sp, 0755)
	}

	return sp
}

func IsInitialized() bool {
	if tools.FileExists(GetBtlcCFGFile()) {
		return true
	}

	return false
}
