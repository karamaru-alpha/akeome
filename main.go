package main

import (
	"fmt"
	"strings"
	"syscall"
	"time"

	"github.com/gosuri/uilive"
	"golang.org/x/term"
)

func main() {
	og := []string{
		"                                                                                    .....-J&uyy&+(...              ",
		"                                                                              ..JY'!                 ?'n,.  ....   ",
		"                                                                     ..ZY6a.,'! ...-...        .(7??7^(. ?WP!  _Ta.",
		"                                                                   .f!    .t  ,'       =,    .=        .1  (h...  b",
		"    (,            ..     ..         .,                   ..        J:  (M#! .r ..       .,  .l....       t   HM'  d",
		"....M&...Jg%     .#      .N         -F  . .Q,      ,#    (F        ,[  .M`  ,.MMMb       +  ,,MM#H       (    N .,3",
		"    M   .,       (F  .....MJgF   d''W#''=   TN,     Mg''TM'Mg.      ,S..l   ./WMR>       [   t?'9^      .r    ,#^  ",
		"  ..MT``MHa,     d%       M         -F ....       .#Tb  d$  .M,        #     (,        ./....,=,       .^      d.  ",
		".d' N .d`  Tb    M).,    .M       ..M#'7!~?TN,   .#  M,(F    (F       .]       ?<....,7_,MMMMM<._7<??7`        ,]  ",
		"d%  M.M'   .#    dLM`    .F      JD -F .p   ,#   -b  .M%    .M`       (:              ,!   `    3               b  ",
		"Ua..M'   .(#!    (M%   ..B`      Wa.-F .N,..M3    TMM'   ..dB'        d`              (...?j:.../               N  ",
		"  `   .j=`        !   ,'!          ?7`   _?!            ?=`           d                 .` (  t                 H  ",
		"                                                                      d.                ,,.?..%                 W  ",
		"                                                                      J;   `  `                                 W  ",
	}

	// ogのwidth取得
	ogw := len(og[0])

	// terminalサイズ取得
	tw, _, err := term.GetSize(syscall.Stdin)
	if err != nil {
		panic(err)
	}

	// terminal分の空白をogの各行に追加
	for i := range og {
		og[i] = strings.Repeat(" ", tw) + og[i]
	}

	// 描画
	writer := uilive.New()
	writer.Start()
	for i := 0; i <= tw+ogw; i++ {
		var out string
		for _, v := range og {
			if i < ogw {
				out += v[i:tw+i] + "\n"
			} else if i < len(v) {
				out += v[i:] + "\n"
			} else {
				out += "\n"
			}
		}
		fmt.Fprint(writer, out)
		time.Sleep(10 * time.Millisecond)
	}
	writer.Stop()
}
