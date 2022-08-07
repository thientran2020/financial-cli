package models

var (
	START_YEAR = 2017
	HEADERS    = []string{"#", "DATE", "DESCRIPTION", "CATEGORY", "COST"}
	LABELS     = map[string]map[string]string{
		"Expense": {
			"Description": "What did you spend money for",
			"Cost":        "Nice. How much did you spend",
		},
		"Income": {
			"Description": "What did you work for today",
			"Cost":        "Awesome. How much did you earn",
		},
	}

	CATEGORY = []string{
		"Income",
		"Rent",
		"Utilities",
		"Insurance",
		"Vehicle Services",
		"Fuel - Car Wash",
		"Subcription",
		"Restaurants",
		"Amazon Shopping",
		"Merchandise",
		"Travel",
		"Personal",
	}
)

const (
	BASE_FILEPATH             = "/finance/finance<YEAR>.csv"
	BASE_FILEPATH_SUBCRIPTION = "/finance/subscription.json"
	DASH                      = "+-----+-----+-----+------------------------------------+--------+-------------------+"

	INSTRUCTION = `
---------------------------------------------------------------------------------
				FINANCIAL CLI
   		 A great tool helps you manage your financial data ^^
---------------------------------------------------------------------------------

COMMANDS
	- data add
	USAGE: add financial data (expense or income)
	FLAGS:
		-s: add subscription/membership data 
		    this will be handled automatically based on its billing cycle
	NOTE: 
		This command will update subscriptions automatically.
		Please run once if you see your subscriptions/memberships are not renewed yet in the financial table.

	- data show
	USAGE: display financial data in table format - current date by default 
	FLAGS:
	  -i: show only income data
		-e: show only expense data
		-m: show only data filtered by specified month
			Ex: 1, 2, 10,....
		-y: show only data filtered by specified year
			Ex: 2021, 2022,...
		-k: show only data filtered by specified keyword 
		-h: show default instruction
	
	- data get
	USAGE: get category mapping table or subscriptions' details
	FLAGS:
		-c: display category map for specific details
		-s: display subscriptions' details
	
	- data search <keyword>
	USAGE: search specific financial key

	- data help
	USAGE: display app's instruction

---------------------------------------------------------------------------------
`

	CATEGORY_TABLE = `
			CATEGORY TABLE
	|------|-----------------------------------------|
	|  0   |  Income                                 |
	|------|-----------------------------------------|
	|  1   |  Rent                                   |
	|------|-----------------------------------------|
	|  2   |  Ultilites (AT&T, PG&E, Water)          |
	|------|-----------------------------------------|
	|  3   |  Insurance (Allstate, Medical, Dental)  |
	|------|-----------------------------------------|
	|  4   |  Vehicle Services                       |
	|------|-----------------------------------------|
	|  5   |  Fuel & Car Wash                        |
	|------|-----------------------------------------|
	|  6   |  Subscription & Membership              |
	|------|-----------------------------------------|
	|  7   |  Restaurants                            |
	|------|-----------------------------------------|
	|  8   |  Amazon Shopping                        |
	|------|-----------------------------------------|
	|  9   |  Merchandise                            |
	|------|-----------------------------------------|
	|  10  |  Travel                                 |
	|------|-----------------------------------------|
	|  11  |  Personal                               |
	|------|-----------------------------------------|

`
)
