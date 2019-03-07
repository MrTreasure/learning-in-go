const jayson = require('jayson')

const client = jayson.Client.tcp({
  port: 3000
})

// client.request('Arith.Multiply', [{
//   A: 10,
//   B: 12
// }], function(err, res) {
//   if (err) {
//     console.log(err)
//     return
//   }
//   console.log(res)
// })

// client.request('Arith.Divide', [{
//   A: 10,
//   B: 0
// }], function(err, res) {
//   if (err) {
//     console.log(err)
//     return
//   }
//   console.log(res)
// })

client.request('add', [2, 3], function(err, response) {
  if(err) throw err;
  console.log(response.result) // 2
})