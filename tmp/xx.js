var msg = `{
"token":"MCCB real-time data 1",
"timestamp":"2022-07-11T09:59:31.918+0800",
"data_row":"single",
"body":[{
"name":"PhV_phsA",
"val":"232.100006",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"PhV_phsB",
"val":"232.199997",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"PhV_phsC",
"val":"232.199997",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"A_phsA",
"val":"  0.000000",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"A_phsB",
"val":"  0.000000",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"A_phsC",
"val":"  0.000000",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}, {
"name":"ResA",
"val":"  0.000000",
"quality":"0",
"timestamp":"2022-07-11T09:59:31.917+0800"
}]
}`

let payload = JSON.parse(msg)
if (payload.body !== undefined && payload.body.length > 0) {
    for (each in payload.body) {
        console.log(payload.body[payload.body].name)
    }
}
console.log(payload)