type options struct {
	token   string
	qrcode  string
	config  string
	help    bool
	version bool
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.help, "help", "h", false, "print this message")      // 自分の必要なものを書く（helpについて）
	flags.BoolVarP(&opts.version, "version", "v", false, "print the massage") // 自分の必要なものを書く（versionについて）
}
func perform(opts *options, args []string) *SleetEror {
	fmt.Println("Hello World")
	return nil
}
func parseOptions(args []string) (*options, []string, *SleetEror) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMassage(args[0]))
		return nil, nil, &SleetEror{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil,
			&SleetEror{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}
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