# 一个 outputs 连接多条线问题

当我们在开发节点时，如果设置了 `outputs: 1` 并且连接了多条线，比如：

js 代码

```js
module.exports = function (RED) {
    function Fn(config) {
        RED.nodes.createNode(this, config); // 创建节点对象
        var node = this;
        // node.send(msg);// 直接将数据对象发送到下一个节点

        node.on('input', function (msg) {// 注册一个侦听器的输入事件流从上游节点接收消息
            msg.payload = {start: -1}
            for (let j = 0; j < 2; j++) {
                msg.payload.start = j
                node.send(msg)
            }
        });
    }

    RED.nodes.registerType("start_topo", Fn); //绑定函数
}

```

html 代码

```html

<script type="text/javascript">
    RED.nodes.registerType('start_topo', {
        category: '开关监测',                      //category: （字符串）节点的分类类别
        color: '#a6bbcf',                     //color: (字符串)节点要使用的背景颜色
        defaults: {                           //defaults: （对象）节点的可编辑属性
            name: {value: "开启topo功能"}
        },
        inputs: 1,                             //inputs: （数字）的节点有多少输入具有，无论是0或1
        outputs: 1,                            //outputs: (数字) 节点有多少个输出。可以0或更多
        icon: "file.png",                     //icon: （字符串）要使用的图标
        paletteLabel: "start_topo",            //paletteLabel: (字符串|函数)控件区中显示的节点名称
        label: function () {                   //label: (字符串|函数)工作区中显示的节点名称
            if (this.name === "") {
                this.name = this._def.paletteLabel
            }
            return this.name || this._def.paletteLabel
        },
        oneditprepare: function () {          //oneditprepare:（函数）在构建编辑对话框时调用的自定义编辑行为
        },
        oneditsave: function () {             //oneditsave: （函数）在编辑对话框正常时调用的自定义编辑行为
        }
    });
</script>

<script type="text/html" data-template-name="start_topo">
    <div class="form-row">
        <label for="node-input-name"><i class="fa fa-tag"></i> Name</label>
        <input type="text" id="node-input-name" placeholder="Name">
    </div>
</script>

<style>
    pre {
        width: 500px;
    }
</style>


<script type="text/markdown" data-help-name="start_topo">

</script>
```

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230421/d64d3cd94b324ecdae8b8f9afe10ba93.png)

[图1]

在注入数据时就会出现如上 “问题”，查看官方文档可知

官方文档 [Sending messages asynchronously](https://nodered.org/docs/user-guide/writing-functions#node)

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230421/9241e5642849434b9096153c1919a612.png)

其中第一条信息没有 clone 也就是**没有深拷贝**，但是剩下的信息进行深拷贝了，从 [图1]
中红色标识，也就是输出到节点`0b6acbef559467b3`
的数据，没有进行深拷贝，所以在 `node.send()` 异步执行后，2 次输出都是 1，而剩下的节点`9f0d6c0cded3cf16`和`f475c8867ec2a0fc`
数据进行了深拷贝，分别输出 0 和 1。

调试的结果也符合预期：

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230421/081cf123573e482baeecf9e88843bc1d.png)

`node.send()`
源码地址 [https://github.com/node-red/node-red/blob/2.0.5/packages/node_modules/%40node-red/runtime/lib/nodes/Node.js#L362-L463](https://github.com/node-red/node-red/blob/2.0.5/packages/node_modules/%40node-red/runtime/lib/nodes/Node.js#L362-L463)

## 建议

开发中根据具体情况，**适时使用深拷贝**，可以直接调用方法`RED.util.cloneMessage()`
，该方法的源代码地址为 [https://github.com/node-red/node-red/blob/2.0.5/packages/node_modules/%40node-red/util/lib/util.js#L87-L107](https://github.com/node-red/node-red/blob/2.0.5/packages/node_modules/%40node-red/util/lib/util.js#L87-L107)

如果上述节点使用深拷贝，改成

```js
node.send(RED.util.cloneMessage(msg))
```

就不会出现上述 “错误”，并且符合我们的开发预期。

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230421/2d4089a42b8a4d9bb5f2e9eb9075c0de.png)
