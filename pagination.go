package coinbase

type APIPagination struct {
 	Order string
	Starting_after string
	Ending_before string
	Total int
	Limit int
	Previous_uri string
	Next_uri string
	Yielded int
	Cursor_range []string
}
