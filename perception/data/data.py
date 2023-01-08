from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport

class Data:
    def __init__(self, url):
        self.transport = AIOHTTPTransport(url=url)
        self.client = Client(transport=self.transport, fetch_schema_from_transport=True)
    
    def get_images(self):
        query = gql(
            """
            query Query {
                images {
                    id
                    filename
                    faceDetected {
                        faceDetected
                    }
                }
            }
            """
        )
        result = self.client.execute(query)
        return result

if __name__ == "__main__":
    data = Data("http://localhost:4000")
    print(data.get_images())
