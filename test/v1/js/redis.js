var redis = require('redis');



// function sleep(delay) {
//     var start = (new Date()).getTime();
//     while ((new Date()).getTime() - start < delay) {
//         continue;
//     }
// }
while (true){
    var client = redis.createClient(6380, '47.75.66.40');
    client.get('future:nw:HUOBI:BTC', function (err, v) {
        if (err) {
            console.log(err)
        } else {
            console.log(v);
            client.end(true);
        }
    });
    
}



