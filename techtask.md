### FEATURES ###
-Should work as an api
-Should be able to scrap data from web pages by link or by a list of links. 
-Should scrap data from elements by css selectors. 
-Should parse a json or csv table with results (preferably json).
-Should have an authentication and handle payments in future


### PLAN ###
-Create one endpoint which will receive text data with links and tags should scrapped.
-Make this endpoint to parse json with : link, tags and their content, timestamp. 
-Make register and auth endpoint. 
-Collect all the data about scrap request and result on db (Postgres)