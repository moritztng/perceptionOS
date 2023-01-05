from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport

class Data:
    def __init__(self):
        self.transport = AIOHTTPTransport(url="https://countries.trevorblades.com/")
        self.client = Client(transport=transport, fetch_schema_from_transport=True)
    def __call__(self, filename):
        query = gql(
            """
            query getContinents {
            continents {
                code
                name
            }
            }
        """
        )
        result = client.execute(query)
        print(result)
