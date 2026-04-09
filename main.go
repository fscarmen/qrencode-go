package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	// 定义命令行参数
	level := flag.String("level", "L", "Error correction level: L/M/Q/H")
	qz := flag.Int("qz", 1, "Quiet zone size (default 1)")
	compact := flag.Bool("compact", true, "Use half-height rendering")
	flag.Parse()

	// 校验参数数量
	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	// 获取输入内容
	content := strings.TrimSpace(flag.Arg(0))
	if content == "" {
		usage()
		os.Exit(1)
	}

	// 解析纠错等级
	qrLevel, err := parseLevel(*level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 生成二维码对象
	qr, err := qrcode.New(content, qrLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 关闭库默认边框，改为自定义控制
	qr.DisableBorder = true

	// 渲染输出
	matrix := qr.Bitmap()
	DrawASCIIBitmap(matrix, *qz, *compact, true)
}

// usage prints command usage information.
func usage() {
	fmt.Fprintln(os.Stderr, "USAGE: qrencode-go [options] <url>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	fmt.Fprintln(os.Stderr, "  -level L/M/Q/H   Error correction level (default L)")
	fmt.Fprintln(os.Stderr, "  -qz int          Quiet zone size (default 1)")
	fmt.Fprintln(os.Stderr, "  -compact         Half-height rendering (default true)")
}

// parseLevel converts string to qrcode.RecoveryLevel.
func parseLevel(s string) (qrcode.RecoveryLevel, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "L":
		return qrcode.Low, nil
	case "M":
		return qrcode.Medium, nil
	case "Q":
		return qrcode.High, nil
	case "H":
		return qrcode.Highest, nil
	default:
		return qrcode.Low, fmt.Errorf("invalid level %q", s)
	}
}