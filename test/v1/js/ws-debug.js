const WebSocket = require('ws')

ws = new WebSocket('ws://47.75.66.40:5002/ws' + "/huobi");

ws.onmessage = evt => {
    let json = JSON.parse(evt.data);
    json.forEach(element => {
        if(element != null) {
            if (element.channel == "HUOBI-HUOBI") {
                console.info(element.data)
                // console.info("future_ts: " + element.data.future_ts + " stock_ts:" + element.data.stock_ts)
            } 
        }
    })
};