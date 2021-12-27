package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(`
	  __
_/  |_  ____   ____   ____
\   __\/  _ \ / ___\ /  _ \
 |  | (  <_> ) /_/  >  <_> )
 |__|  \____/\___  / \____/
            /_____/
	`)
}
