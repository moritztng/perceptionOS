import { PrismaClient } from '@prisma/client'
import { ApolloServer } from '@apollo/server'
import { startStandaloneServer } from '@apollo/server/standalone'

const prisma = new PrismaClient()

const typeDefs = `#graphql
  type Image {
    id: Int!
    filename: String!
    faceDetected: FaceDetected

  }
  type FaceDetected {
    faceDetected: Boolean!
  }
  type Query {
    images: [Image]
    imagesUnprocessed: [Image]
  }
  type Mutation {
    addImage(filename: String!): Image
    addFaceDetected(imageId: Int!, faceDetected: Boolean!): FaceDetected
  }
`

const resolvers = {
  Query: {
    images: () => {
      return prisma.image.findMany({
        include: {
          faceDetected: true,
        },
      });
    },
    imagesUnprocessed: () => {
      return prisma.image.findMany({
        where: {
          faceDetected: null,
        },
      });
    },
  },
  Mutation: {
    addImage: (parent, args, contextValue, infos) => { 
      return prisma.image.create({
        data: {
          filename: args.filename,
        },
      });
    },
    addFaceDetected: (parent, args, contextValue, info) => {
      return prisma.faceDetected.create({
        data: {
          imageId: args.imageId,
          faceDetected: args.faceDetected,
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
