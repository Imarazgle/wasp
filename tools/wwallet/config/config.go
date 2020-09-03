package config

import (
	"fmt"
	"os"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/nodeclient"
	"github.com/iotaledger/wasp/packages/nodeclient/goshimmer"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configPath string
var Verbose bool
var WaitForConfirmation bool
var Utxodb bool

const (
	hostKindApi     = "api"
	hostKindPeering = "peering"
	hostKindNanomsg = "nanomsg"
)

func InitCommands(commands map[string]func([]string), flags *pflag.FlagSet) {
	commands["set"] = setCmd

	fs := pflag.NewFlagSet("config", pflag.ExitOnError)
	fs.StringVarP(&configPath, "config", "c", "wwallet.json", "path to wwallet.json")
	fs.BoolVarP(&Verbose, "verbose", "v", false, "verbose")
	fs.BoolVarP(&WaitForConfirmation, "wait", "w", false, "wait for confirmation")
	fs.BoolVarP(&Utxodb, "utxodb", "u", false, "use utxodb")
	flags.AddFlagSet(fs)
}

func setCmd(args []string) {
	if len(args) != 2 {
		fmt.Printf("Usage: %s set <key> <value>\n", os.Args[0])
		os.Exit(1)
	}
	Set(args[0], args[1])
}

func Read() {
	viper.SetConfigFile(configPath)
	_ = viper.ReadInConfig()
}

func GoshimmerApi() string {
	r := viper.GetString("goshimmer." + hostKindApi)
	if r != "" {
		return r
	}
	return "127.0.0.1:8080"
}

func GoshimmerClient() nodeclient.NodeClient {
	if Utxodb {
		return testutil.NewGoshimmerUtxodbClient(GoshimmerApi())
	}
	return goshimmer.NewGoshimmerClient(GoshimmerApi())
}

func WaspApi() string {
	r := viper.GetString("wasp." + hostKindApi)
	if r != "" {
		return r
	}
	return committeeHost(hostKindApi, 0)
}

func WaspNanomsg() string {
	r := viper.GetString("wasp." + hostKindApi)
	if r != "" {
		return r
	}
	return committeeHost(hostKindNanomsg, 0)
}

func CommitteeApi(indices []int) []string {
	return committee(hostKindApi, indices)
}

func CommitteePeering(indices []int) []string {
	return committee(hostKindPeering, indices)
}

func CommitteeNanomsg(indices []int) []string {
	return committee(hostKindNanomsg, indices)
}

func committee(kind string, indices []int) []string {
	hosts := make([]string, 0)
	for _, i := range indices {
		hosts = append(hosts, committeeHost(kind, i))
	}
	return hosts
}

func committeeHost(kind string, i int) string {
	r := viper.GetString(fmt.Sprintf("wasp.%d.%s", i, kind))
	if r != "" {
		return r
	}
	defaultPort := defaultWaspPort(kind, i)
	return fmt.Sprintf("127.0.0.1:%d", defaultPort)
}

func defaultWaspPort(kind string, i int) int {
	switch kind {
	case hostKindNanomsg:
		return 5550 + i
	case hostKindPeering:
		return 4000 + i
	case hostKindApi:
		return 9090 + i
	}
	panic(fmt.Sprintf("no handler for kind %s", kind))
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	check(viper.WriteConfig())
}

func SetSCAddress(scName string, address string) {
	Set(scName+".address", address)
}

func TrySCAddress(scName string) *address.Address {
	b58 := viper.GetString(scName + ".address")
	if len(b58) == 0 {
		return nil
	}
	address, err := address.FromBase58(b58)
	check(err)
	return &address
}

func GetSCAddress(scName string) *address.Address {
	address := TrySCAddress(scName)
	if address == nil {
		check(fmt.Errorf("call `set <sc>.address` or `<sc> admin init` first"))
	}
	return address
}

func check(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
