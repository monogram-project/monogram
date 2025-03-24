# A simple example showing how a JSON example might be adapted in Monongram.
#
# {
#   "database": {
#     "type": "postgres",
#     "host": "db.example.com",
#     "port": 5432,
#     "username": "admin",
#     "password": "s3cret",
#     "database": "example_db",
#     "ssl": true
#   }
# }
#

# In our new version we avoid the outer superfluous braces in favour 
# of something that looks like a function call - suggesting an implementation
# strategy! The string quotes around the field names are discarded in favour
# of a form that looks like keyword arguments. 

database(
  type=postgres,          # Hints that the RDBMS types are defined elsewhere.
  host="db.example.com",
  port=5432,
  username="admin",
  secret="s3cret"
  database: "example_db",
  ssl=true
)

