const WebSocket = require('ws')

ws = new WebSocket('ws://47.75.66.40:5002/ws' + "/ihuobi");
ws.onopen = evt => {
    console.info("connection");
    // console.info(evt);
};
ws.onmessage = evt => {
    let json = JSON.parse(evt.data);
    if (json.channel == "OKEX-HUOBI") {
        console.info("OKEX-HUOBI")
    } else if (json.channel == "OKEX-OKEX") {
        console.info("OKEX-OKEX")
    } else if (json.channel == "OKEX-BFX") {
        console.info("OKEX-BFX");
    } else if (json.channel == "HUOBI-OKEX") {
        console.info("HUOBI-OKEX")
    } else if (json.channel == "HUOBI-BFX") {
        console.info("HUOBI-BFX");
    } else if (json.channel == "HUOBI-HUOBI") {
        console.info("HUOBI-BFX");
    }

    else if (json.channel == "IHUOBI-OKEX") {
        console.info("IHUOBI-OKEX")
    } else if (json.channel == "IHUOBI-BFX") {
        console.info("IHUOBI-BFX");
    } else if (json.channel == "IHUOBI-HUOBI") {
        console.info("IHUOBI-BFX");
    }

    else if (json.channel == "IOKEX-HUOBI") {
        console.info("IOKEX-HUOBI")
    } else if (json.channel == "IOKEX-OKEX") {
        console.info("IOKEX-OKEX")
    } else if (json.channel == "IOKEX-BFX") {
        console.info("IOKEX-BFX");
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