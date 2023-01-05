import { PrismaClient } from '@prisma/client'
import { ApolloServer } from '@apollo/server'
import { startStandaloneServer } from '@apollo/server/standalone'

const prisma = new PrismaClient()

const typeDefs = `#graphql
  type Image {
    fileName: String
    faceDetected: FaceDetected
  }
  type FaceDetected {
    faceDetected: Boolean
  }
  type Query {
    images: [Image]
  }
`

const resolvers = {
  Query: {
    images: () => {
      return prisma.image.findMany()
    },
  },
  Image: {
    faceDetected: (parent) => {
      return prisma.faceDetected.findMany({
        where: {
          filename: {
            equals: parent.filename,
          },
        },
      })
    },
  },
}

const server = new ApolloServer({
  typeDefs,
  resolvers,
})

const { url } = await startStandaloneServer(server, {
  listen: { port: 4000 },
})
console.log(`🚀  Server ready at: ${url}`)
