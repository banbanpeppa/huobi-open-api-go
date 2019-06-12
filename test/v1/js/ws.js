const WebSocket = require('ws')

ws = new WebSocket('ws://47.75.66.40:5001/ws' + "/okex");
ws.onopen = evt => {
    console.info("connection");
    // console.info(evt);
};
ws.onmessage = evt => {
let json = JSON.parse(evt.data);
// console.info(json)
if(json.channel=="OKEX-HUOBI"){
    console.info(json.data)
}else if(json.channel =="HUOBI-HUOBI"){
    console.info(json.data)
} else if (json.channel == "HUOBI-BFX") {
    //console.info(json.data);
}
};
ws.onclose = evt => {
    // console.log("okokws close" + evt);
    if (ws != null) {
        ws.close();
    }
};
ws.onerror = evt => {
    // console.info(evt.code);
};