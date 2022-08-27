module.exports = function (RED) {
    function Fn(config) {
        RED.nodes.createNode(this, config); // 创建节点对象
        var node = this;
        // node.send(msg);// 直接将数据对象发送到下一个节点
        let counts = 0;
        let systemTimeLast;
        var systemTimeNow;
        let worldTimeLast;
        let worldTimeNow = new Date();
        let timeDifferents = {
            "name": "tvdiff", "val": -1,
        }
        let newSeconds = {
            "name": "isNewSeconds", "val": false,
        }
        node.on('input', function (msg) { // 注册一个侦听器的输入事件流从上游节点接收消息
            let data = JSON.parse(JSON.stringify(msg.payload));
            counts++;
            let systemTime = new Date(data.timestamp);
            if (systemTimeNow !== undefined) {
                systemTimeLast = systemTimeNow;
            }

            systemTimeNow = systemTime;

            if (systemTimeLast !== systemTimeNow) {
                if (worldTimeNow !== undefined) {
                    worldTimeLast = worldTimeNow;
                }
                worldTimeNow = new Date();

                if (counts < 2) {
                    timeDifferents.val = -2;
                } else {
                    tvdiff = (worldTimeNow.getSeconds() > worldTimeLast.getSeconds()) ? (worldTimeNow.getSeconds() - worldTimeLast.getSeconds()) : 0;
                    timeDifferents.val = tvdiff;
                    newSeconds.val = true;
                }
            }

            data.body.push(timeDifferents, newSeconds);
            node.send({
                payload: {
                    "data": data,
                    "time": data.timestamp,
                    "syst": Date.parse(systemTime.toString()),
                    "stl": systemTimeLast,
                    "stn": systemTimeNow,
                    "wtl": worldTimeLast,
                    "wtn": worldTimeNow,
                }
            })
        });
    }

    RED.nodes.registerType("check_time", Fn); //绑定函数
}
