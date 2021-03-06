package operators
import(
	"strings"
	"sync"
	"fmt"
	ahocorasick"github.com/bobusumisu/aho-corasick"
	"github.com/jptosso/coraza-waf/pkg/engine"
	"github.com/jptosso/coraza-waf/pkg/utils"
)


type PmFromFile struct{
	Data []string
	mux *sync.RWMutex
}

func (o *PmFromFile) Init(data string){
	o.Data = []string{}
	o.mux = &sync.RWMutex{}
    b, err := utils.OpenFile(data)
    content := string(b)
    if err != nil{
    	fmt.Println("Error parsing path " + data)
    	return
    }
    sp := strings.Split(string(content), "\n")
    for _, l := range sp {
    	if len(l) == 0{
    		continue
    	}
    	l = strings.ReplaceAll(l, "\r", "") //CLF
    	if l[0] != '#'{
    		o.Data = append(o.Data, l)
    	}
    }
}

func (o *PmFromFile) Evaluate(tx *engine.Transaction, value string) bool{
	o.mux.RLock()
	defer o.mux.RUnlock()
	trie := ahocorasick.NewTrieBuilder().
	    AddStrings(o.Data).
	    Build()
	matches := trie.MatchString(value)
	return len(matches) > 0
}

func (o *PmFromFile) GetType() string{
	return ""
}