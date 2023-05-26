package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kimura-hytaae/sleet"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.0"

func versionString(args []string) string {
	prog := "sleet"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

/*
ヘルプメッセージ
*/
func helpMessage(args []string) string {
	prog := "sleet"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [LOCATIONs...] [DAYs...] [WEATHERs]
	OPTIONS
	-v, --version            print the version and exit.
	-h, --help               print this mesasge and exit.

	LOCATION
	specify the location in the following ways. 次の方法などで指定します。
	- 緯度経度
	- ポスタルコード
	- 市の名前
	DAY
	specify the location in the following ways. 次の方法などで指定します。
	- 日付指定
	- 何日かの指定`, prog)
}

type SleetError struct {
	statusCode int
	message    string
}

func (e SleetError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag    bool
	listGroupFlag bool
	helpFlag      bool
	versionFlag   bool
}

type runOpts struct {
	token  string
	qrcode string
	config string
	group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

/*
sleet.goの中身を変えて必要な部分だけ残す必要あり
*/
func (opts *options) mode(args []string) sleet.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return sleet.ListGroup
	case len(args) == 0:
		return sleet.List
	case opts.flagSet.deleteFlag:
		return sleet.Delete
	case opts.runOpt.qrcode != "":
		return sleet.QRCode
	default:
		return sleet.Shorten
	}
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "print this mesasge and exit.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *SleetError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &SleetError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &SleetError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &SleetError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

/*
修正する必要がある
とりあえず、bitlyの部分を open_meteo に変更
urleapをsleetに変更
Shortenを Weather
*/
func shortenEach(open_meteo *sleet.Open_meteo, config *sleet.Config, localcite string) error {
	result, err := open_meteo.Weather(config, localcite)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

/*
urleapの部分をsleetに変更
urlの部分をlocalciteに変更
*/
func deleteEach(open_meteo *sleet.Open_meteo, config *sleet.Config, localcite string) error {
	return open_meteo.Delete(config, localcite)
}

/*
bitlyとurleap、url、urlsをopen_meteoとsleet、localcite、localcitesに変更
*/
func listUrls(open_meteo *sleet.Open_meteo, config *sleet.Config) error {
	urls, err := open_meteo.List(config)
	if err != nil {
		return err
	}
	for _, localcite := range localcites {
		fmt.Println(localcite)
	}
	return nil
}

/*
bitlyとurleap、url、urlsをopen_meteoとsleet、localcite、localcitesに変更
*/
func listGroups(open_meteo *sleet.Open_meteo, config *sleet.Config) error {
	groups, err := open_meteo.Groups(config)
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}

/*
UrleapErrorをSleetErrorに変更
urlをlocalciteに変更
*/
func performImpl(args []string, executor func(url string) error) *SleetError {
	for _, localcite := range args {
		err := executor(localcite)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

/*
bitlyとurleap、url、urlsをopen_meteoとsleet、localcite、localcitesに変更
UrleapErrorをSleetError
*/
func perform(opts *options, args []string) *SleetError {
	/*
	open_meteo := sleet.NewOpen_meteo(opts.runOpt.group)
	config := sleet.NewConfig(opts.runOpt.config, opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case sleet.List:
		err := listUrls(open_meteo, config)
		return makeError(err, 1)
	case sleet.ListGroup:
		err := listGroups(open_meteo, config)
		return makeError(err, 2)
	case sleet.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(open_meteo, config, url)
		})
	case sleet.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(open_meteo, config, url)
		})
	}
	*/
	fmt.Println("Hello World.")
	return nil
}

/*
UrleapErrorをSleetErrorに変更
*/
func makeError(err error, status int) *SleetError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*SleetError)
	if ok {
		return ue
	}
	return &SleetError{statusCode: status, message: err.Error()}
}

/*
mainから
*/
func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
