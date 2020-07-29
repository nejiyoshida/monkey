package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/nejiyoshida/monkey/lexer"
	"github.com/nejiyoshida/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // 一行ずつ読み込むやつ。
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan() // scanner.Scan()で次の行へ読み進めるかどうか戻す。
		if !scanned {
			return // 何も読めない（終わり）なのでreturn
		}

		line := scanner.Text() // 直近で読み込んだtokenを戻す
		l := lexer.New(line)   // ↑で読み込んだ部分を字句解析にかける

		// EOFがくるまで読み進める
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok) // 字句解析器が返したものを表示
		}
	}
}
