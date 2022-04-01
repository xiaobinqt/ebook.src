module.exports = function (RED) {
    function Sum(config) {
        RED.nodes.createNode(this, config);
        var node = this;

        // 获取输入的参数
        let add1 = parseInt(config.add1)
        let add2 = parseInt(config.add2)
        node.send({ // 向下一个节点输出信息
            payload: `${add1} + ${add2} 结果为 ` + (add1 + add2)
        });

        node.on('input', function (msg) { // 接收上游节点接收消息

        });
    }

    RED.nodes.registerType("sum", Sum);
}