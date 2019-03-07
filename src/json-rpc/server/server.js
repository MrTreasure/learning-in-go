const jayson = require('jayson')

const server = jayson.server({
  add(args, callback) {
    console.log(args)
    callback(null, 11)
  }
})

server.tcp().listen(3000)