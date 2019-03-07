const jayson = require('jayson')

const client = jayson.Client.tcp({
  port: 8000
})

client.request('Arith.Multiply', [{
  A: 10,
  B: 12
}], function(err, res) {
  if (err) {
    console.log(err)
    return
  }
  console.log(res)
})

client.request('Arith.Divide', [{
  A: 10,
  B: 0
}], function(err, res) {
  if (err) {
    console.log(err)
    return
  }
  console.log(res)
})