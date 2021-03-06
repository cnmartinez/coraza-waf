package operators
import(
	"github.com/jptosso/coraza-waf/pkg/engine"
)

type ValidateUrlEncoding struct{
	data string
}

func (o *ValidateUrlEncoding) Init(data string){
	
}

func (o *ValidateUrlEncoding) Evaluate(tx *engine.Transaction, value string) bool{
    if len(value) == 0 {
        return true
    }

    rc := validateUrlEncoding(value)
    switch (rc) {
        case 1 :
            /* Encoding is valid */
            return true
        case -2 :
            // Invalid URL Encoding: Non-hexadecimal 
            return false
        case -3 :
        	//Invalid URL Encoding: Not enough characters at the end of input
            return false
        default :
            //Invalid URL Encoding: Internal error
            return false
    }
    return true
}

func validateUrlEncoding(input string) int {
    i := 0
    input_length := len(input)

    for i < input_length {
        if (input[i] == '%') {
            if i + 2 >= input_length {
                /* Not enough bytes. */
                return -3
            }
            /* Here we only decode a %xx combination if it is valid,
             * leaving it as is otherwise.
             */
            c1 := input[i + 1];
            c2 := input[i + 2];

            if (((c1 >= '0') && (c1 <= '9')) || ((c1 >= 'a') && (c1 <= 'f')) || ((c1 >= 'A') && (c1 <= 'F'))) && (((c2 >= '0') && (c2 <= '9')) || ((c2 >= 'a') && (c2 <= 'f')) || ((c2 >= 'A') && (c2 <= 'F'))) {
                i += 3
            } else {
                /* Non-hexadecimal characters used in encoding. */
                return -2
            }
        } else {
            i++
        }
    }
    return 1
}