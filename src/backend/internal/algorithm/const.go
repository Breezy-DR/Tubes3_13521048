package algorithm

const CalculatorRegexString string = "\\s*(([\\+\\-\\*\\/()^]|\\d)+\\s*)*\\s*"
const DateRegexString string = "((\\d{4}[\\/-]\\d{1,2}[\\/-]\\d{1,2})|(\\d{1,2}[\\/-]\\d{1,2}[\\/-]\\d{4}))"
const AddRegexString string = "(?i)tambahkan pertanyaan\\s+[\\x00-\\x7F]+[\\x00-\\x7F\\s]*\\s+dengan\\s+jawaban\\s+[\\x00-\\x7F]+[\\x00-\\x7F\\s]*"
const RemoveRegexString string = "(?i)\\s*hapus\\s+pertanyaan\\s+[\\x00-\\x7F\\s]+[\\x00-\\x7F\\s\\s]*"

var dates = []string{"(\\d{4}[\\/-]\\d{1,2}[\\/-]\\d{1,2})", "(\\d{1,2}[\\/-]\\d{1,2}[\\/-]\\d{4})"}

const KMP = "kmp"
const BM = "bm"
